# Go-Helix
Implementation of backend services for HelixPay.

## Features
- Bundle UserOps and submit to multiple blockchains.
- Search for all transactions related to a viewing key.
- Private pay gas for operation.

## Services:
1. Bundler: 
Bundle UserOperations and submit to the blockchains.
- Submit user ops
- Query user ops
- Centralized bundler (for now / Decentralized in future)
- Key-value store for UserOps
- Compatible with ERC-4337 Bundler-rpc

2. Searcher:
Take viewing key and search for all transaction related to the viewing key.

## Packages: 
1. Mempool: Mempool provide read-write access to the mempool. It is used by the Bundler to store UserOps before submitting to the blockchain. It uses a key-value store to store UserOps.