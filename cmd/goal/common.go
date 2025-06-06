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

package main

import (
	"github.com/spf13/cobra"

	cmdutil "github.com/algorand/go-algorand/cmd/util"
	"github.com/algorand/go-algorand/data/basics"
)

const (
	stdoutFilenameValue = "-"
	stdinFileNameValue  = "-"
)

// validateNoPosArgsFn is a reusable cobra positional argument validation function
// for generating proper error messages when commands see unexpected arguments when they expect no args.
// We don't use cobra.NoArgs directly, in case we want to customize behavior later.
var validateNoPosArgsFn = cobra.NoArgs

// transaction validity period margins
var firstValid basics.Round
var lastValid basics.Round

// numValidRounds specifies validity period for a transaction and used to calculate last valid round
var numValidRounds basics.Round // also used in account and asset

var (
	fee                uint64
	outFilename        string
	sign               bool
	noteBase64         string
	noteText           string
	lease              string
	noWaitAfterSend    bool
	dumpForDryrun      bool
	dumpForDryrunAccts []string
)

var dumpForDryrunFormat cmdutil.CobraStringValue = *cmdutil.MakeCobraStringValue("json", []string{"msgp"})

func addTxnFlags(cmd *cobra.Command) {
	cmd.Flags().Uint64Var(&fee, "fee", 0, "The transaction fee (automatically determined by default), in microAlgos")
	cmd.Flags().Uint64Var((*uint64)(&firstValid), "firstvalid", 0, "The first round where the transaction may be committed to the ledger")
	cmd.Flags().Uint64Var((*uint64)(&numValidRounds), "validrounds", 0, "The number of rounds for which the transaction will be valid")
	cmd.Flags().Uint64Var((*uint64)(&lastValid), "lastvalid", 0, "The last round where the transaction may be committed to the ledger")
	cmd.Flags().StringVarP(&outFilename, "out", "o", "", "Write transaction to this file")
	cmd.Flags().BoolVarP(&sign, "sign", "s", false, "Use with -o to indicate that the dumped transaction should be signed")
	cmd.Flags().StringVar(&noteBase64, "noteb64", "", "Note (URL-base64 encoded)")
	cmd.Flags().StringVarP(&noteText, "note", "n", "", "Note text (ignored if --noteb64 used also)")
	cmd.Flags().StringVarP(&lease, "lease", "x", "", "Lease value (base64, optional): no transaction may also acquire this lease until lastvalid")
	cmd.Flags().BoolVarP(&noWaitAfterSend, "no-wait", "N", false, "Don't wait for transaction to commit")
	cmd.Flags().BoolVar(&dumpForDryrun, "dryrun-dump", false, "Dump in dryrun format acceptable by dryrun REST api")
	cmd.Flags().Var(&dumpForDryrunFormat, "dryrun-dump-format", "Dryrun dump format: "+dumpForDryrunFormat.AllowedString())
	cmd.Flags().StringSliceVar(&dumpForDryrunAccts, "dryrun-accounts", nil, "additional accounts to include into dryrun request obj")
	cmd.Flags().StringVarP(&signerAddress, "signer", "S", "", "Address of key to sign with, if different from transaction \"from\" address due to rekeying")
}
