package assets

import (
	"embed"
	"io"

	resource "github.com/quasilyte/ebitengine-resource"
)

//go:embed all:_data
var gameAssets embed.FS

func OpenAsset(path string) io.ReadCloser {
	f, err := gameAssets.Open("_data/" + path)
	if err != nil {
		panic(err)
	}
	return f
}

func RegisterResources(loader *resource.Loader) {
	RegisterImageResource(loader)
}
