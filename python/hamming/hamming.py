def distance(strand_a, strand_b):
    if len(strand_a) != len(strand_b):
        raise ValueError("Value must have the same length")

    return len([i for i in range(len(strand_a)) if strand_a[i] != strand_b[i]])


# def distance(strand_a, strand_b):
#     if len(strand_a) != len(strand_b):
#         raise ValueError("Value must have the same length")
#
#     diff = 0
#     for i in range(len(strand_a)):
#         if strand_a[i] != strand_b[i]:
#             diff += 1
#
#     return diff
