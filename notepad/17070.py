import sys
readline = sys.stdin.readline

N = int(input())

lat = [[*map(int, readline().split())] for _ in range(N)]
 # 가로 대각선 세로
memo = [[[0, 0, 0] for _ in range(N)] for _ in range(N)]
memo[0][1] = [1, 0, 0]
# print(memo)

for i in range(0, N):
    for j in range(2, N):
        if lat[i][j] == 1:
            continue
        
        memo[i][j][0] += (memo[i][j-1][0] + memo[i][j-1][1])

        if i >= 1 and lat[i-1][j] != 1 and lat[i][j-1] != 1:
            memo[i][j][1] += (memo[i-1][j-1][0] + memo[i-1][j-1][1] + memo[i-1][j-1][2])

        if i >= 1:
            memo[i][j][2] += (memo[i-1][j][1] + memo[i-1][j][2])    

print(memo[N-1][N-1][0] + memo[N-1][N-1][1] + memo[N-1][N-1][2])