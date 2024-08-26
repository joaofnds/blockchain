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

const (
	minDifficulty = 1
	maxDifficulty = 6
)

func BenchmarkMineSeq(b *testing.B) {
	now, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	clock := clock.NewFixedClock(now)
	miner := mine.NewSeq(hash.NewSHA256())
	chain := blockchain.New(clock, miner)
	chain.AddGenesisBlock()

	blk := block.NewBlock([]byte{}, clock.Now(), chain.LastBlock().Hash)

	for difficulty := minDifficulty; difficulty <= maxDifficulty; difficulty++ {
		b.Run(fmt.Sprintf("difficulty %d", difficulty), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				miner.Mine(&blk, difficulty)
			}
		})
	}
}
