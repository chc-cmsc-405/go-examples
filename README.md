# Go Examples

**CMSC-405 Programming Languages — Go**

Companion code for the Go module Learn pages. Each file is a standalone Go program — open it, read it, run it.

## Running Examples

Each folder has its own `go.mod`. Run individual files from within the folder:

```bash
cd collections
go run maps-basics.go
```

**Important:** Use `go run filename.go` (not `go run .`). Each folder has multiple standalone programs — running `go run .` will fail because Go sees multiple `main` functions. This also causes red underlines in VS Code — those are cosmetic, not real errors. Your code runs fine with `go run filename.go`.

## Folders

```
basics/          # Variables, types, printing, hello world
functions/       # Functions, multiple returns, error handling, defer
http-server/     # HTTP handlers, JSON encode/decode, method routing, timeouts
collections/     # Maps, slices, arrays, filtering, iteration
structs/         # Struct definitions, methods, value vs pointer receivers
interfaces/      # Implicit satisfaction, polymorphism, analyzer patterns
composition/     # Struct embedding, method promotion, factory functions
```

New files are added to existing folders each week. Pull the latest to get new examples:

```bash
git pull
```
