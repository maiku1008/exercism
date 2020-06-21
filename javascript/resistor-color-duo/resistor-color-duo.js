const resistors = new Map([
    ['black', '0'],
    ['brown', '1'],
    ['red', '2'],
    ['orange', '3'],
    ['yellow', '4'],
    ['green', '5'],
    ['blue', '6'],
    ['violet', '7'],
    ['grey', '8'],
    ['white', '9'],
]);

export const decodedValue = (colors) => {
    let result = "";
    for(var i in colors.slice(0, 2)){
        result += resistors.get(colors[i]);
    }
    return parseInt(result);
}
