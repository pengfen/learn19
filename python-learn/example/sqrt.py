#!/usr/bin/python

num = float(input('请输入一个数字: '))
num_sqrt = num ** 0.5
print(' %0.3f 的平方根为 %0.3f' % (num , num_sqrt))

import cmath
sqrt = cmath.sqrt(num)
print(' {0} 的平方根为 {1}'.format(num , sqrt)) # 4 2+0j