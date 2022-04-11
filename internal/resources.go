package resources

import "embed"

var (
	//go:embed templates
	Templates embed.FS

	//go:embed static
	Statics embed.FS
)
