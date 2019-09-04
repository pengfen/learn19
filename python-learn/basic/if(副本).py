#!/usr/bin/python

x = int(input("请输入一个整数:"))
if x < 0:
	x = 0
	print("负数")
elif x == 0:
	print("0")
elif x == 1:
	print("1")
else:
	print("正数")

# 真值检测
if x:
	print("pass")
if x == True:
	print("pass")