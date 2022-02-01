package xorm

type Config struct {
	DSN                  string `json:"dsn" yaml:"dsn"`
	EnableSSL            bool   `json:"enable_ssl" yaml:"enable_ssl"`
	CertificateWritePath string `json:"certificate_write_path" yaml:"certificate_write_path"`
	ClientKey            string `json:"client_key" yaml:"client_key"`
	ClientSecret         string `json:"client_secret" yaml:"client_secret"`
	ServerCA             string `json:"server_ca" yaml:"server_ca"`
}
