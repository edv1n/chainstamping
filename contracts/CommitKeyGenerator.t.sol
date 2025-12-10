pragma solidity ^0.8.0;

import {CommitKeyGenerator} from "./CommitKeyGenerator.sol";
import {Test} from "forge-std/Test.sol";

contract CommitKeyGeneratorTest is Test {
    CommitKeyGenerator _CommitKeyGenerator = new CommitKeyGenerator();

    function test_generate(
        string calldata commit,
        string[] calldata parents,
        string calldata tree
    ) public view {
        if (bytes(commit).length == 0) {
            // Skip commit with empty hash
            return;
        }

        assertEq(
            keccak256(abi.encode(commit, tree, parents)),
            _CommitKeyGenerator.generate(commit, tree, parents),
            "Generated key does not match expected value"
        );
    }

    function test_generate_empty_commit(
        string[] calldata parents,
        string calldata tree
    ) public view {
        string memory commit = "";

        // Attempt to generate key with empty commit hash
        try _CommitKeyGenerator.generate(commit, tree, parents) {
            revert("Generating key with empty commit hash should have failed");
        } catch Error(string memory reason) {
            assertEq(reason, "Invalid commit hash", "Unexpected error message");
        }
    }
}
