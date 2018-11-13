def is_isogram(string):
    lowerString = string.lower().replace(" ", "").replace("-", "")
    return len(set(lowerString)) == len(lowerString)
