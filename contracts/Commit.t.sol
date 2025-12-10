pragma solidity ^0.8.0;

import "hardhat/console.sol";

import {Commit} from "./Commit.sol";
import {Test} from "forge-std/Test.sol";

contract CommitTest is Test {
    function test_key(Commit memory commit) public pure {
        console.log("Testing commit:", commit.hash);

        if (bytes(commit.hash).length == 0) {
            // Skip commit with empty hash
            console.log("Skipping commit with empty hash");

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
            assertEq(reason, "Invalid commit hash", "Unexpected error message");
        }
    }
}
