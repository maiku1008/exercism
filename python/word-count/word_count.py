from collections import Counter
from string import punctuation


def clean(sentence):
    return [w.strip(punctuation) for w in sentence.lower().split()]


def count_words(sentence):
    return dict(Counter(clean(sentence)))
