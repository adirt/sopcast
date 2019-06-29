// Copyright © 2019 Adir Tzuberi <adir85@gmail.com>
package cmd

import (
	"errors"
	"fmt"
	"github.com/adirt/sopcast/sopcast"
	"github.com/spf13/cobra"
	"strconv"
)

var channelCmd = &cobra.Command{
	Use:     "channel",
	Short:   "SopCast channel to stream",
	Example: "sopcast channel 123456",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		channel := args[0]
		if err := validateChannel(channel); err != nil {
			return err
		}
		url := fmt.Sprintf("%s/%s", sopcast.Broker, channel)
		return sopcast.Stream(url)
	},
}

func validateChannel(channel string) error {
	channelNumber, err := strconv.Atoi(channel)
	if err != nil || channelNumber < 100000 || channelNumber > 999999 {
		return errors.New("A SopCast channel is a 6-digit number")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(channelCmd)
}
