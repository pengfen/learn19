class LoginRoom(Room):
	"""
	处理登录用户
	"""
	def add(self, session):
		# 用户连接成功的回应
		Room.add(self, session)
		# 使用 asynchat.asyn_chat.push 方法发送数据
		session.push(b'连接成功')

	def do_login(self, session, line):
		# 用户登录逻辑
		name = line.strip()
		# 获取用户名称
		if not name:
			session.push(b'用户名为空')
		# 检测是否有同名用户
	    elif name in self.server.users:
	    	session.push(b'用户名存在')
	    # 用户名检测成功后 进入主聊天室
	    else:
	    	session.name = name
	    	session.enter(self.server.main_room)
		