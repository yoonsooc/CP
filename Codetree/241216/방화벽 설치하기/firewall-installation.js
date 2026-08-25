const readline = require('readline')

const rl = readline.createInterface({
    input: process.stdin
})

let n, m
const lines = []
const isIn = (r, c) => 0 <= r && r < n && 0 <= c && c < m
const rdir = [-1, 0, 1, 0]
const cdir = [0, -1, 0, 1]

let maps
let fireWalls = 0
const fires = []

rl.on('line', line => lines.push(line))


rl.on('close', () => {
    [n, m] = splitAsNumber(lines[0])

    const maps = lines.slice(1).map(l => splitAsNumber(l))
    const spaces = maps.flatMap((row, rIdx) => {
        return row.map((col, cIdx) => {
            if (col === 2) {
                fires.push([rIdx, cIdx])
            } else if (col === 1) {
                fireWalls += 1
            }

            return [col, rIdx, cIdx]
        })
            .filter(info => info[0] === 0)
            .map(info => ([info[1], info[2]]))
    })

    // 불이 퍼진 영역의 최솟값을 구하면 퍼지지 않은 영역의 최댓값을 알 수 있음
    const spacesCount = spaces.length
    let minBurnedSpaces = n * m - fireWalls - 3

    for (let i = 0; i < spacesCount - 2; i++) {
        for (let j = i + 1; j < spacesCount - 1; j++) {
            for (let k = j + 1; k < spacesCount; k++) {
                const copiedMap = maps.map(r => [...r])
                const [r1, c1] = spaces[i]
                const [r2, c2] = spaces[j]
                const [r3, c3] = spaces[k]

                copiedMap[r1][c1] = 1
                copiedMap[r2][c2] = 1
                copiedMap[r3][c3] = 1

                // console.log(r1, c1, '|', r2, c2, '|', r3, c3)

                const burnedSpaces = simulate(copiedMap)
                minBurnedSpaces = Math.min(minBurnedSpaces, burnedSpaces)
            }
        }
    }

    const additionalFireWalls = 3
    console.log(n * m - fireWalls - additionalFireWalls - minBurnedSpaces)
})

const simulate = (m, minBurnedSpaces) => {
    const queue = []
    fires.forEach(([r, c]) => queue.push([r, c]))
    let burned = fires.length

    while (queue.length > 0) {
        const [r, c] = queue.shift()

        for (let i = 0; i < 4; i++) {
            const nr = r + rdir[i]
            const nc = c + cdir[i]

            if (isIn(nr, nc) && m[nr][nc] === 0) {
                queue.push([nr, nc])
                m[nr][nc] = 2
                burned += 1

                if (burned > minBurnedSpaces) {
                    return minBurnedSpaces
                }
            }
        }
    }

    // console.log(m)
    // console.log('burned', burned)
    return burned
}

const splitAsNumber = l => l.split(' ').map(e => Number(e))