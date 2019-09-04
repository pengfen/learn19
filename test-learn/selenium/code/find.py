from selenium import webdriver

# 查找元素
browser = webdriver.Chrome()
browser.get("https://www.baidu.com")

# 单个元素
input_first = browser.find_element_by_id('kw')
input_second = browser.find_element_by_css_selector('#kw')
input_third = browser.find_element_by_xpath('//*[@id="kw"]')
print(input_first, input_second, input_third)
browser.close()

# 常用的查找方法
# find_element_by_name
# find_element_by_xpath
# find_element_by_link_text
# find_element_by_partial_link_text
# find_element_by_tag_name
# find_element_by_class_name
# find_element_by_css_selector

# 通用方法 find_element(BY.ID, 'kw') 第一个参数传入名称 第二个传入具体的参数
# 多个元素 elements多个s   input_first = browser.find_elements_by_id('kw')
