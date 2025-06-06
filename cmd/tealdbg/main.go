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
	"log"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	cmdutil "github.com/algorand/go-algorand/cmd/util"
	"github.com/algorand/go-algorand/data/basics"
)

func main() {
	// Hidden command to generate docs in a given directory
	// tealdbg generate-docs [path]
	if len(os.Args) == 3 && os.Args[1] == "generate-docs" {
		err := doc.GenMarkdownTree(rootCmd, os.Args[2])
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "tealdbg",
	Short: "Algorand TEAL Debugger",
	Long: `Debug a local or remote TEAL code in controlled environment
with Web or Chrome DevTools frontends`,
	Run: func(cmd *cobra.Command, args []string) {
		//If no arguments passed, we should fallback to help
		cmd.HelpFunc()(cmd, args)
	},
}

var debugCmd = &cobra.Command{
	Use:   "debug [program.tok [program.teal ...]]",
	Short: "Debug TEAL program(s) off-chain",
	Long:  `Debug TEAL program(s) in controlled environment using a local TEAL evaluator`,
	Run: func(cmd *cobra.Command, args []string) {
		debugLocal(args)
		// //If no arguments passed, we should fallback to help
		// cmd.HelpFunc()(cmd, args)
	},
}

var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Debug TEAL program on-chain",
	Long:  `Start the server and wait for upcoming debug connections from remote TEAL evaluator`,
	Run: func(cmd *cobra.Command, args []string) {
		debugRemote()
	},
}

type frontendValue struct {
	*cmdutil.CobraStringValue
}

func (f *frontendValue) value() string {
	return f.CobraStringValue.String()
}

func (f *frontendValue) Make(router *mux.Router, appAddress string) (da DebugAdapter) {
	switch f.value() {
	case "web":
		wa := MakeWebPageFrontend(&WebPageFrontendParams{router, appAddress})
		return wa
	case "cdt":
		fallthrough
	default:
		cdt := MakeCdtFrontend(&CdtFrontendParams{router, appAddress, verbose})
		return cdt
	}
}

type runModeValue struct {
	*cmdutil.CobraStringValue
}

var frontend frontendValue = frontendValue{cmdutil.MakeCobraStringValue("cdt", []string{"web"})}
var proto string
var txnFile string
var groupIndex int
var balanceFile string
var ddrFile string
var indexerURL string
var indexerToken string
var roundNumber uint64
var timestamp int64
var runMode runModeValue = runModeValue{cmdutil.MakeCobraStringValue("auto", []string{"signature", "application"})}
var port int
var iface string
var noFirstRun bool
var noBrowserCheck bool
var noSourceMap bool
var verbose bool
var painless bool
var appID basics.AppIndex
var listenForDrReq bool

func init() {
	rootCmd.PersistentFlags().VarP(&frontend, "frontend", "f", "Frontend to use: "+frontend.AllowedString())
	rootCmd.PersistentFlags().IntVar(&port, "remote-debugging-port", 9392, "Port to listen on")
	rootCmd.PersistentFlags().StringVar(&iface, "listen", "127.0.0.1", "Network interface to listen on")
	rootCmd.PersistentFlags().BoolVar(&noFirstRun, "no-first-run", false, "")
	rootCmd.PersistentFlags().MarkHidden("no-first-run")
	rootCmd.PersistentFlags().BoolVar(&noBrowserCheck, "no-default-browser-check", false, "")
	rootCmd.PersistentFlags().MarkHidden("no-default-browser-check")
	rootCmd.PersistentFlags().BoolVar(&noSourceMap, "no-source-map", false, "Do not generate source maps")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Verbose output")

	debugCmd.Flags().StringVarP(&proto, "proto", "p", "", "Consensus protocol version for TEAL evaluation")
	debugCmd.Flags().StringVarP(&txnFile, "txn", "t", "", "Transaction(s) to evaluate TEAL on in form of json or msgpack file")
	debugCmd.Flags().IntVarP(&groupIndex, "group-index", "g", 0, "Transaction index in a txn group")
	debugCmd.Flags().StringVarP(&balanceFile, "balance", "b", "", "Balance records to evaluate stateful TEAL on in form of json or msgpack file")
	debugCmd.Flags().StringVarP(&ddrFile, "dryrun-req", "d", "", "Program(s) and state(s) in dryrun REST request format")
	debugCmd.Flags().Uint64VarP((*uint64)(&appID), "app-id", "a", 1380011588, "Application ID for stateful TEAL if not set in transaction(s)")
	debugCmd.Flags().Uint64VarP(&roundNumber, "round", "r", 0, "Ledger round number to evaluate stateful TEAL on")
	debugCmd.Flags().Int64VarP(&timestamp, "latest-timestamp", "l", 0, "Latest confirmed timestamp to evaluate stateful TEAL on")
	debugCmd.Flags().VarP(&runMode, "mode", "m", "TEAL evaluation mode: "+runMode.AllowedString())
	debugCmd.Flags().BoolVar(&painless, "painless", false, "Automatically create balance record for all accounts and applications")
	debugCmd.Flags().StringVarP(&indexerURL, "indexer-url", "i", "", "URL for indexer to fetch Balance records from to evaluate stateful TEAL")
	debugCmd.Flags().StringVarP(&indexerToken, "indexer-token", "", "", "API token for indexer to fetch Balance records from to evaluate stateful TEAL")
	debugCmd.Flags().BoolVarP(&listenForDrReq, "listen-dr-req", "q", false, "Listen for upcoming debugging dryrun request objects instead of taking program(s) from command line")

	rootCmd.AddCommand(debugCmd)
	rootCmd.AddCommand(remoteCmd)
}

