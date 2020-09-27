const ALPHABET = new Set("abcdefghijklmnopqrstuvwxyz")


export const isPangram = (sentence) => {
    const sentence_letters = new Set(sentence.toLowerCase().replace(/[^a-z]/g, ''))
    return ALPHABET.size === sentence_letters.size
};
