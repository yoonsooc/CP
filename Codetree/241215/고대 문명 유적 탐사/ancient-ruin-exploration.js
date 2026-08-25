const readline = require('readline')

const rl = readline.createInterface({
    input: process.stdin
})

const answer = []
const lines = []
let K
let M
let mIdx = 0 
let maps
let pieces

// 90, 180, 270
// 0~7 : N, NE, E, SE... , W, NW
const adjs = [ [-1, 0],[-1,1],[0,1],[1,1],[1,0],[1,-1],[0,-1],[-1,-1] ]
const moves = [ 
            [ [0,1], [1,1], [1,0], [1, -1], [0,-1], [-1,-1], [-1,0],[-1, 1] ],
            [ [1,0], [1, -1], [0,-1], [-1,-1], [-1,0],[-1, 1], [0,1], [1,1] ],
            [ [0,-1], [-1,-1], [-1,0],[-1, 1],[0,1], [1,1], [1,0], [1, -1] ]
]

const isIn = (r,c) => 0<=r && r<=4  && 0<=c && c<=4
const rd = [-1, 0, 1, 0]
const cd = [0, -1, 0, 1]

rl.on('line', line => lines.push(line))

rl.on('close', () => {
    [K, M] = lines[0].split(' ').map(e => Number(e))
    maps = lines.slice(1, 6).map(l => l.split(' ').map(e=>Number(e)))
    pieces = lines[lines.length-1].split(' ').map(e=>Number(e))

    for (let turn=0; turn < K; turn++) {
        let totalRewards = 0
        
        // console.log('before turn', maps) 

        const { reward, degree, i, j, resultMap, erasePoints} = search()
        if (reward < 0) {
            // console.log('no rewards')
            continue
        }

        totalRewards += reward
        maps = resultMap

        // console.log('after turn', maps,i, j, degree, reward)
        // console.log('erase points', erasePoints)
        erasePoints.forEach(([i, j]) => erase(i, j, maps))

        // console.log('after erase', maps)
        while (true) {
            fill(maps)

            // console.log('after fill', maps)


            const {rewards:f, erasePoints: ep} = countRewards(maps)
            if (f > 0) {
                totalRewards += f
                ep.forEach(([i,j]) => erase(i, j, maps))

                // console.log('after fill - erase', maps)

            } else {
                break
            }
        }

        if (totalRewards > 0) {
            answer.push(totalRewards)
        }

    
    }
    console.log(answer.join(' '))
})

const fill = (maps) => {
    for (let j=0; j<=4; j++) {
        for (let i=4; i>=0; i--) {
            if (maps[i][j] === 0) {
                maps[i][j] = pieces[mIdx]
                mIdx += 1
            }
        }
    }
}

const turnAndCount = (copiedMap, degreeIndex, i, j) => {
    turn(copiedMap, degreeIndex, i, j)
    return countRewards(copiedMap)
}

const turn = (copiedMap, degreeIndex, i, j) => {
    const m = moves[degreeIndex]
    adjs.forEach((p, index) => {
        const [ni, nj] = [i+p[0], j+p[1]]
        const [di, dj] = m[index]
        const ci = i+di
        const cj = j+dj
        copiedMap[ci][cj] = maps[ni][nj]
    })
}


const deepCopy = (m) => m.map(v => [...v])

const search = () => {
    let reward = 0
    let degree = 0
    let eraseTargets 

    let [ri, rj] = [5, 5]
    let resultMap

    for (let i=1; i<=3; i++) {
        for (let j=1; j<=3; j++) {
            const degs = [90, 180, 270]
            let copiedMap 
            const results = degs.map((deg, index) => {
                copiedMap = deepCopy(maps)
                const {rewards: count, erasePoints} = turnAndCount(copiedMap, index, i, j)
                return {
                    count,
                    erasePoints,
                    deg,
                    copiedMap
                }
            })

            results.sort((a, b) => b.count-a.count || (b.count===a.count && a.deg-b.deg))
            // console.log('sorted', i, j)
            // results.forEach(r => {
            //     console.log(r.deg, r.count)
            // })

            const { count: rewIJ, deg: degIJ, erasePoints, copiedMap: cmp } = results[0]

            if (reward < rewIJ) {
                reward = rewIJ
                degree = degIJ
                ri = i
                rj = j
                resultMap = cmp
                eraseTargets = erasePoints
            } else if (reward === rewIJ) {
                if (degree > degIJ) {
                    degree = degIJ
                    ri = i
                    rj = j
                    resultMap = cmp
                    eraseTargets = erasePoints
                }                
                else if (degree === degIJ && rj > j) {
                    degree = degIJ
                    ri = i
                    rj = j
                    resultMap = cmp
                    eraseTargets = erasePoints
                } 
                else if (degree === degIJ && rj === j && ri > i) {
                    degree = degIJ
                    ri = i
                    rj = j
                    resultMap = cmp
                    eraseTargets = erasePoints
                }   
            }
        }
    }

    return {
        reward: reward === 0 ? -1 : reward,
        degree,
        i: ri,
        j: rj,
        resultMap,
        erasePoints: eraseTargets
    }
}

const countRewards = (map) => {
    let rewards = 0
    const erasePoints = []

    let queue = []
    const vst = initVst(5, 5)

    for (let i=0; i<5; i++) {
        for (let j=0; j<5; j++) { 
            if (vst[i][j]) {
                continue
            }
            const pieceNum = map[i][j]
            
            queue.push([i,j])
            vst[i][j] = true
            let count = 1
            while (queue.length > 0) {
                const [ii, jj] = queue.shift()

                for (let d=0; d<=3; d++) {
                    const ni = ii + rd[d]
                    const nj = jj + cd[d]

                    if (isIn(ni, nj) && vst[ni][nj] === false && pieceNum === map[ni][nj]) {
                        queue.push([ni, nj])
                        vst[ni][nj] = true
                        count += 1
                    }
                }
            }
            if (count >= 3) {
                rewards += count
                erasePoints.push([i, j])
            }
        }
    }

    return {rewards, erasePoints}
}

const erase = (i, j, map) => {
    const pieceNum = map[i][j]
    
    const queue = []
    queue.push([i,j])
    
    while (queue.length > 0) {
        const [ii, jj] = queue.shift()

        for (let d=0; d<=3; d++) {
            const ni = ii + rd[d]
            const nj = jj + cd[d]

            if (isIn(ni, nj)  && pieceNum === map[ni][nj]) {
                queue.push([ni, nj])
                map[ni][nj] = 0
            }
        }
    }
}

const initVst = (r, c) => Array.from({length: r}, () => Array.from({length: c}, () => false))