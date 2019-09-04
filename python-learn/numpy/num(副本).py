#!/usr/bin/python

import numpy as np

print(np.__version__)

print(np.array([1, 2, 3])) # 通过列表创建一维数组

print(np.array([(1, 2, 3), (4, 5, 6)])) # 通过列表创建二维数组

print(np.zeros((3, 3))) # 创建全为0的二维数组

print(np.ones((2, 3, 4))) # 创建全为1的三维数组

print(np.arange(5)) # 创建一维等差数组

print(np.arange(6).reshape(2, 3)) # 创建二维等差数组
