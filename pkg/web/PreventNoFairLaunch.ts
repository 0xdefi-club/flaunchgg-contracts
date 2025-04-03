export const PreventNoFairLaunchABI = [
  {
    "type": "constructor",
    "inputs": [
      {
        "name": "_notifier",
        "type": "address",
        "internalType": "address"
      }
    ],
    "stateMutability": "nonpayable"
  },
  {
    "type": "function",
    "name": "MINIMUM_INITIAL_TOKENS",
    "inputs": [],
    "outputs": [
      {
        "name": "",
        "type": "uint256",
        "internalType": "uint256"
      }
    ],
    "stateMutability": "view"
  },
  {
    "type": "function",
    "name": "notifier",
    "inputs": [],
    "outputs": [
      {
        "name": "",
        "type": "address",
        "internalType": "address"
      }
    ],
    "stateMutability": "view"
  },
  {
    "type": "function",
    "name": "notify",
    "inputs": [
      {
        "name": "",
        "type": "bytes32",
        "internalType": "PoolId"
      },
      {
        "name": "_key",
        "type": "bytes4",
        "internalType": "bytes4"
      },
      {
        "name": "_data",
        "type": "bytes",
        "internalType": "bytes"
      }
    ],
    "outputs": [],
    "stateMutability": "view"
  },
  {
    "type": "function",
    "name": "subscribe",
    "inputs": [
      {
        "name": "",
        "type": "bytes",
        "internalType": "bytes"
      }
    ],
    "outputs": [
      {
        "name": "",
        "type": "bool",
        "internalType": "bool"
      }
    ],
    "stateMutability": "view"
  },
  {
    "type": "function",
    "name": "unsubscribe",
    "inputs": [],
    "outputs": [],
    "stateMutability": "nonpayable"
  },
  {
    "type": "error",
    "name": "InvalidInitialTokenFairLaunch",
    "inputs": [
      {
        "name": "_invalidAmount",
        "type": "uint256",
        "internalType": "uint256"
      },
      {
        "name": "_minTokens",
        "type": "uint256",
        "internalType": "uint256"
      }
    ]
  },
  {
    "type": "error",
    "name": "InvalidNotifier",
    "inputs": [
      {
        "name": "_sender",
        "type": "address",
        "internalType": "address"
      },
      {
        "name": "_validNotifier",
        "type": "address",
        "internalType": "address"
      }
    ]
  }
] as const;
