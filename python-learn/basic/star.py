#!/usr/bin/python

row = int(input("请输入行号:"))
n = row
while n >= 0:
	x = "*" * n
	print(x)
	n -= 1

i = 1;
while i <= n:
	print("*" * i)
	i += 1

while n >= 0:
	x = "*" * n
	y = " " * (row - n)
	print(y + x)
	n -= 1

