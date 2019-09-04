# Solidity 合约中属性和行为的访问权限

## 属性的访问权限

* 属性: 状态变量。
    
    ```bash
    pragma solidity ^0.4.4;

	// public internal private
	contract Person {
	  // internal 是合约属性默认的访问权限
	  uint internal _age;
	  uint _weight;
	  uint private _height;
	  uint public _money; // 等价于_money() 方法 但如果编写_money() 则_money() 方法会覆盖此

	  function _money() constant returns (uint) {
	    return _money;
	  }
	}
	```
    
* 说明:
  > 属性默认类型为`internal`,`internal`和`private`类型的属性都不能被外部访问，当属性的类型为`public`类型时,会生成一个和属性名相同并且返回值就是当前属性的get函数
    比如: uint public _money;
    会生成一个函数
      
	  function _money() constant returns (uint) {
	  	return _money;
	  }

   >当然 生成的这个get函灵敏会覆盖掉public类型的属性自动生成的get函数.
      比如:

	    pragma solidity ^0.4.4;

		// public internal private
		contract Person {
		  // internal 是合约属性默认的访问权限
		  uint internal _age;
		  uint _weight;
		  uint private _height;
		  uint public _money; // 等价于_money() 方法 但如果编写_money() 则_money() 方法会覆盖此

		  function _money() constant returns (uint) {
		    return 120;
		  }
		}

  >在这个代码里面, _money函数返回值为120, 而不是返回0.
	
## 方法/行为访问权限
    
* 方法/行为: 合约里面的函数

    ```bash
    pragma solidity ^0.4.4;

	contract Person {

	  function age() constant returns (uint) {
	    return 55;
	  }

	  function weight() constant returns (uint) {
	    return 180;
	  }

	  function height() constant returns (uint) {
	    return 172;
	  }

	  function money() constant returns (uint) {
	    return 32000; 
	  }
	}
    ```

    >合约中的方法默认为`public`类型

    ```bash
    pragma solidity ^0.4.4;

	contract Person {

	  function age() constant returns (uint) {
	    return 55;
	  }

	  function weight() constant public returns (uint) {
	    return 180;
	  }

	  function height() constant internal returns (uint) {
	    return 172;
	  }

	  function money() constant private returns (uint) {
	    return 32000;
	  }
	}
    ```



## 属性和方法自己合约内部的访问权限总结


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

  function getWeight1() constant returns (uint) {
    return this.getWeight();
  }

  function getHeight1() constant returns (uint) {
    return this.getHeight();
  }

  function getMoney1() constant returns (uint) {
    return getMoney();
  }

  function getAge1() constant returns (uint) {
    return getAge(); 
  }
}
```

* 属性默认权限为`internal`,只有`public`类型的属性才可能供外部访问 `internal`和`private`类型的属性只能在合约内部使用

* 函数的权限默认是`public`类型 `public`类型的函数可供外部访问 而`internal`和`private`类型的函数不能够通过指针进行访问 哪怕是在内部通过this访问都会报错 在合约内部访问 可直接访问函数

* 不管是属性还是方法 只有是public类型时 才可以通过合约地址进行访问 合约内部的this就是当前合约地址 在合约倍如果要访问`internal`和private类型的属性或者是函数 直接访问即可 不要试图通过this去访问