# 点餐系统服务端

这是一个基于 Go 语言和 Gin 框架开发的点餐系统后端服务。

## 项目结构

```
orderFood-server-cus/
├── cmd/                 # 主程序入口
│   └── main.go          # 主程序文件
│
├── common/              # 公共代码目录
│   ├── db/              # 数据库相关
│   │   ├── database.go  # 数据库连接管理
│   │   ├── instance.go  # 数据库实例管理
│   │   └── migrate.go   # 数据库迁移
│   │
│   ├── middleware/      # 中间件
│   │   ├── cors.go      # 跨域中间件
│   │   └── auth.go      # 认证中间件
│   │
│   └── utils/           # 工具函数
│       └── env.go       # 环境变量工具
│
├── pkg/                 # 按功能模块划分的包
│   ├── account/         # 账户模块
│   │   ├── account.model.go     # 数据模型
│   │   ├── account.service.go   # 业务逻辑
│   │   └── account.router.go    # 路由处理
│   │
│   ├── category/        # 分类模块
│   │   ├── category.model.go
│   │   ├── category.service.go
│   │   └── category.router.go
│   │
│   ├── dish/            # 菜品模块
│   │   ├── dish.model.go
│   │   ├── dish.service.go
│   │   └── dish.router.go
│   │
│   ├── order/           # 订单模块
│   │   ├── order.model.go
│   │   ├── order.service.go
│   │   └── order.router.go
│   │
│   └── token/           # 认证模块
│       ├── token.model.go
│       ├── token.service.go
│       └── token.router.go
│
├── .env                 # 环境变量配置
└── .env.example         # 环境变量示例配置
```

## 设计思路

项目采用模块化结构设计，每个功能模块负责自己的领域：

1. **common 目录** - 包含所有模块共享的基础设施代码

   - 数据库连接和管理
   - 中间件
   - 工具函数

2. **pkg 目录** - 包含按功能划分的业务模块
   - 每个模块都有自己的模型(model)、服务层(service)和路由(router)

## 环境要求

- Go 1.16+
- PostgreSQL
- 设置正确的环境变量

## 如何运行

1. 克隆仓库
2. 复制`.env.example`为`.env`并填写正确的配置
3. 运行`go run cmd/main.go`

## API 文档

- 账户管理: `/api/v1/accounts`
- 认证: `/api/v1/auth`
- 分类管理: `/api/v1/categories`
- 菜品管理: `/api/v1/dishes`
- 订单管理: `/api/v1/orders`
