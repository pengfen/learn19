# solidity数组入门

* 定长数组 数组长度固定 int[3] bool[2]
* 动态数组 数组长度动态调整 int[] bool[]

uint[] public myArray;
myArray.push(i);
myArray.push(30)
myArray.push(20)
myArray.length
myArray[0];

## solidity的map和结构体
* 定长数组 数组长度固定 int[3] bool[2]
* 动态数组 数组长度动态调整 int[] bool[]
* mapping 键值对 类似java的map  mapping(string=>string) mapping(int=>bool)
* struct 结构体 struct Student{string name; string id}
* mapping 像词典 根据key可以获取到value 描述很多事物 struct是用来描述复杂的数据类型 描述一个事物

## solidity 的二维数组的小细节
* solidity 支持二维/多维数组
* const myArray = [1,2,3],[4,5,6],[7,8,9]
* ABI 对多维数组支持不完备 string是字符数组 string[] 不支持