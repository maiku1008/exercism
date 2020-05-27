def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("Value must have the same length")

    return len([strand for strand in zip(strand_a, strand_b) if strand[0] != strand[1]])
