// Copyright © 2019 Adir Tzuberi <adir85@gmail.com>
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "sopcast",
	Short: "A wrapper application for running SopCast using Wine",
	Long: `SopCast is a desktop p2p streaming application that uses channel numbers.
This wrapper application runs the latest Windows SopCast binary (v.4.2.0) over Wine
since the latest Linux binary is broken.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
