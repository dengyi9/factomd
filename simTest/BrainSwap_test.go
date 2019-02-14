package simtest

import (
	. "github.com/FactomProject/factomd/testHelper"
	"os"
	"strconv"
	"testing"
)

var logName string = "simTest"

func TestBrainSwap(t *testing.T) {

	t.Run("Run sim to create entries", func(t *testing.T) {
		givenNodes := os.Getenv("GIVEN_NODES")
		factomHome := os.Getenv("FACTOM_HOME")
		maxBlocks, _ := strconv.ParseInt(os.Getenv("MAX_BLOCKS"), 10, 64)
		peers := os.Getenv("PEERS")

		if factomHome == "" {
			factomHome = "."
		}

		if maxBlocks == 0 {
			maxBlocks = 30
		}

		if peers == "" {
			peers = "127.0.0.1:37003"
		}

		if givenNodes == "" {
			givenNodes = "LLLLAAA"
		}

		// FIXME update to match test data
		params := map[string]string{
			"--db":                  "LDB", // NOTE: using LEVELDB
			"--network":             "LOCAL",
			"--net":                 "alot+",
			"--enablenet":           "true",
			"--blktime":             "10",
			"--startdelay":          "1",
			"--stdoutlog":           "out.txt",
			"--stderrlog":           "out.txt",
			"--checkheads":          "false",
			"--controlpanelsetting": "readwrite",
			//"--debuglog":            ".",
			"--logPort":             "38000",
			"--port":                "38001",
			"--controlpanelport":    "38002",
			"--networkport":         "38003",
			"--peers":               peers,
			"--factomhome": 		 factomHome,
		}

		state0 := SetupSim(givenNodes, params, int(maxBlocks), 0, 0, t)
		state0.LogPrintf(logName, "GIVEN_NODES:%v", givenNodes)

		t.Run("Wait For Identity Swap", func(t *testing.T) {
			// NOTE: external scripts swap config files
			// during this time
			WaitForBlock(state0, 20)
		})

		t.Run("Verify Network", func(t *testing.T) {
			WaitBlocks(state0, 1)
			ShutDownEverything(t)
			WaitForAllNodes(state0)
		})

	})
}