def to_rna(dna_strand):
  gene_map = {
    "G": "C",
    "C": "G",
    "T": "A",
    "A": "U"
  }

  return "".join([gene_map[gene] for gene in dna_strand])