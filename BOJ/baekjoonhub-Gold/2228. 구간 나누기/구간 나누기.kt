package pg

fun main(args: Array<String>) {
    val (n, m) = readNumbers()
    val memo = Array(n) { IntArray(m + 1) { -32768 * 100 } }
    val partSums = IntArray(n) { 0 }

    for (i in 0 until n) {
        val num = readNumber()
        partSums[i] = if (i == 0) num else (partSums[i - 1] + num)

        if (i == 0) {
            memo[i][1] = num
            continue
        }

        memo[i][1] = memo[i - 1][1]
        for (j in 0..<i) {
            memo[i][1] = memo[i][1].coerceAtLeast(partSums[i] - partSums[j])
        }
        memo[i][1] = memo[i][1].coerceAtLeast(partSums[i])

        if (i < 2) {
            continue
        }

        for (r in 2..m) {
            memo[i][r] = memo[i - 1][r]

            for (k in 0..i - 2) {
                memo[i][r] = memo[i][r].coerceAtLeast(memo[k][r-1] + partSums[i] - partSums[k+1])
            }
        }
    }

    println(memo[n - 1][m])
}

fun readNumbers(): IntArray {
    return readln().split(" ").map { it.toInt() }.toIntArray()
}

fun readNumber(): Int {
    return readln().toInt()
}