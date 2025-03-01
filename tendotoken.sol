// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract TendoToken is ERC20 {
    constructor(uint256 initialSupply) ERC20("TendoToken", "TDT") {
        _mint(msg.sender, initialSupply * (10 ** decimals()));
    }
}
