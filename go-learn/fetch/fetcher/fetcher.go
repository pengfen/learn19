package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status code")
	}

	// 对resp.Body 进行封装
	bodyReader := bufio.NewReader(resp.Body)
	// 如果网页编码为gbk 则转成utf8
	//utf8Reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	//all, err := ioutil.ReadAll(utf8Reader)

	//e := determineEncoding(resp.Body)
	//utf8Reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(bodyReader)
}


// 获取网页编码
func determineEncoding(r *bufio.Reader) encoding.Encoding  {
	//bytes, err := bufio.NewReader(r).Peek(1024)
	// io.Reader ---> *bufio.Reader
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
