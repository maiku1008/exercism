#include "grade_school.h"
#include <stdbool.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>

static roster_t roster = { 0 };

roster_t get_roster()
{
    qsort(roster.students, roster.count, sizeof(student_t), cmpstudents);
    return roster;
}

int cmpstudents(const void *student_a, const void *student_b)
{
    student_t *l = (student_t *)student_a;
    student_t *r = (student_t *)student_b;
    if (l->grade == r->grade)
    {
        return strcmp(l->name, r->name);
    }
    return (l->grade > r->grade) ? (1) : (-1);
}

roster_t get_grade(uint8_t desired_grade)
{
    // Here I chose to create a separate roster with only the grade we're interested in,
    // sort it, and return it. This seems wasteful however.
    // What is a better way?
    roster_t filtered_roster = {0};
    for (size_t i = 0; i < roster.count; i++)
    {
        if (roster.students[i].grade == desired_grade)
        {
            filtered_roster.students[i] = roster.students[i];
            filtered_roster.count++;
        }
    }
    qsort(filtered_roster.students, filtered_roster.count, sizeof(student_t), cmpstudents);
    return filtered_roster;
}

void clear_roster()
{
    roster.count = 0;
    // Dear mentor: what is an acceptable way to zero out an array with custom types?
    // I have a hunch this is wrong and that it could possibly lead to memory leaks
    for (int i = 0; i < MAX_STUDENTS; i++)
    {
        roster.students[i].name = NULL;
        roster.students[i].grade = 0;
    }
    return;
}

bool add_student(char *name, uint8_t grade)
{
    if (roster.count >= MAX_STUDENTS)
    {
        return false;
    }
    if (strlen(name) > MAX_NAME_LENGTH)
    {
        return false;
    }

    student_t s = { .name = name, .grade = grade};

    roster.students[roster.count] = s;
    roster.count++;

    return true;
}


