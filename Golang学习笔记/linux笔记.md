# 目录结构

- linux只有一个根目录`/`
- root : 该目录为系统管理员目录，root是具有超级权限的用户。
- ==bin ->usr/bin== : 存放系统预装的可执行程序，这里存放的可执行文件可以在系统的任何目录下执行。
- usr是linux的系统资源目录，里边存放的都是一些系统可执行文件或者系统以来的一些文件库。
- ==usr/local/bin==：存放用户自己的可执行文件，同样这里存放的可执行文件可以在系统的任何目录下执行。
- lib->usr/lib: 这个目录存放着系统最基本的动态连接共享库，其作用类似于Windows里的DLL文件，几乎所有的应用程序都需要用到这些共享库。
- boot : 这个目录存放启动Linux时使用的一些核心文件，包括一些连接文件以及镜像文件。
- dev: dev是Device(设备)的缩写, 该目录下存放的是Linux的外部设备，Linux中的设备也是以文件的形式存在。
- ==etc==: 这个目录存放所有的系统管理所需要的配置文件（/etc/profile配置环境变量）
- ==home==：用户的主目录，在Linux中，每个用户都有一个自己的目录，一般该目录名以用户的账号命名，叫作用户的根目录；用户登录以后，默认打开自己的根目录。
- var : 这个目录存放着在不断扩充着的东西，我们习惯将那些经常被修改的文件存放在该目录下，比如运行的各种日志文件。
- mnt : 系统提供该目录是为了让用户临时挂载别的文件系统，我们可以将光驱挂载在/mnt/上，然后进入该目录就可以查看光驱里的内容
- ==opt==: 这是给linux额外安装软件所存放的目录。比如你安装一个Oracle数据库则就可以放到这个目录下，默认为空。
- tmp: 这个目录是用来存放一些临时文件的。

# vim快捷键

- 一般模式拷贝当前行(yy) , 拷贝当前行向下的5行(5yy)，并粘贴(p)

- 一般模式删除当前行(dd) , 删除当前行向下的5行(5dd)。

- **查找**：一般模式下，在文件中查找某个单词，[命令模式下：(/关键字)，回车查找, 输入(n) 就是查找下一个]

  - 取消高亮：命令模式下，输入:nohlsearch，当然，可以简写，noh

- **替换**

  - `:s/str1/str2/`：替换当前行第一个str1为str2
  - `:s/str1/str2/`g：替换当前行所有str1为str2
  - `:n,$s/str1/str2/`：替换第 n 行开始到最后一行中，每一行的第一个str1为str2
  - `:n,$s/str1/str2/g`：替换第 n 行开始到最后一行中，每一行的所有str1为str2
  - `:%s/str1/str2/`：替换每一行的第一个str1为str2
  - `:%s/str1/str2/g`：替换每一行的所有str1为str2

- 一般模式下，使用快捷键到达文档的最首行[gg]和最末行[G]

- 一般模式下，在一个文件中输入"xxxx" ,然后又撤销这个动作(u)

- 一般模式下，并将光标移动到10行，输入10，输入shift+g，回车

- 命令行模式下，设置文件的行号，取消文件的行号.[命令行下(: set nu) 和(:set nonu)]

