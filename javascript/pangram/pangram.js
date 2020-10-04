const ALPHABET = [..."abcdefghijklmnopqrstuvwxyz"]

export const isPangram = (sentence) => {
  return ALPHABET.every((character) => sentence.toLowerCase().includes(character))
}
