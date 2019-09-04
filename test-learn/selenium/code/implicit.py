from selenium import webdriver

# 隐式等待
# 当使用了隐式等待执行测试的时候 如果webdriver 没有在DOM中找到元素 将继续等待
# 超出设定时间后则抛出找不到元素的异常
# 当查找元素或元素并没有立即出现的时候 隐式等待将等待一段时间再查找 DOM 默认时间是0
browser = webdriver.Chrome()
browser.implicitly_wait(10)#等待十秒加载不出来就会抛出异常，10秒内加载出来正常返回
browser.get('https://www.zhihu.com/explore')
input = browser.find_element_by_class_name('zu-top-add-question')
print(input)
browser.close()
