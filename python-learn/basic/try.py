#!/usr/bin/python

def get_number():
	"返回浮点数"
	number = float(input("请输入一个浮点数: "))
	return number

while True:
	try:
		print(get_number())
	except ValueError:
		print*("你输入了一个错误的值.")