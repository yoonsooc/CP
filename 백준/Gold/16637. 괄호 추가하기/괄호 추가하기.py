import sys
readline = sys.stdin.readline

N = int(input())
equation = readline().strip()

# print(equation)
num_count = len(equation) // 2 + 1
max_res_mem = [0] * num_count
min_res_mem = [0] * num_count

def oper(a, op, b):
    if op == '+':
        return a+b
    elif op == '-':
        return a-b
    else:
        return a*b

op = None

for i_idx in range(0, num_count * 2 - 1, 2):
    c_num = int(equation[i_idx])

    i_num = i_idx // 2

    if i_num == 0:
        max_res_mem[i_num] = c_num
        min_res_mem[i_num] = c_num
        continue

    if i_num == 1:
        max_res_mem[i_num] = oper(int(equation[0]), equation[1], c_num)
        min_res_mem[i_num] = oper(int(equation[0]), equation[1], c_num)
        continue
    
    not_include_cur_with_max = oper(max_res_mem[i_num-1], equation[i_idx - 1], c_num)
    not_include_cur_with_min = oper(min_res_mem[i_num-1], equation[i_idx - 1], c_num)
    
    include_cur_with_max = oper(max_res_mem[i_num-2], equation[i_idx - 3], oper(int(equation[i_idx-2]), equation[i_idx - 1], c_num))
    include_cur_with_min = oper(min_res_mem[i_num-2], equation[i_idx - 3], oper(int(equation[i_idx-2]), equation[i_idx - 1], c_num))
    
    max_res_mem[i_num] = max(not_include_cur_with_max, not_include_cur_with_min, include_cur_with_max, include_cur_with_min)
    min_res_mem[i_num] = min(not_include_cur_with_max, not_include_cur_with_min, include_cur_with_max, include_cur_with_min)

print(max_res_mem[num_count-1])