#!/usr/bin/python

# list定义
li = ["a", "b", "mpilgrim", "z", "example"]
print(li)
print(li[1])

# list负数索引
print(li[-1]) # example
print(li[-3]) # mpilgrim
print(li[1:3])
print(li[1:-1])
print(li[0:3])

# list增加元素
li.append("new")
print(li)

li.insert(2, "new")
print(li)

li.extend(["two", "elements"])
print(li)

# list搜索
print(li.index("example"))
#print(li.index("c"))

# list 删除元素
li.remove("a")
print(li)
li.remove("new")
print(li)
# print(li.pop()) # 删除list的最后一个元素 然后返回删除元素的值

# list运算符
li = ['a', 'b', 'mpilgrim']
li = li += ["example", "new"]