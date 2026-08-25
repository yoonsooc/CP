from itertools import combinations


def solution(n, q, ans):
    answer = 0

    integers = [i+1 for i in range(n)]
    combs = combinations(integers, 5)
    
    for c in combs:
        n = 0
        for (idx, quiz) in enumerate(q):
            if len(set(quiz) & set(c)) == ans[idx]:
                n += 1
        if n == len(q):
            answer += 1  
    return answer