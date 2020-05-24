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

func GetConfigDir() string {
	if configDir := os.Getenv("FILESTASH_CONFIG_DIR"); len(configDir) != 0 {
		return configDir
	} else {
		return GetExecutableDir()
	}
}

func GetCacheDir() string {
	if cacheDir := os.Getenv("FILESTASH_CACHE_DIR"); len(cacheDir) != 0 {
		return cacheDir
	} else {
		return GetExecutableDir()
	}
}

func GetLogDir() string {
	if logDir := os.Getenv("FILESTASH_LOG_DIR"); len(logDir) != 0 {
		return logDir
	} else {
		return GetExecutableDir()
	}
}

func GetPublicDataDir() string {
	if logDir := os.Getenv("FILESTASH_PUBLIC_DATA_DIR"); len(logDir) != 0 {
		return logDir
	} else {
		return GetExecutableDir()
	}
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
