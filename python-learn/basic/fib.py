#!/usr/bin/python

a, b = 0, 1
while b < 100:
	print(b)
	a, b = b, a + b

a, b = 0, 1
while b < 100:
	print(b, end=' ') # end 替换换行符
	a, b = b, a + b
print() #会换行