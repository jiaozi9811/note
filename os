os.sep    取代操作系统特定的路径分割符
os.name   字符串指示正在使用的平台。比如对于Windows,它是'nt',而对于Linux/Unix,它是'posix'
os.getcwd() 当前工作目录,即Python脚本工作的目录路径
os.getenv()和os.putenv()   分别用来读取和设置环境变量
os.listdir()  返回指定目录下的所有文件和目录名
os.remove()   用来删除一个文件
os.system()   运行shell命令
os.linesep    给出当前平台使用的行终止符。例如，Windows使用'\r\n'，Linux使用'\n'而Mac使用'\r'
os.path.split('path/file')   返回一个路径的目录名和文件名
os.path.isfile()  检验给出的路径是文件
os.path.isdir()   检验给出的路径是目录
os.path.existe()  用来检验给出的路径是否真地存在

os.getcwd()   获得当前工作目录
os.curdir     返回当前目录（'.')
os.chdir(dirname) 改变工作目录到dirname
os.path.getsize(name) 获得文件大小，如果name是目录返回
os.path.abspath(name) 获得绝对路径
os.path.split(name)   分割文件名与目录（事实上，如果你完全使用目录，它也会将最后一个目录作为文件名而分离，同时它不会判断文件或目录是否存在）
os.path.splitext():分离文件名与扩展名
os.path.join(path,name):连接目录与文件名或目录
os.path.basename(path):返回文件名
os.path.dirname(path):返回文件路径
