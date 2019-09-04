class Counter(object):
	"""
	迭代器
	"""
	def __init__(self, low, high):
		self.current = low
		self.high = high

	def __iter__(self): # 返回迭代器对象自身 用在for和in语句中
		return self

	def __next__(self): # 返回迭代器的下一个值 如果没有下一个值返回 抛出 StopIteration异常
		if self.current > self.high:
			raise StopIteration
		else:
			self.current += 1
			return self.current - 1

"""
c = Counter(5, 10)
for i in c:
	print(i, end='')
"""			