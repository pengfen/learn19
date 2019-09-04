#!/usr/bin/python

class Account(object):
	"""
	账号类
	amount 是美元金额
	"""
	def __init__(self, rate):
		self.__amt = 0
		self.rate = rate

	@property
	def amount(self):
		""" 账号金额(美元) """
	    return self.__amt

	@property
	def cny(self):
		""" 账号余额(人民币) """
		return self.__amt * self.rate

	@amount.setter
	def amount(self, value):
		if value < 0:
			print("对不起 你的账号没金额")
			return
		self.__amt = value

if __name__ == '__main__':
	acc = Account(rate=6.6) # 汇率
	acc.amount = 20
	print("Dollar :", acc.amount)
	print("CNY :", acc.cny)
	acc.amount = -100
	print("Dollar :", acc.amount)