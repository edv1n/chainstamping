pragma solidity ^0.8.0;

struct Commit {
    string hash;
    string tree;
    string[] parents;
}

using CommitLibrary for Commit global;

library CommitLibrary {
    function key(Commit calldata commit) public pure returns (bytes32) {
        require(bytes(commit.hash).length != 0, "Invalid commit hash");

        return keccak256(abi.encode(commit));
    }
}
