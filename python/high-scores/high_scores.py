def latest(scores):
    ''' Returns the latest score(int) '''
    return scores[-1]


def personal_best(scores):
    ''' Returns a personal best score(int) '''
    sorted_scores = sorted(scores, reverse=True)
    return sorted_scores[0]


def personal_top_three(scores):
    '''
    Find the top three scores in a list.

    Returns a list of scores(int).
    '''
    sorted_scores = sorted(scores, reverse=True)
    return sorted_scores[:3]
