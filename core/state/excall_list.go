// Copyright 2020 The go-ethereum Authors
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

// Hacked together by Joshua Ellul and Gordon J. Pace 
// We've modified the the geth code to demonstrate that
// blockchain systems can perform calls to external systems (and use) their
// response in associated smart contract computation. This prototype only
// serves as a demonstration to prove the concept but further work is needed
// to fully support this. See https://joshuaellul.github.io/excalls
// The concept is presented within a pre-print paper titled
// "Towards External Calls for Blockchain and Distributed Ledger Technology"
// and is available from: https://www.researchgate.net/publication/350959398

package state

import (
	"math/big"
)

type excallList struct {
	msgs [][]byte
	sigr []*big.Int
	sigs []*big.Int
	pubx []*big.Int
	puby []*big.Int
}

func newExcallList() *excallList {
	return &excallList{}
}

func (e *excallList) Copy() *excallList {
	cp := newExcallList()
	cp.msgs = make([][]byte, len(e.msgs))
	cp.sigr = make([]*big.Int, len(e.sigr))
	cp.sigs = make([]*big.Int, len(e.sigr))
	cp.pubx = make([]*big.Int, len(e.sigr))
	cp.puby = make([]*big.Int, len(e.sigr))

	for i, m := range e.msgs {
		cp.msgs[i] = make([]byte, len(m))
		cp.msgs[i] = m
		cp.sigr[i] = e.sigr[i]
		cp.sigs[i] = e.sigs[i]
		cp.pubx[i] = e.pubx[i]
		cp.puby[i] = e.puby[i]
	}

	return cp
}
