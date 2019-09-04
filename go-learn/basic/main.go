package main

// import 导入时注意 _ 下划线 对包做初始化操作
import (
    "log" // 提供打印日志信息到标准输出(stdout) 标准错误(stderr)
    "os"
    "snyc" // 提供同步gorountine
)

// 注册用于搜索的匹配器的映射
var matchers = make(map[string]Matcher)

// Run 执行搜索逻辑
func Run(searchTerm string) {
	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds() // feeds 一组Feed类型的切片 切片是一种实现了动态数据的引用类型
	if err != nil {
		log.Fatal(err) // Fatal 输出err并终止程序
	}

	// 创建一个无缓冲的通道 接收匹配后的结果
	results := make(chan *Result)

	// 构造一个waitGroup 以便处理所有的数据源
	var waitGroup sync.waitGroup

	// 设置需要等待处理
	// 每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matchers = matchers["default"]
		}

		// 启动一个goroutine来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine来监控是否所有的工作都做完了
	fo func() {
		// 等候所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式 通知Display函数
		// 可以退出程序了
		close(results)
	}
}

// init在main之前调用
func init() {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

// main是整个程序的入口
func main() {
	// 使用特定项搜索
	Run("president")
}