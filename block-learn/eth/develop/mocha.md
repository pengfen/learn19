# 测试用例 mocha

* 安装mocha  npm install --save mocha
* 修改package.json  "test":"mocha"
* Welcome.test.js

```bash
const assert = require("assert"); // 引入断言库

class Dog {
    say() {
        return "wangwang";
    }
    happy() {
        return "miaomiao";
    }
}

let dog; // 每个测试函数it需要使用到dog 声明类的全局变量let
beforeEach(()=>{ // 抽取new Dog()的过程到beforeEach()
    dog = new Dog();
});
describe("测试dog",()=>{
    it("测试dog的say方法", ()=>{
        assert.equal(dog.say(),"wangwang");
    });
    it("测试dog的happy方法", ()=>{
        assert.equal(dog.happy(),"miaomiao");
    });
});
```

* npm install --save ganache-cli (testrpc升级版 local test network)
* npm install --save web3 (bytecode ---> 部署到ganache-cli ABI ---> 使用web3.js进行调用)
* mocha 测试框架的使用
* it 跑一个测试或者断言
* describe it函数分组
* beforeEach 执行一些初始化代码

## mocha代码测试流程
1. mocha start
2. 部署智能合约 beforeEach
3. 调用智能合约 it
4. 进行断言 it