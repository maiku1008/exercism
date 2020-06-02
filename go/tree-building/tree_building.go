package tree

type Node struct {
	ID       int
	Children []*Node
}

type Record struct {
	ID     int
	Parent int
}

func Build(records []Record) (*Node, error)  {
	var n = &Node{}
	var err error

	for _, record := range records {
		n.ID = record.ID
	}

	return n, err
}