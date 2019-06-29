// Copyright © 2019 Adir Tzuberi <adir85@gmail.com>
package sopcast

import (
	"errors"
	"os/exec"
	"os/user"
	"path"
)

const (
	Emulator = "wine"
	Sopcast  = ".wine/drive_c/Program Files (x86)/SopCast/SopCast.exe"
	Broker   = "sop://broker.sopcast.com:3912"
)

func Stream(url string) error {
	currentUser, err := user.Current()
	if err != nil {
		return errors.New("Failed to get current user's home directory: " + err.Error())
	}
	sopcastPath := path.Join(currentUser.HomeDir, Sopcast)
	cmd := exec.Command(Emulator, sopcastPath, url)
	if err := cmd.Run(); err != nil {
		return errors.New("Failed to run SopCast: " + err.Error())
	}
	return nil
}
