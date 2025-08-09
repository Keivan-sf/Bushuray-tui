package utils

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func SpawnBushurayCore() error {
	paths := []string{
		"./bushuray-core",
		filepath.Join(os.Getenv("HOME"), ".local", "bin", "bushuray-core"),
	}

	var cmdPath string
	for _, p := range paths {
		if _, err := os.Stat(p); err == nil {
			cmdPath = p
			break
		}
	}

	if cmdPath == "" {
		return errors.New("bushuray-core not found in any known paths")
	}

	cmd := exec.Command(cmdPath)

	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	cmd.SysProcAttr = detachedSysProcAttr()

	// Start (do not wait)
	return cmd.Start()
}

func detachedSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setsid: true, // Start in new session
	}
}
