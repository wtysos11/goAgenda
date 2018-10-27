# Agenda设计文档

## 需求

子命令：help、register、cm
要求用户信息存储在curUser.txt中，User和Meeting实体要用json进行存储。
要求实现日志服务

### agenda help

子命令`agenda help`，可以列出命令说明。此外还要实现`agenda help register`列出regsiter命令的描述

### agenda user

用户相关的子命令

相关参数：
1. --username/-u，用户名
2. -password/-p，密码
3. -email/-e，邮箱
4. -telphone/-t，电话

#### agenda user register

子命令`agenda user register`，注册用户
接受四个参数（必选）：
1. --username/-u，用户名
2. -password/-p，密码
3. -email/-e，邮箱
4. -telphone/-t，电话

此外用户名是唯一的，要查重。
邮箱、电话要检查有效性。

反馈：
* 如果登记成功，要返回成功注册的信息
* 如果登记失败，要反馈错误信息（哪里错了，等等）

#### agenda user login

子命令`agenda user login`，用户登录
1. --username/-u，用户名
2. -password/-p，密码

这两个参数为必须
反馈：
* 用户名与密码都正确的时候返回一个成功登录的信息
* 登录失败，返回登录失败的信息

#### agenda user logout

子命令`agenda user logout`，用户登出
登出后只能够使用用户注册和用户登录功能

不接受参数
反馈：
* 反馈登出信息

#### agenda user lookup

子命令`agenda user lookup`，在创建的用户中查找用户
只有在已经登录的状态才可以使用
反馈：
* 如果找到用户，返回已注册的所有用户的用户名、邮箱以及电话信息（打表）

#### agenda user delete

子命令`agenda user delete`，删除用户
操作为删除自己的账号，自动注销

反馈:
* 返回成功注销的信息
* 失败要返回注销失败（有什么情况会导致注销失败吗？）

用户账号删除之后：
1. 以该用户为发起者的会议将会被删除
2. 以该用户为参与者的会议将会从参与者列表中移除该用户。如果此操作造成参与者数量为0，则会议也会被移除。

#### agenda meeting create

子命令`agenda meeting create`，创建会议
参数：
1. --start/-s，开始时间，格式为(YYYY-MM-DD/HH:mm:ss)
2. --end/-e，结束时间
3. --title/-t，会议主题
4. --participant/-p,会议参与者

反馈：
* 成功则返回适当信息
* 失败也返回相关信息

注意：用户无法分身参与多个狐疑，如果用户已有的会议安排与将要创建的会议在时间上有重叠，则会无法创建这个会议。
会议主题是唯一的。

#### agenda meeting addUser

添加会议参与者
参数：
1. --title/-t，会议主题
2. --username/-u，要添加的用户名（可以为多个）

反馈：
* 成功返回相关信息
* 失败也返回相关信息

#### agenda meeting deleteUser

删除会议参与者
同上

### agenda meeting lookup

查找会议，仅登录可用
参数：
1. --start/-s，开始时间
2. --end/-e，结束时间

已登录的用户可以查询自己的议程在某一时间段内的所有会议安排

#### agenda meeting cancel

取消会议，已经登录的用户可以取消自己发起的某一会议安排
参数：
1. --title/-t，会议主题

#### agenda meeting exit

退出会议，已经登录的用户可以退出自己参与的某一会议安排
参数：
1. --title/-t，会议主题

如果此操作造成会议参与者人数为0，则会议将会被删除

#### agenda meeting clear

清空会议，已经登录的用户可以清空自己发起的所有会议安排