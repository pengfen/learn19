#!/bin/bash

# 复制目录(排除node_modules)
cp -rp `ls | grep -v github.com | grep -v golang.org |  xargs` /media/ricky/项目/learn/go-learn/
