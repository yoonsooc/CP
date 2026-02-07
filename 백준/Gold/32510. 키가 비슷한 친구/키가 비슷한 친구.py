import sys
input = sys.stdin.readline

MAX_H = 500001
INF = 10**9

def main():
    tree = [INF] * (2 * MAX_H)
    modified = []

    T = int(input())
    out = []

    for t in range(1, T + 1):
        # reset
        for idx in modified:
            tree[idx] = INF
        modified = []

        N, K = map(int, input().split())
        A = list(map(int, input().split()))

        ans = 0
        for i in range(N):
            h = A[i]
            lo = h - K if h >= K else 0

            # inline query
            res = INF
            l = lo + MAX_H
            r = h + MAX_H + 1
            while l < r:
                if l & 1:
                    if tree[l] < res:
                        res = tree[l]
                    l += 1
                if r & 1:
                    r -= 1
                    if tree[r] < res:
                        res = tree[r]
                l >>= 1
                r >>= 1

            if res <= i:
                ans += i - res

            # inline update
            p = h + MAX_H
            modified.append(p)
            if tree[p] > i:
                tree[p] = i
                while p > 1:
                    p >>= 1
                    modified.append(p)
                    left = tree[p << 1]
                    right = tree[(p << 1) | 1]
                    tree[p] = left if left < right else right

        out.append(f"Case #{t}\n{ans}")

    print('\n'.join(out))

main()
