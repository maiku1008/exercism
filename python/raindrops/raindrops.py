def convert(number):
    '''
    Take a number and return a string depending on its factors.
    '''
    result = ""
    if number % 3 == 0:
        result += "Pling"

    if number % 5 == 0:
        result += "Plang"

    if number % 7 == 0:
        result += "Plong"

    if len(result) == 0:
        result += f"{number}"

    return result
