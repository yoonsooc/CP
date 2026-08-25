class Solution {
      fun solution(n: Int, lighthouse: Array<IntArray>): Int {
        val links = Array(n + 1) { mutableListOf<Int>() }

        lighthouse.forEach { (f, t) ->
            links[f].add(t)
            links[t].add(f)
        }

        val (on, off) = search(1, 0, links)
        return on.coerceAtMost(off)

    }

    private fun search(node: Int, parent: Int, links: Array<MutableList<Int>>): MutableList<Int> {
        val answer = mutableListOf(1, 0)

        if (links[node].size == 1 && links[node][0] == parent) {
            return answer
        }

        for (next in links[node]) {
            if (next == parent) {
                continue
            }

            val (on, off) = search(next, node, links)
            answer[0] += off.coerceAtMost(on)
            answer[1] += on
        }

        return answer

    }
}

