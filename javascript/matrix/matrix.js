const getColumnsFromRows = (rows) => {
  return rows[0].map((_, i) => rows.map((row) => row[i]))
}

export class Matrix {
  constructor(matrix) {
    this._rows = matrix.split("\n").map((row) => row.split(" ").map(Number));
    this._columns = getColumnsFromRows(this._rows);
  }

  get rows() {
    return this._rows;
  }

  get columns() {
    return this._columns;
  }
}
