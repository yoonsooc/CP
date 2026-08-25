import sys
input = sys.stdin.readline

def solve():
    N, M = map(int, input().split())
    S = input().strip()

    MOD = 998244353

    # M >= 25*N이면 모든 26^N 문자열 도달 가능
    if M >= 25 * N:
        print(pow(26, N, MOD))
        return

    # 팩토리얼, 역팩토리얼 전처리
    fact = [1] * (N + 1)
    for i in range(1, N + 1):
        fact[i] = fact[i - 1] * i % MOD

    inv_fact = [1] * (N + 1)
    inv_fact[N] = pow(fact[N], MOD - 2, MOD)
    for i in range(N - 1, -1, -1):
        inv_fact[i] = inv_fact[i + 1] * (i + 1) % MOD

    inv_N_fact = inv_fact[N]

    # 포함-배제:
    # sum_{j=0}^{min(N, M//26)} (-1)^j * C(N,j) * C(M-26j+N, N)
    ans = 0
    for j in range(min(N, M // 26) + 1):
        # C(N, j) — 둘 다 작은 수
        cnj = fact[N] * inv_fact[j] % MOD * inv_fact[N - j] % MOD

        # C(M - 26j + N, N) — 위가 클 수 있으므로 직접 곱셈
        m = M - 26 * j + N
        prod = 1
        for i in range(N):
            prod = prod * ((m - i) % MOD) % MOD
        c_large = prod * inv_N_fact % MOD

        term = cnj * c_large % MOD
        if j % 2 == 0:
            ans = (ans + term) % MOD
        else:
            ans = (ans - term + MOD) % MOD

    print(ans)

solve()
