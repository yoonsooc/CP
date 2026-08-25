const readline = require('readline')

const rl = readline.createInterface({
    input: process.stdin
})

const lines = []
const isIn = (r, c) => 0 <= r && r < n && 0 <= c && c < n
const rdir = [0, -1, -1, -1, 0, 1, 1, 1]
const cdir = [1, 1, 0, -1, -1, -1, 0, 1]

let A, trees, n, m, k, juices

rl.on('line', line => lines.push(line))

rl.on('close', () => {
    [n, m, k] = toNumberArr(lines[0])
    A = lines.slice(1, 1 + n).map(l => toNumberArr(l))
    juices = Array.from({ length: n }, () => Array.from({ length: n }, () => 5))

    trees = lines.slice(1 + n, 1 + n + m)
        .map(l => toNumberArr(l))
        .map(l => {
            const [x, y, age] = l
            return [x - 1, y - 1, age]
        })

    for (let i = 0; i < k; i++) {
        breed()
    }
    console.log(trees.length)
})

const breed = () => {
    // spring
    trees.sort((a, b) => a[2] - b[2])
    trees.forEach(tree => {
        const [r, c, age] = tree
        let juice = juices[r][c]

        if (age <= juice) {
            juices[r][c] -= age
            tree[2] += 1
        } else {
            tree[2] = -age
        }
    })
    // summer
    trees = trees.filter(tree => {
        const [r, c, age] = tree
        if (age < 0) {
            juices[r][c] += Math.floor((-age) / 2)
            return false
        }
        return true
    })

    // autumn
    trees.filter(tree => tree[2] % 5 === 0)
        .forEach(tree => {
            const [r, c] = tree
            for (let d = 0; d < 8; d++) {
                const [nr, nc] = [r + rdir[d], c + cdir[d]]
                if (isIn(nr, nc)) {
                    trees.push([r + rdir[d], c + cdir[d], 1])
                }
            }
        })

    // winter
    for (let r = 0; r < n; r++) {
        for (let c = 0; c < n; c++) {
            juices[r][c] += A[r][c]
        }
    }

}

const toNumberArr = l => l.trim().split(' ').map(e => Number(e))