## External Calls for go-ethereum
by Joshua Ellul and Gordon J. Pace

go-ethereum (geth) was modified to demonstrate that blockchain systems can execute calls to external systems. For further details regarding our work on external calls for blockchain visit [this page](https://joshuaellul.github.io/excalls/)

See the official [go-ethereum github repository](https://github.com/ethereum/go-ethereum) in relation to geth.


## Building the source - original from geth

Since this is a modified version of geth, instructions for building geth will be similar to this modified version. The below instructions in this section are from the go-ethereum repository:

For prerequisites and detailed build instructions please read the [Installation Instructions](https://geth.ethereum.org/docs/install-and-build/installing-geth).

The following 
Building `geth` requires both a Go (version 1.14 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make geth
```

or, to build the full suite of utilities:

```shell
make all
```



## Steps to try out external call functionality

1. Build the geth binary as above (i.e. `make geth').

2. Set up a private network. You can use a template provided consisting of two nodes (1 miner, and 1 verifier):
a) Allocate funds to your test account: Open geth-network/genesis.json and add an entry in the `alloc` section to allocate some amount of ETH to an account you will use to test with.

b) Initialise the two nodes:
Initialise the miner: Execute `./build/bin/geth init --datadir ./geth-network/miner ./geth-network/genesis.json'

Initialise the verifier: Execute `./build/bin/geth init --datadir ./geth-network/verifier ./geth-network/genesis.json'

c) Run the miner (set to only accept connections from localhost)
Execute `./build/bin/geth --datadir ./geth-network/miner/ --nodiscover --verbosity 3 --networkid 15 --netrestrict 127.0.0.1/8 --rpc --rpcaddr 0.0.0.0 --rpcport 22004 --rpccorsdomain "*" --rpc.allow-unprotected-txs --allow-insecure-unlock --keystore ./geth-network/miner/keystore/ --unlock 0x8305ab65E82F383041b878021Cf08712994cE1AF --password ./geth-network/miner/keystore/pw1.txt --mine'

d) Run the verifier
In a separate terminal, execute `./build/bin/geth --datadir ./geth-network/verifier --networkid 15 --nodiscover --port 30305 --verbosity 3 --netrestrict 127.0.0.1/8 --rpc --rpcaddr 0.0.0.0 --rpcport 22009 --rpccorsdomain "*" --rpc.allow-unprotected-txs --allow-insecure-unlock --notxpool'
The last flag `--notxpool' is a flag we have added to be able to test out a node that does not process transactions from the transaction pool but only verifies blocks that have been mined. 

e) Find out the miner's enode:
In a new terminal, attach to the miner node: `./build/bin/geth attach ./geth-network/miner/geth.ipc'

and take a copy of the enode string output when executing the following command in the geth console:
```admin.nodeInfo.enode'''

Edit the enode string copied so that (i) it does not contain any parameteres (text after and including the ? should be removed); and (ii) change the IP address to 127.0.0.1.  For example:
"enode://1f58d049e2029a8014d5318726be2fb627b1e8a70e0c6ebaac4daccef3b30d26c055d61ab939f47fd370baa0fc66af55eb659a5af4dac1f8e7c81b2b05867d78@127.0.0.1:30303"

f) Tell the verifier to connect to the miner:
In a new terminal, attach to the verifier node: `./build/bin/geth attach ./geth-network/verifier/geth.ipc'

and execute the following command with the edited enode string from above in the new geth console:
```admin.addPeer(<EDITED ENODE STRING>)'''


3. Test the EXCALL functionality:
Configure whatever tools you use to inertact with Ethereum to connect to the local network. You could use MetaMask and add a network with the following settings:
RPC URL: http://127.0.0.1:22004
Chain ID: 15

Deploy the following smart contract:
```
// SPDX-License-Identifier: GPL-3.0

pragma solidity 0.5.10;

contract TestEXCALL {
    
    string public response;
    
    function getECT() public returns(string memory) {
        response = "http://joshuaellul.github.io/ect";
    }
    
    function getAAA() public returns(string memory) {
        response = "http://joshuaellul.github.io/aaa";
    }
    
    
}
'''

Then after calling getECT(), if you check the value of the response string, it should return back a response from the URL that has been verified.

In the miner's geth console outtput you should see a log entry along the following lines: ```EXCALL: made to http://joshuaellul.github.io/aaa'''
This means that the miner performed the external call.

In the verifier's geth console output you should see a log entry along the following lines: EXCALL: retrieved from trsansaction: http://joshuaellul.github.io/aaa 
This means that the verifier did not undertake an external calls. Yet it relied on the transaction information and following additional external call data items (proposed in the paper): response, public key and the response digital signature.

The external call is only made on the miner.  The verifier does verify that the response is valid, however it does this by verifying the signed response, and not by undertaking a call itself.

Since we did not want to change the virtual machine for this prototype so as to not have to implement a lot of tooling, this prototype is limited to undertaking an external call to urls that: (i) are exactly 32 bytes; (ii) start with 'http'; and (iii) respond with text in the following format:
Line 1: <A 32 byte message>
Line 2: <Signature R value>
Line 3: <Signature S value>
Line 4: <Public Key X value>
Line 5: <Public Key Y value>

In this particular implementation we are simulating a public key that could be retrived from a public certificate registry, or perhaps from the external system using it's domain name; however in a future implementation the public key could also be encoded within the smart contract. To do this a lot more work needs to be implemented (i.e. compiler, VM instructions, and modifications to various tools to support it).

As discussed in the paper, it is envisaged that signed responses will become more prevalent in the future. 


## Contact

Get in touch with us if you're interested in this work and want to collaborate or explore other ideas


## License

The go-ethereum library (i.e. all code outside of the `cmd` directory) is licensed under the
[GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html),
also included in our repository in the `COPYING.LESSER` file.

The go-ethereum binaries (i.e. all code inside of the `cmd` directory) is licensed under the
[GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also
included in our repository in the `COPYING` file.
