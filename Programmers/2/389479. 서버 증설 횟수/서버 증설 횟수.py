def solution(players, m, k):
    currentServerCount = 0
    serverShutdownTimeMap = {}
    totalAddCount = 0

    for startTime, player in enumerate(players):
        minimumServerCount = player // m
        
        if startTime in serverShutdownTimeMap:
            shutdownCount = serverShutdownTimeMap[startTime]
            del(serverShutdownTimeMap[startTime])
            currentServerCount -= shutdownCount
        
        if minimumServerCount > currentServerCount:
            addCount = minimumServerCount - currentServerCount
            serverShutdownTimeMap[startTime + k] = addCount
            currentServerCount += addCount
            totalAddCount += addCount

       # print(startTime, player, minimumServerCount, currentServerCount, totalAddCount)

    return totalAddCount