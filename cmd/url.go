// Copyright © 2019 Adir Tzuberi <adir85@gmail.com>
package cmd

import (
	"errors"
	"fmt"
	"github.com/adirt/sopcast/sopcast"
	"github.com/spf13/cobra"
	"strings"
)

var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "SopCast URL to stream",
	Example: fmt.Sprintf("sopcast url %s/123456", sopcast.Broker),
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
        url := args[0]
        if err := validateUrl(url); err != nil {
        	return err
		}
		return sopcast.Stream(url)
	},
}

func validateUrl(url string) error {
	if !strings.HasPrefix(url, sopcast.Broker) {
		return errors.New("A SopCast URL is a broker address followed by a 6-digit channel number")
	}
	channelOffset := strings.LastIndex(url, "/") + 1
	channel := url[channelOffset:]
	return validateChannel(channel)
}

func init() {
	rootCmd.AddCommand(urlCmd)
}
