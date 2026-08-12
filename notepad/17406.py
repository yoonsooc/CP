import sys

readline = sys.stdin.readline


N, M, K = [*map(int, readline().split())]
A = [[*map(int, readline().split())] for _ in range(N)]
rcs = [[*map(int, readline().split())] for _ in range(K)]

answer = 100 * 50
visited = [False] * K


def turn(ri, ci, s):
    """시계 방향 회전"""
    for d in range(s, 0, -1):
        top, bottom = ri - d, ri + d
        left, right = ci - d, ci + d

        temp = A[top][right]
        # 위쪽 행: 오른쪽으로
        for j in range(right, left, -1):
            A[top][j] = A[top][j - 1]
        # 왼쪽 열: 위로
        for i in range(top, bottom):
            A[i][left] = A[i + 1][left]
        # 아래쪽 행: 왼쪽으로
        for j in range(left, right):
            A[bottom][j] = A[bottom][j + 1]
        # 오른쪽 열: 아래로
        for i in range(bottom, top + 1, -1):
            A[i][right] = A[i - 1][right]
        A[top + 1][right] = temp


def rollback(ri, ci, s):
    """반시계 방향 회전 (turn의 역연산)"""
    for d in range(1, s + 1):
        top, bottom = ri - d, ri + d
        left, right = ci - d, ci + d

        temp = A[top + 1][right]
        # 오른쪽 열: 위로
        for i in range(top + 1, bottom):
            A[i][right] = A[i + 1][right]
        # 아래쪽 행: 오른쪽으로
        for j in range(right, left, -1):
            A[bottom][j] = A[bottom][j - 1]
        # 왼쪽 열: 아래로
        for i in range(bottom, top, -1):
            A[i][left] = A[i - 1][left]
        # 위쪽 행: 왼쪽으로
        for j in range(left, right):
            A[top][j] = A[top][j + 1]
        A[top][right] = temp


def backtrack(k: int, t: int):
    global answer
    r, c, s = rcs[k]
    ri = r - 1
    ci = c - 1
    turn(ri, ci, s)
    if t == K:
        answer = min(answer, min(map(sum, A)))
    else:
        for nk in range(K):
            if not visited[nk]:
                visited[nk] = True
                backtrack(nk, t + 1)
                visited[nk] = False

    rollback(ri, ci, s)


for k in range(K):
    visited[k] = True
    backtrack(k, 1)
    visited[k] = False
print(answer)
