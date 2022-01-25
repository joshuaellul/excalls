// SPDX-License-Identifier: GPL-3.0

pragma solidity >= 0.5.10;

contract TestEXCALL {
    
    string public response;
    bytes public excall;
    
    address public oracleAddress;
    
    mapping(address => mapping(uint => bool)) public pending;
    mapping(address => int) public winnings;
    
    event BetPlaced(address indexed _from, uint oracleRefNr);
    
    uint public pendingOracleNr;
    
    function setOracle(address add) public {
        oracleAddress = add;
    }
    
    function getECT() public returns(string memory) {
        response = "http://joshuaellul.github.io/ect";
    }
    
    function getAAA() public returns(string memory) {
        response = "http://joshuaellul.github.io/aaa";
    }
    
    function get0s() public returns(string memory) {
        response = "http://localhost:8080/excall0str";
    }
    
    function get1s() public returns(string memory) {
        response = "http://localhost:8080/excall1str";
    }
    
    function alwaysWinGithub() public {
        //excall = "http://localhost:8080/excall1str";
        excall = "http://joshuaellul.github.io/win";
        if (excall[0] == '1') {
            winnings[msg.sender] += 1;
        } else {
            winnings[msg.sender] -= 1;
        }
    }
    
    function alwaysWinLocal() public {
        excall = "http://localhost:8080/excall1str";
        if (excall[0] == '1') {
            winnings[msg.sender] += 1;
        } else {
            winnings[msg.sender] -= 1;
        }
    }
    
    function beginAlwaysWinOracle() public {
        pending[msg.sender][pendingOracleNr++] = true;
        emit BetPlaced(msg.sender, pendingOracleNr);
    }
    
    function continueAlwaysWinOracle(address forPunter, uint oracleRef, bool answer) public returns(bool) {
        require(msg.sender == oracleAddress, "You are not the oracle!");
        bool isPending = pending[forPunter][oracleRef];
        if (isPending) {
            if (answer == true) {
                winnings[forPunter] += 1;
            } else {
                winnings[forPunter] -= 1;
            }
            pending[forPunter][oracleRef] = false;
            return true;
        }
        return false;
    }
    
    function sometimesWin() public {
        excall = "http://localhost:8080/excallrand";
        if (excall[0] == '1') {
            winnings[msg.sender] += 1;
        } else {
            winnings[msg.sender] -= 1;
        }
    }
    
}
