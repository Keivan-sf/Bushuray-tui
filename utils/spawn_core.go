package utils

import (
	"os/exec"
	"syscall"
)

func SpawnBushurayCore() error {
	cmdPath, err := GetBinPath("bushuray-core")
	if err != nil {
		return err
	}

	cmd := exec.Command(cmdPath)

	cmd.Stdin = nil
	cmd.Stdout = nil
	cmd.Stderr = nil

	cmd.SysProcAttr = detachedSysProcAttr()

	return cmd.Start()
}

func detachedSysProcAttr() *syscall.SysProcAttr {
	return &syscall.SysProcAttr{
		Setsid: true,
	}
}
