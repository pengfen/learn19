# 显式等待
# 指定一个等待条件，和一个最长等待时间，程序会判断在等待时间内条件是否满足，如果满足则返回，如果不满足会继续等待，超过时间就会抛出异常
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

browser = webdriver.Chrome()
browser.get('https://www.taobao.com/')
wait = WebDriverWait(browser, 10)
input = wait.until(EC.presence_of_element_located((By.ID, 'q')))
button = wait.until(EC.element_to_be_clickable((By.CSS_SELECTOR, '.btn-search')))
print(input, button)
browser.close()

# title_is 标题是某内容
# title_contains 标题包含某内容
# presence_of_element_located 元素加载出，传入定位元组，如(By.ID, 'p')
# visibility_of_element_located 元素可见，传入定位元组
# visibility_of 可见，传入元素对象
# presence_of_all_elements_located 所有元素加载出
# text_to_be_present_in_element 某个元素文本包含某文字
# text_to_be_present_in_element_value 某个元素值包含某文字
# frame_to_be_available_and_switch_to_it frame加载并切换
# invisibility_of_element_located 元素不可见
# element_to_be_clickable 元素可点击
# staleness_of 判断一个元素是否仍在DOM，可判断页面是否已经刷新
# element_to_be_selected 元素可选择，传元素对象
# element_located_to_be_selected 元素可选择，传入定位元组
# element_selection_state_to_be 传入元素对象以及状态，相等返回True，否则返回False
# element_located_selection_state_to_be 传入定位元组以及状态，相等返回True，否则返回False
# alert_is_present 是否出现Alert
