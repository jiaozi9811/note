# Gradle

https://gradle.org/
https://gradle.org/releases/(download)

## 安装 
Gradle 需要运行在一个 Java 环境里
- 安装一个 Java JDK 或者 JRE. 而且 Java 版本必须至少是 6 以上.
- Gradle 自带 Groovy 库, 所以没必要安装 Groovy. 任何已经安装的 Groovy 会被 Gradle 忽略

### 设置环境变量 
- 添加一个 GRADLE_HOME 环境变量来指明 Gradle 的安装路径
- 添加 GRADLE_HOME/bin 到您的 PATH 环境变量中. 通常, 这样已经足够运行Gradle了
- JAVA_OPTS 是一个用于 JAVA 应用的环境变量. 一个典型的用例是在 JAVA_OPTS 里设置HTTP代理服务器(proxy)
- GRADLE_OPTS 是内存选项. 这些变量可以在 gradle 的一开始就设置或者通过 gradlew 脚本来设置
```
export GRADLE_HOME = /usr/local/gradle
export export PATH=$PATH:$GRADLE_HOME/bin
gradle -v
```

## 构建脚本基础 
Gradle 里的任何东西都是基于这两个基础概念 
- projects (项目) 
- tasks (任务) 

- 每一个构建都是由一个或多个 projects 构成的. 一个 project 到底代表什么取决于你想用 Gradle 做什么 
- 每一个 project 是由一个或多个 tasks 构成的. 一个 task 代表一些更加细化的构建

