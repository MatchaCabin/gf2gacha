# 少女前线2:追放 抽卡导出分析工具

## 为什么做这个工具？

[追放采购记录导出工具](https://github.com/EtherealAO/exilium-recruit-export)作者在其[NGA发布贴](https://bbs.nga.cn/read.php?tid=38812531)中表示不搞了，希望有人接盘

但我并不熟悉`Electron`，也不太喜欢`Electron`直接包一个巨大浏览器的做法

于是把之前自用的导出程序用更加轻量的`Wails`包了个界面传了上来

![image](sample.png)

## 主要功能

- [x] 增量更新 `拉取服务端最新的抽卡数据，与本地数据合并`
- [x] 全量更新 `拉取服务端所有抽卡数据，与本地数据合并(目前合并算法准备重写，完成后再做)`
- [x] 一键社区 `一键完成社区签到与任务，并根据自定义设置兑换道具`
- [x] 导入Ere数据 `导入EreJson或EreExcel格式的数据，与本地数据合并(合并前会自动备份数据库)`
- [x] 导入/导出RawJson `导入/导出服务器原始格式的数据的map`
- [x] 导入/导出MccExcel `导入/导出Mcc格式的Excel`

## 基本原理

本软件通过读取日志获取`游戏路径`、`抽卡链接`、`UID`、`AccessToken`，进而从服务器获取抽卡记录

本软件会在当前目录下的生成名为`gf2gacha.db`的`sqlite`数据库，这就是你的抽卡数据了，你可以使用任何支持sqlite的数据库管理工具查看其内容

## 特别提醒

本软件不会将用户数据传输到任何第三方服务上，数据仅存储在你本地

进行高危操作时(如合并数据)，数据库会自动备份

如果你有云端托管数据的需求，请点击软件右上角的`回形针图标`查看你当前日志信息，将其复制给第三方服务提供方

注意！`AccessToken`是你的临时登录凭证，请勿随意泄露

## FAQ

### 为什么看不到新手池数据？为什么角色池/武器池/常驻池缺失数据？

追放服务器只保留最近180天的数据，大多数抽卡游戏都有类似的限制，所以有需求的用户才会采取多种办法来保存抽卡记录

### 如何导入来自[追放采购记录导出工具](https://github.com/EtherealAO/exilium-recruit-export)的数据?

由于ERE的UID使用的UID并不是游戏内的UID，我无法判断原数据是属于哪个用户

**请务必先在主界面手动选择你要导入旧数据的用户对应的UID**

然后点击`导入ERE数据`按钮，在ERE程序目录中找到`userData`文件夹，选择里面类似`gacha-list-1234567890.json`的文件,点击确定即可

## 附录

### 卡池类型

| PoolType | 备注    |
|----------|-------|
| 1        | 常驻池   |
| 2        | 暂无    |
| 3        | 角色池   |
| 4        | 武器池   |
| 5        | 新手池   |
| 6        | 自选角色池 |
| 7        | 自选武器池 |
| 8        | 神秘箱   |

### 卡池信息

| PoolId | 备注           | 时间戳起       | 时间戳止       | 时间起                           | 时间止                           |
|--------|--------------|------------|------------|-------------------------------|-------------------------------|
| 1001   | 常规采购         |            |            |                               |                               |
| 2001   | 军备提升β-托洛洛武器  | 1703124000 | 1704358799 | 2023-12-21 10:00:00 +0800 CST | 2024-01-04 16:59:59 +0800 CST |
| 3001   | 定向采购β-托洛洛    | 1703124000 | 1704358799 | 2023-12-21 10:00:00 +0800 CST | 2024-01-04 16:59:59 +0800 CST |
| 4001   | 初始采购-新手池     |            |            |                               |                               |
| 5001   | 军备提升α-黛烟武器   | 1703494800 | 1705913999 | 2023-12-25 17:00:00 +0800 CST | 2024-01-22 16:59:59 +0800 CST |
| 6001   | 定向采购α-黛烟     | 1703494800 | 1705913999 | 2023-12-25 17:00:00 +0800 CST | 2024-01-22 16:59:59 +0800 CST |
| 7001   | 定向采购β-塞布丽娜   | 1705395600 | 1706605199 | 2024-01-16 17:00:00 +0800 CST | 2024-01-30 16:59:59 +0800 CST |
| 8001   | 军备提升β-塞布丽娜武器 | 1705395600 | 1706605199 | 2024-01-16 17:00:00 +0800 CST | 2024-01-30 16:59:59 +0800 CST |
| 9001   | 定向采购α-桑朵莱希   | 1705978800 | 1708657199 | 2024-01-23 11:00:00 +0800 CST | 2024-02-23 10:59:59 +0800 CST |
| 10001  | 军备提升α-桑朵莱希武器 | 1705978800 | 1708657199 | 2024-01-23 11:00:00 +0800 CST | 2024-02-23 10:59:59 +0800 CST |
| 11001  | 定向采购β-琼玖     | 1707382800 | 1708657199 | 2024-02-08 17:00:00 +0800 CST | 2024-02-23 10:59:59 +0800 CST |
| 12001  | 军备提升β-琼玖武器   | 1707382800 | 1708657199 | 2024-02-08 17:00:00 +0800 CST | 2024-02-23 10:59:59 +0800 CST |
| 13001  | 定向采购α-莱娜     | 1708657200 | 1711421999 | 2024-02-23 11:00:00 +0800 CST | 2024-03-26 10:59:59 +0800 CST |
| 14001  | 军备提升α-莱娜武器   | 1708657200 | 1711421999 | 2024-02-23 11:00:00 +0800 CST | 2024-03-26 10:59:59 +0800 CST |
| 15001  | 定向采购-黛烟      | 1712653200 | 1714445999 | 2024-04-09 17:00:00 +0800 CST | 2024-04-30 10:59:59 +0800 CST |
| 16001  | 军备提升-黛烟武器    | 1712653200 | 1714445999 | 2024-04-09 17:00:00 +0800 CST | 2024-04-30 10:59:59 +0800 CST |
| 17001  | 定向采购-绛雨      | 1711422000 | 1714445999 | 2024-03-26 11:00:00 +0800 CST | 2024-04-30 10:59:59 +0800 CST |
| 18001  | 军备提升-绛雨武器    | 1711422000 | 1714445999 | 2024-03-26 11:00:00 +0800 CST | 2024-04-30 10:59:59 +0800 CST |
| 19001  | 定向采购β-莫辛纳甘   | 1709629200 | 1711421999 | 2024-03-05 17:00:00 +0800 CST | 2024-03-26 10:59:59 +0800 CST |
| 20001  | 军备提升β-莫辛纳甘武器 | 1709629200 | 1711421999 | 2024-03-05 17:00:00 +0800 CST | 2024-03-26 10:59:59 +0800 CST |
| 21001  | 定向采购-佩里缇亚    | 1711422000 | 1712653199 | 2024-03-26 11:00:00 +0800 CST | 2024-04-09 16:59:59 +0800 CST |
| 22001  | 军备提升-佩里缇亚武器  | 1711422000 | 1712653199 | 2024-03-26 11:00:00 +0800 CST | 2024-04-09 16:59:59 +0800 CST |
| 23001  | 定向采购-玛绮朵     | 1714446000 | 1717556399 | 2024-04-30 11:00:00 +0800 CST | 2024-06-05 10:59:59 +0800 CST |
| 24001  | 军备提升-玛绮朵武器   | 1714446000 | 1717556399 | 2024-04-30 11:00:00 +0800 CST | 2024-06-05 10:59:59 +0800 CST |
| 25001  | 定向采购-桑朵莱希    | 1715677200 | 1717556399 | 2024-05-14 17:00:00 +0800 CST | 2024-06-05 10:59:59 +0800 CST |
| 26001  | 军备提升-桑朵莱希武器  | 1715677200 | 1717556399 | 2024-05-14 17:00:00 +0800 CST | 2024-06-05 10:59:59 +0800 CST |
| 27001  | 定向采购-乌尔丽德    | 1717556400 | 1719284399 | 2024-06-05 11:00:00 +0800 CST | 2024-06-25 10:59:59 +0800 CST |
| 28001  | 军备提升-乌尔丽德武器  | 1717556400 | 1719284399 | 2024-06-05 11:00:00 +0800 CST | 2024-06-25 10:59:59 +0800 CST |
| 29001  | 定向采购-莱娜      | 1717556400 | 1719284399 | 2024-06-05 11:00:00 +0800 CST | 2024-06-25 10:59:59 +0800 CST |
| 30001  | 军备提升-莱娜武器    | 1717556400 | 1719284399 | 2024-06-05 11:00:00 +0800 CST | 2024-06-25 10:59:59 +0800 CST |
| 31001  | 定向采购-索米      | 1719284400 | 1721098799 | 2024-06-25 11:00:00 +0800 CST | 2024-07-16 10:59:59 +0800 CST |
| 32001  | 军备提升-索米武器    | 1719284400 | 1721098799 | 2024-06-25 11:00:00 +0800 CST | 2024-07-16 10:59:59 +0800 CST |
| 33001  | 定向采购-绛雨      | 1719284400 | 1721098799 | 2024-06-25 11:00:00 +0800 CST | 2024-07-16 10:59:59 +0800 CST |
| 34001  | 军备提升-绛雨武器    | 1719284400 | 1721098799 | 2024-06-25 11:00:00 +0800 CST | 2024-07-16 10:59:59 +0800 CST |
| 35001  | 定向采购-杜莎妮     | 1721098800 | 1722913199 | 2024-07-16 11:00:00 +0800 CST | 2024-08-06 10:59:59 +0800 CST |
| 36001  | 军备提升-杜莎妮武器   | 1721098800 | 1722913199 | 2024-07-16 11:00:00 +0800 CST | 2024-08-06 10:59:59 +0800 CST |
| 37001  | 定向采购-玛绮朵     | 1721098800 | 1722913199 | 2024-07-16 11:00:00 +0800 CST | 2024-08-06 10:59:59 +0800 CST |
| 38001  | 军备提升-玛绮朵武器   | 1721098800 | 1722913199 | 2024-07-16 11:00:00 +0800 CST | 2024-08-06 10:59:59 +0800 CST |
| 39001  | 定向采购-朝晖      | 1722913200 | 1724705999 | 2024-08-06 11:00:00 +0800 CST | 2024-08-27 04:59:59 +0800 CST |
| 40001  | 军备提升-朝晖武器    | 1722913200 | 1724705999 | 2024-08-06 11:00:00 +0800 CST | 2024-08-27 04:59:59 +0800 CST |
| 41001  | 定向采购-乌尔丽德    | 1722913200 | 1724705999 | 2024-08-06 11:00:00 +0800 CST | 2024-08-27 04:59:59 +0800 CST |
| 42001  | 军备提升-乌尔丽德武器  | 1722913200 | 1724705999 | 2024-08-06 11:00:00 +0800 CST | 2024-08-27 04:59:59 +0800 CST |
| 43001  | 定向采购-可露凯     | 1724706000 | 1726714799 | 2024-08-27 05:00:00 +0800 CST | 2024-09-19 10:59:59 +0800 CST |
| 44001  | 军备提升-可露凯武器   | 1724706000 | 1726714799 | 2024-08-27 05:00:00 +0800 CST | 2024-09-19 10:59:59 +0800 CST |
| 45001  | 定向采购-索米      | 1724706000 | 1726714799 | 2024-08-27 05:00:00 +0800 CST | 2024-09-19 10:59:59 +0800 CST |
| 46001  | 军备提升-索米武器    | 1724706000 | 1726714799 | 2024-08-27 05:00:00 +0800 CST | 2024-09-19 10:59:59 +0800 CST |
| 90001  | 自选采购·人形      | 1724706000 | 1728442799 | 2024-08-27 05:00:00 +0800 CST | 2024-10-09 10:59:59 +0800 CST |
| 91001  | 自选采购·军备      | 1724706000 | 1728442799 | 2024-08-27 05:00:00 +0800 CST | 2024-10-09 10:59:59 +0800 CST |
| 99001  | 神秘箱          | 1689476400 | 4089668399 | 2023-07-16 11:00:00 +0800 CST | 2099-08-06 10:59:59 +0800 CST |

### 人形信息

| ItemId                          | Rank                         | 备注                              |
|---------------------------------|------------------------------|---------------------------------|
| <font color=#B288DD>1001</font> | <font color=#B288DD>4</font> | <font color=#B288DD>克罗丽科</font> |
| <font color=#B288DD>1008</font> | <font color=#B288DD>4</font> | <font color=#B288DD>纳美西丝</font> |
| <font color=#B288DD>1009</font> | <font color=#B288DD>4</font> | <font color=#B288DD>寇尔芙</font>  |
| <font color=#FFA500>1013</font> | <font color=#FFA500>5</font> | <font color=#FFA500>莱娜</font>   |
| <font color=#FFA500>1015</font> | <font color=#FFA500>5</font> | <font color=#FFA500>维普蕾</font>  |
| <font color=#B288DD>1017</font> | <font color=#B288DD>4</font> | <font color=#B288DD>闪电</font>   |
| <font color=#FFA500>1021</font> | <font color=#FFA500>5</font> | <font color=#FFA500>佩里缇亚</font> |
| <font color=#B288DD>1022</font> | <font color=#B288DD>4</font> | <font color=#B288DD>夏克里</font>  |
| <font color=#FFA500>1023</font> | <font color=#FFA500>5</font> | <font color=#FFA500>杜莎妮</font>  |
| <font color=#B288DD>1024</font> | <font color=#B288DD>4</font> | <font color=#B288DD>奇塔</font>   |
| <font color=#FFA500>1025</font> | <font color=#FFA500>5</font> | <font color=#FFA500>托洛洛</font>  |
| <font color=#B288DD>1026</font> | <font color=#B288DD>4</font> | <font color=#B288DD>纳甘</font>   |
| <font color=#FFA500>1027</font> | <font color=#FFA500>5</font> | <font color=#FFA500>琼玖</font>   |
| <font color=#FFA500>1028</font> | <font color=#FFA500>5</font> | <font color=#FFA500>桑朵莱希</font> |
| <font color=#FFA500>1029</font> | <font color=#FFA500>5</font> | <font color=#FFA500>塞布丽娜</font> |
| <font color=#FFA500>1032</font> | <font color=#FFA500>5</font> | <font color=#FFA500>黛烟</font>   |
| <font color=#FFA500>1033</font> | <font color=#FFA500>5</font> | <font color=#FFA500>莫辛纳甘</font> |
| <font color=#FFA500>1034</font> | <font color=#FFA500>5</font> | <font color=#FFA500>玛绮朵</font>  |
| <font color=#FFA500>1035</font> | <font color=#FFA500>5</font> | <font color=#FFA500>绛雨</font>   |
| <font color=#B288DD>1036</font> | <font color=#B288DD>4</font> | <font color=#B288DD>科谢尼娅</font> |
| <font color=#FFA500>1037</font> | <font color=#FFA500>5</font> | <font color=#FFA500>乌尔丽德</font> |
| <font color=#B288DD>1038</font> | <font color=#B288DD>4</font> | <font color=#B288DD>莉塔拉</font>  |
| <font color=#FFA500>1039</font> | <font color=#FFA500>5</font> | <font color=#FFA500>索米</font>   |
| <font color=#FFA500>1040</font> | <font color=#FFA500>5</font> | <font color=#FFA500>波波沙</font>  |
| <font color=#B288DD>1041</font> | <font color=#B288DD>4</font> | <font color=#B288DD>洛塔</font>   |
| <font color=#FFA500>1050</font> | <font color=#FFA500>5</font> | <font color=#FFA500>朝晖</font>   |
| <font color=#FFA500>1052</font> | <font color=#FFA500>5</font> | <font color=#FFA500>可露凯</font>  |

### 武器信息

| ItemId                           | Rank                         | 备注                                      |
|----------------------------------|------------------------------|-----------------------------------------|
| <font color=#FFA500>10001</font> | <font color=#FFA500>5</font> | <font color=#FFA500>喧闹恶灵</font>         |
| <font color=#FFA500>10002</font> | <font color=#FFA500>5</font> | <font color=#FFA500>绝密手稿</font>         |
| <font color=#FFA500>10003</font> | <font color=#FFA500>5</font> | <font color=#FFA500>阿尔克纳</font>         |
| <font color=#FFA500>10004</font> | <font color=#FFA500>5</font> | <font color=#FFA500>盖尔诺</font>          |
| <font color=#FFA500>10005</font> | <font color=#FFA500>5</font> | <font color=#FFA500>王冠鹿角兔</font>        |
| <font color=#FFA500>10006</font> | <font color=#FFA500>5</font> | <font color=#FFA500>妙尔尼尔</font>         |
| <font color=#FFA500>10007</font> | <font color=#FFA500>5</font> | <font color=#FFA500>远行游鸽</font>         |
| <font color=#87CEEB>10131</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式通用冲锋枪9</font>     |
| <font color=#B288DD>10132</font> | <font color=#B288DD>4</font> | <font color=#B288DD>通用冲锋枪9</font>       |
| <font color=#FFA500>10133</font> | <font color=#FFA500>5</font> | <font color=#FFA500>幼狮</font>           |
| <font color=#87CEEB>10231</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式科夫罗夫</font>       |
| <font color=#B288DD>10232</font> | <font color=#B288DD>4</font> | <font color=#B288DD>科夫罗夫</font>         |
| <font color=#FFA500>10233</font> | <font color=#FFA500>5</font> | <font color=#FFA500>传颂之诗</font>         |
| <font color=#87CEEB>10331</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式莫辛-纳甘</font>      |
| <font color=#B288DD>10332</font> | <font color=#B288DD>4</font> | <font color=#B288DD>莫辛-纳甘</font>        |
| <font color=#FFA500>10333</font> | <font color=#FFA500>5</font> | <font color=#FFA500>斯摩希克</font>         |
| <font color=#87CEEB>10341</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式瓦尔特2000</font>    |
| <font color=#B288DD>10342</font> | <font color=#B288DD>4</font> | <font color=#B288DD>瓦尔特2000</font>      |
| <font color=#FFA500>10343</font> | <font color=#FFA500>5</font> | <font color=#FFA500>苦涩焦糖</font>         |
| <font color=#87CEEB>10351</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式九七式</font>        |
| <font color=#B288DD>10352</font> | <font color=#B288DD>4</font> | <font color=#B288DD>九七式</font>          |
| <font color=#FFA500>10353</font> | <font color=#FFA500>5</font> | <font color=#FFA500>跃虎</font>           |
| <font color=#87CEEB>10361</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式斯捷奇金</font>       |
| <font color=#B288DD>10362</font> | <font color=#B288DD>4</font> | <font color=#B288DD>斯捷奇金</font>         |
| <font color=#87CEEB>10371</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式羽锋</font>         |
| <font color=#B288DD>10372</font> | <font color=#B288DD>4</font> | <font color=#B288DD>羽锋</font>           |
| <font color=#FFA500>10373</font> | <font color=#FFA500>5</font> | <font color=#FFA500>流羽白英</font>         |
| <font color=#87CEEB>10381</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式加利尔轻机枪</font>     |
| <font color=#B288DD>10382</font> | <font color=#B288DD>4</font> | <font color=#B288DD>加利尔轻机枪</font>       |
| <font color=#87CEEB>10391</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式索米</font>         |
| <font color=#B288DD>10392</font> | <font color=#B288DD>4</font> | <font color=#B288DD>索米</font>           |
| <font color=#FFA500>10393</font> | <font color=#FFA500>5</font> | <font color=#FFA500>未言使命</font>         |
| <font color=#87CEEB>10401</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式波波沙冲锋枪</font>     |
| <font color=#B288DD>10402</font> | <font color=#B288DD>4</font> | <font color=#B288DD>波波沙冲锋枪</font>       |
| <font color=#FFA500>10403</font> | <font color=#FFA500>5</font> | <font color=#FFA500>斯瓦罗格</font>         |
| <font color=#87CEEB>10411</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式超级90</font>       |
| <font color=#B288DD>10412</font> | <font color=#B288DD>4</font> | <font color=#B288DD>超级90</font>         |
| <font color=#87CEEB>10501</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式长风零六</font>       |
| <font color=#B288DD>10502</font> | <font color=#B288DD>4</font> | <font color=#B288DD>长风零六</font>         |
| <font color=#FFA500>10503</font> | <font color=#FFA500>5</font> | <font color=#FFA500>不留行</font>          |
| <font color=#87CEEB>10521</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式黑克勒科赫416</font>   |
| <font color=#B288DD>10522</font> | <font color=#B288DD>4</font> | <font color=#B288DD>黑克勒科赫416</font>     |
| <font color=#FFA500>10523</font> | <font color=#FFA500>5</font> | <font color=#FFA500>斯库拉</font>          |
| <font color=#B288DD>11007</font> | <font color=#B288DD>4</font> | <font color=#B288DD>野兔</font>           |
| <font color=#87CEEB>11008</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式金牛座曲线</font>      |
| <font color=#87CEEB>11009</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式复仇女神</font>       |
| <font color=#87CEEB>11010</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式格洛利娅</font>       |
| <font color=#B288DD>11014</font> | <font color=#B288DD>4</font> | <font color=#B288DD>复仇女神</font>         |
| <font color=#B288DD>11015</font> | <font color=#B288DD>4</font> | <font color=#B288DD>金牛座曲线</font>        |
| <font color=#FFA500>11016</font> | <font color=#FFA500>5</font> | <font color=#FFA500>猎心者</font>          |
| <font color=#87CEEB>11017</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式莫洛12</font>       |
| <font color=#FFA500>11020</font> | <font color=#FFA500>5</font> | <font color=#FFA500>光学幻境</font>         |
| <font color=#B288DD>11021</font> | <font color=#B288DD>4</font> | <font color=#B288DD>莫洛12</font>         |
| <font color=#87CEEB>11022</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式野兔</font>         |
| <font color=#B288DD>11023</font> | <font color=#B288DD>4</font> | <font color=#B288DD>格洛利娅</font>         |
| <font color=#87CEEB>11024</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式佩切涅</font>        |
| <font color=#B288DD>11026</font> | <font color=#B288DD>4</font> | <font color=#B288DD>佩切涅</font>          |
| <font color=#87CEEB>11030</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式罗宾逊先进步枪</font>    |
| <font color=#B288DD>11031</font> | <font color=#B288DD>4</font> | <font color=#B288DD>罗宾逊先进步枪</font>      |
| <font color=#87CEEB>11036</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式卡拉什-阿尔法</font>    |
| <font color=#B288DD>11037</font> | <font color=#B288DD>4</font> | <font color=#B288DD>卡拉什-阿尔法</font>      |
| <font color=#FFA500>11038</font> | <font color=#FFA500>5</font> | <font color=#FFA500>游星</font>           |
| <font color=#87CEEB>11039</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式黑科赫7</font>       |
| <font color=#B288DD>11040</font> | <font color=#B288DD>4</font> | <font color=#B288DD>黑科赫7</font>         |
| <font color=#87CEEB>11042</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式一九一式</font>       |
| <font color=#B288DD>11043</font> | <font color=#B288DD>4</font> | <font color=#B288DD>一九一式</font>         |
| <font color=#FFA500>11044</font> | <font color=#FFA500>5</font> | <font color=#FFA500>金石奏</font>          |
| <font color=#87CEEB>11045</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式特殊用途自动型霰弹枪</font> |
| <font color=#B288DD>11046</font> | <font color=#B288DD>4</font> | <font color=#B288DD>特殊用途自动型霰弹枪</font>   |
| <font color=#FFA500>11047</font> | <font color=#FFA500>5</font> | <font color=#FFA500>梅扎露娜</font>         |
| <font color=#87CEEB>11048</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式纳甘左轮</font>       |
| <font color=#B288DD>11049</font> | <font color=#B288DD>4</font> | <font color=#B288DD>纳甘左轮</font>         |
| <font color=#87CEEB>11051</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式九五式</font>        |
| <font color=#B288DD>11052</font> | <font color=#B288DD>4</font> | <font color=#B288DD>九五式</font>          |
| <font color=#FFA500>11053</font> | <font color=#FFA500>5</font> | <font color=#FFA500>重弦</font>           |
| <font color=#87CEEB>11054</font> | <font color=#87CEEB>3</font> | <font color=#87CEEB>旧式格威尔36</font>      |
| <font color=#B288DD>11055</font> | <font color=#B288DD>4</font> | <font color=#B288DD>格威尔36</font>        |
| <font color=#FFA500>11056</font> | <font color=#FFA500>5</font> | <font color=#FFA500>女仆准则</font>         |