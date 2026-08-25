from collections import deque

rdir = [-1, 0, 1, 0]
cdir = [0, -1, 0, 1]
hidden = {}
revealed = {}

def isInShell(r, c, N, M):
    return 1 <= r and r < N-1 and 1 <= c and c < M-1

def isIn(r, c, N, M):
    return 0 <= r and r < N and 0 <= c and c < M

def putOut(storage, char, N, M):
    if char not in revealed:
        return  
        
    targets = revealed[char]
    q = deque()
    for r, c in targets:
        q.append((r, c))
        storage[r][c] = '-'
    del revealed[char]

    while len(q) > 0:
        r, c = q.popleft()
        for d in range(4):
            nr, nc = r+rdir[d], c+cdir[d]
            if not isInShell(nr, nc, N, M):
                continue

            nextChar = storage[nr][nc]
            if nextChar == '+':
                q.append((nr, nc))
                storage[nr][nc] = '-'

            elif nextChar != '-':
                if nextChar not in revealed:
                    revealed[nextChar] = set()
                revealed[nextChar].add((nr, nc))
                if nextChar in hidden and (nr, nc) in hidden[nextChar]:
                    hidden[nextChar].remove((nr, nc))

# char -> '+'
def putOutHidden(storage, char):
    if char not in hidden:
        return  

    targets = hidden[char]
    for r, c in targets:
        storage[r][c] = '+'

    del hidden[char]

def solution(storage, requests):
    answer = 0
    N = len(storage)
    M = len(storage[0])

    for r, line in enumerate(storage):
        for c, char in enumerate(line):
            if (r == 0 or r == N-1) or (c == 0 or c == M-1):
                if char not in revealed:
                    revealed[char] = set()
                revealed[char].add((r, c))
            else:
                if char not in hidden:
                    hidden[char] = set()
                hidden[char].add((r, c))
        storage[r] = list(storage[r])

    for turn, req in enumerate(requests):
        if len(req) == 2:
            putOutHidden(storage, req[0])
            putOut(storage, req[0], N, M)
            
        else:
            putOut(storage, req[0], N, M)
        
    hiddenBlocks = sum(len(value) for value in hidden.values())
    revBlocks = sum(len(value) for value in revealed.values()) 
    answer += (hiddenBlocks + revBlocks)
    return answer
