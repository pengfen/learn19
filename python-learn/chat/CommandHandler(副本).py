class CommandHandler:
	"""
	命令处理类
	"""

    # 响应未知命令 通过 aynchat.async_chat.push 方法发送消息
	def unknown(self, session, cmd):
		session.push(("未知命令 {} \n".fomat(cmd)).encode("utf-8"))

	def handler(self, session, line):
		line = line.decode()

		# 命令处理
		if not line.strip(): 
			return

		parts = line.splite(' ', 1)
		cmd = parts[0]
		try:
		    line = parts[1].strip()
		except IndexError:
		    line = ''

		# 通过协议代码执行相应的方法
		method = getattr(self, 'db_' + cmd, None)
		try:
		    method(session, line)
		except TypeError:
		    self.unknown(session, cmd) 
