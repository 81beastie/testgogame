package assets

import (
	_ "image/png"

	resource "github.com/quasilyte/ebitengine-resource"
)

const (
	ImageNone resource.ImageID = iota
	ImageGopher
)

func RegisterImageResource(loader *resource.Loader) {

	imageResources := map[resource.ImageID]resource.ImageInfo{
		ImageGopher: {Path: "images/gopher.png"},
	}

	for id, res := range imageResources {

		loader.ImageRegistry.Set(id, res)
	}

}
