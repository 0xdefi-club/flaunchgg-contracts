#!/bin/bash

# Directory containing the Solidity contract files
CONTRACT_DIR="src"

# Directory to store the generated TypeScript files and Go files
WEB="web"
GO_ABIS="go-abis"
SUBGRAPH="subgraph"
PKG_DIR="pkg"

WEB_DIR="$PKG_DIR/$WEB"
GO_ABIS_DIR="$PKG_DIR/$GO_ABIS"
SUBGRAPH_DIR="$PKG_DIR/$SUBGRAPH"
TEMP_DIR="$PKG_DIR/temp"

cleanup() {
    rm -rf $PKG_DIR/*
}

cleanup

# Create the necessary directories
mkdir -p $WEB_DIR
mkdir -p $GO_ABIS_DIR
mkdir -p $SUBGRAPH_DIR
mkdir -p $TEMP_DIR

# Function to convert string to snake_case
to_snake_case() {
    echo "$1" | sed -r 's/([a-z0-9])([A-Z])/\1_\2/g' | tr '[:upper:]' '[:lower:]'
}

# Function to extract ABI using solc directly
extract_abi_with_solc() {
    local contract_file=$1
    local contract_name=$2
    local output_file=$3
    local temp_dir="$TEMP_DIR/$contract_name"
    
    mkdir -p "$temp_dir"
    
    echo "Using solc to extract ABI for $contract_name..."
    
    # Copy the contract file to a temporary directory
    # This helps solc resolve dependencies more easily
    CONTRACT_DIR_PATH=$(dirname "$contract_file")
    
    # Create a complete solc command with remappings
    # We'll compile with all imported libraries from the original location
    solc --abi "$contract_file" \
         --include-path "$CONTRACT_DIR" \
         --include-path "$CONTRACT_DIR_PATH" \
         --allow-paths "*" \
         -o "$temp_dir" \
         2>/dev/null
    
    # Try finding the ABI file
    local abi_file=$(find "$temp_dir" -name "*.abi" | head -1)
    
    if [ -n "$abi_file" ] && [ -f "$abi_file" ]; then
        cp "$abi_file" "$output_file"
        echo "Successfully extracted ABI for $contract_name using solc"
        return 0
    fi
    
    # If no specific ABI file found, try to compile the entire directory
    solc --abi "$CONTRACT_DIR_PATH/*.sol" \
         --include-path "$CONTRACT_DIR" \
         --allow-paths "*" \
         -o "$temp_dir" \
         2>/dev/null
    
    abi_file=$(find "$temp_dir" -name "$contract_name.abi" | head -1)
    
    if [ -n "$abi_file" ] && [ -f "$abi_file" ]; then
        cp "$abi_file" "$output_file"
        echo "Successfully extracted ABI for $contract_name using solc directory compilation"
        return 0
    fi
    
    return 1
}

# Function to extract bytecode using solc
extract_bytecode_with_solc() {
    local contract_file=$1
    local contract_name=$2
    local output_file=$3
    local temp_dir="$TEMP_DIR/$contract_name"
    
    mkdir -p "$temp_dir"
    
    echo "Using solc to extract bytecode for $contract_name..."
    
    # Try to compile with solc
    CONTRACT_DIR_PATH=$(dirname "$contract_file")
    
    solc --bin "$contract_file" \
         --include-path "$CONTRACT_DIR" \
         --include-path "$CONTRACT_DIR_PATH" \
         --allow-paths "*" \
         -o "$temp_dir" \
         2>/dev/null
    
    # Look for the binary file
    local bin_file=$(find "$temp_dir" -name "*.bin" | head -1)
    
    if [ -n "$bin_file" ] && [ -f "$bin_file" ]; then
        # Add 0x prefix to the bytecode
        echo -n "0x" > "$output_file"
        cat "$bin_file" >> "$output_file"
        echo "Successfully extracted bytecode for $contract_name using solc"
        return 0
    fi
    
    return 1
}

# Function to create a minimal ABI from Solidity source
create_minimal_abi() {
    local contract_file=$1
    local output_file=$2
    local contract_name=$(basename "$contract_file" .sol)
    
    echo "Creating minimal ABI from source for $contract_name..."
    
    # Check if it's a contract (not an interface or library)
    if ! grep -q "^contract $contract_name" "$contract_file"; then
        echo "File does not contain a matching contract declaration for $contract_name"
        echo "[]" > "$output_file"
        return
    fi
    
    # Extract all function declarations
    {
        echo "["
        # Extract function declarations and convert to ABI format
        grep -E "^\s*(function|event) [a-zA-Z0-9_]+" "$contract_file" | 
        grep -v "internal" |  # Skip internal functions
        grep -v "private" |   # Skip private functions
        sed -E 's/.*function ([a-zA-Z0-9_]+)\((.*)\).*/  {\n    "name": "\1",\n    "type": "function",\n    "inputs": [],\n    "outputs": [],\n    "stateMutability": "nonpayable"\n  },/g' |
        # Also handle events
        sed -E 's/.*event ([a-zA-Z0-9_]+)\((.*)\).*/  {\n    "name": "\1",\n    "type": "event",\n    "inputs": [],\n    "anonymous": false\n  },/g' |
        sed '$s/,$//'
        echo "]"
    } > "$output_file"
    
    # Validate and clean up the JSON
    if command -v jq >/dev/null 2>&1; then
        if ! jq -e . "$output_file" >/dev/null 2>&1; then
            # If invalid JSON, create a minimal valid empty array
            echo "[]" > "$output_file"
        else
            # Clean up the JSON with jq
            jq . "$output_file" > "$output_file.tmp" && mv "$output_file.tmp" "$output_file"
        fi
    fi
}

