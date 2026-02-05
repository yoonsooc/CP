import sys
readline = sys.stdin.readline

def solve():
    N = int(input())
    origin_range = [*map(int, readline().split())]
    lost_1_life_range = [[0, 1001]]

    for i in range(1, N):
        L, R = map(int, readline().split())
        if origin_range: # has elements
            oL, oR = origin_range
            nextL = max(oL - 1, L)
            nextR = min(oR + 1, R)
            
            updated_origin = [nextL, nextR] if nextL <= nextR else None
        else:
            updated_origin = None
        
        survivors = []
        for oL, oR in lost_1_life_range:
            nextL = max(oL - 1, L)
            nextR = min(oR + 1, R)
            if nextL <= nextR:
                survivors.append([nextL, nextR])

        if origin_range:
            survivors.append([origin_range[0] - 1, origin_range[1] + 1]) 
        
        if not survivors:
            print("Surrender")
            return
        
        survivors.sort()
        updated_1_lost = [survivors[0]]
        for i in range(1, len(survivors)):
            L, R = survivors[i]
            if updated_1_lost[-1][1] + 1 >= L:
                updated_1_lost[-1][1] = max(updated_1_lost[-1][1], R)
            else:
                updated_1_lost.append([L, R])

        origin_range, lost_1_life_range = updated_origin, updated_1_lost

    print("World Champion")
 
solve()   