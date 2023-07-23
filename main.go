package main

import (
	"os"

	"github.com/ismtabo/magus/context"
	"github.com/ismtabo/magus/filetree"
	"github.com/ismtabo/magus/fs"
	"github.com/ismtabo/magus/magic"
	"github.com/ismtabo/magus/manifest"
	"github.com/kr/pretty"
)

func main() {
	os.Setenv("WORLD", "World")
	defer os.Unsetenv("WORLD")

	ctx := context.New()
	ctx = ctx.WithCwd("./output")
	mpath := "./docs/examples/template-only.yaml"
	mf := manifest.Manifest{}
	err := manifest.Unmarshal(ctx, mpath, &mf)
	if err != nil {
		panic(err)
	}
	mgc := magic.FromManifest(mf)

	files, err := mgc.Render(ctx)
	if err != nil {
		panic(err)
	}

	pretty.Printf("Rendered %d files\n", len(files))
	for _, file := range files {
		pretty.Printf("File: %s\n", file.Path())
		pretty.Printf("Value: %s\n", file.Value())
	}

	if err := filetree.AssertNotHaveWriteConficts(ctx, files); err != nil {
		panic(err)
	}

	if err := fs.WriteFiles(ctx, files); err != nil {
		panic(err)
	}
}
