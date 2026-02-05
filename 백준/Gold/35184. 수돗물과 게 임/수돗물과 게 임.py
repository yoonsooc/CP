import sys
import heapq

input = sys.stdin.readline

N, M, T = map(int, input().split())
lat = [input().strip() for _ in range(N)]

DX = (0, -1, 0, 1)
DY = (-1, 0, 1, 0)

# 게 찾기
start = None
for y in range(N):
    for x in range(M):
        if lat[y][x] in '0123':
            start = (x, y, int(lat[y][x]))
            break
    if start:
        break

sx, sy, sd = start
INF = 10**9

# 수돗물 결과 미리 계산
# tap_result[y][x][d] = (fx, fy, fd) or None
tap_result = [[[None] * 4 for _ in range(M)] for _ in range(N)]

for y in range(N):
    for x in range(M):
        if lat[y][x] == 'T':
            for d in range(4):
                fx, fy, fd = x, y, d
                valid = True
                visited = set()
                while lat[fy][fx] == 'T':
                    state = (fx, fy, fd)
                    if state in visited:  # 순환 감지!
                        valid = False
                        break
                    visited.add(state)
                    fd = (fd + 1) & 3
                    fx, fy = fx + DX[fd], fy + DY[fd]
                    if not (0 <= fx < M and 0 <= fy < N):
                        valid = False
                        break
                if valid:
                    tap_result[y][x][d] = (fx, fy, fd)

# 다익스트라
dist = [[[INF] * 4 for _ in range(M)] for _ in range(N)]
dist[sy][sx][sd] = 0

pq = [(0, sx, sy, sd)]

while pq:
    cost, cx, cy, cd = heapq.heappop(pq)

    if cost > dist[cy][cx][cd]:
        continue

    if lat[cy][cx] == 'S':
        print(cost)
        sys.exit()

    # 회전 (비용 T)
    nd = (cd + 1) & 3
    nc = cost + T
    if nc < dist[cy][cx][nd]:
        dist[cy][cx][nd] = nc
        heapq.heappush(pq, (nc, cx, cy, nd))

    # 게걸음 (비용 1)
    for side in (-1, 1):
        md = (cd + side) & 3
        nx, ny = cx + DX[md], cy + DY[md]

        if not (0 <= nx < M and 0 <= ny < N):
            continue

        if lat[ny][nx] == 'T':
            result = tap_result[ny][nx][cd]
            if result is None:
                continue
            fx, fy, fd = result
        else:
            fx, fy, fd = nx, ny, cd

        nc = cost + 1
        if nc < dist[fy][fx][fd]:
            dist[fy][fx][fd] = nc
            heapq.heappush(pq, (nc, fx, fy, fd))

print(-1)
