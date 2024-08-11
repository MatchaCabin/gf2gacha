# 少女前线2:追放 抽卡导出分析工具

## 为什么做这个工具？

[追放采购记录导出工具](https://github.com/EtherealAO/exilium-recruit-export)作者在其[NGA发布贴](https://bbs.nga.cn/read.php?tid=38812531)中表示不搞了，希望有人接盘

但我并不熟悉`Electron`，也不太喜欢`Electron`直接包一个巨大浏览器的做法

于是把之前自用的导出程序用更加轻量的`Wails`包了个界面传了上来

## 功能计划

根据优先级排序
- [x] 多用户
- [x] 增量更新
- [ ] 全量更新
- [x] 导入[追放采购记录导出工具](https://github.com/EtherealAO/exilium-recruit-export)数据

## 使用方法

打开软件，点击`增量更新`，等待读取服务端数据，即可显示抽卡数据

## 基本原理

本软件通过读取日志获取`游戏路径`、`抽卡链接`、`UID`、`AccessToken`，进而从服务器获取抽卡记录

本软件会在当前目录下的生成名为`gf2gacha.db`的`sqlite`数据库，这就是你的抽卡数据了，你可以使用任何支持sqlite的数据库管理工具查看其内容

## 特别提醒

本软件数据不会将用户数据传输到任何第三方服务上，仅存储在你本地

如果你有云端托管数据的需求，请点击软件右上角的`回形针图标`查看你当前日志信息，将其复制给第三方服务提供方

注意！`AccessToken`是你的临时登录凭证，请勿随意泄露

## FAQ

### 为什么看不到新手池数据？为什么角色池/武器池/常驻池缺失数据？
追放服务器只保留最近180天的数据，大多数抽卡游戏都有类似的限制，所以有需求的用户才会采取多种办法来保存抽卡记录

### 如何导入来自[追放采购记录导出工具](https://github.com/EtherealAO/exilium-recruit-export)的数据?
由于ERE的UID使用的UID并不是游戏内的UID，我无法判断原数据是属于哪个用户

**请务必先在主界面手动选择你要导入旧数据的用户对应的UID**

然后点击`导入ERE数据`按钮，在ERE程序目录中找到`userData`文件夹，选择里面类似`gacha-list-1234567890.json`的文件,点击确定即可