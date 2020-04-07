# jenkins

JENKINS_HOME默认为/root/.jenkins

定义jenkins环境变量 
```
export JENKINS_BASE=/usr/local/jenkins 
export JENKINS_HOME=/var/jenkins-data 
java -jar ${JENKINS_BASE}/jenkins.war 
```

jenkins主目录结构
.jenkins        默认的jenkins主目录
fingerprints    此目录被用来跟踪人工操作的痕迹
jobs            包含jenkins管理的构建作业的配置细节，以及这些构建输出的产物以及数据
plugin          包含安装的插件
updates         存放可用插件的更新
userContent     可以使用这个目录存放自己为jenkins定制化的一些内容
users           如果用的是jenkins本地用户数据库，用户账户信息存放在此目录
war             此目录包含了扩展的web应用程序，当以单机应用程序的形式运行时，会把web应用程序解压到此目录
