import sys
readline = sys.stdin.readline
from collections import deque


N = int(input())
populs = [0] + [*map(int, readline().split())]
links = [[0]] + [ [*map(int, readline().split())] for _ in range(N) ]

answer = 1000

def search(group):
    if not group:
        return 0
    visit = [ False ] * (N+1)
    q = deque([group[0]])
    visit[group[0]] = True
    total = populs[group[0]]

    while q:
        cur = q.popleft()
        if links[cur][0] == 0:
            continue
        for l in links[cur][1:]:
            if l in group and not visit[l]:
                visit[l] = True
                q.append(l)
                total += populs[l]

    return total if (sum(visit[i] for i in group) == len(group)) else -1

def backtrack(start, group):
    global answer       

    if len(group) > 0:
        another_group = [i for i in range(1, N+1) if i not in group]
        
        sum_group = search(group)
        sum_another_group = search(another_group)

        if sum_group != -1 and sum_another_group != -1:
            answer = min(answer, abs(sum_group - sum_another_group))
        
    for i in range(start, N+1):
        if i not in group:
            backtrack(i+1, group + [i])

backtrack(1, [])
print(-1 if answer == 1000 else answer)