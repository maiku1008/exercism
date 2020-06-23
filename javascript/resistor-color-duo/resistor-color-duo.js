export const COLORS = ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white'];

export const decodedValue = (colors) => {
    let a = []
    for (var i in colors.slice(0, 2)) {
        a.push(COLORS.indexOf(colors[i]))
    }
    if (a.length > 1) {
        a[0] *= 10
    }
    return a.reduce((total, amount) => total + amount)
}
