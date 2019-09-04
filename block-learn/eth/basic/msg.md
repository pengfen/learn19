# msg全局变量

* msg 全称message 描述了transaction和call的所有细节信息
* msg.data transaction或call的data的信息
* msg.gas transaction的调用花了多少gas
* msg.sender 调用者的address
* msg.value 调用者花了多少钱(ether)

```bash
function Lottery() public {
	manager = msg.sender;
}
```