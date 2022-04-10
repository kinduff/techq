package resources

import "embed"

var (
	//go:embed views/*
	Views embed.FS
)
