LETTERS = set("abcdefghijklmnopqrstuvwxyz")


def is_pangram(sentence: str) -> bool:
    return LETTERS.issubset(sentence.lower())
