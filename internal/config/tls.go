package config

type TLSConfig struct {
	ReqSSL   bool   `json:"require_ssl" toml:"require_ssl" yaml:"require_ssl"`
	KeyFile  string `json:"key_file" toml:"key_file" yaml:"key_file"`
	CertFile string `json:"cert_file" toml:"cert_file" yaml:"cert_file"`
}