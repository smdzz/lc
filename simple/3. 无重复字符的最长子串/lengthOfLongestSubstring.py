# -*- coding: utf-8 -*-
# @File  : lengthOfLongestSubstring.py
# @Author: mengde
# @Date  : 2021/11/11

def lengthOfLongestSubstring(s):
    occ = dict()
    length = len(s)
    cur, max_len = -1, 0
    for i in range(length):
        if i != 0:
            occ.pop(s[i - 1])
        while cur + 1 < length and s[cur + 1] not in occ:
            occ[s[cur + 1]] = cur + 1
            cur += 1
        max_len = max(max_len, cur + 1 - i)
    return max_len

def lengthOfLongestSubstring(s):
    occ = set()
    n = len(s)
    # 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
    rk, ans = -1, 0
    for i in range(n):
        if i != 0:
            # 左指针向右移动一格，移除一个字符
            occ.remove(s[i - 1])
        while rk + 1 < n and s[rk + 1] not in occ:
            # 不断地移动右指针
            occ.add(s[rk + 1])
            rk += 1
        # 第 i 到 rk 个字符是一个极长的无重复字符子串
        ans = max(ans, rk - i + 1)
    return ans





# print(lengthOfLongestSubstring("pwwkew"))
print(lengthOfLongestSubstring("pwwkew"))