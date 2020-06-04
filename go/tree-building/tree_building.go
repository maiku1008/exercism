package tree

import "errors"

// Record contains information about a possible node
type Record struct {
	ID     int
	Parent int
}

// Node represents a node in our tree
type Node struct {
	ID       int
	Children []*Node
}

// Build builds the tree
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	root, err := setRoot(records)
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		// Is root, skip it
		if record.ID == 0 {
			continue
		}

		// Check if the parent ID is the same as root node
		if record.Parent == root.ID {
			adoptChild(root, &Node{ID: record.ID})
		}

		// Check if the parent ID is in the root node's children ?
	}
	return root, nil
}

// A helper function that adopts a child and would check for errors at some point
func adoptChild(parent, child *Node) {
	// TODO: Check if child is already there
	parent.Children = append(parent.Children, child)
}

// A helper function that sets the root before dealing with the children
func setRoot(records []Record) (*Node, error) {
	var root *Node
	for _, record := range records {
		switch {
		case record.ID == 0 && record.Parent != 0:
			return nil, errors.New("node with id 0 has parent")
		case record.ID == 0 && record.Parent == 0:
			root = &Node{
				ID: record.ID,
			}
		}
	}
	return root, nil
}
