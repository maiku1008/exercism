from collections import Counter

# Custom punctuation string without single quote '
punctuation = r"""!"#$%&()*+,-./:;<=>?@[\]^_`{|}~"""

def count_words(sentence):
    # Translate anything found in constant punctuation into whitespace
    table = str.maketrans({k: " " for k in punctuation})
    sentence = sentence.translate(table)

    # Strip single quotes surrounding words
    s = [w.strip("'") for w in sentence.lower().split()]

    return dict(Counter(s))
