N = int(input())
LIVES = 2
memo = [[LIVES] * 1001 for _ in range(N)]

LEFT_LIMIT = 1
RIGHT_LIMIT = 1000
isSurrendered = False
for i in range(0, N): # Turn 하나씩 순서대로 수행
    [L, R] = list(map(int, input().split()))
    for x, life in enumerate(memo[i]):      
        if x==0:
            memo[i][x] = 0
            continue      

        leftCanSurvive =  0
        rightCanSurvive =  0
        notMoveSurvive =  0

        # hasLeft
        if x-1 >= LEFT_LIMIT:
            leftCanSurvive = memo[i][x] if i==0 else memo[i-1][x-1]

        # hasRight
        if x+1 <= RIGHT_LIMIT:
            rightCanSurvive = memo[i][x] if i==0 else memo[i-1][x+1]
        
        # not move
        notMoveSurvive = memo[i][x] if i==0 else memo[i-1][x]

        realLife = max(leftCanSurvive, rightCanSurvive, notMoveSurvive)
        if realLife == 0:
            memo[i][x]= 0
            continue

        canEvade = (x >= L and x <= R)
        if not canEvade:
            memo[i][x] = realLife - 1
        else:
            memo[i][x] = realLife

    if sum(memo[i]) == 0:
        isSurrendered = True
        break
        
print("Surrender" if isSurrendered else "World Champion")