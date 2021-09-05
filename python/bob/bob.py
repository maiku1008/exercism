def response(hey_bob):
    hey_bob = "".join(hey_bob.split())
    print(hey_bob)
    print(all(c.isupper() for c in hey_bob if c.isalpha()))
    if all(c.isupper() for c in hey_bob if c.isalpha()):
        return "Whoa, chill out!"
    if hey_bob.endswith("?"):
        return "Sure."
    if hey_bob == "":
        return "Fine. Be that way!"
    return "Whatever."
