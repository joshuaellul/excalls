package main

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joshuaellul/usecase-client/github.com/joshuaellul/usecase-client/testexcall"
)

const LOOP_COUNT = 1000

var totalOracleReplies = 0

const CONTRACT_ADDRESS = "0x622A4B186a31406858c0909a03E333b3140e9b28"
const ORACLE_PUBLIC_ADDRESS = "0x8144bE3b6F0767C9782fB020BCB4A60340247475"
const ORACLE_PRIVATE_KEY = "852fc068439ff78023ada2a321ddc9f0983c602722798149d8b6a1546fecdca9"

var privateKeys []string = []string{
	"2a295e31f2c116cc5df1245692a0cf0d22f3d4ffa0457be9ffb1f08bb0fe1d7d",
	"3c4580ee900b1e79a26a58fec18f647831654d21cb8842a27e1ba85ad6c77793",
	"1031fcea3cfbf6ffa7bb5bd155c43a5c84192bbf359e3e801ae43401a8e093ba",
	"07e5c10e4633a49ccce4c8019e4c47eca0c3f864faa2936dc3dff61e2430e9ac",
}

var winningsBefore []int = []int{0, 0, 0, 0}

func initClient() *ethclient.Client {
	//client, err := ethclient.Dial("http://127.0.0.1:22004")
	client, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatalf("Error on init!", err)
	}
	return client
}

func getAccountFromStrings(privateKeyString string) (*ecdsa.PrivateKey, common.Address) {
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, fromAddress
}

/*func getAccount(accountIndex int) (*ecdsa.PrivateKey, common.Address) {
	privateKeyString := privateKeys[accountIndex]
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, fromAddress
}*/

func callAlwaysWinEXCALL(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, lastNonce uint64) (*types.Transaction, error) {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	if nonce <= lastNonce && nonce != 0 {
		return nil, errors.New("Same as last nonce")
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(15))
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.Context = context.Background()

	//addr := common.HexToAddress("0x95f17B8fE8eb6c0678DD75b19A5e0f3477b094e5")
	addr := common.HexToAddress(CONTRACT_ADDRESS)
	instance, err := testexcall.NewTestEXCALL(addr, client)
	if err != nil {
		log.Fatal(err)
	}

	//tx, err := instance.SometimesWin(auth)
	//tx, err := instance.AlwaysWinLocal(auth)
	tx, err := instance.AlwaysWinGithub(auth)
	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
		return nil, err
	}

	return tx, nil
}

func callAlwaysWinOracle(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, lastNonce uint64) (*types.Transaction, error) {
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	if nonce <= lastNonce && nonce != 0 {
		return nil, errors.New("Same as last nonce")
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(15))
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.Context = context.Background()

	//addr := common.HexToAddress("0x95f17B8fE8eb6c0678DD75b19A5e0f3477b094e5")
	addr := common.HexToAddress(CONTRACT_ADDRESS)
	instance, err := testexcall.NewTestEXCALL(addr, client)
	if err != nil {
		log.Fatal(err)
	}

	//tx, err := instance.AlwaysWin(auth)
	tx, err := instance.BeginAlwaysWinOracle(auth)
	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
		return nil, err
	}

	return tx, nil
}

func callContinueWinOracle(client *ethclient.Client, instance *testexcall.TestEXCALL, privateKey *ecdsa.PrivateKey, fromAddress common.Address, punterAddress common.Address, oracleRefNr *big.Int) (*types.Transaction, error) {
	var pending bool
	pending, _ = instance.Pending(nil, punterAddress, oracleRefNr)
	if !pending {
		return nil, nil
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	/*if nonce <= lastNonce && nonce != 0 {
		return nil, errors.New("Same as last nonce")
	}*/

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(15))
	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	auth.Context = context.Background()

	answer := false

	replyFromGithubOracle := getStringFromUrl("https://joshuaellul.github.io/win")
	if replyFromGithubOracle[0] == '1' {
		answer = true
	} //*/
	answer = true //this process is the oracle... and the oracle is replying 'true'

	tx, err := instance.ContinueAlwaysWinOracle(auth, punterAddress, oracleRefNr, answer)

	if err != nil {
		log.Fatalf("Oops! There was a problem", err)
		return nil, err
	}

	return tx, nil
}

