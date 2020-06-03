package tree


type Node struct {
	ID       int
	Children []*Node
}

type Record struct {
	ID     int
	Parent int
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	var n = &Node{}
	var err error

	for _, record := range records {
		if record.ID == 0 {
			n.ID = record.ID
		} else {
			n.Children = append(n.Children, &Node{record.ID, nil})
		}

	}
	return n, err

}
