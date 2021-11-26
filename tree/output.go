package tree

import (
	"fmt"
	"io"
)

// Printf writes a tree on "out" in the format that CreateFromString or
// CreateNumericFromString can turn into a tree.
func Printf(out io.Writer, node Node) {
	if node.IsNil() {
		fmt.Fprintf(out, "()")
		return
	}
	fmt.Fprintf(out, "(%s", node) // relies on node fitting fmt.Stringer
	if !node.LeftChild().IsNil() || !node.RightChild().IsNil() {
		Printf(out, node.LeftChild())
		Printf(out, node.RightChild())
	}
	fmt.Fprintf(out, ")")
}
