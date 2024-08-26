package mine_test // bench

import (
	"fmt"
	"testing"
	"time"

	"github.com/joaofnds/blockchain/block"
	"github.com/joaofnds/blockchain/blockchain"
	"github.com/joaofnds/blockchain/clock"
	"github.com/joaofnds/blockchain/hash"
	"github.com/joaofnds/blockchain/mine"
)

func BenchmarkConcurrent(b *testing.B) {
	now, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	clock := clock.NewFixedClock(now)
	miner := mine.NewConcurrent(hash.NewSHA256())
	chain := blockchain.New(clock, miner)
	chain.AddGenesisBlock()

	blk := block.NewBlock([]byte{}, clock.Now(), chain.LastBlock().Hash)

	for difficulty := 0; difficulty <= 7; difficulty++ {
		b.Run(fmt.Sprintf("difficulty %d", difficulty), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				miner.Mine(&blk, difficulty)
			}
		})
	}
}

func TestConcurrent(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	clock := clock.NewFixedClock(now)
	miner := mine.NewConcurrent(hash.NewSHA256())
	chain := blockchain.New(clock, miner)
	chain.AddGenesisBlock()

	testTable := []struct {
		difficulty    int
		expectedNonce uint64
	}{
		{4, 40018007},
		{5, 30033096},
		{6, 12186394},
		{7, 137747328},
		// {8, 3810002515},
	}

	for _, testCase := range testTable {
		t.Run(fmt.Sprintf("difficulty %d", testCase.difficulty), func(t *testing.T) {
			blk := block.NewBlock([]byte{}, clock.Now(), chain.LastBlock().Hash)

			miner.Mine(&blk, testCase.difficulty)

			if blk.Nonce != testCase.expectedNonce {
				t.Errorf("unexpected nonce: %d", blk.Nonce)
			}
		})
	}
}
