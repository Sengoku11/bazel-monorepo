# Multi-language Monorepo with Bazel

Simple template of a monorepo using Bazel. 

Capabilities:
* Generate BUILD files from `go.mod`
* Runs `vet` and other checks on build
* Supports multiple versions of `go` and avoids excessive `replace ... => ...` by using `go.work`
* Platform independent builds

### Gazelle & Go configuration doc:
https://github.com/bazel-contrib/rules_go/blob/master/docs/go/core/bzlmod.md

```bash
bazel mod tidy
```

Generate BUILD files:
```bash
bazel run //:gazelle  
```

Note that go command arguments starting with - require the use of the double dash separator with bazel run:
````bash
bazel run @rules_go//go -- mod tidy -v
bazel run @rules_go//go work sync
````

```bash
bazel build //... 
```

```bash
bazel clean --async
```

```bash
bazel run //apps/go/bar:bar
```