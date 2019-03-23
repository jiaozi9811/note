
解决被墙问题
https://packagecontrol.io/channel_v3.json

prefrences->package setting->Package control->setting-user

{
	"channels":
	[
	"C:/Users/jiao/AppData/Roaming/Sublime Text 3/channel_v3.json"
	],
}

PlainTasks


All Autocomplete
	搜索所有打开的文件来寻找匹配的提示词

SublimeCodeIntel
	代码自动完成引擎

AutoFileName
	快速帮助你在文件中写路径
	
BufferScroll
	不同的位置可以有相同的文件名
ChineseLocalization
plainnote
Color Highlighter
DocBlockr
	自动补全注释插件
Emmet
	使用仿CSS选择器的语法来生成代码，大大提高了HTML/CSS代码编写的速度
Keymaps
	快速查找所有插件的快捷键
SideBarEnhancements
	增强右键菜单文件操作功能
SublimeLinter
	代码校验插件，支持多种语言，这个是主插件，如果想检测特定的文件需要单独下载
Alignment
	代码对齐
Ctags
	函数跳转，Alt+点击函数名称，会跳转到相应的函数
PackageResourceViewer
	查看和编辑SublimeText附带的包
Terminal
	在Sublime中直接使用终端打开你的项目文件夹，并支持使用快捷键
Trmmer
	自动删除不必要的空格
TrailingSpaces
	检测并一键去除代码中多余的空格
FileDiffs
	比较两个文件的差异
Better Completion 
	js库的自动补全插件
SublimeServer 
	让Sublime成为静态WEB服务器
ConvertToUTF8 
	解决Sublime不支持GBK、GB2312编码的问题，支持打开GBK编码的文件
SublimeREPL
	允许在Sublime中运行各种语言
SFTP
WordCount
Sublime CodeIntel 	代码自动提示
Color​Picker		调色板
IMESupport		sublime中文输入法


## gosublime
gotools
gofmt
godef
进入插件packeage目录
git clone https://margo.sh/GoSublime

修改gosublime user设置
{
	"env": {
		"GOPATH":"D:/pyth/script",
		"GOROOT":"D:/envpath/go",
		"PATH":"$PATH;$GOPATH/bin"
	},
	"autocomplete_builtins": true,
	"autocomplete_closures": true,
}