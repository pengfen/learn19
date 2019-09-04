class LogoutRoom(Room):
	"""
	处理退出用户
	"""

	# 从服务器上移除用户
	def add(self, session):
		try:
			del self.server.users[session.name]
		except KeyError:
			pass