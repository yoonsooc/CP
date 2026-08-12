import sys


readline = sys.stdin.readline

lat = [[*map(int, readline().split())] for _ in range(10)]
visited = [[False] * 10 for _ in range(10)]
pre_sum = [0] * 100
answer = 100


def is_in(r, c):
    return 0 <= r and r <= 9 and 0 <= c and c <= 9


def toggle(r, c, s, bool_val):
    for i in range(r, r + s):
        for j in range(c, c + s):
            visited[i][j] = bool_val


def get_sum(r1, c1, r2, c2):
    result = pre_sum[r2 * 10 + c2]
    if r1 > 0:
        result -= pre_sum[(r1 - 1) * 10 + c2]
    if c1 > 0:
        result -= pre_sum[r2 * 10 + (c1 - 1)]
    if r1 > 0 and c1 > 0:
        result += pre_sum[(r1 - 1) * 10 + (c1 - 1)]
    return result


def can_cover(r, c, s):
    for i in range(r, r + s):
        for j in range(c, c + s):
            if visited[i][j]:
                return False
    return True


def backtrack(remain_papers, cr, cc, used):
    global answer
    if used >= answer:
        return

    for s in reversed(range(1, 6)):
        sr, sc = cr + s - 1, cc + s - 1
        if (
            is_in(sr, sc)
            and can_cover(cr, cc, s)
            and get_sum(cr, cc, sr, sc) == s * s
            and remain_papers[s] > 0
        ):
            remain_papers[s] -= 1
            toggle(cr, cc, s, True)

            nr, nc = -1, -1
            found = False
            for i in range(cr, 10):
                start_j = cc if i == cr else 0
                for j in range(start_j, 10):
                    if lat[i][j] == 1 and not visited[i][j]:
                        nr, nc = i, j
                        backtrack(remain_papers, nr, nc, used + 1)
                        found = True
                        break
                if found:
                    break

            if nr == -1:
                answer = min(answer, used + 1)

            toggle(cr, cc, s, False)
            remain_papers[s] += 1


# 누적합 구해놓기
for r in range(10):
    for c in range(10):
        pos = r * 10 + c
        up = (r - 1) * 10 + c
        left = r * 10 + (c - 1)
        diag = (r - 1) * 10 + (c - 1)
        if r > 0:
            pre_sum[pos] += pre_sum[up]
        if c > 0:
            pre_sum[pos] += pre_sum[left]
        if r > 0 and c > 0:
            pre_sum[pos] -= pre_sum[diag]
        pre_sum[pos] += lat[r][c]

found = False
for i in range(10):
    for j in range(10):
        if lat[i][j] == 1:
            backtrack([0, 5, 5, 5, 5, 5], i, j, 0)
            found = True
            break
    if found:
        break

if not found:
    print(0)
else:
    print(answer if answer <= 25 else -1)
