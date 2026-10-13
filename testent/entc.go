//go:build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/ubgo/entx/entcascade"
)

func main() {
	config := &gen.Config{
		Target:  "./ent",
		Package: "github.com/ubgo/entx/testent/ent",
	}
	opts := []entc.Option{
		entc.Extensions(
			entcascade.NewExtension(),
		),
	}
	if err := entc.Generate("./schema", config, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
