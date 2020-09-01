const COMPLEMENTS = {C: 'G', G: 'C', A:'U', T: 'A'};

export const toRna = (dna_string) => {
    return dna_string.split('').map(i => COMPLEMENTS[i]).join('')
};
