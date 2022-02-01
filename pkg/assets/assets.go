package assets

import "embed"

//go:embed *.html js/* css/*
var Content embed.FS
