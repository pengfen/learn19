# Solidity 合约继承

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
}
```

## 子合约继承的属性和方法的权限总结

* 子合约只可以继承`public`类型的函数 而子合约可以继承`public`和`internal`类型的属性

## 合约多继承
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
}
```