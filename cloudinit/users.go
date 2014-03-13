package cloudinit

type User struct {
	Name              string   `yaml:"name"`
	PasswordHash      string   `yaml:"passwd"`
	SSHAuthorizedKeys []string `yaml:"ssh-authorized-keys"`
}
