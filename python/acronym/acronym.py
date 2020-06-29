from string import punctuation as punc


def abbreviate(words):
    return "".join(
        [
            word[0].upper()
            for word in words.replace("'", "")
            .translate(str.maketrans({k: " " for k in punc}))
            .split()
        ]
    )
