# 方法重写

```bash
pragma solidity ^0.4.4;

contract Animal {
  uint _weight;
  uint private _height;
  uint internal _age;
  uint public _money;

  function getWeight() constant returns (uint) {
    return _weight;
  }

  function getHeight() constant public returns (uint) {
    return _height;
  }

  function getAge() constant internal returns (uint) {
    return _age;
  }

  function getMoney() constant private returns (uint) {
    return _money;
  }
}

contract Animal1 {
  uint _sex; //

  function Animal1() {
    _sex = 1;
  }

  function sex() constant returns (uint) {
    return _sex;
  }
}

contract Dog is Animal {

  function testWeight() constant returns (uint) {
    return _weight;
  }

  function testHeight() constant returns (uint) {
    return _height;
  }

  function testAge() constant returns (uint) {
    return _age;
  }

  function testMoney() constant returns (uint) {
    return _money;
  }

  // 方法重写 
  function sex() constant returns (uint) {
    return 100;
  }
}
```

# 值传递
* 值类型包含
* 布尔 booleans
* 整型 Integer
* 地址 Address
* 定长字节数组 (fixed byte arrays)
* 有理数和整型 (Rational and Integer Literals, String literals)
* 枚举类型 Enums
* 函数 Function Types

```bash
pragma solidity ^0.4.4;

contract Person {

  uint _age;

  function Person(uint age) {
    _age = age;
  }

  function f() {
    modify(_age);
  }

  function modify(uint age) {
    age = 100;
  }

  function age() constant returns (uint) {
    return _age; 
  }
}
```

# 引用传递
* 引用类型包含
* 不定长字节数组
* 字符串 string
* 数组 array
* 结构体 structs

```bash
pragma solidity ^0.4.4;

contract Person {

  string _name;

  function Person(string name) {
    _name = name;
  }

  function f() {
    modify(_name);
  }

  // 值传递(modify(string name)) ---> 引用传递 modify(string storage name)
  // 注意 引用传递类型必须使用 private或internel
  function modify(string storage name) private {
    bytes(name)[0] = "L";
  }

  function name() constant returns (string) {
    return _name;
  }
}
```