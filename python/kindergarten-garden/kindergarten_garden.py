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
        self.plant_table = {student: [] for student in students}

        for line in diagram.splitlines():
            student_index, increment = 0, False
            for char in line:
                self.plant_table[students[student_index]].append(PLANT_TYPES[char])
                if increment:
                    student_index += 1
                    increment = False
                else:
                    increment = True

    def plants(self, name) -> list:
        return self.plant_table[name]
