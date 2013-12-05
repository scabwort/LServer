package utils

import (
	"os"
	"os/exec"
	"path/filepath"
)

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	return path
}
