pragma solidity ^0.8.0;

contract CommitKeyGenerator {
    function generate(
        string memory commit,
        string memory tree,
        string[] memory parents
    ) public pure returns (bytes32) {
        require(bytes(commit).length != 0, "Invalid commit hash");

        return keccak256(abi.encode(commit, tree, parents));
    }
}
