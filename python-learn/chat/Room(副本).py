class Room(CommandHandler):
	"""
	房间 包含多个用户的环境 负责基本的命令处理和广播
	"""
	def __init__(self, server):
		self.server = server
		self.sessions = []

    # 添加用户进房间
	def add(self, session):
		self.session.append(session)

	# 从房间删除用户
	def remove(self, session):
		self.session.remove(session)

	# 广播 身所有用户发送指定消息 使用asynchat.asyn_chat.push 方法发送消息
	def broadcast(self, line):
		for session in self.sessions:
			session.push(line)

    # 退出房间
	def do_logout(self, session, line):
		raise EndSession
		