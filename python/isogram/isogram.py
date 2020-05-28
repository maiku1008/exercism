def is_isogram(word):
    seen = set()
    for c in word:
        if c != "-" or c != " ":
            c = c.lower()
            if c.lower() in seen:
                return False
            seen.add(c)
    return True