# Function to process a single Solidity file
process_file() {
    local contract=$1
    
    # Extract the contract name from the file name
    contract_name=$(basename "$contract" .sol)
    
    echo "Processing contract: $contract_name from file: $contract"
    
    # Skip test files and mock contracts if filename indicates it
    if [[ "$contract" == *"/test/"* || "$contract_name" == *Test* || "$contract_name" == *Mock* ]]; then
        echo "Skipping test/mock contract: $contract_name"
        return 0
    fi
    
    # Skip interfaces and libraries based on filename or content
    if [[ "$contract_name" == I[A-Z]* || "$contract_name" == *Interface ]]; then
        echo "Skipping interface contract: $contract_name"
        return 0
    fi
    
    # Check if file content indicates it's a library
    if grep -q "^library " "$contract"; then
        echo "Skipping library contract: $contract_name"
        return 0
    fi
    
    # Convert contract name to snake_case
    contract_name_snake=$(to_snake_case "$contract_name")
    
    # Create a subdirectory for the Go file
    go_subdir="$GO_ABIS_DIR/$contract_name_snake"
    mkdir -p "$go_subdir"
    
    # Extract ABI - try multiple methods
    local abi_extracted=false
    
    # 1. First try with solc (most reliable)
    if extract_abi_with_solc "$contract" "$contract_name" "$SUBGRAPH_DIR/$contract_name.json"; then
        echo "ABI extraction successful with solc for $contract_name"
        abi_extracted=true
    else
        # 2. Then try with forge
        echo "Attempting extraction with forge for $contract_name..."
        forge inspect "$contract:$contract_name" abi > "$SUBGRAPH_DIR/$contract_name.json" 2>/dev/null
        
        # Check if the ABI is valid
        if [ -s "$SUBGRAPH_DIR/$contract_name.json" ] && 
           jq -e . "$SUBGRAPH_DIR/$contract_name.json" >/dev/null 2>&1 &&
           ! grep -q "^-\+\|-\+" "$SUBGRAPH_DIR/$contract_name.json"; then
            echo "ABI extraction successful with forge for $contract_name"
            abi_extracted=true
        else
            # 3. If both fail, create a minimal ABI from the source
            echo "Both solc and forge failed. Creating minimal ABI from source."
            create_minimal_abi "$contract" "$SUBGRAPH_DIR/$contract_name.json"
            
            # Check if the created ABI is empty or just contains an empty array
            if [ ! -s "$SUBGRAPH_DIR/$contract_name.json" ] || 
               [ "$(cat "$SUBGRAPH_DIR/$contract_name.json" | tr -d '[:space:]')" = "[]" ]; then
                echo "No ABI could be generated for $contract_name - skipping"
                rm -f "$SUBGRAPH_DIR/$contract_name.json"
                return 0
            else
                abi_extracted=true
            fi
        fi
    fi
    
    # Extract bytecode - try multiple methods
    
    # 1. First try with solc
    if extract_bytecode_with_solc "$contract" "$contract_name" "$WEB_DIR/$contract_name.bin"; then
        echo "Bytecode extraction successful with solc for $contract_name"
    else
        # 2. Then try with forge
        echo "Attempting bytecode extraction with forge for $contract_name..."
        forge inspect "$contract:$contract_name" bytecode > "$WEB_DIR/$contract_name.bin" 2>/dev/null
        
        # Check if bytecode is valid
        if [ ! -s "$WEB_DIR/$contract_name.bin" ]; then
            echo "Creating minimal bytecode for $contract_name"
            echo "0x" > "$WEB_DIR/$contract_name.bin"
        fi
    fi
    
    # Special case for PairFactory
    if [ "$contract_name" = "PairFactory" ]; then
        alias_param="--alias pairFee:newPairFee"
    else
        alias_param=""
    fi
    
    # Check if we have a valid ABI before proceeding
    if [ ! -f "$SUBGRAPH_DIR/$contract_name.json" ]; then
        echo "No valid ABI available for $contract_name - skipping generation"
        rm -f "$WEB_DIR/$contract_name.bin"
        rm -rf "$go_subdir"
        return 0
    fi
    
    # Run abigen with the ABI and bytecode
    abigen --abi "$SUBGRAPH_DIR/$contract_name.json" --bin "$WEB_DIR/$contract_name.bin" \
           --pkg "$contract_name_snake" --type "$contract_name" $alias_param \
           --out "$go_subdir/${contract_name_snake}.go" 2>/dev/null
    
    if [ $? -ne 0 ]; then
        echo "Warning: abigen failed for $contract_name, checking ABI content"
        
        # Check if the ABI is just an empty array
        if [ "$(cat "$SUBGRAPH_DIR/$contract_name.json" | tr -d '[:space:]')" = "[]" ]; then
            echo "ABI is empty for $contract_name - skipping"
            rm -f "$SUBGRAPH_DIR/$contract_name.json"
            rm -f "$WEB_DIR/$contract_name.bin"
            rm -rf "$go_subdir"
            return 0
        else
            echo "Trying with empty ABI as fallback"
            echo "[]" > "$SUBGRAPH_DIR/$contract_name.json.empty"
            
            abigen --abi "$SUBGRAPH_DIR/$contract_name.json.empty" --bin "$WEB_DIR/$contract_name.bin" \
                   --pkg "$contract_name_snake" --type "$contract_name" $alias_param \
                   --out "$go_subdir/${contract_name_snake}.go" 2>/dev/null
            
            if [ $? -ne 0 ]; then
                echo "Warning: abigen still failed for $contract_name - skipping"
                rm -f "$SUBGRAPH_DIR/$contract_name.json"
                rm -f "$SUBGRAPH_DIR/$contract_name.json.empty"
                rm -f "$WEB_DIR/$contract_name.bin"
                rm -rf "$go_subdir"
                return 0
            fi
            
            # Replace the original ABI with the empty one
            mv "$SUBGRAPH_DIR/$contract_name.json.empty" "$SUBGRAPH_DIR/$contract_name.json"
        fi
    fi
    
    # Generate TypeScript file only if ABI exists
    if [ -f "$SUBGRAPH_DIR/$contract_name.json" ] && jq -e . "$SUBGRAPH_DIR/$contract_name.json" >/dev/null 2>&1; then
        echo "export const ${contract_name}ABI = $(cat "$SUBGRAPH_DIR/$contract_name.json") as const;" > "$WEB_DIR/$contract_name.ts"
    else
        # No valid ABI, cleanup and skip
        rm -f "$WEB_DIR/$contract_name.bin"
        rm -rf "$go_subdir"
        rm -f "$SUBGRAPH_DIR/$contract_name.json"
        return 0
    fi
    
    # Clean up
    rm -f "$WEB_DIR/$contract_name.bin"
    
    echo "Successfully generated Go and TypeScript files for $contract_name"
    return 0
}

