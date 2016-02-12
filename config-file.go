package nodester

import "os"

// ConfigDir returns the configuration directory for Packer.
func ConfigDir() (string, error) {

	if nodesterRoot := os.Getenv("NODESTER_ROOT"); nodesterRoot != "" {
		return nodesterRoot, nil
	}
	return configDir()

}
