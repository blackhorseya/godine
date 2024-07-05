package snowflakex

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

// NewNode create and return a new snowflake node
func NewNode() (*snowflake.Node, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, fmt.Errorf("new snowflake node failed: %w", err)
	}

	return node, nil
}
