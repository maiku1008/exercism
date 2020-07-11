CHILDREN = [
    "Alice",
    "Bob",
    "Charlie",
    "David",
    "Eve",
    "Fred",
    "Ginny",
    "Harriet",
    "Ileana",
    "Joseph",
    "Kincaid",
    "Larry",
]

PLANT_TYPES = {
    "G": "Grass",
    "C": "Clover",
    "V": "Violets",
    "R": "Radishes",
}


class Garden:
    def __init__(self, diagram: str, students=None):

        students = sorted(students if students is not None else CHILDREN)
        self.plant_table = {student: [None] * 4 for student in students}

        diagrams = [diag for diag in diagram.splitlines()]
        size = len(diagrams[0])
        zipped = list(
            zip(*[[dg[i : i + 2] for i in range(0, size, 2)] for dg in diagrams]),
        )
        grouped = ["".join(group) for group in zipped]

        for i, group in enumerate(grouped):
            for j, plant in enumerate(group):
                self.plant_table[students[i]][j] = PLANT_TYPES[plant]

    def plants(self, name) -> list:
        return self.plant_table[name]
