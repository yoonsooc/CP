import sys
from collections import deque, defaultdict
input = sys.stdin.readline

def main():
    T = int(input())
    out = []

    for t in range(1, T + 1):
        N, K = map(int, input().split())
        A = list(map(int, input().split()))

        # 1단계: 각 키의 최소 인덱스 저장
        mins = defaultdict(lambda: 10**9)
        for i in range(N):
            if i < mins[A[i]]:
                mins[A[i]] = i

        # 2단계: 모노토닉 덱으로 각 키별 결과 계산
        results = {}
        q = deque()
        for h in sorted(A):
            # 범위 벗어난 키 제거 (h - K 미만)
            while q and q[0] < h - K:
                q.popleft()

            # 덱 뒤에서 현재보다 인덱스 큰 것 제거
            # (인덱스 큰 키는 앞으로 절대 최적해가 될 수 없음)
            while q and mins[q[-1]] > mins[h]:
                q.pop()

            q.append(h)
            results[h] = mins[q[0]]  # 범위 내 최소 인덱스

        # 3단계: 결과 합산
        ans = 0
        for i in range(N):
            ans += i - results[A[i]]

        out.append(f"Case #{t}\n{ans}")

    print('\n'.join(out))

main()
