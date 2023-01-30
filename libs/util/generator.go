package util

import (
	"bytes"
	"os/exec"

	"github.com/sirupsen/logrus"
)

func GanerateUUID() string {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		logrus.Fatal(err)
	}
	newUUID = bytes.TrimSuffix(newUUID, []byte{10})
	return string(newUUID)
}
