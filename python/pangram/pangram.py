def is_pangram(sentence):
    alphabet = "abcdefghijklmnopqrstuvxyz"

    for letter in sentence:
      if (letter.isalpha() and letter.lower() in alphabet):
        alphabet = alphabet.replace(letter.lower(), "")

    return len(alphabet) == 0
