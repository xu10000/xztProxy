package config

type Config struct {
	Host        string   `json:"host"`
	BeginPort   int      `json:"beginPort"`
	PasswordArr []string `json:"passwordArr"`
	PortNumber  int      `json:"portNumber"`
}
