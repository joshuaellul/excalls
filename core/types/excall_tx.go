// Copyright 2021 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Hacked together by Joshua Ellul and Gordon J. Pace to demonstrate that
// blockchain systems can perform calls to external systems (and use) their
// response in associated smart contract computation. This prototype only
// serves as a demonstration to prove the concept but further work is needed
// to fully support this. See https://joshuaellul.github.io/excalls
// The concept is presented within a pre-print paper titled
// "Towards External Calls for Blockchain and Distributed Ledger Technology"
// and is available from: https://www.researchgate.net/publication/350959398

package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//go:generate gencodec -type ExcallTuple -out gen_excall_tuple.go

type ExcallList []ExcallTuple

type ExcallTuple struct {
	Msg  []byte   `json:"msg"     gencodec:"required"`
	SigR *big.Int `json:"sigr"    gencodec:"required"`
	SigS *big.Int `json:"sigx"    gencodec:"required"`
	PubX *big.Int `json:"pubx"    gencodec:"required"`
	PubY *big.Int `json:"puby"    gencodec:"required"`
}

type ExcallListTx struct {
	ChainID    *big.Int        // destination chain ID
	Nonce      uint64          // nonce of sender account
	GasPrice   *big.Int        // wei per gas
	Gas        uint64          // gas limit
	To         *common.Address `rlp:"nil"` // nil means contract creation
	Value      *big.Int        // wei amount
	Data       []byte          // contract invocation input data
	ExcallList ExcallList
	V, R, S    *big.Int // signature values
}

func (el ExcallList) GetEXCALLTuple() *ExcallTuple {
	for _, tuple := range el {
		return &tuple
	}
	return nil
}

func (tx *ExcallListTx) SetEXCALLTuple(msg []byte, sigR *big.Int, sigS *big.Int, pubX *big.Int, pubY *big.Int) {
	tx.ExcallList = make(ExcallList, 1)
	tx.ExcallList[0] = ExcallTuple{
		Msg:  common.CopyBytes(msg),
		SigR: new(big.Int).Set(sigR),
		SigS: new(big.Int).Set(sigS),
		PubX: new(big.Int).Set(pubX),
		PubY: new(big.Int).Set(pubY),
	}
}

func (tx *ExcallListTx) CopyWithoutExcalls() TxData {
	cpy := &ExcallListTx{
		Nonce: tx.Nonce,
		To:    tx.To, // TODO: copy pointed-to address
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are copied below.
		ExcallList: make(ExcallList, 0),
		Value:      new(big.Int),
		ChainID:    new(big.Int),
		GasPrice:   new(big.Int),
		V:          new(big.Int),
		R:          new(big.Int),
		S:          new(big.Int),
	}
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// copy creates a deep copy of the transaction data and initializes all fields.
func (tx *ExcallListTx) copy() TxData {
	cpy := &ExcallListTx{
		Nonce: tx.Nonce,
		To:    tx.To, // TODO: copy pointed-to address
		Data:  common.CopyBytes(tx.Data),
		Gas:   tx.Gas,
		// These are copied below.
		ExcallList: make(ExcallList, len(tx.ExcallList)),
		Value:      new(big.Int),
		ChainID:    new(big.Int),
		GasPrice:   new(big.Int),
		V:          new(big.Int),
		R:          new(big.Int),
		S:          new(big.Int),
	}
	copy(cpy.ExcallList, tx.ExcallList)
	if tx.Value != nil {
		cpy.Value.Set(tx.Value)
	}
	if tx.ChainID != nil {
		cpy.ChainID.Set(tx.ChainID)
	}
	if tx.GasPrice != nil {
		cpy.GasPrice.Set(tx.GasPrice)
	}
	if tx.V != nil {
		cpy.V.Set(tx.V)
	}
	if tx.R != nil {
		cpy.R.Set(tx.R)
	}
	if tx.S != nil {
		cpy.S.Set(tx.S)
	}
	return cpy
}

// accessors for innerTx.

func (tx *ExcallListTx) txType() byte           { return ExcallListTxType }
func (tx *ExcallListTx) chainID() *big.Int      { return tx.ChainID }
func (tx *ExcallListTx) protected() bool        { return true }
func (tx *ExcallListTx) accessList() AccessList { return nil }
func (tx *ExcallListTx) excallList() ExcallList { return tx.ExcallList }
func (tx *ExcallListTx) data() []byte           { return tx.Data }
func (tx *ExcallListTx) gas() uint64            { return tx.Gas }
func (tx *ExcallListTx) gasPrice() *big.Int     { return tx.GasPrice }
func (tx *ExcallListTx) value() *big.Int        { return tx.Value }
func (tx *ExcallListTx) nonce() uint64          { return tx.Nonce }
func (tx *ExcallListTx) to() *common.Address    { return tx.To }

func (tx *ExcallListTx) rawSignatureValues() (v, r, s *big.Int) {
	return tx.V, tx.R, tx.S
}

func (tx *ExcallListTx) setSignatureValues(chainID, v, r, s *big.Int) {
	tx.ChainID, tx.V, tx.R, tx.S = chainID, v, r, s
}
