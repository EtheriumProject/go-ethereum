// Copyright 2016 The go-etherium Authors
// This file is part of the go-etherium library.
//
// The go-etherium library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-etherium library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-etherium library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "github.com/etherium/go-etherium"

// Verify that Client implements the etherium interfaces.
var (
	_ = etherium.ChainReader(&Client{})
	_ = etherium.TransactionReader(&Client{})
	_ = etherium.ChainStateReader(&Client{})
	_ = etherium.ChainSyncReader(&Client{})
	_ = etherium.ContractCaller(&Client{})
	_ = etherium.GasEstimator(&Client{})
	_ = etherium.GasPricer(&Client{})
	_ = etherium.LogFilterer(&Client{})
	_ = etherium.PendingStateReader(&Client{})
	// _ = etherium.PendingStateEventer(&Client{})
	_ = etherium.PendingContractCaller(&Client{})
)
