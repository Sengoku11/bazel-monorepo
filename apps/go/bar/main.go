package main

import f "github.com/Sengoku11/bazel-monorepo/libs/go/foo/pkg"
import bf "github.com/Sengoku11/bazel-monorepo/libs/go/baz/foo/pkg"

func main() {
	f.Shout()
	bf.Shout()
}
