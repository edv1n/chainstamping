pragma solidity ^0.8.0;

import {Chainstamper} from "./Chainstamper.sol";
import {Test} from "forge-std/Test.sol";

import {console} from "hardhat/console.sol";

contract ChainstamperTest is Test, Chainstamper {
    Chainstamper chainstamper;

    function setUp() public {
        chainstamper = new Chainstamper();
    }

    function test_stampCommit() public {
        Chainstamper.Commit memory commit;
        commit.hash = "a1b2c3d4e5f60718293a4b5c6d7e8f9012345678";
        commit.tree = "a1b2c3d4e5f60718293a4b5c6d7e8f9012345678";

        bytes32 key = _commitKey(commit);

        uint256 _now = block.timestamp;

        if (_validCommit(commit)) {
            vm.expectEmit();
            emit Chainstamper.CommitTimestamped(key, commit.hash, commit, _now);
        }

        uint256 timestamped = chainstamper.stampCommit(commit);

        assertLt(0, timestamped, "Timestamp should be greater than 0");
        assertEq(_now, timestamped, "Timestamp should match block timestamp");

        console.log("Timestamped commit", commit.hash, "at time", timestamped);
    }

    function test_stampCommit_WithParents() public {
        Chainstamper.Commit memory commit;
        commit.hash = "b1c2d3e4f5061728394a5b6c7d8e9f0012345678";
        commit.tree = "b1c2d3e4f5061728394a5b6c7d8e9f0012345678";
        commit.parents = new string[](2);
        commit.parents[0] = "a1b2c3d4e5f60718293a4b5c6d7e8f9012345678";
        commit.parents[1] = "c1d2e3f40516273849a5b6c7d8e9f0012345678";

        bytes32 key = _commitKey(commit);

        uint256 _now = block.timestamp;

        vm.expectEmit();
        emit Chainstamper.CommitTimestamped(key, commit.hash, commit, _now);

        uint256 timestamped = chainstamper.stampCommit(commit);

        assertLt(0, timestamped, "Timestamp should be greater than 0");
        assertEq(_now, timestamped, "Timestamp should match block timestamp");

        console.log(
            "Timestamped commit with parents",
            commit.hash,
            "at time",
            timestamped
        );
    }

    function test_stampCommit_Fuzz(Chainstamper.Commit memory commit) public {
        console.log("Testing commit:", commit.hash);

        if (!_validCommit(commit)) {
            // Skip invalid commit
            console.log("Skipping invalid commit");

            vm.expectRevert("Invalid commit");
            chainstamper.stampCommit(commit);
            return;
        }

        bytes32 key = _commitKey(commit);

        uint256 _now = block.timestamp;

        vm.expectEmit();
        emit Chainstamper.CommitTimestamped(key, commit.hash, commit, _now);

        uint256 timestamped = chainstamper.stampCommit(commit);

        assertLt(0, timestamped, "Timestamp should be greater than 0");
        assertEq(_now, timestamped, "Timestamp should match block timestamp");

        console.log("Timestamped commit", commit.hash, "at time", timestamped);
    }

    function test_getTimestamp() public {
        Chainstamper.Commit memory commit;
        commit.hash = "a1b2c3d4e5f60718293a4b5c6d7e8f9012345678";
        commit.tree = "a1b2c3d4e5f60718293a4b5c6d7e8f9012345678";

        bytes32 key = _commitKey(commit);

        vm.expectEmit();
        emit Chainstamper.CommitTimestamped(
            key,
            commit.hash,
            commit,
            block.timestamp
        );

        uint256 timestamped = chainstamper.stampCommit(commit);
        console.log("Timestamped commit", commit.hash, "at time", timestamped);

        uint256 retrievedTimestamp = chainstamper.getTimestamp(commit);

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

    function test_getTimestamp_Fuzz(Chainstamper.Commit memory commit) public {
        console.log("Testing commit:", commit.hash);

        if (!_validCommit(commit)) {
            // Skip invalid commit
            console.log("Skipping invalid commit");

            vm.expectRevert("Invalid commit");
            chainstamper.getTimestamp(commit);
            return;
        }

        uint256 timestamped = chainstamper.stampCommit(commit);
        console.log("Timestamped commit", commit.hash, "at time", timestamped);

        uint256 retrievedTimestamp = chainstamper.getTimestamp(commit);

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

    function test_stampCommit_EmptyCommitHashShouldFail() public {
        Chainstamper.Commit memory commit;
        commit.hash = "";

        console.log("Testing commit:", commit.hash);

        vm.expectRevert("Invalid commit");
        chainstamper.stampCommit(commit);
    }

    function test_stampCommit_EmptyTreeHashShouldFail() public {
        Chainstamper.Commit memory commit;
        commit.tree = "";

        console.log("Testing commit:", commit.hash);

        vm.expectRevert("Invalid commit");
        chainstamper.stampCommit(commit);
    }

    function test_stampCommit_DoubleTimestampingShouldFail() public {
        Chainstamper.Commit memory commit;
        commit.hash = "a1b2c3d4e5f60718293a4b5c6d7e8f9012345678";
        commit.tree = "a1b2c3d4e5f60718293a4b5c6d7e8f9012345678";

        chainstamper.stampCommit(commit);

        vm.expectRevert("Commit already timestamped");

        chainstamper.stampCommit(commit);
    }

    function test_getTimestamp_NonexistentCommitShouldFail() public {
        Chainstamper.Commit memory commit;
        commit.hash = "a1b2c3d4e5f60718293a4b5c6d7e8f9012345678";
        commit.tree = "a1b2c3d4e5f60718293a4b5c6d7e8f9012345678";

        vm.expectRevert("Commit not timestamped");

        chainstamper.getTimestamp(commit);
    }

    function test_getTimestamp_InvalidCommitShouldFail(
        Chainstamper.Commit memory commit
    ) public {
        if (_validCommit(commit)) {
            return;
        }

        vm.expectRevert("Invalid commit");
        chainstamper.getTimestamp(commit);
    }

    /*
    function test_commitKey(Chainstamper.Commit memory commit) public {
        console.log("Testing commit:", commit.hash);

        if (!_validCommit(commit)) {
            // Skip invalid commit
            console.log("Skipping invalid commit");

            vm.expectRevert("Invalid commit");
        }

        bytes32 key = _commitKey(commit);

        if (!_validCommit(commit)) {
            return;
        }
        assertEq(
            keccak256(abi.encode(commit)),
            key,
            "Generated key does not match expected value"
        );
    }

    function test_commitKey_WithEmptyCommitHashShouldFail(
        Chainstamper.Commit memory commit
    ) public {
        commit.hash = "";

        // Attempt to generate key with empty commit hash
        vm.expectRevert("Invalid commit");
        _commitKey(commit);
    }

    function test_commitKey_WithEmptyTreeHashShouldFail(
        Chainstamper.Commit memory commit
    ) public {
        commit.tree = "";

        // Attempt to generate key with empty tree hash
        vm.expectRevert("Invalid commit");
        _commitKey(commit);
    }
    */
}
