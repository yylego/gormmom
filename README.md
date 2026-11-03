<!-- TEMPLATE (EN) BEGIN: BADGES -->

[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yylego/gormmom/release.yml?branch=main&label=BUILD)](https://github.com/yylego/gormmom/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yylego/gormmom)](https://pkg.go.dev/github.com/yylego/gormmom)
[![Coverage Status](https://img.shields.io/coveralls/github/yylego/gormmom/main.svg)](https://coveralls.io/github/yylego/gormmom?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.27%2B-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/yylego/gormmom.svg)](https://github.com/yylego/gormmom/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yylego/gormmom)](https://goreportcard.com/report/github.com/yylego/gormmom)

<!-- TEMPLATE (EN) CLOSE: BADGES -->

# 🌍 GORMMOM - Native Language GORM Naming Guardian

**gormmom** lets you write GORM structs with fields in a **native language** (Chinese, Japanese, Korean and more), then checks that each database name the struct gives — the table name, each column name, each index name — stays database-safe. It runs as a plain unit test and does not touch a database.

> 🎯 **Write the fields in the language you think in, ship database-safe schemas.** gormmom catches the native-language names a database would reject, before those names reach the database.

---

## 🤖 Design Intent

gormmom used to generate the column names and index names (pinyin, romaji and so on). Now an AI assistant writes a good database-safe name the moment you define the field, so that generation is gone.

What can go wrong is a name that slips through: a field with no column name, and an index with no name, both give a non-ASCII name that the database rejects late, at migration / at runtime. gormmom does that one check, in a unit test with no database.

---

## Ecosystem

![GORM Type-Safe Ecosystem](https://github.com/yylego/gormcnm/raw/main/assets/gormcnm-ecosystem.svg)

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->

## CHINESE README

[中文说明](README.zh.md)

<!-- TEMPLATE (EN) CLOSE: LANGUAGE NAVIGATION -->

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
type T学生 struct {
    ID  uint   `gorm:"primaryKey"`
    X姓名 string `gorm:"uniqueIndex"` // ❌ no column name / index name → GORM builds both from the field name → non-ASCII → rejected
    B班级 string `gorm:""`            // ❌ no column name → column falls back to a non-ASCII name → rejected
}
```

### ✅ The gormmom way

Keep the native-language field names, but give explicit database-safe column / index / table names — and let gormmom guard them for you in a unit test:

```go
type T学生 struct {
    ID  uint   `gorm:"primaryKey"`
    X姓名 string `gorm:"column:name;uniqueIndex:udx_students_name"`
    B班级 string `gorm:"column:grade;index:idx_students_grade"`
    S社团 string `gorm:"column:club"`
}

func (*T学生) TableName() string {
    return "students" // database-safe ASCII table name
}
```

---

## 🛠️ Usage — validate in a unit test

gormmom never connects to a database; it runs as a plain unit test:

```go
func TestValidateNaming(t *testing.T) {
    objects := []any{
        &T学生{},
    }

    // Build the validation configs straight from the models (reflection only, reads nothing from disk).
    configs := gormmom.NewConfigs(objects, gormmom.NewOptions())

    // Guard: table / column / index names must all be database-safe.
    require.NoError(t, configs.Validate())
}
```

The rule is one regexp: each name must match `^[a-zA-Z0-9_]{1,N}$`. The check stops at the first bad name, and the error gives the struct, the name, the field when one applies, and what to write to fix it.

The options start from `gormmom.NewOptions()`: the length cap N defaults to 63, as in Postgres, and `WithMaxNameLength(n)` sets a different cap; the names come from GORM's default naming strategy, and `WithNamingStrategy(namer)` is just needed when the project runs GORM with a custom strategy.

---

## 🌍 Multi-Language Examples

The pattern is the same in each language: native-language field names, plus explicit database-safe column / index / table names. Each language has its own example, with comments in that language.

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

Spanish letters have a case, so the fields need no ASCII prefix. Most words are plain ASCII, but a field with an accent such as `Año` gives a non-ASCII column, so give it an explicit name too:

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

A field with an accent such as `Série` gives a non-ASCII column, so give it an explicit name:

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

A field with an accent such as `Prénom` gives a non-ASCII column, so give it an explicit name:

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

A struct name with an umlaut such as `Schüler` gives a non-ASCII table, so give it a `TableName()` method:

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

Most words have a tone mark, such as `Tên`, so give each column an explicit name:

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

> 💡 The language needs letters that Go accepts in a name. Chinese, Japanese, Korean, Spanish, Portuguese, French, German and Vietnamese work. Scripts that write the vowels as combining marks, such as ខ្មែរ and ไทย, can not be written as field names, since Go rejects those marks in a name.

---

## 🤝 Use with the ecosystem

gormmom only guards naming. Generating type-safe column accessors and running repo-style CRUD are the jobs of sibling packages:

- **[gormcngen](https://github.com/yylego/gormcngen)** generates the `XxxColumns` type-safe column accessors from your model.
- **[gormrepo](https://github.com/yylego/gormrepo)** uses those columns for repo-style CRUD.

The common flow: write a native-language struct → let gormmom validate the naming → let gormcngen generate the columns → do CRUD with gormrepo. The generation test runs the check first and generates the columns just when it passes, see [example1](internal/examples/example1/internal/models/models.gen_test.go).

---

## 📝 Complete Examples

| Example                                | Topic                                                                                               |
| -------------------------------------- | --------------------------------------------------------------------------------------------------- |
| [example1](internal/examples/example1) | The basics with zh-Hans (简体中文) fields: pass the check, then use the generated columns with GORM |
| [example2](internal/examples/example2) | The same with zh-Hant (繁體中文) fields                                                             |
| [example3](internal/examples/example3) | The same with Japanese (日本語) fields                                                              |
| [example4](internal/examples/example4) | The same with Korean (한국어) fields                                                                |
| [example5](internal/examples/example5) | The same with Spanish (Español) fields                                                              |
| [example6](internal/examples/example6) | The same with Portuguese (Português) fields                                                         |
| [example7](internal/examples/example7) | The same with French (Français) fields                                                              |
| [example8](internal/examples/example8) | The same with German (Deutsch) fields                                                               |
| [example9](internal/examples/example9) | The same with Vietnamese (Tiếng Việt) fields                                                        |

---

## Related Projects

Explore the complete GORM ecosystem with these integrated packages:

### Core Ecosystem

- **[gormcnm](https://github.com/yylego/gormcnm)** - GORM foundation providing type-safe column operations and query builders
- **[gormcngen](https://github.com/yylego/gormcngen)** - AST-based code generation engine with type-safe GORM operations
- **[gormrepo](https://github.com/yylego/gormrepo)** - Repo pattern implementation with GORM best practices
- **[gormmom](https://github.com/yylego/gormmom)** - Native language GORM naming guardian — validates table/column/index names are database-safe (this project)
- **[gormzhcn](https://github.com/yylego/gormzhcn)** - Complete Chinese programming interface with GORM

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

<!-- TEMPLATE (EN) CLOSE: STANDARD PROJECT FOOTER -->
