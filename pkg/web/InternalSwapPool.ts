export const InternalSwapPoolABI = [
  {
    "type": "function",
    "name": "poolFees",
    "inputs": [
      {
        "name": "_poolKey",
        "type": "tuple",
        "internalType": "struct PoolKey",
        "components": [
          {
            "name": "currency0",
            "type": "address",
            "internalType": "Currency"
          },
          {
            "name": "currency1",
            "type": "address",
            "internalType": "Currency"
          },
          {
            "name": "fee",
            "type": "uint24",
            "internalType": "uint24"
          },
          {
            "name": "tickSpacing",
            "type": "int24",
            "internalType": "int24"
          },
          {
            "name": "hooks",
            "type": "address",
            "internalType": "contract IHooks"
          }
        ]
      }
    ],
    "outputs": [
      {
        "name": "",
        "type": "tuple",
        "internalType": "struct InternalSwapPool.ClaimableFees",
        "components": [
          {
            "name": "amount0",
            "type": "uint256",
            "internalType": "uint256"
          },
          {
            "name": "amount1",
            "type": "uint256",
            "internalType": "uint256"
          }
        ]
      }
    ],
    "stateMutability": "view"
  },
  {
    "type": "event",
    "name": "PoolFeesDistributed",
    "inputs": [
      {
        "name": "_poolId",
        "type": "bytes32",
        "indexed": true,
        "internalType": "PoolId"
      },
      {
        "name": "_donateAmount",
        "type": "uint256",
        "indexed": false,
        "internalType": "uint256"
      },
      {
        "name": "_creatorAmount",
        "type": "uint256",
        "indexed": false,
        "internalType": "uint256"
      },
      {
        "name": "_bidWallAmount",
        "type": "uint256",
        "indexed": false,
        "internalType": "uint256"
      },
      {
        "name": "_governanceAmount",
        "type": "uint256",
        "indexed": false,
        "internalType": "uint256"
      },
      {
        "name": "_protocolAmount",
        "type": "uint256",
        "indexed": false,
        "internalType": "uint256"
      }
    ],
    "anonymous": false
  },
  {
    "type": "event",
    "name": "PoolFeesReceived",
    "inputs": [
      {
        "name": "_poolId",
        "type": "bytes32",
        "indexed": true,
        "internalType": "PoolId"
      },
      {
        "name": "_amount0",
        "type": "uint256",
        "indexed": false,
        "internalType": "uint256"
      },
      {
        "name": "_amount1",
        "type": "uint256",
        "indexed": false,
        "internalType": "uint256"
      }
    ],
    "anonymous": false
  },
  {
    "type": "event",
    "name": "PoolFeesSwapped",
    "inputs": [
      {
        "name": "_poolId",
        "type": "bytes32",
        "indexed": true,
        "internalType": "PoolId"
      },
      {
        "name": "zeroForOne",
        "type": "bool",
        "indexed": false,
        "internalType": "bool"
      },
      {
        "name": "_amount0",
        "type": "uint256",
        "indexed": false,
        "internalType": "uint256"
      },
      {
        "name": "_amount1",
        "type": "uint256",
        "indexed": false,
        "internalType": "uint256"
      }
    ],
    "anonymous": false
  }
] as const;
