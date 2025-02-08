# Multi-language Monorepo with Bazel

### Gazelle & Go configuration doc:
https://github.com/bazel-contrib/rules_go/blob/master/docs/go/core/bzlmod.md

```bash
bazel mod tidy
```

Generate BUILD files:
```bash
bazel run //:gazelle  
```

````bash
bazel run @rules_go//go -- mod tidy -v
````

```bash
bazel build //... 
```

```bash
bazel clean
```

```bash
bazel run //apps/go/bar:bar
```