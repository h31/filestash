package common

import (
	"io/ioutil"
	slog "log"
	"os"
	"path/filepath"
)

const (
	APP_VERSION       = "v0.4rc"
	LOG_PATH          = "data/state/log/"
	CONFIG_PATH       = "data/state/config/"
	DB_PATH           = "data/state/db/"
	FTS_PATH          = "data/state/search/"
	CERT_PATH         = "data/state/certs/"
	TMP_PATH          = "data/cache/tmp/"
	COOKIE_NAME_AUTH  = "auth"
	COOKIE_NAME_PROOF = "proof"
	COOKIE_NAME_ADMIN = "admin"
	COOKIE_PATH_ADMIN = "/admin/api/"
	COOKIE_PATH       = "/api/"
	FILE_INDEX        = "./data/public/index.html"
	FILE_ASSETS       = "./data/public/"
	URL_SETUP         = "/admin/setup"
)

func init() {
	os.MkdirAll(filepath.Join(GetLogDir(), LOG_PATH), os.ModePerm)
	os.MkdirAll(filepath.Join(GetConfigDir(), FTS_PATH), os.ModePerm)
	os.RemoveAll(filepath.Join(GetCacheDir(), TMP_PATH))
	os.MkdirAll(filepath.Join(GetCacheDir(), TMP_PATH), os.ModePerm)

	if _, err := os.Stat(filepath.Join(GetConfigDir(), CONFIG_PATH)); os.IsNotExist(err) {
		CopyDefaultConfig()
	}
}

var (
	BUILD_REF                     string
	BUILD_DATE                    string
	SECRET_KEY                    string
	SECRET_KEY_DERIVATE_FOR_PROOF string
	SECRET_KEY_DERIVATE_FOR_ADMIN string
	SECRET_KEY_DERIVATE_FOR_USER  string
	SECRET_KEY_DERIVATE_FOR_HASH  string
)

/*
 * Improve security by calculating derivative of the secret key to restrict the attack surface
 * in the worst case scenario with one compromise secret key
 */
func InitSecretDerivate(secret string) {
	SECRET_KEY = secret
	SECRET_KEY_DERIVATE_FOR_PROOF = Hash("PROOF_"+SECRET_KEY, len(SECRET_KEY))
	SECRET_KEY_DERIVATE_FOR_ADMIN = Hash("ADMIN_"+SECRET_KEY, len(SECRET_KEY))
	SECRET_KEY_DERIVATE_FOR_USER = Hash("USER_"+SECRET_KEY, len(SECRET_KEY))
	SECRET_KEY_DERIVATE_FOR_HASH = Hash("HASH_"+SECRET_KEY, len(SECRET_KEY))
}

func CopyDefaultConfig() {
	destDir := filepath.Join(GetConfigDir(), CONFIG_PATH)
	srcDir := filepath.Join(GetDefaultConfigDir(), CONFIG_PATH)
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		slog.Printf("ERROR creating config directory: %+v", err)
		return
	}
	defaultConfigEntries, err := ioutil.ReadDir(srcDir)
	if err != nil {
		slog.Printf("ERROR accessing default config directory: %+v", err)
		return
	}
	for _, entry := range defaultConfigEntries {
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(destDir, entry.Name())

		content, err := ioutil.ReadFile(srcPath)
		if err != nil {
			slog.Printf("ERROR reading default config: %+v", err)
			return
		}
		err = ioutil.WriteFile(dstPath, content, os.ModePerm)
		if err != nil {
			slog.Printf("ERROR writing default config: %+v", err)
			return
		}
	}
}
