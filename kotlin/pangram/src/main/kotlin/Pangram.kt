object Pangram {
  fun isPangram(phrase: String): Boolean {

    return ('a'..'z').all { l -> phrase.contains(l, true) }
  }
}