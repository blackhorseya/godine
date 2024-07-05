package snowflakex

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/bwmarrin/snowflake"
)

// NewNode create and return a new snowflake node
func NewNode() (*snowflake.Node, error) {
	// Define the range (2^10)
	maxInt := big.NewInt(1 << 10)

	// Generate a random number in the range [0, 2^10)
	num, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		return nil, err
	}

	// Convert to an integer
	id := num.Int64()

	node, err := snowflake.NewNode(id)
	if err != nil {
		return nil, fmt.Errorf("new snowflake node failed: %w", err)
	}

	return node, nil
}
