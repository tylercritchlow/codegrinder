package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func CommandConfig(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		mustLoadConfig(cmd)
		if Config.Root != "" {
			fmt.Printf("root: %s\n", Config.Root)
		} else {
			fmt.Println("root: (default: ~)")
		}
		return
	}

	if len(args) != 2 || args[0] != "root" {
		cmd.Help()
		os.Exit(1)
	}

	path, err := filepath.Abs(args[1])
	if err != nil {
		log.Fatalf("invalid path: %v", err)
	}

	mustLoadConfig(cmd)
	Config.Root = path
	mustWriteConfig()
	fmt.Printf("root directory set to: %s\n", Config.Root)
}