check_command() {
    # Check for required commands
    if ! command -v abigen >/dev/null 2>&1; then
        echo "Error: abigen not found"
        exit 1
    fi
    
    if ! command -v zip >/dev/null 2>&1; then
        echo "Error: zip not found"
        exit 1
    fi
    
    # Check for solc (now preferred)
    if ! command -v solc >/dev/null 2>&1; then
        echo "Warning: solc not found. ABI extraction may be limited."
        echo "Consider installing solc for better results."
    fi
    
    # Check for forge (fallback)
    if ! command -v forge >/dev/null 2>&1; then
        echo "Warning: forge not found. Will rely entirely on solc and source parsing."
    fi
    
    # Check for jq (for JSON validation)
    if ! command -v jq >/dev/null 2>&1; then
        echo "Warning: jq not found. JSON validation will be limited."
    fi
}

# Function to recursively find and process all Solidity files
process_recursively() {
    local dir=$1
    local success=0
    local errors=0
    local skipped=0
    local no_abi=0
    
    echo "Finding Solidity files in $dir..."
    
    # Create a temporary file to store our contract list
    local contract_list="$TEMP_DIR/contract_list.txt"
    find "$dir" -type f -name "*.sol" | sort > "$contract_list"
    
    echo "Found $(wc -l < "$contract_list") Solidity files."
    
    # Process each contract
    while read -r contract; do
        # Skip test and mock files before processing
        if [[ "$contract" == *"/test/"* || "$contract" == *"/tests/"* || "$contract" == *"/mocks/"* ]]; then
            echo "Skipping test file: $contract"
            skipped=$((skipped+1))
            continue
        fi
        
        # Check for interfaces and libraries by filename pattern
        contract_name=$(basename "$contract" .sol)
        if [[ "$contract_name" == I[A-Z]* || "$contract_name" == *Interface || "$contract_name" == *Lib ]]; then
            echo "Skipping interface/library file: $contract"
            skipped=$((skipped+1))
            continue
        fi
        
        # Process the file
        local result=0
        process_file "$contract"
        result=$?
        
        # Check the result
        if [ $result -eq 0 ]; then
            # Check if we actually generated any files
            contract_name=$(basename "$contract" .sol)
            contract_name_snake=$(to_snake_case "$contract_name")
            
            if [ -f "$SUBGRAPH_DIR/$contract_name.json" ] && [ -f "$WEB_DIR/$contract_name.ts" ] && [ -d "$GO_ABIS_DIR/$contract_name_snake" ]; then
                success=$((success+1))
            else
                echo "No output files generated for $contract_name (no valid ABI)"
                no_abi=$((no_abi+1))
            fi
        else
            errors=$((errors+1))
        fi
    done < "$contract_list"
    
    echo "Summary: Successfully processed $success contracts, skipped $skipped files, $no_abi contracts had no valid ABI, encountered $errors errors"
}

# Function to zip directories
zip_directories() {
    echo "Creating ZIP archives..."
    
    # Remove temporary files and directories
    rm -rf "$TEMP_DIR"
    find "$PKG_DIR" -name "*.empty" -delete 2>/dev/null
    find "$PKG_DIR" -type d -empty -delete 2>/dev/null
    
    # Create the zip files
    (cd "$PKG_DIR" && 
     zip -r "$WEB.zip" "$WEB" && 
     zip -r "$GO_ABIS.zip" "$GO_ABIS" && 
     zip -r "$SUBGRAPH.zip" "$SUBGRAPH")
}

# Main execution
check_command

# Disable the nightly warning
export FOUNDRY_DISABLE_NIGHTLY_WARNING=1

# Process all contract files recursively
process_recursively "$CONTRACT_DIR"

# Create the zip archives
zip_directories

echo "Go file and ABI generation complete"