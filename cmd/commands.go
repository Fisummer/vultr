package cmd

import (
	"fmt"

	"github.com/jawher/mow.cli"
)

func (c *CLI) RegisterCommands() {
	// sshkeys
	c.Command("sshkey", "control SSH public keys on Vultr account", func(cmd *cli.Cmd) {
		cmd.Command("create", "upload and add new SSH public key to your Vultr account", printApiKey)
		cmd.Command("update", "update an existing SSH public key in your Vultr account", printApiKey)
		cmd.Command("delete", "remove an existing SSH public key from your Vultr account", printApiKey)
		cmd.Command("list", "list all SSH public keys in Vultr account", printApiKey)
	})
}

// for debugging..
func printApiKey(cmd *cli.Cmd) {
	cmd.Action = func() {
		fmt.Println(*apiKey)
	}
}