func WaitTillNoPendingTranstactions(client *ethclient.Client) error {
	ctx := context.Background()
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		pending, err := client.PendingTransactionCount(context.Background())
		if pending == 0 {
			return nil
		}
		if err != nil {
			return err
		}

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-queryTicker.C:
		}
	}
}

func callContractLoop(accountIndex int, client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	privateKey, fromAddress := getAccountFromStrings(privateKeys[accountIndex])

	//_ = callSometimesWin(client, privateKey, fromAddress)
	var lastNonce uint64
	for i := 0; i < LOOP_COUNT; i++ {
		tx, err := callAlwaysWinEXCALL(client, privateKey, fromAddress, lastNonce)
		if err != nil {
			//fmt.Println("ERROR: ", err)
			i -= 1
		} else {
			lastNonce = tx.Nonce()
		}
	}

	err := WaitTillNoPendingTranstactions(client)
	if err != nil {
		panic("Call failed")
	}
	fmt.Println("loop " + strconv.Itoa(accountIndex) + " done")
}

func callContractLoopOracle(accountIndex int, client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	privateKey, fromAddress := getAccountFromStrings(privateKeys[accountIndex])

	//_ = callSometimesWin(client, privateKey, fromAddress)
	var lastNonce uint64
	for i := 0; i < LOOP_COUNT; i++ {
		tx, err := callAlwaysWinOracle(client, privateKey, fromAddress, lastNonce)
		if err != nil {
			//fmt.Println("ERROR: ", err)
			i -= 1
		} else {
			lastNonce = tx.Nonce()
		}
	}

	err := WaitTillNoPendingTranstactions(client)
	if err != nil {
		panic("Call failed")
	}
	fmt.Println("loop " + strconv.Itoa(accountIndex) + " done")
}

func mainEXCALL(client *ethclient.Client) {

	var wg sync.WaitGroup

	for i := range privateKeys {
		wg.Add(1)
		go callContractLoop(i, client, &wg)
	}

	wg.Wait()
}

//var pendingContinues = []PendingContinue{}
//var mu = &sync.Mutex{}

/*func waitForContinueTransactionReceipts(client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	soc := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), soc)
	if err != nil {
		log.Fatal(err)
	}

	var totProcessed = 0

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-soc:
			if header != nil {
				//mu.Lock()
				//localCopyOfPCs := pendingContinues
				//mu.Unlock()
				b, err := client.BlockByNumber(context.Background(), header.Number)
				if err != nil {
					panic("Call failed")
				}

				for _, tx := range b.Transactions() {
					mu.Lock()

					//for x, realValue := range pendingContinues {
					for x := 0; x < len(pendingContinues); x++ {
						realValue := pendingContinues[x]
						if tx.Nonce() == realValue.nonce {
							pendingContinues[x] = pendingContinues[len(pendingContinues)-1]
							pendingContinues[len(pendingContinues)-1] = PendingContinue{}
							pendingContinues = pendingContinues[:len(pendingContinues)-1]
							totProcessed++
							break
						}
					}
					mu.Unlock()
					if totProcessed >= LOOP_COUNT {
						return
					}
				}

				/*for _, value := range localCopyOfPCs {
					tx, err := client.TransactionReceipt(context.Background(), value.txhash)
					if err != nil {
						//fmt.Println("transaction pending", err)
					} else {
						if tx.Status == 0 {
							//try again
							fmt.Println("transaction pending", value.txhash)
						} else {
							mu.Lock()
							for x, realValue := range pendingContinues {
								if realValue.txhash == value.txhash {
									pendingContinues[x] = pendingContinues[len(pendingContinues)-1]
									pendingContinues[len(pendingContinues)-1] = PendingContinue{}
									pendingContinues = pendingContinues[:len(pendingContinues)-1]
									totProcessed++
									break
								}
							}
							mu.Unlock()
							if totProcessed >= LOOP_COUNT {
								fmt.Println("PENDING CHECKER DONE")
								return
							}
						}
					}
				}
			}
		}
	}
}*/

