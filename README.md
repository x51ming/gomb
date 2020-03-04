# 致谢

Fork 自 github.com/selfplan/gowebui

进行了部分修改，移除了gonode.dll

基于github.com/weolar/miniblink49 封装的GOLANG使用的miniblink

# 关于本项目

基于github.com/weolar/miniblink49 封装的GOLANG使用的WebUI。

直接调用DLL，未使用CGO。

# 使用前准备

- 暂时只支持32位dll，

- 64位dll可以使用，但是loadURL会出错，可以载入html，这个很奇怪（暂时不知道问题出在哪）

# 目前已实现公开接口

- 写代码时，考虑到升级64位的需要，使用了uintptr，且大部分代码在编写时都是32位64位兼容

- 通过dumpbin，将node.dll导出的函数（wke\*\*，js\*\*）全部用dll调用包装了，但是只实现了一部分函数，具体见 generatedFun.go

- 模块构成：

  - api.go：暂时只有freeMBDLL()，释放dll

  - apicall.go：接受api调用命令chan，以及事件循环
    （来自原selfplan的代码）

  - constants.go：从wke.h中摘取的部分常数，后面用到别的再添加

  - defines.go：从wke.h中摘取的类型定义，包含了大部分常用的回调函数类型，以及一些结构体

  - fastdll.go：对部分具有特定类型的参数列表的函数写了一些快速调用的函数，如func (mb *WebView) f_void(fun uintptr) uintptr是指函数只使用到第一个WkeWebView类型的参数；该部分主要用于减小代码量

  - generatedFun.go：使用工具从dll中将wke，js开头的全部函数导出成了dll调用，对其中较为常用的一些函数进行了实现，其他的用到再实现

  - generatedLoad.go：载入dll，载入函数

  - generatedVar.go：存放函数地址的变量

  - gomb.go：几个小函数

    ```go
    func Run(dll string) bool
    //dll是node.dll的路径
    //载入dll，开始事件循环，初始化mb，返回初始化结果（是否成功）
    func Finalize()
    //结束事件循环
    //还有几个小函数，没高兴删，在generatedFun.go中都有
    ```

  - tools.go：一些小工具，如类型转换等

  - usefulltools.go：常用工具

    ```go
    func MyGetDevtoolsPath(s string) string
    //相对路径，返回file：///+绝对路径
    //这里到底应该用几个/搞不清楚，暂时没出问题
    func (mb *WebView) MyBindClosing()
    //监听窗口关闭事件，关闭窗口时，弹出确认窗口
    func (mb *WebView) MyDisablePic()
    //禁止载入图片（傻傻的，根据后缀名判断）
    //有个WkeWillSendRequestInfo结构体中有ResourceType，但是没搞清楚怎么用，WkeOnOtherLoadCallback回调中用到了这个，但是，不知道怎么取消netjob，所以暂时使用的是后缀名判断
    //WkeOnOtherLoad已经实现，可以尝试一下
    func (mb *WebView) MyBindDestroy() chan bool
    //监听窗口销毁事件，返回chan，销毁时向chan发送true
    func (mb *WebView) MyBindNavi()
    //监听导航事件，发生导航时，输出导航类型以及webview
    func (mb *WebView) MyBindReady()
    //监听Ready，Ready时输出webview以及frameid
    func (mb *WebView) MyWaitReady(timeout int) chan bool
    //返回的chan，当ready时向其发送true
    //使用方法<-MyWaitReady(timeout)即会阻塞，超时单位秒
    
    
    //后几个函数，取得js参数
    //用于在绑定的js调用的回调里获取参数，实现了string，bol
    ```

  - webview.go：最开始自己写的一些函数，在generatedFun中都有对应函数（generatedFun是后来搞的，原因是要用到的函数太多），是自己最常用的一批函数，函数名没有wke前缀

- 示例：

  - node.dll放在examples目录

  - 基本使用：examples/first

  - 绑定js函数：examples/bindjs

