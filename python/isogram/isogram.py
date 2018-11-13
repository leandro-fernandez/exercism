def is_isogram(string):
    lowerString = string.lower()
    dict = {}

    for char in lowerString:
      if char.isalpha():
        if (char in dict): return False
        else: dict[char] = 1
    
    return True
