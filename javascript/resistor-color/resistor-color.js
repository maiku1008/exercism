const resistors = new Map([
  ['black', 0],
  ['brown', 1],
  ['red', 2],
  ['orange', 3],
  ['yellow', 4],
  ['green', 5],
  ['blue', 6],
  ['violet', 7],
  ['grey', 8],
  ['white', 9],
]);

export const colorCode = (color) => resistors.get(color);

export const COLORS = [...resistors.keys()];
