#!/usr/bin/python

s = input("请输入一个字符串: ")
z = s[::-1]
if s == z:
	print("是回文")
else:
	print("不是回文")