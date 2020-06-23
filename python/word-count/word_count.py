from collections import Counter
import string


def count_words(sentence):
    # Translate anything found punctuation into whitespace
    table = str.maketrans({k: " " for k in string.punctuation.replace("'", "")})
    sentence = sentence.translate(table)

    # Strip single quotes surrounding words
    stripped = [w.strip("'") for w in sentence.lower().split()]

    return dict(Counter(stripped))
