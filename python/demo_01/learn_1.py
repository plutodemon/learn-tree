import random

a = b = 1
c, d = 2, 3
e = c / 4  # float
f = c // 4  # int
s = "abcd"
print(s[1:-1])  # 左闭右开
print(r"abc\a")  # 原始字符串
h = [1, 2, 3, 4]
h[1:3] = [5, 6]
h.count(1)
k = (1, h,)  # 元素不可变

s = {1, 2, 3, 4}  # 集合无序 不重复 可变
t = {1, 2, 5, 6}
p = set()  # 空集合
print(s & t)  # 交集
print(s | t)  # 并集
print(s - t)  # 差集
print(s ^ t)  # 对称差集(copilot提示的对称插件？)

di = {"a": 1, "b": 2}
空字典 = {}
推导式 = {x: x ** 2 for x in (2, 4, 6)}

bh = b"haha"  # bytes

#  ! /usr/bin/env python3  修改权限后可以直接运行

# 传统写法
n = 10
if n > 5:
    print(n)
# 使用海象运算符
if (n := 10) > 5:
    print(n)

print(random.choice(range(10)))  # 随机选择

print("\a")  # 响铃
print("\000")  # 空字符
print("\v")  # 垂直制表符
print("\t")  # 水平制表符
