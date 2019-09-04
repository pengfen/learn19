class ChatServer(asyncore.dispatcher):
	"""
	聊天服务器
	"""
	def __init__(self, port):
		asyncore.dispatcher.__init__(self)
		self.create_socket() # 创建socket
		self.set_reuse_addr() # 设置socket为可重用
		self.bind("", port) # 绑定端口
		self.listen(5)
		self.users = {}
		self.main_room = ChatRoom(self) # 绑定房间

	def handle_accept(self):
		conn, addr = self.accept()
		ChatSession(self, conn)