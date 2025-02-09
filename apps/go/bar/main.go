package main

import (
	bf "github.com/Sengoku11/bazel-monorepo/libs/go/baz/foo/pkg"
	f "github.com/Sengoku11/bazel-monorepo/libs/go/foo/pkg"
)

func main() {
	f.Shout()
	bf.Shout()

	// Call this to trigger linter:
	//fmt.Printf("This should trigger a vet warning: %d\n", "not a number")
}
