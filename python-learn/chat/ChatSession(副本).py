class ChatSession(asynchat.async_chat):
	"""
	负责和客户端通信
	"""

	def __init__(self, server, sock):
		asynchat.async_chat.__init__(self, sock)
		self.server = server
		self.set_terminator(b'\n')
		self.data = []
		self.name = None
		self.center(LoginRoom(server)) # 登录用户

	# 从当前房间移除自身 然后添加到指定房间	
	def enter(self, room):
		try:
			cur = self.room
		except AttributeError:
			pass
		else:
			cur.remove(self)
		self.room = room
	    room.add(self)

	# 接收客户端的数据
	def collect_incoming_data(self, data):
		self.data.append(data.decode("utf-8"))

	def found_terminator(self):
		line = ''.join(self.data)
		self.data = []
		try:
			self.room.handle(self, line.encode("utf-8"))
		except EndSession: # 退出聊天室处理
			self.handle_close()

    # 当session关闭时 将退出房间 LogoutRoom 
	def handle_close():
		asynchat.async_chat.handle_close(self)
		self.enter(LogoutRoom(self.server))