func getStringFromUrl(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	bodys := string(body)
	resp.Body.Close()
	return bodys
}

func processRequestForOracleReply(vLog types.Log, instance *testexcall.TestEXCALL, client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address) {
	bpLog, err := instance.ParseBetPlaced(vLog)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	if err == nil {
	tryagain:
		tx, err := callContinueWinOracle(client, instance, privateKey, fromAddress, bpLog.From, bpLog.OracleRefNr)
		if err == nil && tx == nil {
			//this has already been processed
		} else if err != nil {
			//fmt.Println("ERROR: ", err)
			goto tryagain
		} else {
			//lastNonce = tx.Nonce()
			totalOracleReplies++
		}
	}
}

func oracleListenAndReply(latestBlock uint64, client *ethclient.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	addr := common.HexToAddress(CONTRACT_ADDRESS)
	instance, err := testexcall.NewTestEXCALL(addr, client)
	if err != nil {
		panic("Call failed")
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
		FromBlock: big.NewInt(int64(latestBlock)),
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, fromAddress := getAccountFromStrings(ORACLE_PRIVATE_KEY)

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
			//goto startagain
		case vLog := <-logs:
			processRequestForOracleReply(vLog, instance, client, privateKey, fromAddress)
		}
		if totalOracleReplies >= LOOP_COUNT*len(privateKeys) {
			//make sure the transactions are not marked as pending in the smart contract

			fmt.Println("ORACLE LISTENER DONE-")
			err := WaitTillNoPendingTranstactions(client)
			if err != nil {
				panic("Call failed")
			}
			fmt.Println("ORACLE LISTENER DONE----")
			return
		}
	}
}

func mainOracle(client *ethclient.Client, latestBlock uint64) {

	var wg sync.WaitGroup

	wg.Add(1)
	go oracleListenAndReply(latestBlock, client, &wg)
	//wg.Add(1)
	//go waitForContinueTransactionReceipts(client, &wg)

	for i := range privateKeys {
		wg.Add(1)
		go callContractLoopOracle(i, client, &wg)
	}

	wg.Wait()

}

func waitUntilWinningsReached(client *ethclient.Client) {
	addr := common.HexToAddress(CONTRACT_ADDRESS)
	instance, err := testexcall.NewTestEXCALL(addr, client)
	if err != nil {
		log.Fatal(err)
	}

	for i := range privateKeys {

		_, fromAddress := getAccountFromStrings(privateKeys[i])
		ok := false

		for !ok {
			bi, err := instance.Winnings(nil, fromAddress)
			if err != nil {
				log.Fatal(err)
			}
			if winningsBefore[i]+LOOP_COUNT <= int(bi.Int64()) {
				ok = true
			}
		}

	}
}

func main() {
	client := initClient()

	err := WaitTillNoPendingTranstactions(client)
	if err != nil {
		panic("Call failed")
	}

	//get winnings before starting
	addr := common.HexToAddress(CONTRACT_ADDRESS)
	instance, err := testexcall.NewTestEXCALL(addr, client)
	if err != nil {
		log.Fatal(err)
	}
	for i := range privateKeys {
		_, fromAddress := getAccountFromStrings(privateKeys[i])
		bi, err := instance.Winnings(nil, fromAddress)
		if err != nil {
			log.Fatal(err)
		}
		winningsBefore[i] = int(bi.Int64())
	}

	latestBlock, err := client.BlockNumber(context.Background())
	fmt.Println("First block: ", latestBlock)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Loop count: ", LOOP_COUNT)
	fmt.Println("Routines: ", len(privateKeys))
	start := time.Now()
	fmt.Println("Time with NanoSeconds: ", start.Format("2006-01-02 15:04:05.000000000"))

	mainOracle(client, latestBlock)
	mainEXCALL(client)

	err = WaitTillNoPendingTranstactions(client)
	if err != nil {
		panic("Call failed")
	}

	waitUntilWinningsReached(client)

	end := time.Now()
	fmt.Println("Time with NanoSeconds: ", end.Format("2006-01-02 15:04:05.000000000"))
	fmt.Println("Time taken: ", end.Sub(start))
}
