import sys
readline = sys.stdin.readline

from itertools import combinations

N, M, D = list(map(int, readline().split()))
lat = [[*map(int, readline().split())] for _ in range(N)]
lat.append([-1] * M)

is_in = lambda r, c: 0 <= r and r < N and 0 <= c and c < M

result = 0
def run(arcs):
    visit = [[False] * M for _ in range(N)]
    killed = 0
    for archer_r in reversed(range(1, N+1)):
        total_targets = set()
        for archer_c in arcs:
            targets = []
            for r_d in range(1, D+1):
                for c_d in range(-D + r_d, D - r_d +1):
                    tr = archer_r - r_d
                    tc = archer_c + c_d
                    abs_d = abs(r_d) + abs(c_d)
                    if is_in(tr, tc) and not visit[tr][tc] and lat[tr][tc] == 1:
                        targets.append((abs_d, tc, tr))
            targets.sort()
            # print('sorted', targets)
            if len(targets) > 0:
                total_targets.add((targets[0][1], targets[0][2]))
        # 타겟 삼은 애들 처리    
        # print('archer R', archer_r)                 
        # print('targets', total_targets)
        for t in total_targets:
            c, r = t
            visit[r][c] = True
        # print('now :', killed, '+', len(total_targets), '=', killed + len(total_targets))
        killed += len(total_targets)   
    return killed

archer_cases = list(combinations([*(range(0, M))], 3))
# print(archer_cases)
for arcs in archer_cases:
    # print('THIS TURN', arcs)
    result = max(result, run(arcs))

print(result)