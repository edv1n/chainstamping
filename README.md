# Chainstamping

Timestamping digital fingerprint on blockchain. Git commits are timestamped in a blockchain to provide tamper-proof proof of their existence at a specific point in time.

## Overview

Protecting your ideas and intellectual property is crucial in today's fast-paced world. Timestamping your work provides undeniable proof of when you created it, which can be invaluable in legal disputes or when establishing priority.

Timestamping physically (e.g., mailing a copy to yourself) can be cumbersome and unreliable. Ensuring the integrity of the timestamped document is also difficult, as physical documents can be altered or damaged over time.

Patenting an idea is often time-consuming and expensive. It may not be feasible for every piece of work, especially when the idea is better kept confidential or when rapid iteration is required.

Uploading your work to the Internet often requires trusting a third party, which may not be ideal for sensitive information. It is also challenging to prove the exact time of creation when relying on third-party services.

Chainstamping is a smart contract-based solution that allows developers to timestamp their git commits on a blockchain. By recording commit metadata on-chain, Chainstamping provides a tamper-proof record of when a commit was created, protecting your ideas and intellectual property without revealing them publicly. Git commits are uniquely identified by their metadata hash, which includes the commit hash and the tree hash, ensuring integrity and authenticity.

## Features

- Timestamp git commits in a blockchain (currently Ethereum)
- Verify commit timestamps
- Temper-proof commit "seen" records on-chain
- Prevent commit hash forgery by staming commit hash with tree hash

## Usage

### To timestamp a git commit

1. Use the `timestamp` function to timestamp a commit by providing its metadata (commit hash **required**, tree hash **required**, and parent hashes **optional**).

### To verify a git commit timestamp

1. Use the `timestamped` function to retrieve the timestamp of a commit by providing its metadata.
1. Verify the commit hash and the related metadata using git tools.

## Design

Chainstamping uses smart contracts to store commit metadata along with their timestamps. Each commit is uniquely identified by its metadata hash, ensuring integrity and authenticity.

### Commit metadata structure

Each commit metadata consists of the following fields:

- **hash**: The hash of the git commit (SHA-1).
- **tree**: The hash of the git tree associated with the commit (SHA-1).
- **parents**: An array of parent commit hashes.

### Commit metadata hash

The commit metadata hash is computed by hashing the concatenation of the commit hash and the tree hash using the Keccak-256 hashing algorithm.

This design choice mitigates the risk of commit hash forgery, as the tree hash is inherently linked to the commit's content and structure.

### Timestamping process

When a commit is timestamped, the following steps occur:

1. The commit metadata hash is computed.
2. The current block timestamp is recorded alongside the commit metadata hash in the smart contract.

### Verification process

To verify a commit's timestamp, the following steps are performed:

1. The commit metadata hash is computed.
2. The mapping in the smart contract is queried using the commit metadata hash to retrieve the stored timestamp.

### Security considerations

- The use of both commit hash and tree hash in the metadata hash computation prevents attackers from timestamping abitrary commit hashes.
- The smart contract enforces validation checks to ensure that neither the commit hash nor the tree hash are empty, further enhancing security.
- No direct write access to the timestamp storage is provided, ensuring that only commits with sufficient metadata can be timestamped through the defined interface.

## Contracts

Smart contracts are implemented in Solidity and can be found in the `contracts` directory. 

The main contracts is `ChainstampingCommits.sol`, which handles the timestamping and verification of git commits.

## Testing

The project includes comprehensive tests to ensure the correctness and security of the timestamping and verification processes. The tests cover various scenarios, including valid commits, commits with missing metadata, and edge cases.

The tests are implemented using the hardhat framework and can be found in the `contracts` directory, specifically in the `ChainstampingCommits.t.sol` and `Commit.t.sol` files.

### Running Tests

To run the tests, one need to have `Hardhat` installed. The installation instructions can be found in the [Hardhat documentation](https://hardhat.org/getting-started/).

Once Hardhat is installed, the tests can be run using the following command on the root of the repo.

```bash
npx hardhat test
```

## Related Works

- [OriginStamp](https://originstamp.com/en/solutions/timestamp): A service that provides blockchain-based timestamping for documents and files.
- [OpenTimestamps](https://opentimestamps.org/): A protocol for creating and verifying timestamps using Bitcoin.