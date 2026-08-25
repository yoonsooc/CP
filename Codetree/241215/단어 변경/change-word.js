const readline = require('readline')

const rl = readline.createInterface({
    input: process.stdin
})

const lines = []

rl.on('line', line => lines.push(line))

// a~z 

rl.on('close', () => {
    const [n, m] = lines[0].split(' ').map(e => Number(e))
    const [origin, goal] = lines.slice(1).map(str => str.split(''))
    origin.unshift(' ')
    goal.unshift(' ')

    const dp = Array.from({ length: n + 1 }, () => Array.from({ length: m + 1 }, () => 0))

    for (let j = 0; j <= m; j++) {
        dp[0][j] = j
    }
    for (let i = 0; i <= n; i++) {
        dp[i][0] = i
    }

    const ijl = ['i', 'j', 'l']
    const vw = ['v', 'w']

    for (let i = 1; i <= n; i++) { // 세로: origin
        for (let j = 1; j <= m; j++) {     // 가로: goal
            const goalChar = goal[j]
            const originChar = origin[i]

            if (goalChar === originChar
                || (originChar === 'i' && ijl.includes(goalChar))
                || (originChar === 'v' && vw.includes(goalChar))) {
                dp[i][j] = dp[i - 1][j - 1]
            } else {
                dp[i][j] = Math.min(dp[i - 1][j] + 1, dp[i][j - 1] + 1, dp[i - 1][j - 1] + 1)
            }
        }
    }

    console.log(dp[n][m])

})

