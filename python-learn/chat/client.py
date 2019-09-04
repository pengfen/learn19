import wx
import telnetlib
from time import sleep
import _thread as thread

if __name__ == '__main__':
	app = wx.App()
	con = telnetlib.Telnet()
	LoginFrame(None, -1, title="Login", size=(320, 250))
	app.MainLoop()