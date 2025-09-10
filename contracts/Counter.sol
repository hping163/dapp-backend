// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
// 计数器合约
contract Counter {
    uint256 private count;

    // 增加计数
    function increment() public {
        count++;
    }

    // 获取当前计数
    function getCount() public view returns (uint256) {
        return count;
    }
}
