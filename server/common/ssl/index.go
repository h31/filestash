package ssl

import (
	. "github.com/mickael-kerjean/filestash/server/common"
	"os"
	"path/filepath"
)

var keyPEMPath string = filepath.Join(GetConfigDir(), CERT_PATH, "key.pem")
var certPEMPath string = filepath.Join(GetConfigDir(), CERT_PATH, "cert.pem")

func init() {
	os.MkdirAll(filepath.Join(GetConfigDir(), CERT_PATH), os.ModePerm)
}

func Clear() {
	clearPrivateKey()
	clearCert()
}
