class Matrix:
    def __init__(self, matrix_string):
        self.list = [
            [int(i) for i in row.split(" ")] for row in matrix_string.split("\n")
        ]

    def row(self, index):
        return self.list[index - 1]

    def column(self, index):
        return [i[index - 1] for i in self.list]
