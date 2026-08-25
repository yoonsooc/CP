def solution(diffs, times, limit):
    upperLimit = (10**5)
    lowerLimit = 1
    answer = upperLimit

    
    while lowerLimit <= upperLimit:
        mid = (lowerLimit + upperLimit) // 2

        stack = 0 
        for idx, diff in enumerate(diffs):
            isDifficult = (diff - mid > 0)
            timeCur = times[idx]
            timePrev = 0 if idx == 0 else times[idx-1]
            
            if isDifficult:
                stack += ((diff - mid) * (timeCur + timePrev) + timeCur) 
            else:
                stack += timeCur

            if stack > limit:
                break

        if stack <= limit:
            answer = mid
            upperLimit = mid - 1
            
        else:
            lowerLimit = mid + 1
            
    return answer