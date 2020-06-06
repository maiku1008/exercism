package tree

import (
	"errors"
	"sort"
)

// Record contains information about a node
type Record struct {
	ID     int
	Parent int
}

// Node represents a node in the tree
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

	nodes := make([]*Node, len(records))
	for i, record := range records {
		switch {
		case i != record.ID:
			return nil, errors.New("record ids are not contiguous")
		case record.Parent > record.ID:
			return nil, errors.New("parent id is bigger than child id")
		case record.ID > 0 && record.Parent == record.ID:
			return nil, errors.New("child cannot be its own parent")
		}

		// Create the node and store it using the index number as its id
		node := &Node{ID: record.ID}
		nodes[i] = node

		// Adopt the child
		if record.ID > 0 {
			parent := nodes[record.Parent]
			parent.Children = append(parent.Children, nodes[record.ID])
		}
	}
	return nodes[0], nil
}
