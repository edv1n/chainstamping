pragma solidity ^0.8.0;

import {CommitKeyGenerator} from "./CommitKeyGenerator.sol";

contract ChainstampingCommits {
    CommitKeyGenerator _CommitKeyGenerator = new CommitKeyGenerator();

    event CommitTimestamped(
        bytes32 indexed key,
        string indexed commit,
        string tree,
        string[] parents,
        uint256 timestamp
    );

    // mapping from commit with metadata to its timestamp
    mapping(bytes32 => uint256) internal _stamped;

    // timestamp timestamps a commit with its metadata
    function timestamp(
        string calldata commit,
        string[] calldata parents,
        string calldata tree
    ) public returns (uint256) {
        bytes32 key = _CommitKeyGenerator.generate(commit, tree, parents);

        require(_stamped[key] == 0, "Commit already timestamped");

        uint256 _now = block.timestamp;

        _stamped[key] = _now;

        emit CommitTimestamped(key, commit, tree, parents, _now);

        return _now;
    }

    // timestamped returns the timestamp of a previously timestamped commit
    function timestamped(
        string calldata commit,
        string[] calldata parents,
        string calldata tree
    ) public view returns (uint256) {
        bytes32 key = _CommitKeyGenerator.generate(commit, tree, parents);

        uint256 _timestamp = _stamped[key];

        require(_timestamp != 0, "Commit not timestamped");

        return _timestamp;
    }
}
