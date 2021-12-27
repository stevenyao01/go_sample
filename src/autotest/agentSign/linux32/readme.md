EdgeAgent服务安装说明:(注:安装agent服务需要使用root用户)
	1、解压压缩包：
		unzip EdgeAgent_linux_64-1.6-A161XXXX.zip
		若系统无unzip命令，则先安装unzip
		解压后文件：EdgeAgent_linux_64-1.6-A161XXXX	mqtt.conf	server.crt	agent.sh	readme.md
	2、修改mqtt.conf文件:
		将ip改为EdgeServer的ip地址，如：broker=172.17.172.12:4567
	3、在Leapiot平台中的设置中选择Edge连接密钥，点击下载密钥文件，下载device.sk文件后放到此目录下
	4、给agent.sh赋执行权限:	chmod +x agent.sh
	5、执行agent.sh:(注:安装完成后agent服务会自动启动)	./agent.sh install
	6、安装完agent服务后可执行的操作有:service agent start(stop|status|uninstall)
