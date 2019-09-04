#!/usr/bin/python

import datetime

# 获取昨天日期
def getYesterday():
	today = datetime.date.today();
	print(today) #2018-06-26
	oneday = datetime.timedelta(days=1)
	print(oneday) #1 day, 0:00:00
	yesterday = today - oneday
	print(yesterday)

getYesterday()

# 计算每个月天数
import calendar
monthRange = calendar.monthrange(2016, 9)
print(monthRange) # 输出的是元组
print(monthRange[0]) # 当月第一天星期几
print(monthRange[1]) # 当前共有多少天

# 输入指定年月
y = int(input("请输入年份: "))
m = int(input("请输入月份: "))

print(calendar.month(y, m))