const getTriangleInequality = (sides) => {
  let [a, b, c] = sides;
  return (a + b >= c) && (b + c >= a) && (c + a >= b);
}

const getDegenerate = (sides) => {
  let [a, b, c] = sides;
  return (a + b == c) || (b + c == a) || (c + a == b);
}

export class Triangle {
  constructor(...sides) {
    this.uniqueSideLengths = new Set(sides);
    this.isValidTriangle = sides.every(side => side != 0) && getTriangleInequality(sides);
    this._isDegenerate = getDegenerate(sides);
  }

  get isEquilateral() {
    return this.isValidTriangle && this.uniqueSideLengths.size == 1;
  }

  get isIsosceles() {
    return this.isValidTriangle && this.uniqueSideLengths.size <= 2;
  }

  get isScalene() {
    return this.isValidTriangle && this.uniqueSideLengths.size == 3;
  }

  get isDegenerate() {
    return this._isDegenerate;
  }
}
