package handler

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
)

//go:embed www/public
var publicDir embed.FS

func NewStaticHandler() (http.Handler, error) {
	files, err := fs.Sub(publicDir, "www/public")
	if err != nil {
		return nil, fmt.Errorf("failed to open embedded public files: %w", err)
	}

	return http.FileServerFS(files), nil
}
