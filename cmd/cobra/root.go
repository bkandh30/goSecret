package cobra

import (
	"path/filepath"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
)

var RootCmd = &cobra.Command{
	Use:   "goSecret",
	Short: "goSecret is an API key and other secrets manager like Hashicorp Vault",
}

var encodingKey string

func init() {
	RootCmd.PersistentFlags().StringVarP(&encodingKey, "key", "k", "", "Key for encoding and decoding secrets")
}

func secretsPath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, ".secrets")
}
