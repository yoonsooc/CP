const readline = require('readline')

const rl = readline.createInterface({
    input: process.stdin
})

const lines = []
let maps
let mapsClockwise90
let n, L

let readCount = -1
rl.on('line', line => {
    if (readCount < 0) {
        [n, L] = line.trim().split(' ').map(e => parseInt(e))

        maps = init2DimArr(n)
        mapsClockwise90 = init2DimArr(n)
        readCount += 1

    } else {
        lines.push(line.trim())
    }
})

rl.on('close', () => {
    let answer = 0

    lines.forEach((line, lineIndex) => {
        const row = line.split(' ').map(e => Number(e))
        row.forEach((it, index) => {
            maps[lineIndex][index] = it
            mapsClockwise90[index][n - 1 - lineIndex] = it
        })
    })

    maps.forEach(row => answer += isEnable(row) ? 1 : 0)
    mapsClockwise90.forEach(row => answer += isEnable(row) ? 1 : 0)

    console.log(answer)
})

const init2DimArr = n => Array.from({ length: n }, () => Array.from({ length: n }, () => 0))

const isEnable = (row) => {
    const charAppender = [row[0]]
    for (let i = 1; i < n; i++) {
        if (Math.abs(row[i] - row[i - 1] > 1)) {
            return false
        }

        if (row[i] !== row[i - 1]) {
            charAppender.push('|')
        }
        charAppender.push(row[i])
    }

    const blocks = charAppender.join('').split('|')
    let usedBlockLen = 0

    for (let i = 0; i < blocks.length - 1; i++) {
        const pb = blocks[i]
        const nb = blocks[i + 1]

        const nbLen = nb.length
        const pbLen = pb.length

        if (Number(pb[0]) > Number(nb[0])) {
            if (nbLen >= L) {
                usedBlockLen = L
            } else {
                return false
            }
        } else {
            if (usedBlockLen + L <= pbLen) {
                usedBlockLen = 0
            } else {
                return false
            }
        }
    }
    return true
}