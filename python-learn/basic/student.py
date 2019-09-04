#!/usr/bin/python

n = int(input("请输入学生数:"))
data = {} # 用来存储数据的
Subjects = ('Physics', 'Maths', 'History') # 所有科目

for i in range(0, n):
	name = input("请学生名称 {}:".format(i + 1))
	marks = []
	for x in Subjects:
		marks.append(int(input("每一科分数 {}:".format(x))))
	data[name] = marks

for x, y in data.items():
	total = sum(y)
	print("{}'s total marks {}".format(x, total))
	if total < 120:
		print(x, "失败 :(")
	else:
		print(x, "成功 :)")
		