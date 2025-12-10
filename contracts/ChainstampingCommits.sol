pragma solidity ^0.8.0;

import {Commit} from "./Commit.sol";

contract ChainstampingCommits {
    event CommitTimestamped(
        bytes32 indexed key,
        Commit indexed commit,
        uint256 timestamp
    );

    // mapping from commit with metadata to its timestamp
    mapping(bytes32 => uint256) internal _stamped;

    // timestamp timestamps a commit with its metadata
    function timestamp(Commit calldata commit) public returns (uint256) {
        bytes32 key = commit.key();

        require(_stamped[key] == 0, "Commit already timestamped");

        uint256 _now = block.timestamp;

        _stamped[key] = _now;

        emit CommitTimestamped(key, commit, _now);

        return _now;
    }

    // timestamped returns the timestamp of a previously timestamped commit
    function timestamped(Commit calldata commit) public view returns (uint256) {
        bytes32 key = commit.key();

        uint256 _timestamp = _stamped[key];

        require(_timestamp != 0, "Commit not timestamped");

        return _timestamp;
    }
}
