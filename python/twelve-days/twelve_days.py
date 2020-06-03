verses = {
    "first": "a Partridge in a Pear Tree.",
    "second": "two Turtle Doves, ",
    "third": "three French Hens, ",
    "fourth": "four Calling Birds, ",
    "fifth": "five Gold Rings, ",
    "sixth": "six Geese-a-Laying, ",
    "seventh": "seven Swans-a-Swimming, ",
    "eighth": "eight Maids-a-Milking, ",
    "ninth": "nine Ladies Dancing, ",
    "tenth": "ten Lords-a-Leaping, ",
    "eleventh": "eleven Pipers Piping, ",
    "twelfth": "twelve Drummers Drumming, ",
}
day = list(verses)


def recite(start_verse, end_verse):
    # Adjust for indexing day
    start_verse -= 1

    if end_verse - start_verse == 0:
        return [single_verse(start_verse)]

    return [single_verse(i) for i in range(start_verse, end_verse)]


def single_verse(start_verse):
    recited = [f"On the {day[start_verse]} day of Christmas my true love gave to me: "]
    for i in range(start_verse, -1, -1):
        recited.append(verses[day[i]])

    if len(recited) > 2:
        recited[-1] = "and " + recited[-1]

    return "".join(recited)
