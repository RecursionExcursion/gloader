# gloader (v.1.0.0)

A lightweight, flexible environment variable loader for Go.

Supports loading from `.env` files or any `io.Reader` source (such as strings, network streams, or in-memory data). 

[![Go](https://pkg.go.dev/badge/github.com/yourusername/gloader.svg)](https://pkg.go.dev/github.com/yourusername/gloader)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

---

## Features

- Load env vars from one or more `io.Reader`s
- Automatically falls back to `.env` file in none are provided
- Supports panicking or fallback lookups
- Simple API for flexible integration
- No external dependencies

---

## Installation

```bash
go get github.com/RecursionExcursion/gloader
```

---

## Usage

### Load from default `.env` file

```go
loader := &gloader.EnvLoader{}
_ = loader.LoadEnv()

foo := loader.MustGet("FOO")
```

### Load from custom readers (in-memory)

```go
r1 := strings.NewReader("FOO=bar")
r2 := strings.NewReader("BAR=baz")

loader := &gloader.EnvLoader{}
_ = loader.LoadEnv(r1, r2)

fmt.Println(os.Getenv("FOO")) // bar
```

---

## Get API

```go
loader.Get("FOO")                  // string, error
loader.MustGet("FOO")             // string, panics if missing
loader.GetOrFallback("FOO", "x")  // string, returns fallback
loader.GetOrDefault("FOO")        // string, returns ""
```

---

## .env Format

Each line should follow:

```
KEY=VALUE
```

Lines starting with `#` are ignored. Whitespace is trimmed.

---

## 🪪 License

MIT License © 2025 RecursionExcursion(https://github.com/RecursionExcursion)

See [LICENSE](LICENSE) for details.
