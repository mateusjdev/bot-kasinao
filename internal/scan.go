package internal

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func ScanFiles(file_path string) (map[string]bool, error) {
	files := map[string]bool{}
	err := filepath.WalkDir(
		file_path,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			// todo: verificar se é vídeo
			if !d.IsDir() {
				files[path] = true
			}

			return nil
		},
	)
	fmt.Println(files)
	if err != nil {
		return nil, err
	}
	return files, nil
}
