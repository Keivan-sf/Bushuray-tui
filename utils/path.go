package utils

import (
	"fmt"
	"log"
	"os"
	osuser "os/user"
	"strconv"
)

func GetHomeDir() (string, error) {
	uid := os.Getuid()
	if uid == 0 {
		real_uid, err := strconv.Atoi(os.Getenv("SUDO_UID"))
		if err != nil {
			return "", fmt.Errorf("failed to get user id outside of sudo %w", err)
		}
		uid = real_uid
	}

	user, err := osuser.LookupId(strconv.Itoa(uid))
	if err != nil {
		log.Fatal("failed to get user from uid")
		return "", fmt.Errorf("failed to get user from uid %d: %w", uid, err)
	}
	log.Printf("using home directory:%v\n", user.HomeDir)
	return user.HomeDir, nil
}
