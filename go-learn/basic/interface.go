package main

/*
python
def download(retriever):
    return retriever.get("www.baidu.com")
运行时才知道传入的retriever有没有get
需要注释来说明接口

c++
template <class R>
string download(const R& retriever) {
    return retriever.get("...")
}
运行时才知道传入的retriever有没有get
需要注释来说明接口

java
<R extends Retriever>
String download(R r) {
    return r.get("...")
}
// 传入的参数必须实现Retriever接口

go
同时需要Readable Appendable
*/