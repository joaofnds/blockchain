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

func BenchmarkSeq(b *testing.B) {
	now, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	clock := clock.NewFixedClock(now)
	miner := mine.NewSeq(hash.NewSHA256())
	chain := blockchain.New(clock, miner)
	chain.AddGenesisBlock()

	blk := block.NewBlock([]byte{}, clock.Now(), chain.LastBlock().Hash)

	for difficulty := 1; difficulty <= 6; difficulty++ {
		b.Run(fmt.Sprintf("difficulty %d", difficulty), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				miner.Mine(&blk, difficulty)
			}
		})
	}
}

func TestSeq(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2021-01-01T00:00:00Z")
	clock := clock.NewFixedClock(now)
	miner := mine.NewSeq(hash.NewSHA256())
	chain := blockchain.New(clock, miner)
	chain.AddGenesisBlock()

	testTable := []struct {
		difficulty    int
		expectedNonce uint64
	}{
		{1, 17},
		{2, 61},
		{3, 4910},
		{4, 34551},
		{5, 379280},
		{6, 8921088},
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
