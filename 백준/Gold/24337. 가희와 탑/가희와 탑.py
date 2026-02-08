import sys
readline = sys.stdin.readline

N, a, b = [*map(int, readline().split())]

arr = []
rest = 0
if a >= b:
    l_arr = list(range(1, a+1))
    r_arr = list(reversed(range(1, b)))
else:
    l_arr = list(range(1, a))
    r_arr = list(reversed(range(1, b+1)))

rest = [1] * (N - (a+b-1))
if a == 1:
    comb = [b] + rest + list(range(b-1, 0, -1))
else:
    comb = rest + l_arr + r_arr

# print(arr)
if len(comb) > N:
    print(-1)
else:
    print(' '.join(map(str, comb)))