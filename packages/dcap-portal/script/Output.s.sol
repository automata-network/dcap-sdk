// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.13;

import {Script} from "forge-std/Script.sol";

contract Output is Script {
    string outputName;
    string env;

    function __Output_init(string memory _name, string memory _env) public {
        outputName = _name;
        env = _env;
    }

    function __Output_init(string memory _name) public {
        __Output_init(_name, vm.envString("ENV"));
    }

    function getOutputFilePath() private view returns (string memory) {
        return string.concat(vm.projectRoot(), "/out/deploy/", outputName, "_", env, ".json");
    }

    function ensureDirectoryExists() private {
        string memory dirPath = string.concat(vm.projectRoot(), "/out/deploy");
        if (!vm.exists(dirPath)) {
            vm.createDir(dirPath, true); // true表示递归创建
        }
    }

    function saveJson(string memory json) public {
        string memory finalJson = vm.serializeString(
            json,
            "remark",
            outputName
        );
        vm.writeJson(finalJson, getOutputFilePath());
    }

    function readJson() public returns (string memory) {
        ensureDirectoryExists();
        if (!vm.exists(getOutputFilePath())) {
            return "{}"; // 如果文件不存在，返回空的JSON对象
        }
        bytes32 remark = keccak256(abi.encodePacked("remark"));
        string memory output = vm.readFile(getOutputFilePath());
        string[] memory keys = vm.parseJsonKeys(output, ".");
        for (uint256 i = 0; i < keys.length; i++) {
            if (keccak256(abi.encodePacked(keys[i])) == remark) {
                continue;
            }
            string memory keyPath = string(abi.encodePacked(".", keys[i]));
            vm.serializeAddress(output, keys[i], vm.parseJsonAddress(output, keyPath));
        }
        return output;
    }
}
