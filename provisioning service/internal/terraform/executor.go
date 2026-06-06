package terraform

import "os/exec"

func Apply() error {

	cmd := exec.Command(
		"terraform",
		"apply",
		"-auto-approve",
	)

	return cmd.Run()
}
