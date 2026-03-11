[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/gormmom/release.yml?branch=main&label=BUILD)](https://github.com/yylego/gormmom/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/gormmom)](https://pkg.go.dev/github.com/yylego/gormmom)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/gormmom/main.svg)](https://coveralls.io/github/yylego/gormmom?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/gormmom.svg)](https://github.com/yylego/gormmom/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/gormmom)](https://goreportcard.com/report/github.com/yylego/gormmom)

# 🌍 GORMMOM - Native Language Programming Revolution with GORM

**gormmom** is the **native language programming engine** that breaks down language barriers in database development. As the **smart tag generation engine** of the GORM ecosystem, it empowers teams worldwide to write database models in native languages while auto generating database-compatible GORM tags and column names.

> 🎯 **Language Liberation**: Code in Chinese, Arabic, Japanese, and various languages - gormmom bridges the gap between human expression and database requirements.

---

## Ecosystem

![GORM Type-Safe Ecosystem](https://github.com/yylego/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

---

## 🚀 Installation

```bash
go get github.com/yylego/gormmom
```

---

## 🔄 Tech Comparison

| Ecosystem             | Java MyBatis Plus  | Python SQLAlchemy | Go GORM Ecosystem  |
|-----------------------|--------------------|-------------------|--------------------|
| **Type-Safe Columns** | `Example::getName` | `Example.name`    | `cls.Name.Eq()`    |
| **Code Generation**   | ✅ Plugin support   | ✅ Reflection      | ✅ AST precision    |
| **Repo Pattern**      | ✅ BaseMapper       | ✅ Session API     | ✅ GormRepo         |
| **Native Language**   | 🟡 Limited         | 🟡 Limited        | ✅ Complete support |

---

## 🌟 The Problem & Solution

### ⚡ Standard Approach
```go
// ❌ Common approach: Developers constrained to English naming
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

### ✅ GORMMOM Solution
```go
// ✅ GORMMOM: Program in native language!
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
    return "accounts" // Database-compatible table name
}
```

---

## 🌍 Multi-Language Examples

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

## 🛠️ Usage

### 1. Auto Tag Generation

Once gormmom executes, the struct gets database-compatible column tags:

```go
// Generated with database-compatible column names
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

### 2. Generate Commands

```bash
# Step 1: Generate GORM tags with native language fields
go test -v -run TestGen/GenGormMom

# Step 2: Generate type-safe column methods (with gormcngen)
go test -v -run TestGen/GenGormCnm
```

### 3. Use with gormrepo

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

## 📝 Complete Examples

Check [examples](internal/examples/) DIR with complete integration examples

---

## Related Projects

Explore the complete GORM ecosystem with these integrated packages:

### Core Ecosystem

- **[gormcnm](https://github.com/yylego/gormcnm)** - GORM foundation providing type-safe column operations and query builders
- **[gormcngen](https://github.com/yylego/gormcngen)** - AST-based code generation engine with type-safe GORM operations
- **[gormrepo](https://github.com/yylego/gormrepo)** - Repo pattern implementation with GORM best practices
- **[gormmom](https://github.com/yylego/gormmom)** - Native language GORM tag generation engine with smart column naming (this project)
- **[gormzhcn](https://github.com/go-zwbc/gormzhcn)** - Complete Chinese programming interface with GORM

Each package targets different aspects of GORM development, including localization, type-safe operations, and code generation.

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## 📈 GitHub Stars

[![Stargazers](https://starchart.cc/yylego/gormmom.svg?variant=adaptive)](https://starchart.cc/yylego/gormmom)
