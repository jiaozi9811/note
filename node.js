Node.js环境中，也有唯一的全局对象，而叫global


process
process也是Node.js提供的一个对象，它代表当前Node.js进程

> process === global.process;
true
> process.version;
'v5.2.0'
> process.platform;
'darwin'
> process.arch;
'x64'
> process.cwd(); //返回当前工作目录
'/Users/michael'
> process.chdir('/private/tmp'); // 切换当前工作目录



fs  内置模块就是文件系统模块，负责读写文件
  fs.readFile()
  fs.readFileSync()
  fs.writeFile()
  fs.writeFileSync()
  fs.stat
  statSync()
  
  

stream(流)

  标准输入流（stdin）
  标准输出流（stdout）
  
  事件
  data事件表示流的数据已经可以读取了
  end事件表示这个流已经到末尾了，没有数据可以读取了
  error事件表示出错了
  
  
pipe
  Readable流
  Writable流
