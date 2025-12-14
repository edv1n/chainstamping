pragma solidity ^0.8.0;

import {Chainstamper} from "./Chainstamper.sol";
import {Commit} from "./Commit.sol";
import {Test} from "forge-std/Test.sol";

import {console} from "hardhat/console.sol";

contract ChainstamperTest is Test {
    Chainstamper c;

    function setUp() public {
        c = new Chainstamper();
    }

    function test_stampCommit(Commit memory commit) public {
        console.log("Testing commit:", commit.hash);

        if (!commit.valid()) {
            // Skip invalid commit
            console.log("Skipping invalid commit");

            return;
        }

        bytes32 key = commit.key();

        uint256 _now = block.timestamp;

        vm.expectEmit();
        emit Chainstamper.CommitTimestamped(key, commit.hash, commit, _now);

        uint256 timestamped = c.stampCommit(commit);

        assertLt(0, timestamped, "Timestamp should be greater than 0");
        assertEq(_now, timestamped, "Timestamp should match block timestamp");

        console.log("Timestamped commit", commit.hash, "at time", timestamped);
    }

    function test_getTimestamp(Commit memory commit) public {
        console.log("Testing commit:", commit.hash);

        if (!commit.valid()) {
            // Skip invalid commit
            console.log("Skipping invalid commit");

            return;
        }

        bytes32 key = commit.key();

        vm.expectEmit();
        emit Chainstamper.CommitTimestamped(
            key,
            commit.hash,
            commit,
            block.timestamp
        );

        uint256 timestamped = c.stampCommit(commit);
        console.log("Timestamped commit", commit.hash, "at time", timestamped);

        uint256 retrievedTimestamp = c.getTimestamp(commit);

        assertEq(
            retrievedTimestamp,
            timestamped,
            "Retrieved timestamp should match the original timestamp"
        );

        console.log(
            "Retrieved timestamp for commit",
            commit.hash,
            ":",
            retrievedTimestamp
        );
    }

    function test_stampCommit_EmptyCommitHashShouldFail(
        Commit memory commit
    ) public {
        commit.hash = "";

        console.log("Testing commit:", commit.hash);

        // Attempt to timestamp with empty commit hash
        try c.stampCommit(commit) {
            revert("Timestamping with empty commit hash should have failed");
        } catch Error(string memory reason) {
            assertEq(reason, "Invalid commit", "Unexpected error message");
        }
    }

    function test_stampCommit_EmptyTreeHashShouldFail(
        Commit memory commit
    ) public {
        commit.tree = "";

        console.log("Testing commit:", commit.hash);

        // Attempt to timestamp with empty tree hash
        try c.stampCommit(commit) {
            revert("Timestamping with empty tree hash should have failed");
        } catch Error(string memory reason) {
            assertEq(reason, "Invalid commit", "Unexpected error message");
        }
    }

    function test_stampCommit_DoubleTimestampingShouldFail(
        Commit memory commit
    ) public {
        console.log("Testing commit:", commit.hash);

        if (!commit.valid()) {
            // Skip invalid commit
            console.log("Skipping invalid commit");

            return;
        }

        c.stampCommit(commit);

        // Attempt to timestamp the same commit again
        try c.stampCommit(commit) {
            revert("Double timestamping should have failed");
        } catch Error(string memory reason) {
            assertEq(
                reason,
                "Commit already timestamped",
                "Unexpected error message"
            );
        }
    }

    function test_getTimestamp_NonexistentCommitShouldFail(
        Commit memory commit
    ) public view {
        console.log("Testing commit:", commit.hash);

        if (!commit.valid()) {
            // Skip invalid commit
            console.log("Skipping invalid commit");

            return;
        }

        // Attempt to retrieve timestamp for a commit that hasn't been timestamped
        try c.getTimestamp(commit) {
            revert(
                "Retrieving timestamp for nonexistent commit should have failed"
            );
        } catch Error(string memory reason) {
            assertEq(
                reason,
                "Commit not timestamped",
                "Unexpected error message"
            );
        }
    }
}
