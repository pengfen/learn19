class ChatFrame(wx.Frame):
    """
    聊天窗口
    """

    def __init__(self, parent, id, title, size):
        # 初始化，添加控件并绑定事件
        wx.Frame.__init__(self, parent, id, title)
        self.SetSize(size)
        self.Center()
        self.chatFrame = wx.TextCtrl(self, pos=(5, 5), size=(490, 310), style=wx.TE_MULTILINE | wx.TE_READONLY)
        self.message = wx.TextCtrl(self, pos=(5, 320), size=(300, 25))
        self.sendButton = wx.Button(self, label="Send", pos=(310, 320), size=(58, 25))
        self.usersButton = wx.Button(self, label="Users", pos=(373, 320), size=(58, 25))
        self.closeButton = wx.Button(self, label="Close", pos=(436, 320), size=(58, 25))
        # 发送按钮绑定发送消息方法
        self.sendButton.Bind(wx.EVT_BUTTON, self.send)
        # Users按钮绑定获取在线用户数量方法
        self.usersButton.Bind(wx.EVT_BUTTON, self.lookUsers)
        # 关闭按钮绑定关闭方法
        self.closeButton.Bind(wx.EVT_BUTTON, self.close)
        thread.start_new_thread(self.receive, ())
        self.Show()

    def send(self, event):
        # 发送消息
        message = str(self.message.GetLineText(0)).strip()
        if message != '':
            con.write(('say ' + message + '\n').encode("utf-8"))
            self.message.Clear()

    def lookUsers(self, event):
        # 查看当前在线用户
        con.write(b'look\n')

    def close(self, event):
        # 关闭窗口
        con.write(b'logout\n')
        con.close()
        self.Close()

    def receive(self):
        # 接受服务器的消息
        while True:
            sleep(0.6)
            result = con.read_very_eager()
            if result != '':
                self.chatFrame.AppendText(result)
