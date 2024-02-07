#### 一个基于Golang的剑灵刺客输出辅助

##### 使用方法

1. 使用管理员权限运行本程序
2. 程序只在剑灵里面生效，其余位置不生效
3. 程序运行使用的按住输出松开停止的策略
4. 程序打开后状态栏会出现一个小图标，右键点击图标会出现一些功能，点击按键打印后会实时输出你按下的按键包括鼠标或者键盘
5. 使用输出的吗值填入config.ini文件的`快捷键值`填入（修改）你想要设置的启动快捷键
6. 在config.ini中`使用鼠标`这个值是一个布尔值，填入`true`或者`false`，如果填入`true`则会使用鼠标，如果填入`false`则会使用键盘
7. 程序需要剑灵运行在不缩放的模式下，否则会导致雷决无法展开，或者使用属性修改剑灵client.exe的属性，勾选`禁用高DPI缩放`
8. 分辨率不同识别精度非不一样，请参考static文件中的图片进行截图替换即可，替换的时候需要注意一般截图不能在边框以外，截图尽量覆盖文字如`x`尽量覆盖左上角的1/4处
9. BOS中毒检测比较困难，截图最好在10S倒计时一开始就截图，选择左上角开始到毒素左边中间有一把刀白色亮起的地方，大概是3/5的左侧不包含,另一个BosZdR这张图是截图左边已经读条以后的灰色部分
10. 本程序自带的图片是1080P分辨率下窗口比例缩放特效全开的效果，不满足的发现有技能不打的就自己重新截图一下，截图推荐使用`Snipaste`

#### 开发说明
1. 不要使用`go mod vendor`
2. 将源码下载到本地以后使用`Goland`或者`VSCode`打开，使用`go mod tidy`下载依赖
3. 使用`go build -ldflags="-H windowsgui"`编译成exe文件，或者使用`build.bat`进行编译
4. 需要注意的事Go版本在1.21及其以后不支持win7，如果需要支持win7请使用1.20.13版本