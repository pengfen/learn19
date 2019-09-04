from selenium import webdriver
from selenium.webdriver.common.keys import Keys
import random
import time

#driver = webdriver.Chrome(r"chromedriver.exe")
driver = webdriver.Chrome()
driver.get("https://abcc.com/signin")

emailInput = driver.find_element_by_xpath("//*[@name='auth_key']")
emailInput.send_keys('jiagate001@sina.com')
emailInput = driver.find_element_by_xpath("//*[@name='password']")
emailInput.send_keys('Huweizhi')
loginBTN = driver.find_element_by_xpath("//*[@class='semi-bold primary']")
loginBTN.click()

input('输入验证码登陆后按回车继续')
driver.get("https://abcc.com/markets/ethbtc")
assert "ETH/BTC" in driver.title
count = 0
while True:
    try:
        ask = driver.find_element_by_xpath("//*[@class='scrollStyle ask-table']") 
        ask1 = ask.text.split('\n')[-1].split(' ')[0]
        bid = driver.find_element_by_xpath("//*[@class='scrollStyle bid-table']") 
        bid1 = bid.text.split('\n')[0].split(' ')[0]
        ask1,bid1 = float(ask1),float(bid1)
        # price = bid1 + 0.000001 + (ask1 - bid1 - 0.000002) * (random.random())
        if ask1 - bid1 > 0.000006:
            price = (ask1 + bid1) / 2
            amount = 5 + random.random()*5
            market_price = "%.6f"%price

            print("%.6f"%price,"%.4f"%amount)

            buyPriceInput = driver.find_element_by_xpath("//*[@class='order-submit order-form']/div[1]/div[1]/input")
            buyPriceInput.clear()
            #buyPriceInput.send_keys("%.6f"%price)
            buyPriceInput.send_keys(market_price)

            buyAmountInput =  driver.find_element_by_xpath("//*[@class='order-submit order-form']/div[1]/div[2]/input")
            buyAmountInput.clear()
            buyAmountInput.send_keys("%.4f"%amount)

            buyBtn = driver.find_element_by_xpath("//*[@class='btn buy fm']")
            buyBtn.click()
            time.sleep(4)

            sellPriceInput = driver.find_element_by_xpath("//*[@class='order-submit order-form']/div[2]/div[1]/input")
            sellPriceInput.clear()
            #sellPriceInput.send_keys("%.6f"%price)
            sellPriceInput.send_keys(market_price)

            sellAmountInput =  driver.find_element_by_xpath("//*[@class='order-submit order-form']/div[2]/div[2]/input")
            sellAmountInput.clear()
            sellAmountInput.send_keys("%.4f"%amount)

            sellBtn = driver.find_element_by_xpath("//*[@class='btn sell fm']")
            sellBtn.click()
            time.sleep(4)

            count += 1
            if count > 20:
                driver.find_element_by_xpath("//*[@class='btn el-popover__reference']").click()
                time.sleep(1)
                driver.find_element_by_xpath("//*[@role='tooltip' and @aria-hidden='false']/div[1]/div[2]/button[2]").click()
                count = 0


        time.sleep(5)
    except Exception as e:
        print(e)
        time.sleep(10)

driver.close()
