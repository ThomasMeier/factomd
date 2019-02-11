package simtest

import (
	. "github.com/FactomProject/factomd/testHelper"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

var logName string = "simTest"

func TestBrainSwap(t *testing.T) {

	t.Run("Run sim to create entries", func(t *testing.T) {
		givenNodes := os.Getenv("GIVEN_NODES")
		maxBlocks, _ := strconv.ParseInt(os.Getenv("MAX_BLOCKS"), 10, 64)

		if maxBlocks == 0 {
			maxBlocks = 20
		}

		if givenNodes == "" {
			givenNodes = "L"
		}

		params := map[string]string{
			"--debuglog": "",
			"--peers": "127.0.0.1:34341",
			"--prefix": "x",
		}

		state0 := SetupSim(givenNodes, params, int(maxBlocks), 0, 0, t)
		state0.LogPrintf(logName, "GIVEN_NODES:%v", givenNodes)

		t.Run("Swap Identities", func(t *testing.T) {

			t.Run("Remove Leader", func(t *testing.T) {
				RunCmd("0")
				RunCmd("z")
				WaitBlocks(state0, 1)

				l, _, _ := CountAuthoritySet()
				assert.Equal(t, 3, l, "should be 3 leaders")
			})

			/// TODO: do a brainswap instead
			t.Run("Promote Follower", func(t *testing.T) {
				RunCmd("4")
				RunCmd("l")
				WaitBlocks(state0, 1)
				l, _, _ := CountAuthoritySet()
				assert.Equal(t, 4, l, "should be 4 leaders again")
			})

		})

		t.Run("Verify Network", func(t *testing.T) {
			WaitBlocks(state0, 1)
			ShutDownEverything(t)
			WaitForAllNodes(state0)
		})

	})
}
