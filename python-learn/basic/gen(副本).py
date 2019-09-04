#!/usr/bin/python

def my_generator():
	print("生成器")
	yield 'a'
	yield 'b'
	yield 'c'

for char in my_generator():
	print(char)

def counter_generator(low, high):
	while low <= high:
		yield low
		low += 1

for i in counter_generator(5, 10):
	print(i, end=' ')

def infinite_generator(start=0):
	while True:
		yield start
		start += 1

for num in infinite_generator(4):
	print(num, end=' ')
	if num > 20:
		break
