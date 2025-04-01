# make deploy CHAIN_ID=<chainID>
-include .env

# Define network parameters dynamically based on chain ID
# If CHAIN_ID is not provided, NETWORK_PARAMS will be empty
NETWORK_PARAMS = $(if $(CHAIN_ID),\
	$(if $(and $(RPC_$(CHAIN_ID)),$(PRIVATE_KEY_$(CHAIN_ID))),\
		--rpc-url $(RPC_$(CHAIN_ID)) --private-key $(PRIVATE_KEY_$(CHAIN_ID)) --broadcast,\
		$(error Configuration for CHAIN_ID=$(CHAIN_ID) not found. Please check RPC_$(CHAIN_ID) and PRIVATE_KEY_$(CHAIN_ID) are set)),)

# Clean the repo
clean:
	@forge clean && rm -rf out cache

# coverage :; forge coverage --report debug > coverage-report.txt

pkg_build:
	@./pkg.sh

call:
	@time forge script script/CallContract.sol:CallContract $(NETWORK_PARAMS) --slow -vvvv

deploy:
	@time forge script script/Deploy.s.sol:DeployScript $(NETWORK_PARAMS) --slow -vvvv

deployVerify:
	@time forge script script/Deploy.s.sol:DeployScript $(NETWORK_PARAMS) --slow -vvvv --verifier blockscout --verifier-url https://unichain-sepolia.blockscout.com/api/ --verify

deployERC20:
	@time forge script script/DeployERC20.sol:DeployERC20 $(NETWORK_PARAMS) --slow -vvvv

deployPool:
	@time forge script script/DeployPool.sol:DeployPool $(NETWORK_PARAMS) --slow -vvvv

swap:
	@time forge script script/Swap.s.sol:RouterSwap $(NETWORK_PARAMS) --slow -vvvv

mint:
	@time forge script script/Mint.s.sol:Mint $(NETWORK_PARAMS) --slow -vvvv

createGauge:
	@time forge script script/CreateGauge.sol:CreateGauge $(NETWORK_PARAMS) --slow -vvvv

createGaugeReward:
	@time forge script script/CreateGaugeReward.sol:CreateGaugeReward $(NETWORK_PARAMS) --slow -vvvv


