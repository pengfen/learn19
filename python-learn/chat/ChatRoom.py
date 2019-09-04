class ChatRoom(Room):
	"""
	聊天房间
	"""

    # 广播用户进入
	def add(self, session):
		session.push(b'登录成功')
		self.broadcast((session.name + '进入房间\n').encode('utf-8'))
		self.server.users[session.name] = session
		Room.add(self, session)

	# 广播用户离开
	def remove(self, session):
		Room.remove(self, session)
		self.broadcast((session.name + '离开房间\n').encode('utf-8'))

	# 客户端发送消息
	def do_say(self, session, line):
		self.broadcast((session.name + ":" + line + "\n").encode('utf-8'))

	# 查看在线用户
	def do_look(self, session, line):
		session.push(b'在线用户\n')
		for other in self.sessions:
			session.push((other.name + '\n').encode('utf-8'))