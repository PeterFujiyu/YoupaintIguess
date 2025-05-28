package main

import "embed"

//go:embed public/* public/i18n/* themes/*.json
var content embed.FS
