# python3.6 中使用 asyncio 代替asynchat asyncore
import asynchat 
import asyncore

PORT = 6666; # 端口号

class EndSession(Exception): # 定义结束异常类
	pass



if __name__ == '__main__':

	s = ChatServer(PORT)
	try:
		print("聊天服务运行在:'0.0.0.0:{0}".format(PORT))
		asyncore.loop()
	except KeyboardInterrupt:
		print("聊天服务退出")