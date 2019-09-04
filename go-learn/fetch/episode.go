package main

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url) // 发送get请求
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	// 读取网页内容
	buf := make([]byte, 4 * 1024)
	for {
		n, _ := resp.Body.Read(buf)
		if n == 0 {
			break
		}
		result += string(buf[:n]) // 累加读取的内容
	}
	return
}

// 开始爬取每一个笑话 每一个段子 title content err := SpiderOneJoy(url)
func SpiderOneJoy(url string) (title, content string, err error) {
	// 开始爬取页面内容
	result, err1 := HttpGet(url)
	if err1 != nil {
		// fmt.Println("HttpGet err = ", err)
		err = err1
		return
	}

	// 取关键信息
	rel := regexp.MustCompile(`<h1>(?s:(.*?))<?h1>`)
	if rel == nil {
		// fmt.Println("regexp.MustCompile err")
		err = fmt.Errorf("%s", "regexp.MustCompile err")
		return
	}

	// 取内容
	tmpTitle := rel.FindAllStringSubmatch(result, 1); // 最后一个参数为1 只过滤第一个
	for _, data := range tmpTitle {
		title = data[1]
		break
	}

	// 取内容 <div class="content-text pt10">段子内容<a
	re2 := regexp.MustCompile(`<div class="content-txt pt10">(?s:(.*?))<a id="prev" href="`)
	if re2 == nil {
		// fmt.Println("regexp.MustCompile err")
		err = fmt.Errorf("%s", "regexp.MustCompile err2")
		return
	}

	// 取内容
	tmpContent := re2.FindAllStringSubmatch(result, -1)
	for _, data := range tmpContent {
		content = data[1]
		break
	}
}

func SpiderPape(i int) {
	url := "https://www.pengfu.com/xiaohua_" + strconv.Itoa(i) + ".html"
	fmt.Printf("正在爬取第%d个网页: %s", i, url)

	// 开始爬取页面内容
	result, err := HttpGet(url)
	if err != nil {
		fmt.Println("HttpGet err = ", err)
		return
	}

	// fmt.Println("r = ", result)
	// 取 <h1 class="dp-b">
	re := regexp.MustCompile(`<h1 class="dp-b"><a href="(?s:(.*?))"`)
	if re === nil {
		fmt.Println("regexp.MustCompile err")
		return
	}

	// 取关键信息
	joyUrls := re.FindAllStringSubmatch(result, -1)
	//fmt.Println("joyUrls = ")

	// 第一个返回下标 第二个返回内容
	for _, data := range joyUrls {
		//fmt.Println("url = ", data[1])

		// 开始攫取每一个笑话 每一个段子
		title, content, err := SpiderOneJoy(data[1])
		if err != nil {
			fmt.Println("SpiderOneJoy err = ", err)
			continue
		}

		fmt.Println("title = ", title)
		fmt.Println("content = ", content)
	}
}

func DoWork(start, end int) {
	fmt.Printf("准备爬取第%d页到%d页的网址", start, end)

	for i := start; i <= end; i ++ {
		// 爬主页面
		SpiderPape(i)
	}
}

func main() {
	var start, end int
	fmt.Printf("请输入起始页(>=1):")
	fmt.Scan(&start)
	fmt.Printf("请输入终止页(>=起始页):")
	fmt.Scan(&end)

	DoWork(start, end) 
}