func debugRemote() {
	ds := makeDebugServer(iface, port, &frontend, nil)
	err := ds.startRemote()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func debugLocal(args []string) {
	// local debugging works in two modes:
	// - listening for upcoming Dryrun Requests
	// - or taking program, transaction or Dryrun Request from command line
	// they can not be combined
	if listenForDrReq && (len(args) != 0 || len(txnFile) != 0 && len(ddrFile) != 0) {
		log.Fatalln("Can not combine listening for Dryrun Requests and program(s), or transaction(s), or dryrun-req object")
	}

	if !listenForDrReq {
		// program can be set either directly
		// or with SignedTxn.Lsig.Logic,
		// or with BalanceRecord.AppParams.ApprovalProgram
		if len(args) == 0 && len(txnFile) == 0 && len(ddrFile) == 0 {
			log.Fatalln("No program to debug: must specify program(s), or transaction(s), or dryrun-req object")
		}

		if len(args) == 0 && groupIndex != 0 {
			log.Fatalln("Error: group-index may be only set only along with program(s)")
		}

		if len(args) == 0 && runMode.IsSet() {
			log.Fatalln("Error: mode may be only set only along with program(s)")
		}

		if len(txnFile) != 0 && len(ddrFile) != 0 {
			log.Fatalln("Error: cannot specify both transaction(s) and dryrun-req")
		}

		if len(balanceFile) != 0 && len(ddrFile) != 0 {
			log.Fatalln("Error: cannot specify both balance records(s) and dryrun-req")
		}
	}

	var programNames []string
	var programBlobs [][]byte
	if len(args) > 0 {
		programNames = make([]string, len(args))
		programBlobs = make([][]byte, len(args))
		for i, file := range args {
			data, err := os.ReadFile(file)
			if err != nil {
				log.Fatalf("Error program reading %s: %s", file, err)
			}
			programNames[i] = file
			programBlobs[i] = data
		}
	}

	var err error
	var txnBlob []byte
	if len(txnFile) > 0 {
		txnBlob, err = os.ReadFile(txnFile)
		if err != nil {
			log.Fatalf("Error txn reading %s: %s", txnFile, err)
		}
	}

	var balanceBlob []byte
	if len(balanceFile) > 0 {
		balanceBlob, err = os.ReadFile(balanceFile)
		if err != nil {
			log.Fatalf("Error balance reading %s: %s", balanceFile, err)
		}
	}

	var ddrBlob []byte
	if len(ddrFile) > 0 {
		ddrBlob, err = os.ReadFile(ddrFile)
		if err != nil {
			log.Fatalf("Error dryrun-dump reading %s: %s", ddrFile, err)
		}
	}

	dp := DebugParams{
		ProgramNames:     programNames,
		ProgramBlobs:     programBlobs,
		Proto:            proto,
		TxnBlob:          txnBlob,
		GroupIndex:       groupIndex,
		BalanceBlob:      balanceBlob,
		DdrBlob:          ddrBlob,
		IndexerURL:       indexerURL,
		IndexerToken:     indexerToken,
		Round:            basics.Round(roundNumber),
		LatestTimestamp:  timestamp,
		RunMode:          runMode.String(),
		DisableSourceMap: noSourceMap,
		AppID:            appID,
		Painless:         painless,
		ListenForDrReq:   listenForDrReq,
	}

	ds := makeDebugServer(iface, port, &frontend, &dp)

	err = ds.startDebug()
	if err != nil {
		log.Fatalf("Debug error: %s", err.Error())
	}
}
