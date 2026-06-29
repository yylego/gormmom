[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/gormmom/release.yml?branch=main&label=BUILD)](https://github.com/yylego/gormmom/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/gormmom)](https://pkg.go.dev/github.com/yylego/gormmom)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/gormmom/main.svg)](https://coveralls.io/github/yylego/gormmom?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/gormmom.svg)](https://github.com/yylego/gormmom/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/gormmom)](https://goreportcard.com/report/github.com/yylego/gormmom)

# 🌍 GORMMOM - Native Language GORM Naming Guardian

**gormmom** lets you write GORM models in your **native language** (Chinese, Japanese, Korean, and more), then **guards** that every database name your model produces — the table name, each column name, each index name — stays database-safe. It runs as a plain unit test and never touches a database.

> 🎯 **Write in your mother tongue, ship database-safe schemas.** gormmom catches native-language names that a database would reject — before they ever reach the database.

---

## 🤖 Why gormmom is just a guardian now (the AI-era simplification)

gormmom used to be a **generation engine**: you wrote native-language fields with empty tags, and it auto-synthesized the column names and index names for you (pinyin for Chinese, romaji for Japanese, and so on).

**The AI era made that generation pointless.** Today an AI assistant writes a good, database-safe column name for you the moment you define the field — directly, in context, far better than any fixed rule engine could. A machine-synthesized `z_zhang_hao` was never really what anyone wanted. So the whole generation layer lost its reason to exist.

What AI does **not** automatically guarantee is the part that still bites you: **validation**. When you program in a native language, GORM may turn a field into a non-ASCII column name, or build a composite index name out of native-language field names — and the database rejects both. That failure shows up late, at migration or at runtime.

So gormmom dropped generation entirely and now does exactly one thing: **catch those naming problems in a database-free unit test**, so they never blow up later.

- ❌ Gone: code generation, pinyin/romaji column synthesis, pattern enums, the old generate commands.
- ✅ Kept: one job — validate that the table / column / index names are database-safe.
- ✅ Zero database connection — a pure static check that runs as an ordinary unit test.

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

## 🌟 The Problem & Solution

### ⚡ The trap of native-language models

Native-language fields read beautifully, but the names GORM derives from them may be names the database refuses:

```go
type T账户信息 struct {
    ID   uint   `gorm:"primaryKey"`
    Z账号 string `gorm:"uniqueIndex"` // ❌ no index name → GORM builds one from the field name → non-ASCII → rejected
    N昵称 string `gorm:""`            // ❌ no column name → column falls back to a non-ASCII name → rejected
}
```

### ✅ The gormmom way

Keep the native-language field names, but give explicit database-safe column / index / table names — and let gormmom guard them for you in a unit test:

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
    return "accounts" // database-safe ASCII table name
}
```

---

## 🛠️ Usage — validate in a unit test

gormmom never connects to a database; it runs as a plain unit test:

```go
func TestValidateNaming(t *testing.T) {
    objects := []interface{}{
        &T账户信息{},
    }

    // Build the validation configs straight from the models (reflection only, reads nothing from disk).
    configs := gormmom.NewConfigs(objects, gormmom.NewOptions())

    // Guard: table / column / index names must all be database-safe.
    require.NoError(t, configs.Validate())
}
```

The rule is one regexp: every name must match `^[a-zA-Z0-9_]{1,N}$` (N defaults to 63, tunable via `Options.MaxNameLength`). A failing name returns a clear error naming the struct, the offending name, and the field when one applies — caught long before any migration.

By default gormmom derives names with GORM's default naming strategy. If your project runs GORM with a custom strategy, pass it via `Options.NamingStrategy` so the validation matches what your database actually sees.

---

## 🌍 Multi-Language Examples

The pattern is the same in every language: native-language field names, plus explicit database-safe column / index / table names that gormmom validates.

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

## 🤝 Use with the ecosystem

gormmom only guards naming. Generating type-safe column accessors and running repo-style CRUD are the jobs of sibling packages:

- **[gormcngen](https://github.com/yylego/gormcngen)** generates the `XxxColumns` type-safe column accessors from your model.
- **[gormrepo](https://github.com/yylego/gormrepo)** uses those columns for repo-style CRUD.

A typical flow: write a native-language model → let gormmom validate the naming → let gormcngen generate the columns → use gormrepo for CRUD. End-to-end runnable examples (Chinese, Japanese, Korean) live in [examples](internal/examples/).

---

## 📝 Complete Examples

Check [examples](internal/examples/) DIR with complete integration examples

---

## Related Projects

Explore the complete GORM ecosystem with these integrated packages:

### Core Ecosystem

- **[gormcnm](https://github.com/yylego/gormcnm)** - GORM foundation providing type-safe column operations and query builders
- **[gormcngen](https://github.com/yylego/gormcngen)** - AST-based code generation engine with type-safe GORM operations
- **[gormrepo](https://github.com/yylego/gormrepo)** - Repo pattern implementation with GORM best practices
- **[gormmom](https://github.com/yylego/gormmom)** - Native language GORM naming guardian — validates table/column/index names are database-safe (this project)
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
