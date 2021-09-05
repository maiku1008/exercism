def slices(series, length):
    if length < 1 or length > len(series):
        raise ValueError("incorrect length")
    return [series[i : i + length] for i in range(1 + (len(series) - length))]
