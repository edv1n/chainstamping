pragma solidity ^0.8.0;

import {ChainstampingCommits} from "./ChainstampingCommits.sol";
import {CommitKeyGenerator} from "./CommitKeyGenerator.sol";
import {Test} from "forge-std/Test.sol";

contract ChainstampingCommitsTest is Test {
    ChainstampingCommits c = new ChainstampingCommits();
    CommitKeyGenerator _CommitKeyGenerator = new CommitKeyGenerator();

    function setUp() public {
        c = new ChainstampingCommits();
    }

    function test_timestamp(
        string calldata commit,
        string[] calldata parents,
        string calldata tree
    ) public {
        if (bytes(commit).length == 0) {
            // Skip commit with empty hash
            return;
        }

        bytes32 key = _CommitKeyGenerator.generate(commit, tree, parents);

        uint256 _now = block.timestamp;

        vm.expectEmit();
        emit ChainstampingCommits.CommitTimestamped(
            key,
            commit,
            tree,
            parents,
            _now
        );

        uint256 timestamped = c.timestamp(commit, parents, tree);

        assertLt(0, timestamped, "Timestamp should be greater than 0");
        assertEq(_now, timestamped, "Timestamp should match block timestamp");
    }

    function test_timestamped(
        string calldata commit,
        string[] calldata parents,
        string calldata tree
    ) public {
        if (bytes(commit).length == 0) {
            // Skip commit with empty hash
            return;
        }

        bytes32 key = _CommitKeyGenerator.generate(commit, tree, parents);

        vm.expectEmit();
        emit ChainstampingCommits.CommitTimestamped(
            key,
            commit,
            tree,
            parents,
            block.timestamp
        );

        uint256 timestamped = c.timestamp(commit, parents, tree);
        uint256 retrievedTimestamp = c.timestamped(commit, parents, tree);

        assertEq(
            retrievedTimestamp,
            timestamped,
            "Retrieved timestamp should match the original timestamp"
        );
    }

    function test_TimestampWithEmptyCommitHashShouldFail(
        string[] calldata parents,
        string calldata tree
    ) public {
        string memory commit = "";

        // Attempt to timestamp with empty commit hash
        try c.timestamp(commit, parents, tree) {
            revert("Timestamping with empty commit hash should have failed");
        } catch Error(string memory reason) {
            assertEq(reason, "Invalid commit hash", "Unexpected error message");
        }
    }

    function test_DoubleTimestampingShouldFail(
        string calldata commit,
        string[] calldata parents,
        string calldata tree
    ) public {
        if (bytes(commit).length == 0) {
            // Skip commit with empty hash
            return;
        }

        c.timestamp(commit, parents, tree);

        // Attempt to timestamp the same commit again
        try c.timestamp(commit, parents, tree) {
            revert("Double timestamping should have failed");
        } catch Error(string memory reason) {
            assertEq(
                reason,
                "Commit already timestamped",
                "Unexpected error message"
            );
        }
    }

    function test_TimestampedNonexistentCommitShouldFail(
        string calldata commit,
        string[] calldata parents,
        string calldata tree
    ) public view {
        if (bytes(commit).length == 0) {
            // Skip commit with empty hash
            return;
        }

        // Attempt to retrieve timestamp for a commit that hasn't been timestamped
        try c.timestamped(commit, parents, tree) {
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
