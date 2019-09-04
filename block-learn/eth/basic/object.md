# Solidity 面向对象编程

* 面向过程编程
* 面向对象编程
* 类和对象
* 如何通过solidity实现一个类

## 面向过程
C语言是面向过程的编程语言
我想听音乐
1. 打开电脑
2. 打开音乐播放器
3. 搜索一首歌曲
4. 播放歌曲
5. 暂停
6. 下一首
7. 音量调节
8. 关闭音乐播放器
9. 关闭电脑
按照预设的顺序一步一步的去执行 这个过程就是面向过程编程


## 面向对象编程
面向对象编程的语言 C++ Java IOS C# ...

对象 电脑 音乐播放器
电脑 ---> 属性 颜色 尺寸 价格 
电脑 ---> 方法/行为 开机 关机 安装音乐播放器 卸载音乐播放器

面向对象实现(我想听音乐)
1. [电脑 开机]
2. [音乐 启动]
3. [音乐 搜索]
4. [音乐 播放]
5. [音乐 调节音量]
6. [音乐 暂停]
7. [音乐 播放]
8. [音乐 关闭]
9. [电脑 关机]

## 类和对象
平果电脑 小明电脑

编程
int a;
int b;
int c;

Counter counter = new Counter();

Counter / int 相当于 `类` `a` `b` `c` `counter` 相当于对象

类和对象总结
类 类它是一个抽象的概念 它将一堆相似的东西共同的属性和行为抽象出来 
所有抽象出来的属性和行为组成一个类 可以通过类去 实例化 N多个对象
对象 对象是一个具体的东西 有具体的属性值和方法行为

## 通过solidity实现一个类
编写智能合约可以通过Atom来进行编写 也可以直接在(https://remix.ethereum.org)里面编写

```bash
pragma solidity ^0.4.4;
/*
pragma 版本声明
solidity 开发语言
0.4.4 当前合约的版本  0.4代表主版本  .4代表修复bug的升级版
^ 代表向上兼容 0.4.4~0.4.9 可以对我们当前的代码进行编译
*/

// contract Person 类比 class Person extends Contract
contract Person {
  uint _height; // 身高
  uint _age; // 年龄
  address _owner; //代表合约的拥有者

  // 方法名和合约名相同时就属于构造函数 在创建对象时 构造函数会自动
  function Person() {
    _height = 180;
    _age = 29;
    _owner = msg.sender;
  }

  function owner() constant returns (address) {
    return _owner;
  }

  // set 方法 可以修改_height属性
  function setHeight(uint height) {
    _height = height;
  }

  // 读取_height constant代表方法只读
  function height() constant returns (uint) {
    return _height;
  }

  function setAge(uint age) {
    _age = age;
  }

  function age() constant returns (uint) {
    return _age;
  }

  function kill() constant {
    if (_owner == msg.sender) {
      // 析构函数
      selfdestruct(_owner);
    }
  }
}
```

[remix](https://remix.ethereum.org/) 部署运行合约

Deploy 注意Account 中账号eth变化