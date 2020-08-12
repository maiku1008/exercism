const COMPLEMENTS = [['G', 'C'], ['C', 'G'], ['T', 'A'], ['A', 'U']]

export const toRna = (dna_string) => {
    let rna_string = ""
    let m = new Map(COMPLEMENTS);
    for (let c of dna_string) {
        c = m.get(c);
        if (c !== undefined) {
            rna_string += c;
        }
    }
    return rna_string
};
