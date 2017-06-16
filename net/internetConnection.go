package net

import (
  "os/exec"
  "strings"
)

func isInternetConnected () bool {
	out, _ := exec.Command("ping", "google.com", "-n", "1").Output()

  if strings.Contains(fmt.Sprintf("%s", out), "ms") {
    return true
  } else {
    return false
  }
}
