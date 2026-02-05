import sys
input = sys.stdin.readline

N, L = map(int, input().split())
commands = input().strip()


def simulate(start):
    """start 위치에서 시뮬레이션하여 (최종위치, 중간경로집합) 반환"""
    pos = start
    path = {start}

    for i in range(L - 1):
        cmd = commands[i]
        if cmd == 'L':
            if pos > 1:
                pos -= 1
        else:
            if pos < N:
                pos += 1
        path.add(pos)

    # 마지막 명령 실행
    if commands[-1] == 'L':
        if pos > 1:
            pos -= 1
    else:
        if pos < N:
            pos += 1

    return pos, path


def check(B, F, path):
    """B에서 시작해서 F에 도달하는 게 유효한지 확인"""
    return B != F and F not in path


def solve():
    # === Case 1: 벽에 한 번도 안 막히는 경우 ===
    # 상대 위치 계산 (시작=0 기준, 벽 무시)
    rel = [0]
    pos = 0
    min_rel, max_rel = 0, 0

    for cmd in commands:
        if cmd == 'L':
            pos -= 1
        else:
            pos += 1
        rel.append(pos)
        if pos < min_rel:
            min_rel = pos
        if pos > max_rel:
            max_rel = pos

    # 벽에 안 막히는 B의 범위
    # B + min_rel >= 1 and B + max_rel <= N
    B_low = 1 - min_rel
    B_high = N - max_rel

    if B_low <= B_high:
        # 중간 경로 상대 위치 집합 vs 최종 상대 위치
        mid_set = set(rel[:-1])
        final_rel = rel[-1]

        # final_rel이 중간에 없고, final_rel != 0 (B != F)
        if final_rel not in mid_set and final_rel != 0:
            B = B_low
            F = B + final_rel
            print("WIN")
            print(B, F)
            return

    # === Case 2: 벽에 막히는 경우 ===
    # B=1 (왼쪽 벽에 막힐 수 있음)
    F, path = simulate(1)
    if check(1, F, path):
        print("WIN")
        print(1, F)
        return

    # B=N (오른쪽 벽에 막힐 수 있음)
    F, path = simulate(N)
    if check(N, F, path):
        print("WIN")
        print(N, F)
        return

    print("DEFEAT")


solve()
