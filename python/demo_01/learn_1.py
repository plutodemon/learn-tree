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

# 条件判断
if c < b:
    print("1 < 2")
elif a < b:
    print("1 == 2")
# switch case
match c:
    case 1:
        print("1")
    case _:
        print("default")
# 循环
while c < 10:
    c += 1
    print(c)
else:
    print("end")

for i in range(10):
    pass
    print(i)
else:
    print("end")

# 列表推导式
names = ['Bob', 'Tom', 'alice', 'Jerry', 'Wendy', 'Smith']
new_names = [name.upper() for name in names if len(name) > 3]
print(new_names)

# 字典推导式
listDemo = ['Google', 'Runoob', 'Taobao']
newDict = {key: len(key) for key in listDemo}
dic = {x: x ** 2 for x in (2, 4, 6)}

# 元组推导式 生成器表达式
a = (x for x in range(1, 10))

# 迭代器
listD = [1, 2, 3, 4]
it = iter(listD)  # 创建迭代器对象

for x in it:
    print(x, end=" ")
# 或者
# while True:
#     try:
#         print(next(it))
#     except StopIteration:
#         sys.exit()

# 生成器  https://www.runoob.com/python3/python3-iterator-generator.html

匿名函数 = lambda: "Hello, world!"
print(匿名函数())  # 输出: Hello, world!
