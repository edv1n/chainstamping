pragma solidity ^0.8.0;

import {Commit} from "./Commit.sol";
import {Test} from "forge-std/Test.sol";

import {console} from "hardhat/console.sol";

contract CommitTest is Test {
    function test_key(Commit memory commit) public pure {
        console.log("Testing commit:", commit.hash);

        if (!commit.valid()) {
            // Skip invalid commit
            console.log("Skipping invalid commit");

            return;
        }

        bytes32 key = commit.key();

        assertEq(
            keccak256(abi.encode(commit)),
            key,
            "Generated key does not match expected value"
        );
    }

    function test_key_WithEmptyCommitHashShouldFail(
        Commit memory commit
    ) public pure {
        commit.hash = "";

        // Attempt to generate key with empty commit hash
        try commit.key() {
            revert("Generating key with empty commit hash should have failed");
        } catch Error(string memory reason) {
            assertEq(reason, "Invalid commit", "Unexpected error message");
        }
    }

    function test_key_WithEmptyTreeHashShouldFail(
        Commit memory commit
    ) public pure {
        commit.tree = "";

        // Attempt to generate key with empty tree hash
        try commit.key() {
            revert("Generating key with empty tree hash should have failed");
        } catch Error(string memory reason) {
            assertEq(reason, "Invalid commit", "Unexpected error message");
        }
    }
}
