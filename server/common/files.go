package common

import (
	"os"
	"path/filepath"
	"strings"
)

var MOCK_CURRENT_DIR string

func GetExecutableDir() string {
	if MOCK_CURRENT_DIR != "" {
		return MOCK_CURRENT_DIR
	}
	ex, _ := os.Executable()
	return filepath.Dir(ex)
}

func getDirFromEnvironment(key string) string {
	if dir := os.Getenv(key); len(dir) != 0 {
		return dir
	} else {
		return GetExecutableDir()
	}
}

func GetConfigDir() string {
	return getDirFromEnvironment("FILESTASH_CONFIG_DIR")
}

func GetCacheDir() string {
	return getDirFromEnvironment("FILESTASH_CACHE_DIR")
}

func GetLogDir() string {
	return getDirFromEnvironment("FILESTASH_LOG_DIR")
}

func GetPublicDataDir() string {
	return getDirFromEnvironment("FILESTASH_PUBLIC_DATA_DIR")
}

func GetDefaultConfigDir() string {
	return getDirFromEnvironment("FILESTASH_DEFAULT_CONFIG_DIR")
}

func GetAbsolutePath(base, p string) string {
	if path, err := filepath.Abs(filepath.Join(base, p)); err == nil {
		return path
	} else {
		panic(err)
	}
}

func IsDirectory(path string) bool {
	if path == "" {
		return false
	}
	if path[len(path)-1:] != "/" {
		return false
	}
	return true
}

/*
 * Join 2 path together, result has a file
 */
func JoinPath(base, file string) string {
	filePath := filepath.Join(base, file)
	if strings.HasPrefix(filePath, base) == false {
		return base
	}
	return filePath
}

func EnforceDirectory(path string) string {
	if path == "" {
		return "/"
	} else if path[len(path)-1:] == "/" {
		return path
	}
	return path + "/"
}
