<!-- TEMPLATE (ZH) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/gormmom/release.yml?branch=main&label=BUILD)](https://github.com/yylego/gormmom/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/gormmom)](https://pkg.go.dev/github.com/yylego/gormmom)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/gormmom/main.svg)](https://coveralls.io/github/yylego/gormmom?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.27%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/gormmom.svg)](https://github.com/yylego/gormmom/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/gormmom)](https://goreportcard.com/report/github.com/yylego/gormmom)

<!-- TEMPLATE (ZH) CLOSE: BADGES -->

# 🌍 GORMMOM - GORM 原生语言命名把关器

**gormmom** 让你用 **母语**（中文、日文、韩文等）编写 GORM 模型，然后**把关**模型产出的每一个数据库名字——表名、每个列名、每个索引名——都是数据库安全的。它作为一个普通单元测试运行，**完全不连接数据库**。

> 🎯 **用母语编程，交付数据库安全的表结构。** gormmom 在母语命名进入数据库之前，就替你抓出数据库会拒绝的那些名字。

---

## 🤖 为什么 gormmom 现在只是个把关器（AI 时代的简化）

gormmom 过去是个**生成引擎**：你用母语写字段、把 tag 留空，它替你自动拼出列名和索引名（中文拼音、日文罗马音等等）。

**AI 时代让这种生成失去了意义。** 今天，你定义字段的那一刻，AI 助手就能直接、结合上下文，替你写出一个又好又数据库安全的列名——远比任何固定规则引擎拼出来的强。机器生成的 `z_zhang_hao` 从来就不是谁真正想要的。于是整个生成层就没有存在的理由了。

而 AI **不会**自动替你保证的，恰恰是那个仍然会咬你一口的环节：**校验**。当你用母语编程时，GORM 可能把字段变成一个非 ASCII 的列名，或者用母语字段名拼出一个复合索引名——这两种数据库都会拒绝。这种问题暴露得很晚，往往是在迁移或运行时才炸。

所以 gormmom 彻底砍掉了生成，现在只做一件事：**在一个不连数据库的单元测试里，把这些命名问题抓出来**，让它们不会在后面才爆炸。

- ❌ 已移除：代码生成、拼音/罗马音列名拼装、模式枚举、旧的生成命令。
- ✅ 已保留：唯一职责——校验表名 / 列名 / 索引名是否数据库安全。
- ✅ 零数据库连接——一个纯静态检查，作为普通单元测试运行。

---

## 生态系统

![GORM Type-Safe Ecosystem](https://github.com/yylego/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)

<!-- TEMPLATE (ZH) CLOSE: LANGUAGE NAVIGATION -->

---

## 🚀 安装

```bash
go get github.com/yylego/gormmom
```

---

## 🌟 问题与解决方案

### ⚡ 母语模型的陷阱

母语字段读起来很自然，但 GORM 从中推导出的名字，可能正是数据库拒绝的名字：

```go
type T账户信息 struct {
    ID   uint   `gorm:"primaryKey"`
    Z账号 string `gorm:"uniqueIndex"` // ❌ 没写索引名 → GORM 用字段名拼索引 → 非 ASCII → 被拒绝
    N昵称 string `gorm:""`            // ❌ 没写列名 → 列名落到非 ASCII → 被拒绝
}
```

### ✅ gormmom 的做法

保留母语字段名，但显式给出数据库安全的列名 / 索引名 / 表名——然后让 gormmom 在单元测试里替你把关：

```go
type T账户信息 struct {
    ID   uint   `gorm:"primaryKey"`
    Z账号 string `gorm:"column:account;uniqueIndex:udx_account"`
    N昵称 string `gorm:"column:nickname;index:idx_nickname"`
    A年龄 int    `gorm:"column:age"`
    D电话 string `gorm:"column:phone"`
    E邮箱 string `gorm:"column:mailbox"`
    J住址 string `gorm:"column:address"`
    S状态 string `gorm:"column:status;index:idx_status"`
}

func (*T账户信息) TableName() string {
    return "accounts" // 数据库安全的 ASCII 表名
}
```

---

## 🛠️ 使用方法 —— 在单元测试里校验

gormmom 从不连接数据库，它作为一个普通单元测试运行：

```go
func TestValidateNaming(t *testing.T) {
    objects := []any{
        &T账户信息{},
    }

    // 直接从模型组装校验配置（纯反射，不读取任何磁盘文件）
    configs := gormmom.NewConfigs(objects, gormmom.NewOptions())

    // 把关：表名 / 列名 / 索引名都必须数据库安全
    require.NoError(t, configs.Validate())
}
```

规则就是一条正则：每个名字都要匹配 `^[a-zA-Z0-9_]{1,N}$`（N 默认 63，可通过 `Options.MaxNameLength` 调整）。校验失败会返回一个清晰的错误，指出是哪个结构体、哪个名字出了问题（涉及具体字段时也会一并指出）——远在任何迁移发生之前就被抓住。

默认情况下 gormmom 用 GORM 的默认命名策略来推导名字。如果你的项目用了自定义命名策略，请通过 `Options.NamingStrategy` 传进来，让校验结果和数据库实际看到的一致。

---

## 🌍 多语言示例

每种语言的模式都一样：母语字段名，加上显式的、数据库安全的列名 / 索引名 / 表名，由 gormmom 校验。

### 繁體中文

```go
type T賬戶信息 struct {
    ID   uint   `gorm:"primaryKey"`
    Z賬號 string `gorm:"column:account;uniqueIndex:udx_account"`
    N暱稱 string `gorm:"column:nickname;index:idx_nickname"`
    A年齡 int    `gorm:"column:age"`
    S狀態 string `gorm:"column:status;index:idx_status"`
}

func (*T賬戶信息) TableName() string {
    return "accounts"
}
```

### 日本語

```go
type Tアカウント情報 struct {
    ID         uint   `gorm:"primaryKey"`
    Aアカウント  string `gorm:"column:account;uniqueIndex:udx_account"`
    Nニックネーム string `gorm:"column:nickname;index:idx_nickname"`
    N年齢       int    `gorm:"column:age"`
    Sステータス   string `gorm:"column:status;index:idx_status"`
}

func (*Tアカウント情報) TableName() string {
    return "accounts"
}
```

### 한국어

```go
type T계정정보 struct {
    ID    uint   `gorm:"primaryKey"`
    G계정   string `gorm:"column:account;uniqueIndex:udx_account"`
    N닉네임  string `gorm:"column:nickname;index:idx_nickname"`
    N나이   int    `gorm:"column:age"`
    S상태   string `gorm:"column:status;index:idx_status"`
}

func (*T계정정보) TableName() string {
    return "accounts"
}
```

---

## 🤝 配合生态使用

gormmom 只负责把关命名。生成类型安全的列访问器、做仓储式 CRUD，是生态里其它包的职责：

- **[gormcngen](https://github.com/yylego/gormcngen)** 从模型生成 `XxxColumns` 类型安全列访问器。
- **[gormrepo](https://github.com/yylego/gormrepo)** 用这些列做仓储式 CRUD。

典型流程：写母语模型 → 让 gormmom 校验命名 → 让 gormcngen 生成列访问器 → 用 gormrepo 做 CRUD。端到端可运行的示例（中文、日文、韩文）都在 [examples](internal/examples/) 里。

---

## 📝 完整示例

查看 [examples](internal/examples/) 目录，里面有完整的集成示例

---

## 关联项目

通过这些集成包探索完整的 GORM 生态：

### 核心生态

- **[gormcnm](https://github.com/yylego/gormcnm)** - GORM 基础，提供类型安全的列操作和查询构建器
- **[gormcngen](https://github.com/yylego/gormcngen)** - 使用 AST 的代码生成引擎，支持类型安全的 GORM 操作
- **[gormrepo](https://github.com/yylego/gormrepo)** - 遵循 GORM 最佳实践的仓储模式实现
- **[gormmom](https://github.com/yylego/gormmom)** - 原生语言 GORM 命名把关器——校验表名/列名/索引名是否数据库安全（本项目）
- **[gormzhcn](https://github.com/yylego/gormzhcn)** - 完整的中文 GORM 编程接口

每个包针对 GORM 开发的不同方面，包括本地化、类型安全操作和代码生成。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) CLOSE: STANDARD PROJECT FOOTER -->
