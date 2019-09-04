package main

import (
    "github.com/PuerkitoBio/goquery"
    "fmt"
    "os"
//  "log"
    "path"
    "strings"
    "regexp"
    "math/rand"
    "io"
    "net/http"
    "time"
)

func randFileName(dirPath string,suffix string) (fileName string) {
    var randChar string = "abcdefghimnjkqzxyt0123456789ioedkaldncalew0129387iue"
    var name string
    size := len(randChar)
    for i := 0; i < 12; i++ {
        name += string(randChar[rand.Intn(size)])
    }
    fileName = path.Join(dirPath,name)
    fileName += suffix
    _,err := os.Stat(fileName)
    if err != nil {
        if os.IsNotExist(err) {
            return fileName
        } else {
            return randFileName(dirPath,suffix)
        }
    }
    return randFileName(dirPath,suffix)
}

func writeImage(dirPath string,href string) {
    i := strings.LastIndex(href,".")
    suffix := href[i:]
    fileName := randFileName(dirPath,suffix)
    imageFile,err := os.Create(fileName)
    if err != nil {
        fmt.Printf("[writeImage create file]: fileName: %s\n href: %s\nerror: %s\n",fileName,href,err.Error())
        return
    }
    resp,err := http.Get(href)
    if err != nil {
        fmt.Printf("[writeImage http.Get: fileName]: %s\nhref: %s\n error: %s\n",fileName,href,err.Error())
        os.Remove(fileName)
        return
    }
    size,err := io.Copy(imageFile,resp.Body)
    if err != nil {
        fmt.Println("io.Copy:error: %s  href: %s\n",err.Error(),href)
        os.Remove(fileName)
    }
    fmt.Printf("Get From %s: %d bytes\n",href,size)
    return
}

func getCorrectHref(href string,url string) (correctHref string,exist bool){
    completeValidator := regexp.MustCompile(`\A((http)|(https))(.+)((.jpg)|(.png)|(.jpeg))\z`)
    imcompleteValidator := regexp.MustCompile(`\A(||)(.+)((.jpg)|(.jpeg)|(.png))\z`)
    if completeValidator.MatchString(href) {
        //fmt.Println(href)
        return href,true
    } else if imcompleteValidator.MatchString(href) {
        ret := href[1:]
        ret =  url + ret
        //fmt.Println(ret)
        return string(ret),true
    }
    return "",false
}

/**
 * 处理错误
 */
func checkErr(err error)  {
    if err != nil {
        // panic(err)
        fmt.Println("执行失败", err)
        return
    }
}

func main() {
    url := "" // 爬取地址
    doc, err := goquery.NewDocument(url)
    checkErr(err)
    pwd,err := os.Getwd()
    checkErr(err)

    dirPath := path.Join(pwd,"images")
    _,err = os.Stat(dirPath)
    if err != nil {
        if os.IsNotExist(err) {
            err = os.Mkdir(dirPath,os.ModePerm)
            checkErr(err)
        }
    }
    doc.Find("img").Each(func(i int,selection *goquery.Selection) {
            href,exist := selection.Attr("src")
            if exist {
                correctHref,exist := getCorrectHref(href,url)
                if exist {
                    go writeImage(dirPath,correctHref)
                }
            }
        })
    time.Sleep(time.Duration(12) * time.Second)
}
