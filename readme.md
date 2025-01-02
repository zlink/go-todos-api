# Golang API Project Demo

这是一个基于 Golang + Gin 框架开发的 API 项目脚手架。

## 项目特性

- 使用 Gin 框架作为 Web 服务器
- 使用 GORM 作为 ORM 框架
- JWT 认证
- 统一的响应处理
- 配置文件管理
- 日志中间件
- 数据库连接池

## 项目结构

```
├── app # 应用目录
│ ├── bootstrap.go # 应用引导文件
│ ├── handler # 控制器目录
│ ├── middleware # 中间件目录
│ ├── models # 数据模型
│ └── service # 服务类目录
├── config # 配置文件目录
│ ├── config.ini # 生产环境配置
│ └── config_dev.ini # 开发环境配置
├── docs # 文档目录
├── routes # 路由目录
├── runtime # 运行时目录
└── main.go # 入口文件
```

## 配置说明

项目支持多环境配置文件:

- config.ini - 生产环境配置
- config_dev.ini - 开发环境配置

主要配置项包括:

- app: 应用基础配置
- server: HTTP 服务器配置  
- mysql: 数据库配置
- redis: Redis 配置(可选)

## 快速开始

1. 克隆项目
2. 配置 config/config.ini
3. 执行数据库迁移(如果需要)
4. 运行项目:
  make dev # 开发环境
  make prod # 生产环境

## API 文档

主要接口:

- POST /auth/login - 用户登录
- GET /user - 用户信息
- GET /ping - 服务器状态检查
