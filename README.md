# 代码生成器

## 描述

把 MySQL DDL 转换成模板代码

Module:

config: 配置文件处理
ddl: 加载表相关信息
path: 路径处理
template: 模板处理
rules: 类型转换与命名处理

## Installation and Upgrade

```
go get -u github.com/go-sql-driver/mysql

go get -u github.com/magiconair/properties

go get -u github.com/iunsuccessful/generator

```

## 使用说明

安装好 golang 环境，下载好代码之后，在 generator 目录下运行 `go build`，或者直接下载编译好的 `generator.exe`在 test 文件夹，设置到 `path` 或者放到模板目录即可

### 第一步，配置 config.properties

```
jdbc.url = jdbc:mysql://demo.mysql.rds.aliyuncs.com/demo
jdbc.username = demo
jdbc.password= demo
```

### 第二步，编写模板

模板使用 golang [模板语法](https://golang.org/pkg/text/template/) 放到 template 目录下，软件会根据 template 下面的目录生成对应目录到 target 下面

#### 可用变量

| 字段 | 类型 | 释意 |
|:--|:--|:--|
| TableName | string | 表名 |
| BeanName  | string | 别名 |
| Name | string | 有别名就是别名，没有就是表名 |
| TableComment | string | 表注释 |
| Columns | []ColumnInfo | 列字段 |

| 字段 | 类型 | 释意 |
|:--|:--|:--|
| ColumnName| string | 列名 |
| JDBCType | string | 列在数据库中的类型 |
| StdJDBCType | string | VARCHAR, INTEGER |
| DataType | string | java 中的类型 |
| ColumnComment | string | 列有注释 |
| ColumnKey | string | 可判断是不是主键 “PRI” 为主键 |
| UpperCamelName | string | 列名大驼峰命名 |
| LowerCamelName | string | 列名小驼峰命名 |

#### 可用方法

| 方法名 | 参数 | 释意 |
|:--|:--|:--|
| lowerCamelCase| string | 字符转换为小驼峰命名 |
