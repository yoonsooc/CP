import sys
readline = sys.stdin.readline

N, M = [*map(int, readline().split())]
field = [[*map(int, readline().split())] for _ in range(N)]
visited = [[False] * M for _ in range(N)]

islands_map = [[0] * M for _ in range(N)]
islands_nodes = {}
island_count = 0

isIn = lambda y, x: 0<=x and x<M and 0<=y and y<N
rdir = [-1, 0, 1, 0]
cdir = [0, -1, 0, 1]

def compose_island(r, c):
    visited[r][c] = True
    islands_nodes[island_count].append((r, c))
    islands_map[r][c] = island_count
    for d in range(4):
        nr = r + rdir[d]
        nc = c + cdir[d]
        if isIn(nr, nc) and not visited[nr][nc] and field[nr][nc] == 1:
            compose_island(nr, nc)

for r in range(N):
    for c in range(M):
        if not visited[r][c] and field[r][c] == 1:
            island_count += 1
            islands_nodes[island_count] = []
            compose_island(r, c)

# print(islands_map)
# print(islands_nodes)
distances = [[100] * (island_count+1) for _ in range(island_count+1)]

for start_island_num, nodes in islands_nodes.items():
    for node in nodes:
        r, c = node
        for dir in range(4):
            nr, nc = r+rdir[dir], c+cdir[dir] 
            if not isIn(nr, nc) or islands_map[nr][nc] != 0:
                continue
               
            dist = 0
            cr, cc = r, c
            while True:
                cr, cc = cr+rdir[dir], cc+cdir[dir]
                if not isIn(cr, cc):
                    break

                target = islands_map[cr][cc]
                if target == start_island_num:
                    break
                elif target == 0:
                    dist += 1
                else: 
                    if dist >= 2:
                            distances[start_island_num][target] = min(dist, distances[start_island_num][target])
                    break

# print(distances)

# Kruskal MST
parent = list(range(island_count + 1))

def find(x):
    if parent[x] != x:
        parent[x] = find(parent[x])
    return parent[x]

def union(a, b):
    pa, pb = find(a), find(b)
    if pa != pb:
        parent[pa] = pb
        return True
    return False

total = 0
count = 0
edges = []
for i in range(1, island_count + 1):
    for j in range(i+1, island_count + 1):
        if distances[i][j] < 100:
            edges.append((distances[i][j], i, j))
edges.sort()
for cost, a, b in edges:
    if union(a, b):
        total += cost
        count += 1

if count == island_count - 1:
    print(total)
else:
    print(-1)