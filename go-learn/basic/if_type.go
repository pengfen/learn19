package main

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
			case bool: 
			    fmt.Printf("参数%d是布尔型\n", i)
			case float64:
				fmt.Printf("参数%d是浮点型\n", i)
			case int, int64:
				fmt.Printf("参数%d是整型\n", i)
			case nil:
				fmt.Printf("参数%d是空类型\n", i)
			case string:
				fmt.Printf("参数%d是字符型\n", i)
			default:
				fmt.Printf("参数%d未知类型\n", i)
		}
	}
}