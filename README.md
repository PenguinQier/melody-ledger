# MelodyLedger

MelodyLedger 是一个用 Go 写的个人博客系统。它不只是简单展示文章，还支持注册登录、文章发布与编辑、评论、搜索、封面上传，以及带有音乐/媒体内容的文章记录。

这个项目更像是一个可以慢慢生长的个人内容空间：写文章、放作品、存喜欢的音乐，也把学习 Go Web 开发时搭出来的后端结构整理成一个完整应用。

## 功能特性

- 首页文章展示，区分精选内容和故事列表
- 用户注册、登录、退出和个人资料修改
- 文章创建、编辑、删除、详情展示
- 文章封面上传
- 登录用户评论文章
- 文章关键词搜索
- `/api/search` JSON 搜索接口
- 支持文章附件和媒体文件上传
- 基于 MySQL 的数据持久化
- 提供迁移和示例数据填充命令
- 使用 Session 做登录状态管理
- 静态资源、页面模板和业务模块分层组织

## 技术栈

- Go 1.23+
- Gin
- GORM
- MySQL
- Cobra
- Viper
- Gin Sessions
- Bootstrap / Mediumish theme

## 项目结构

```text
.
|-- assets/              # CSS、JS、图片等静态资源
|-- cmd/                 # 命令行入口：serve、migrate、seed
|-- config/              # 应用配置
|-- database/            # 额外的数据迁移定义
|-- internal/            # 业务模块、中间件、路由、模板
|-- pkg/                 # 可复用基础设施封装
|-- public/              # 对外访问的运行时文件目录
|-- go.mod
|-- main.go
`-- README.md
```

## 快速开始

### 1. 创建数据库

先创建一个 MySQL 数据库：

```sql
CREATE DATABASE melody_ledger CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

如果你本地有自己的数据库备份，也可以手动导入：

```bash
mysql -u root -p melody_ledger < your_dump.sql
```

### 2. 修改配置

数据库和服务配置在 `config/config.yaml`：

```yaml
app:
  name: "MelodyLedger"

server:
  host: "localhost"
  port: "8080"

db:
  username: "root"
  password: "123456"
  host: "127.0.0.1"
  port: "3306"
  name: "melody_ledger"
```

如果你的 MySQL 用户名、密码或端口不同，先改这里。

### 3. 安装依赖

```bash
go mod download
```

### 4. 运行迁移

```bash
go run . migrate
```

### 5. 填充示例数据

```bash
go run . seed
```

### 6. 启动项目

```bash
go run . serve
```

浏览器打开：

```text
http://localhost:8080
```

## 常用命令

```bash
go run . serve     # 启动 Web 服务
go run . migrate   # 执行数据库迁移
go run . seed      # 写入示例数据
```

也可以先构建二进制文件：

```bash
go build -o melody-ledger .
./melody-ledger serve
```

Windows PowerShell：

```powershell
go build -o melody-ledger.exe .
.\melody-ledger.exe serve
```

## 路由说明

| Method | Path | 说明 |
| --- | --- | --- |
| GET | `/` | 首页 |
| GET | `/register` | 注册页 |
| POST | `/register` | 创建用户 |
| GET | `/login` | 登录页 |
| POST | `/login` | 登录 |
| POST | `/logout` | 退出登录 |
| GET | `/profile` | 个人资料页 |
| POST | `/profile` | 更新个人资料 |
| GET | `/articles/:id` | 文章详情 |
| GET | `/articles/create` | 创建文章页 |
| POST | `/articles/store` | 保存文章 |
| GET | `/articles/:id/edit` | 编辑文章页 |
| POST | `/articles/:id/update` | 更新文章 |
| POST | `/articles/:id/delete` | 删除文章 |
| POST | `/articles/:id/cover` | 更新文章封面 |
| POST | `/comments` | 发表评论 |
| POST | `/comments/:id/delete` | 删除评论 |
| GET | `/search` | 搜索页 |
| GET | `/api/search` | 搜索接口 |

## 本地媒体文件

运行时上传文件会放在 `public/uploads/`。仓库会保留这个目录，但不会提交其中的本地媒体文件。

```text
public/uploads/
```

`music/`、`public/uploads/` 里的音频、图片、表格等文件，以及本地 SQL 备份都会被 `.gitignore` 忽略。这样可以避免 public 仓库里出现体积很大的媒体资源、版权音频或包含个人账号信息的数据库数据。

## 开发检查

```bash
go test ./...
go build ./...
```

## 后续计划

- 支持环境变量覆盖配置
- 增加 Docker Compose 本地开发环境
- 增加 Markdown 或富文本编辑器
- 给文章和评论增加分页
- 增加后台管理页面
- 优化上传文件类型和大小校验
