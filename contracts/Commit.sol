pragma solidity ^0.8.0;

struct Commit {
    string hash;
    string tree;
    string[] parents;
}

using CommitLibrary for Commit global;

library CommitLibrary {
    function key(Commit calldata commit) public pure returns (bytes32) {
        require(valid(commit), "Invalid commit");

        return keccak256(abi.encode(commit));
    }

    function valid(Commit calldata commit) public pure returns (bool) {
        return bytes(commit.hash).length != 0 && bytes(commit.tree).length != 0;
    }
}
