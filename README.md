# Go Examples

**CMSC-405 Programming Languages — Go**

Companion code for the Go module Learn pages. Each file is a standalone Go program — open it, read it, run it.

## Running Examples

Each folder has its own `go.mod`. Run individual files from within the folder:

```bash
cd collections
go run maps-basics.go
```

## Folders

```
basics/          # Variables, types, printing, hello world
functions/       # Functions, multiple returns, error handling, defer
http-server/     # HTTP handlers, JSON encode/decode, method routing, timeouts
collections/     # Maps, slices, arrays, filtering, iteration
```

New files are added to existing folders each week. Pull the latest to get new examples:

```bash
git pull
```
