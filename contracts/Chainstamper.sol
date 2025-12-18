// SPDX-License-Identifier: GPL-2.0-or-later

pragma solidity ^0.8.0;

contract Chainstamper {
    struct Commit {
        string hash;
        string tree;
        string[] parents;
    }

    event CommitTimestamped(
        bytes32 indexed key,
        string indexed hash,
        Commit commit,
        uint256 timestamp
    );

    // mapping from commit with metadata to its timestamp
    mapping(bytes32 => uint256) internal _timestamps;

    // stampCommit timestamps a commit with its metadata
    function stampCommit(Commit memory commit) public returns (uint256) {
        if (!_validCommit(commit)) {
            revert("Invalid commit");
        }

        bytes32 key = _commitKey(commit);

        if (_timestamps[key] != 0) {
            revert("Commit already timestamped");
        }

        uint256 _now = block.timestamp;

        _timestamps[key] = _now;

        emit CommitTimestamped(key, commit.hash, commit, _now);

        return _now;
    }

    // getTimestamp returns the timestamp of a previously timestamped commit
    function getTimestamp(Commit memory commit) public view returns (uint256) {
        if (!_validCommit(commit)) {
            revert("Invalid commit");
        }

        bytes32 key = _commitKey(commit);

        uint256 _timestamp = _timestamps[key];

        if (_timestamp == 0) {
            revert("Commit not timestamped");
        }

        return _timestamp;
    }

    function _commitKey(Commit memory commit) internal pure returns (bytes32) {
        return keccak256(abi.encode(commit));
    }

    function _validCommit(Commit memory commit) internal pure returns (bool) {
        return
            bytes(commit.hash).length == 40 && bytes(commit.tree).length == 40;
    }
}
