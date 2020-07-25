from itertools import chain
from bisect import insort


class School:
    def __init__(self):
        self.__db = {grade: [] for grade in range(1, 8)}

    def add_student(self, name, grade):
        insort(self.__db[grade], name)

    def roster(self):
        return list(chain(*self.__db.values()))

    def grade(self, grade_number):
        return self.__db[grade_number]
