fun transcribeSingleNucleotide(nucleotide: Char) : Char =  when (nucleotide) {
    'G' -> 'C'
    'C' -> 'G'
    'T' -> 'A'
    'A' -> 'U'
    else -> ' '
}

fun transcribeToRna(dna: String): String {
    return dna.map(::transcribeSingleNucleotide).joinToString("")
}