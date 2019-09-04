tp         TP框架
yii        yii框架
wordpress

windows apache配置虚拟主机

C:\Windows\System32\drivers\etc\hosts

http.conf 配置文件中 vhosts.conf文件打开

<VirtualHost *:80>
    ServerAdmin caopeng8787@163.com
    DocumentRoot "F:/btc/pc"
    ServerName self.btc.com
    <Directory "F:/btc/pc">
	    Options Indexes FollowSymLinks MultiViews
	    AllowOverride All
	    Order allow,deny
	    Allow from all
	    Require all granted
    </Directory>
</VirtualHost>

Options Indexes FollowSymLinks MultiViews
解决添加.htaccess文件后出现403问题

kline 