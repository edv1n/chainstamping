// SPDX-License-Identifier: GPL-2.0-or-later

pragma solidity ^0.8.0;

import {Commit} from "./Commit.sol";

contract Chainstamper {
    event CommitTimestamped(
        bytes32 indexed key,
        string indexed hash,
        Commit commit,
        uint256 timestamp
    );

    // mapping from commit with metadata to its timestamp
    mapping(bytes32 => uint256) internal _timestamps;

    // stampCommit timestamps a commit with its metadata
    function stampCommit(Commit calldata commit) public returns (uint256) {
        bytes32 key = commit.key();

        require(_timestamps[key] == 0, "Commit already timestamped");

        uint256 _now = block.timestamp;

        _timestamps[key] = _now;

        emit CommitTimestamped(key, commit.hash, commit, _now);

        return _now;
    }

    // getTimestamp returns the timestamp of a previously timestamped commit
    function getTimestamp(
        Commit calldata commit
    ) public view returns (uint256) {
        bytes32 key = commit.key();

        uint256 _timestamp = _timestamps[key];

        require(_timestamp != 0, "Commit not timestamped");

        return _timestamp;
    }
}
