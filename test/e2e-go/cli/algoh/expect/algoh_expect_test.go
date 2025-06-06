// Copyright (C) 2019-2025 Algorand, Inc.
// This file is part of go-algorand
//
// go-algorand is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-algorand is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with go-algorand.  If not, see <https://www.gnu.org/licenses/>.
package expect

import (
	"testing"

	"github.com/algorand/go-algorand/test/framework/fixtures"
)

// TestAlgohWithExpect Process all expect script files with suffix Test.exp within the test/e2e-go/cli/algoh/expect directory
func TestAlgohWithExpect(t *testing.T) {
	// partitiontest.PartitionTest(t)
	// Causes double partition, so commented out on purpose
	defer fixtures.ShutdownSynchronizedTest(t)
	et := fixtures.MakeExpectTest(t)
	et.Run()
}
