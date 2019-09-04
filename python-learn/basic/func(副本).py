#!/usr/bin/python

def palindrome(s):
	return s == s[::-1]

if __name__ == '__main__':
	s = input("输入字符串:")
	if palindrome(s):
		print("是回文")
	else:
		print("不是回文")