<!-- TEMPLATE (ZH) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/gormmom/release.yml?branch=main&label=BUILD)](https://github.com/yylego/gormmom/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/gormmom)](https://pkg.go.dev/github.com/yylego/gormmom)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/gormmom/main.svg)](https://coveralls.io/github/yylego/gormmom?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.27%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/gormmom.svg)](https://github.com/yylego/gormmom/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/gormmom)](https://goreportcard.com/report/github.com/yylego/gormmom)

<!-- TEMPLATE (ZH) CLOSE: BADGES -->

# 🌍 GORMMOM - GORM 母语命名把关器

**gormmom** 让你用**母语**（中文、日文、韩文等）编写 GORM 模型，然后检查模型产出的每个数据库名字——表名、每个列名、每个索引名——都是数据库安全的。它作为普通单元测试运行，完全不连接数据库。

> 🎯 **用母语编程，交付数据库安全的表结构。** gormmom 在母语命名进入数据库之前，就替你抓出数据库会拒绝的那些名字。

---

## 🤖 设计意图

gormmom 过去会替你拼出列名和索引名（中文拼音、日文罗马音等）。现在你定义字段的那一刻，AI 助手就能写出又好又数据库安全的名字，所以生成这一层已经去掉了。

仍然会出问题的是漏网的名字：字段没写列名、索引没写名字，GORM 就会产出非 ASCII 的名字，数据库要到迁移或运行时才拒绝。gormmom 只做这一项检查，在不连数据库的单元测试里完成。

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
type T学生 struct {
    ID  uint   `gorm:"primaryKey"`
    X姓名 string `gorm:"uniqueIndex"` // ❌ 没写列名和索引名 → GORM 用字段名拼出两者 → 非 ASCII → 被拒绝
    B班级 string `gorm:""`            // ❌ 没写列名 → 列名落到非 ASCII → 被拒绝
}
```

### ✅ gormmom 的做法

保留母语字段名，但显式给出数据库安全的列名 / 索引名 / 表名——然后让 gormmom 在单元测试里替你把关：

```go
type T学生 struct {
    ID  uint   `gorm:"primaryKey"`
    X姓名 string `gorm:"column:name;uniqueIndex:udx_students_name"`
    B班级 string `gorm:"column:grade;index:idx_students_grade"`
    S社团 string `gorm:"column:club"`
}

func (*T学生) TableName() string {
    return "students" // 数据库安全的 ASCII 表名
}
```

---

## 🛠️ 使用方法 —— 在单元测试里校验

gormmom 从不连接数据库，它作为一个普通单元测试运行：

```go
func TestValidateNaming(t *testing.T) {
    objects := []any{
        &T学生{},
    }

    // 直接从模型组装校验配置（纯反射，不读取任何磁盘文件）
    configs := gormmom.NewConfigs(objects, gormmom.NewOptions())

    // 把关：表名 / 列名 / 索引名都必须数据库安全
    require.NoError(t, configs.Validate())
}
```

规则就是一条正则：每个名字都要匹配 `^[a-zA-Z0-9_]{1,N}$`。查到第一个不合格的名字就报错，错误里写明结构体、名字、涉及的字段，以及该补什么。

选项从 `gormmom.NewOptions()` 开始：长度上限 N 默认 63，跟 Postgres 一样，`WithMaxIdentLength(n)` 可以换成别的上限；名字按 GORM 默认的命名策略推导，`WithNamingStrategy(namer)` 只给用了自定义命名策略的项目用。

---

## 🌍 多语言示例

每种语言的写法都一样：母语字段名，加上显式的、数据库安全的列名 / 索引名 / 表名。每种语言各有一个示例，注释用的就是那门语言。

### 繁體中文 — [example2](internal/examples/example2)

```go
type T學生 struct {
    ID  uint   `gorm:"primaryKey"`
    X姓名 string `gorm:"column:name;uniqueIndex:udx_students_name"`
    N年級 string `gorm:"column:grade;index:idx_students_grade"`
    S社團 string `gorm:"column:club"`
}

func (*T學生) TableName() string {
    return "students"
}
```

### 日本語 — [example3](internal/examples/example3)

```go
type T生徒 struct {
    ID   uint   `gorm:"primaryKey"`
    N名前  string `gorm:"column:name;uniqueIndex:udx_students_name"`
    G学年  string `gorm:"column:grade;index:idx_students_grade"`
    Tクラブ string `gorm:"column:club"`
}

func (*T生徒) TableName() string {
    return "students"
}
```

### 한국어 — [example4](internal/examples/example4)

```go
type T학생 struct {
    ID   uint   `gorm:"primaryKey"`
    I이름  string `gorm:"column:name;uniqueIndex:udx_students_name"`
    H학년  string `gorm:"column:grade;index:idx_students_grade"`
    D동아리 string `gorm:"column:club"`
}

func (*T학생) TableName() string {
    return "students"
}
```

### Español — [example5](internal/examples/example5)

西班牙语的字母有大小写，字段不用加 ASCII 前缀。大部分单词本身就是 ASCII，但带重音的字段（如 `Año`）会产出非 ASCII 的列名，所以同样要写明列名：

```go
type Estudiante struct {
    ID     uint   `gorm:"primaryKey"`
    Nombre string `gorm:"column:name;uniqueIndex:udx_students_name"`
    Año    string `gorm:"column:grade;index:idx_students_grade"`
    Club   string `gorm:"column:club"`
}

func (*Estudiante) TableName() string {
    return "students"
}
```

### Português — [example6](internal/examples/example6)

带重音的字段（如 `Série`）会产出非 ASCII 的列名，所以要写明列名：

```go
type Aluno struct {
    ID    uint   `gorm:"primaryKey"`
    Nome  string `gorm:"column:name;uniqueIndex:udx_students_name"`
    Série string `gorm:"column:grade;index:idx_students_grade"`
    Clube string `gorm:"column:club"`
}

func (*Aluno) TableName() string {
    return "students"
}
```

### Français — [example7](internal/examples/example7)

带重音的字段（如 `Prénom`）会产出非 ASCII 的列名，所以要写明列名：

```go
type Élève struct {
    ID     uint   `gorm:"primaryKey"`
    Prénom string `gorm:"column:name;uniqueIndex:udx_students_name"`
    Classe string `gorm:"column:grade;index:idx_students_grade"`
    Club   string `gorm:"column:club"`
}

func (*Élève) TableName() string {
    return "students"
}
```

### Deutsch — [example8](internal/examples/example8)

带变音符号的结构体名（如 `Schüler`）会产出非 ASCII 的表名，所以要写 `TableName()` 方法：

```go
type Schüler struct {
    ID     uint   `gorm:"primaryKey"`
    Name   string `gorm:"column:name;uniqueIndex:udx_students_name"`
    Klasse string `gorm:"column:grade;index:idx_students_grade"`
    Verein string `gorm:"column:club"`
}

func (*Schüler) TableName() string {
    return "students"
}
```

### Tiếng Việt — [example9](internal/examples/example9)

大多数词都带声调符号（如 `Tên`），所以每一列都要写明列名：

```go
type HọcSinh struct {
    ID       uint   `gorm:"primaryKey"`
    Tên      string `gorm:"column:name;uniqueIndex:udx_students_name"`
    Lớp      string `gorm:"column:grade;index:idx_students_grade"`
    CâuLạcBộ string `gorm:"column:club"`
}

func (*HọcSinh) TableName() string {
    return "students"
}
```

> 💡 这门语言的文字得能写进 Go 的标识符。中文、日文、韩文、西班牙语、葡萄牙语、法语、德语、越南语都可以。元音写成组合符号的文字（如高棉语 ខ្មែរ、泰语 ไทย）写不成字段名，因为 Go 不允许标识符里出现这些符号。

---

## 🤝 配合生态使用

gormmom 只负责把关命名。生成类型安全的列访问器、做仓储式 CRUD，是生态里其它包的职责：

- **[gormcngen](https://github.com/yylego/gormcngen)** 从模型生成 `XxxColumns` 类型安全列访问器。
- **[gormrepo](https://github.com/yylego/gormrepo)** 用这些列做仓储式 CRUD。

典型流程：写母语模型 → 让 gormmom 校验命名 → 让 gormcngen 生成列访问器 → 用 gormrepo 做 CRUD。生成测试会先跑校验，通过了才生成列，见 [example1](internal/examples/example1/internal/models/models.gen_test.go)。

---

## 📝 完整示例

| 示例                                   | 讲什么                                                |
| -------------------------------------- | ----------------------------------------------------- |
| [example1](internal/examples/example1) | 入门（简体中文字段）：通过校验，再用生成的列配合 GORM |
| [example2](internal/examples/example2) | 同样的用法，繁體中文字段                              |
| [example3](internal/examples/example3) | 同样的用法，日语（日本語）字段                        |
| [example4](internal/examples/example4) | 同样的用法，韩语（한국어）字段                        |
| [example5](internal/examples/example5) | 同样的用法，西班牙语（Español）字段                   |
| [example6](internal/examples/example6) | 同样的用法，葡萄牙语（Português）字段                 |
| [example7](internal/examples/example7) | 同样的用法，法语（Français）字段                      |
| [example8](internal/examples/example8) | 同样的用法，德语（Deutsch）字段                       |
| [example9](internal/examples/example9) | 同样的用法，越南语（Tiếng Việt）字段                  |

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