- > 如果想让vim永久显示行号，则需要修改vim配置文件vimrc。如果没有此文件可以创建一个。在启动vim时，当前用户根目录下的vimrc文件会被自动读取，因此一般在当前用户的根目录下创建vimrc文件，即使用下面的命令：
  >
  > ![这里写图片描述](https://img-blog.csdn.net/20180111155902419?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvZWxlY3Ryb2NyYXp5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)
  >
  > 在打开的vimrc文件中最后一行输入：set number ，然后保存退出。再次用vim打开文件时，就会显示行号了。
  >
  > ![这里写图片描述](https://img-blog.csdn.net/20180111155928196?watermark/2/text/aHR0cDovL2Jsb2cuY3Nkbi5uZXQvZWxlY3Ryb2NyYXp5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70/gravity/SouthEast)

# 用户

- 创建用户`useradd <username>`：创建了用户，在/home下创建了和用户名相同的用户根目录，创建了一个和用户名相同的组
- 创建用户的同时指定根目录`useradd -d /home/xxx <username>`
- 设置用户密码`passwd <username>`
- 删除用户`userdel <username>`；删除用户的同时删除用户根目录`userdel -r <username>`
- 查看用户的组信息`id <username>`，省去用户名的话就是查看当前用户的信息；`gid`是主组，`groups`是附加组
- 切换用户`su <username>`

# 组

> 相当于角色的概念，对有共性的用户进行统一的权限管理
>
> 任何用户都至少属于一个组，不能独立于组存在，可以属于多个组；文件或目录只能属于一个组，可以通过组来控制权限
>
> 文件把linux的用户分为三类
>
> - 所有者：默认情况下是创建者，可修改
> - 同组用户：和文件、目录属于同一个组的用户
> - 其他组用户：除了以上两种的用户
>
> 新建用户时如果不指定组，会自动新建一个组，组名和用户名相同，并且把该用户添加到组中
>
> 创建用户后，主组不能变（主组没什么用），附加组可以删除、添加

- 创建组`groupadd <groupname>`
- 删除组`groupdel <groupname>`
- 把用户加入组`gpasswd -a <username> <groupname>`
- 把用户从组移除`gpasswd -d <username> <groupname>`
- 创建用户的时候指定主组`useradd -g <groupname> <username>`
- 值修改文件或目录的所有者`chown newOwner file`
  - 文件的所有者和组可以不一致；默认目录中的子目录和文件的所有者、组不会跟着改
  - 同时修改所有者和组`chown newOwner:newGroup file`
  - `chown -R newOwner:newGroup file`所有者和组，目录中的子目录和文件也会跟着改
- 只修改文件或目录的组`chgrp newGroup file`，也可以加`-R`来递归修改

# 权限

> 文件或目录的三种权限：读（Read）、写（Write）、执行（Execute）（对目录而言，进入目录就是执行）
>
> - 读：cat、more 、less、head 、tail；ls（分号分隔的是文件和目录权限的相关命令）
> - 写：vi/vim；mkdir、mv、touch、rm、cp；chmod对于文件和目录都是写权限
> - 执行：./xxx；cd

- `ll`查看后的权限，从第二个字符开始
  - 第一部分的三个是所有者的权限；剩下的就是同组用户、其他组用户的权限
  - 第一个字符，`d`表示目录，`-`表示文件，`l`表示链接文件，`d`表示可随机存取的设备，如U盘等，`c`表示一次性读取设备，如鼠标、键盘等

- 任何文件或目录都有三部分权限：（和上面的三种用户对应）
  - 所有者权限：r、w、x分别表示所有者对文件或目录的权限，如`rwx`表示所有者拥有读、执行的权限；`r-x`读和执行；`w--`只能写；`---`没有任何权限
  - 同组用户权限：同上
  - 其他组用户权限：同上
- `chmod`修改权限
  - `u-w,u+r`去掉所有者的写权限，加上写权限
  - `g+w`给同组用户加上写权限
  - `o-w`给其他组用户减去写权限
  - `a=rw`给所有用户的权限设置成读写
  - 经测试，阿里云的CentOS7，root用户有所有文件的最高权限
- 用数字修改权限
  - r、w、x分别用4、2、1表示（2的n次方）
  - 三部分权限分别用数字之和来表示，如`766`表示所有者的权限是`rwx`，同组、其他组用户的权限是读写
  - `chmod 777 file`

# 网络（知道配置文件在哪即可）

> /etc/sysconfig/network-scripts/
>
> 修改配置文件，然后重启linux（阿里云暂时不管）
>
> DNS要改成DNS1
>
> ![image-20210509210739421](D:\资料\Go\src\studygo\Golang学习笔记\linux笔记.assets\image-20210509210739421.png)

# 进程

> 线程：一个程序的执行线路
>
> 进程：一个程序的执行，每个进程会占一个端口

- `ps`值显示应用级别的进程状态（用户用到的应用）（process status）

  - `ps -e`查看所有进程，`ps -ef`以全格式显示所有进程，`grep --color=auto redis`是grep本身的进程

    > ![image-20210509213040886](D:\资料\Go\src\studygo\Golang学习笔记\linux笔记.assets\image-20210509213040886.png)
    >
    > ![image-20210509213047524](D:\资料\Go\src\studygo\Golang学习笔记\linux笔记.assets\image-20210509213047524.png)

  - `kill pid`杀掉进程；`killall 进程名称(支持通配符)`杀死指定进程名称的所有进程；`kill -9 pid`强迫进程立即停止

# 服务

> 概念：支撑linux运行的，在后台运行的一些必要进程，也称为守护进程（sshd、firewall等），名字一般都以d结尾

- `systemctl [options] 服务名称`，可选参数如下
  - start|stop、restart、reload、status、enable
  - 启停、重启、重新加载配置、查看状态、设置是否开机启动

# 软件包管理

> 一般是软件的安装包

一般来说linux系统基本上分两大类：

- RedHat系列：Redhat、Centos、Fedora等
- Debian系列：Debian、Ubuntu等

==RedHat系列==

安装包格式：RPM包——linux软件包的打包和安装工具，它操作的软件包都是`.rpm`结尾

- `rpm -qa`：查看当前系统用rpm安装了哪些软件包（基本都是系统软件）
- `rpm -e 软件名字`：卸载rpm方式安装的软件包
- `rpm -ivh xxx.rpm`：通过rpm方式安装软件包。-i=install 安装；-v=verbose 提示；-h=hash 进度条（先去官网下载rpm包）

包管理工具：yum——在rpm安装机制上增加了搜索和分析依赖的功能，去yum的服务器上下载并自动安装

- `yum list installed`：查看当前系统用yum方式已经安装的包
- `yum -y installed/remove/update 软件名`：安装/卸载/更新，不手动确认了

==Debian系列==

安装包格式：deb包

- `dpgk -参数`

包管理工具：apt-get

- `apt-get -y install/remove/update 软件名`（曾经是`apt install`，后来用`apt-get`替换`apt`）

# 环境变量

- env，export，或者echo $path,来查看当前环境变量的值
- export PATH=$PATH:/要添加的路径
- export PATH=/要添加的路径$PATH
- 修改/ect/profile后，`source ~/.bashrc`，`source ~/.profile`

# 日志(待补充)

- journal日志
- logrotate
- Rsyslog

# 常用命令

> 所有的文件名、目录名都有相对和绝对两种写法

- 开关机
  - `shutdown now`立即关机
  - `shutdown -h xxx`定时关机
  - `shutdown -r now`立即重启
  - `reboot`立即重启
  - `sync`同步数据库
  
- `ls -a`把当前目录下的文件和目录（包含隐藏、虚拟文件和目录）查看出来

- `man <command>`，`help <command>`，`<command> --help`

- `uname -r`查看内核版本

- `lsb_release -a`查看发行版本

- `ifconfig`查看ip地址

- `mkdir -p`可以创建多级目录；目录相关的操作（ls、cd等），后面都可以跟绝对、相对目录

- `touch`可以创建一个或多个空文件

- `cp`复制目录时，如果目录非空，需要加`-r`；`cat file1 > file2`复制file1内容到file2并覆盖file2的原内容

- `mv`不需要`-r`，能直接移目录；`mv file1 file2`，就是重命名，file2不在的话会新建

- `cat`文件内容全显示出来，`cat -n`显示行号

- `more`文件内容一次性加载到内存，分页显示，回车下一行，空格下一页（不能查找）

- `less`文件内容分页加载到内存

  - 可以查找
  - `空格`或`pagedown`下一页，`b`或`pageup`上一页；`d`下半页，`u`上半页；`g`第一行，`G`最后一行；`q`退出
  - `v`进入编辑

- `head`查看文件的头10行，加`-n 3`则是显示头3行

- `tail`查看文件的后10行，其余类似

- `echo`输出系统变量或某常量的值，`$`表示引用某变量/常量（`echo $path`）；输出字符串到控制台`echo hello`

- `>`把一个结果输出到指定文件，`command > file1`，file1不存在时会新建

- `>>`追加输出内容到文件

- `date`查看时间，后跟`+%Y`之类的则按指定格式显示时间（`date '+%Y/%m//%d %H-%M-%S'`），不带单引号也可，但不能有空格。跟`-s`设置当前的系统时间，设置时间后恢复方法如下

  ```bash
  #有用的是第二句，但是如果你开着ntp服务，那么需要先关闭下，完整命令如下
  service ntpd stop
  ntpdate ntp.api.bz
  service ntpd start
  ```

- `cal`查看当前月份的日历，跟`2021`查看2021年的日历

- 搜索

  - `find [目录] 关键字  `，关键字必须带通配符`*`、`?`等（任意多个字符、任意一个字符）（其实是默认用了`-name`）
  - 按文件大小搜索文件`find [目录] -size -5k`，搜索小于5k的文件，大于是`+`
  - 按文件所有者搜索文件`find [目录] -user <username>`

- `locate 关键字`去目录树中查找，效率高，但不常用（因为树不是实时更新的，可以先updatedb更新树后再搜索）

- `command | grep 过滤条件`，过滤命令的输出结果（区分大小写，跟`-i`的话不区分；跟`-n`的话会显示行号）

- 压缩和解压

  - `gzip file`压缩单个文件并删除源文件，`gunzip file`解压单个文件。默认得到的压缩包是`.gz`

  - `zip 目标压缩包名称 文件或目录列表`、`unzip 压缩包 [-d 目标目录]`，压缩、解压多个文件或目录，通常把压缩包设为`.zip`

  - `tar -c 目标压缩包名称 文件或目录列表`打包压缩，通常压缩包设为`.tar.gz`

  - `tar -x 压缩包名称 [-C 目标目录] `解压

  - 固定用法：`tar -zcvf`，`tar -zxvf`

    ![image-20210508183228470](D:\资料\Go\src\studygo\Golang学习笔记\linux笔记.assets\image-20210508183228470.png)
  
- 查看服务器的系统信息

  - `uname -a`
  - `cat /etc/issue`，适用于所有的`Linux`发行版；`cat /etc/redhat-release`，只适用于`Redha`t系列
  - `lsb_release -a`

- 查看服务器的硬件信息(待补充)

# lsof神器

lsof abc.txt 显示开启文件abc.txt的进程

lsof -i :22 知道22端口现在运行什么程序

lsof -c abc 显示abc进程现在打开的文件

lsof -g gid 显示归属gid的进程情况

lsof +d /usr/local/ 显示目录下被进程开启的文件

lsof +D /usr/local/ 同上，但是会搜索目录下的目录，时间较长

lsof -d 4 显示使用fd为4的进程 www.2cto.com

lsof -i 用以显示符合条件的进程情况

语法: lsof -i[46] [protocol][@hostname|hostaddr][:service|port]

46 --> IPv4 or IPv6

protocol --> TCP or UDP

hostname --> Internet host name

hostaddr --> IPv4位置

service --> /etc/service中的 service name (可以不只一个)

port --> 端口号 (可以不只一个)

例子: TCP:25 - TCP and port 25

@1.2.3.4 - Internet IPv4 host address 1.2.3.4

tcp@ohaha.ks.edu.tw:ftp - TCP protocol hosthaha.ks.edu.tw service name:ftp

lsof -n 不将IP转换为hostname，缺省是不加上-n参数

例子: lsof -i tcp@ohaha.ks.edu.tw:ftp -n

lsof -p 12 看进程号为12的进程打开了哪些文件

lsof +|-r [t] 控制lsof不断重复执行，缺省是15s刷新

-r，lsof会永远不断的执行，直到收到中断信号

+r，lsof会一直执行，直到没有档案被显示

例子：不断查看目前ftp连接的情况：lsof -i tcp@ohaha.ks.edu.tw:ftp -r

lsof -s 列出打开文件的大小，如果没有大小，则留下空白

lsof -u username 以UID，列出打开的文件 www.2cto.com

关注：
进程调试命令:truss、strace和ltrace
进程无法启动，软件运行速度突然变慢，程序的"SegmentFault"等等都是让每个Unix系统用户头痛的问题，而这些问题都可以通过使用truss、strace和ltrace这三个常用的调试工具来快速诊断软件的"疑难杂症"。