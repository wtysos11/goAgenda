# Agenda设计文档

## 需求

子命令：help、register、cm
要求用户信息存储在curUser.txt中，User和Meeting实体要用json进行存储。
要求实现日志服务

### agenda help

子命令`agenda help`，可以列出命令说明。此外还要实现`agenda help register`列出regsiter命令的描述

### agenda register

子命令`agenda register`，注册用户
接受三个参数（必选）：
1. --username/-u，用户名
2. -password，密码
3. -email，邮箱

此外用户名要有查重功能，密码、邮箱要检查有效性。

### agenda cm

子命令`agenda cm`，创建会议
待讨论

