# kigacli

## 安装命令行
初始化，并构建
```bash
go mod init kiga-tool
go build
```


使用viper配置相关服务，此版本为测试案例

后续补充其他工具包

## Direction for Mysql Use

使用viper配置Mysql相关信息

Database operation

- CreateDataBase
- DropDataBase
- CheckDataBase


## Redis

同上，通过viper读取toml配置

- [x] Redis连接池
- [x] key 存储上限，超时时间
- [x] 查找所有符合给定模式 pattern
- [x] 获取指定 key 的值
- [x] redis获取对象
- [x] 存储数据
- [x] 列表操作
- [x] 元素操作
- [x] 删除键值
- [x] 锁操作

## File

文件操作

## wav

音频类型文件使用方法

## ZIP for data

数据或文件的压缩
