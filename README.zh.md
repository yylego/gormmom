[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/gormmom/release.yml?branch=main&label=BUILD)](https://github.com/yylego/gormmom/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/gormmom)](https://pkg.go.dev/github.com/yylego/gormmom)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/gormmom/main.svg)](https://coveralls.io/github/yylego/gormmom?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/gormmom.svg)](https://github.com/yylego/gormmom/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/gormmom)](https://goreportcard.com/report/github.com/yylego/gormmom)

# 🌍 GORMMOM - GORM 原生语言编程革命

**gormmom** 是 **原生语言编程引擎**，打破数据库开发中的语言壁垒。作为 GORM 生态系统的 **智能标签生成引擎**，它赋能全球团队使用原生语言编写数据库模型，同时自动生成数据库兼容的 GORM 标签和列名。

> 🎯 **语言解放**: 用中文、阿拉伯语、日语和各种语言编程 - gormmom 架起人类表达与数据库需求之间的桥梁。

---

## 生态系统

![GORM Type-Safe Ecosystem](https://github.com/yylego/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->

## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

---

## 🚀 安装

```bash
go get github.com/yylego/gormmom
```

---

## 🔄 技术对比

| 生态系统      | Java MyBatis Plus  | Python SQLAlchemy | Go GORM 生态系统    |
|-----------|--------------------|-------------------|-----------------|
| **类型安全列** | `Example::getName` | `Example.name`    | `cls.Name.Eq()` |
| **代码生成**  | ✅ 插件支持             | ✅ 反射机制            | ✅ AST 精度        |
| **仓储模式**  | ✅ BaseMapper       | ✅ Session API     | ✅ GormRepo      |
| **原生语言**  | 🟡 有限支持            | 🟡 有限支持           | ✅ 完整支持          |

---

## 🌟 问题与解决方案

### ⚡ 标准方法
```go
// ❌ 常见方法：开发者被限制在英语命名
type Account struct {
    ID       uint   `gorm:"primaryKey"`
    Username string `gorm:"column:username;uniqueIndex"`
    Nickname string `gorm:"column:nickname;index"`
    Age      int    `gorm:"column:age"`
    PhoneNum string `gorm:"column:phone_num"`
    Mailbox  string `gorm:"column:mailbox"`
    Address  string `gorm:"column:address"`
    Status   string `gorm:"column:status;index"`
}
```

### ✅ GORMMOM 解决方案
```go
// ✅ GORMMOM: 用原生语言编程！
type T账户信息 struct {
    ID   uint   `gorm:"primaryKey"`
    Z账号 string `gorm:"uniqueIndex"`
    N昵称 string `gorm:"index"`
    A年龄 int    `gorm:""`
    D电话 string `gorm:""`
    E邮箱 string `gorm:""`
    J住址 string `gorm:""`
    S状态 string `gorm:"index"`
}

func (*T账户信息) TableName() string {
    return "accounts" // 数据库兼容的表名
}
```

---

## 🌍 多语言示例

### 繁體中文
```go
type T賬戶信息 struct {
    ID    uint   `gorm:"primaryKey"`
    Z賬號  string `gorm:"uniqueIndex"`
    N暱稱  string `gorm:"index"`
    A年齡  int    `gorm:""`
    D電話  string `gorm:""`
    E郵箱  string `gorm:""`
    J住址  string `gorm:""`
    S狀態  string `gorm:"index"`
}

func (*T賬戶信息) TableName() string {
    return "accounts"
}
```

### 日本語
```go
type Tアカウント情報 struct {
    ID        uint   `gorm:"primaryKey"`
    Aアカウント string `gorm:"uniqueIndex"`
    Nニックネーム string `gorm:"index"`
    N年齢      int    `gorm:""`
    D電話番号   string `gorm:""`
    Eメール    string `gorm:""`
    J住所      string `gorm:""`
    Sステータス  string `gorm:"index"`
}

func (*Tアカウント情報) TableName() string {
    return "accounts"
}
```

### 한국어
```go
type T계정정보 struct {
    ID    uint   `gorm:"primaryKey"`
    G계정   string `gorm:"uniqueIndex"`
    N닉네임  string `gorm:"index"`
    N나이   int    `gorm:""`
    J전화번호 string `gorm:""`
    E이메일  string `gorm:""`
    J주소   string `gorm:""`
    S상태   string `gorm:"index"`
}

func (*T계정정보) TableName() string {
    return "accounts"
}
```

---

## 🛠️ 使用方法

### 1. 自动标签生成

gormmom 执行后，结构体获得数据库兼容的列标签：

```go
// 生成的数据库兼容列名
type T账户信息 struct {
    ID    uint   `gorm:"primaryKey"`
    Z账号  string `gorm:"column:z_zhang_hao;uniqueIndex"`
    N昵称  string `gorm:"column:n_ni_cheng;index"`
    A年龄  int    `gorm:"column:a_nian_ling"`
    D电话  string `gorm:"column:d_dian_hua"`
    E邮箱  string `gorm:"column:e_you_xiang"`
    J住址  string `gorm:"column:j_zhu_zhi"`
    S状态  string `gorm:"column:s_zhuang_tai;index"`
}
```

### 2. 生成命令

```bash
# 步骤 1：生成原生语言字段的 GORM 标签
go test -v -run TestGen/GenGormMom

# 步骤 2：生成类型安全列方法（配合 gormcngen）
go test -v -run TestGen/GenGormCnm
```

### 3. 配合 gormrepo 使用

**English Version:**

```go
// Create repo
repo := gormrepo.NewGormRepo(&Account{}, (&Account{}).Columns())

// Select - First (by username)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Username.Eq("alice"))
})

// Select - First (by nickname)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Nickname.Eq("Alice"))
})

// Select - Find
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.Age.Gte(18))
})

// Select - FindPage
accounts, err := repo.With(ctx, db).FindPage(
    func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
        return db.Where(cls.Age.Gte(18))
    },
    func(cls *AccountColumns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// Create
err := repo.With(ctx, db).Create(&Account{Username: "bob", Nickname: "Bob", Age: 25})

// Update
err := repo.With(ctx, db).Updates(
    func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *AccountColumns) map[string]interface{} {
        return cls.Kw(cls.Age.Kv(26)).AsMap()
    },
)

// Delete
err := repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
    return db.Where(cls.ID.Eq(1))
})
```

**中文（简体）版本:**

```go
// Create repo
repo := gormrepo.NewGormRepo(&T账户信息{}, (&T账户信息{}).Columns())

// Select - First (by username)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.Z账号.Eq("wang-xiao-ming"))
})

// Select - First (by nickname)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.N昵称.Eq("王小明"))
})

// Select - Find
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.A年龄.Gte(18))
})

// Select - FindPage
accounts, err := repo.With(ctx, db).FindPage(
    func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
        return db.Where(cls.A年龄.Gte(18))
    },
    func(cls *T账户信息Columns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// Create
err := repo.With(ctx, db).Create(&T账户信息{Z账号: "han-mei-mei", N昵称: "韩梅梅", A年龄: 25})

// Update
err := repo.With(ctx, db).Updates(
    func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *T账户信息Columns) map[string]interface{} {
        return cls.Kw(cls.A年龄.Kv(26)).AsMap()
    },
)

// Delete
err := repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *T账户信息Columns) *gorm.DB {
    return db.Where(cls.ID.Eq(1))
})
```

**中文（繁體）版本:**

```go
// Create repo
repo := gormrepo.NewGormRepo(&T賬戶信息{}, (&T賬戶信息{}).Columns())

// Select - First (by username)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T賬戶信息Columns) *gorm.DB {
    return db.Where(cls.Z賬號.Eq("wang-xiao-ming"))
})

// Select - First (by nickname)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T賬戶信息Columns) *gorm.DB {
    return db.Where(cls.N暱稱.Eq("王小明"))
})

// Select - Find
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *T賬戶信息Columns) *gorm.DB {
    return db.Where(cls.A年齡.Gte(18))
})

// Select - FindPage
accounts, err := repo.With(ctx, db).FindPage(
    func(db *gorm.DB, cls *T賬戶信息Columns) *gorm.DB {
        return db.Where(cls.A年齡.Gte(18))
    },
    func(cls *T賬戶信息Columns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// Create
err := repo.With(ctx, db).Create(&T賬戶信息{Z賬號: "han-mei-mei", N暱稱: "韓梅梅", A年齡: 25})

// Update
err := repo.With(ctx, db).Updates(
    func(db *gorm.DB, cls *T賬戶信息Columns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *T賬戶信息Columns) map[string]interface{} {
        return cls.Kw(cls.A年齡.Kv(26)).AsMap()
    },
)

// Delete
err := repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *T賬戶信息Columns) *gorm.DB {
    return db.Where(cls.ID.Eq(1))
})
```

**日本語版:**

```go
// Create repo
repo := gormrepo.NewGormRepo(&Tアカウント情報{}, (&Tアカウント情報{}).Columns())

// Select - First (by username)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *Tアカウント情報Columns) *gorm.DB {
    return db.Where(cls.Aアカウント.Eq("tanaka"))
})

// Select - First (by nickname)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *Tアカウント情報Columns) *gorm.DB {
    return db.Where(cls.Nニックネーム.Eq("田中太郎"))
})

// Select - Find
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *Tアカウント情報Columns) *gorm.DB {
    return db.Where(cls.N年齢.Gte(18))
})

// Select - FindPage
accounts, err := repo.With(ctx, db).FindPage(
    func(db *gorm.DB, cls *Tアカウント情報Columns) *gorm.DB {
        return db.Where(cls.N年齢.Gte(18))
    },
    func(cls *Tアカウント情報Columns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// Create
err := repo.With(ctx, db).Create(&Tアカウント情報{Aアカウント: "suzuki", Nニックネーム: "鈴木花子", N年齢: 25})

// Update
err := repo.With(ctx, db).Updates(
    func(db *gorm.DB, cls *Tアカウント情報Columns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *Tアカウント情報Columns) map[string]interface{} {
        return cls.Kw(cls.N年齢.Kv(26)).AsMap()
    },
)

// Delete
err := repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *Tアカウント情報Columns) *gorm.DB {
    return db.Where(cls.ID.Eq(1))
})
```

**한국어판:**

```go
// Create repo
repo := gormrepo.NewGormRepo(&T계정정보{}, (&T계정정보{}).Columns())

// Select - First (by username)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T계정정보Columns) *gorm.DB {
    return db.Where(cls.G계정.Eq("kim-cheol-su"))
})

// Select - First (by nickname)
account, err := repo.With(ctx, db).First(func(db *gorm.DB, cls *T계정정보Columns) *gorm.DB {
    return db.Where(cls.N닉네임.Eq("김철수"))
})

// Select - Find
accounts, err := repo.With(ctx, db).Find(func(db *gorm.DB, cls *T계정정보Columns) *gorm.DB {
    return db.Where(cls.N나이.Gte(18))
})

// Select - FindPage
accounts, err := repo.With(ctx, db).FindPage(
    func(db *gorm.DB, cls *T계정정보Columns) *gorm.DB {
        return db.Where(cls.N나이.Gte(18))
    },
    func(cls *T계정정보Columns) gormcnm.OrderByBottle {
        return cls.ID.OrderByBottle("DESC")
    },
    &gormrepo.Pagination{Limit: 10, Offset: 0},
)

// Create
err := repo.With(ctx, db).Create(&T계정정보{G계정: "lee-young-hee", N닉네임: "이영희", N나이: 25})

// Update
err := repo.With(ctx, db).Updates(
    func(db *gorm.DB, cls *T계정정보Columns) *gorm.DB {
        return db.Where(cls.ID.Eq(1))
    },
    func(cls *T계정정보Columns) map[string]interface{} {
        return cls.Kw(cls.N나이.Kv(26)).AsMap()
    },
)

// Delete
err := repo.With(ctx, db).DeleteW(func(db *gorm.DB, cls *T계정정보Columns) *gorm.DB {
    return db.Where(cls.ID.Eq(1))
})
```

## 📝 完整示例

查看 [examples](internal/examples/) 目录获取完整集成示例。

---

## 关联项目

探索完整的 GORM 生态系统集成包：

### 核心生态

- **[gormcnm](https://github.com/yylego/gormcnm)** - GORM 基础层，提供类型安全的列操作和条件构建
- **[gormcngen](https://github.com/yylego/gormcngen)** - 使用 AST 的代码生成引擎，支持类型安全的 GORM 操作
- **[gormrepo](https://github.com/yylego/gormrepo)** - 仓储模式实现，遵循 GORM 最佳实践
- **[gormmom](https://github.com/yylego/gormmom)** - 原生语言 GORM 标签生成引擎，支持智能列名（本项目）
- **[gormzhcn](https://github.com/go-zwbc/gormzhcn)** - 完整的 GORM 中文编程接口

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

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![Stargazers](https://starchart.cc/yylego/gormmom.svg?variant=adaptive)](https://starchart.cc/yylego/gormmom)
