def is_isogram(word):
    stripped = [c.lower() for c in word if c.isalpha()]
    return len(stripped) == len(set(stripped))
