package tree

import (
	"errors"
	"fmt"
	"sort"
)

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
	// Return if records is empty
	if len(records) == 0 {
		return nil, nil
	}

	// Sort the records
	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	var root *Node
	for _, record := range records {
		// Set Root node
		if record.ID == 0 {
			switch {
			case root != nil:
				return nil, errors.New("root node exists and cannot be duplicated")
			case record.ID == 0 && record.Parent != 0:
				return nil, errors.New("node with id 0 has parent")
			case record.ID == 0 && record.Parent == 0:
				root = &Node{
					ID: record.ID,
				}
			}
			continue
		}

		// Check the record's parent, and match it with the right node.
		parent, err := findNodeByID(record.Parent, root)
		if err != nil {
			return nil, err
		}
		// We need to append the ID of the record to the correct parent node (record.Parent)
		err = adoptChild(parent, &Node{ID: record.ID})
		if err != nil {
			return nil, err
		}
	}
	return root, nil
}

// A helper function that adopts a child if it's not there already
func adoptChild(parent, child *Node) error {
	for _, member := range parent.Children {
		if member.ID == child.ID {
			return fmt.Errorf("duplicate child: %d", child.ID)
		}
	}
	parent.Children = append(parent.Children, child)

	return nil
}

// findNodeByID looks for a node with the id
func findNodeByID(id int, start *Node) (*Node, error) {
	if start == nil {
		return nil, errors.New("starting node is nil")
	}

	// Found
	if id == start.ID {
		return start, nil
	}

	for _, node := range start.Children {
		// Found
		if node.ID == id {
			return node, nil
		}
	}

	return nil, fmt.Errorf("could not find node with id: %d ", id)
}
