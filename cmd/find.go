package cmd

import (
	"fmt"
	"os/exec"
	"runtime"
)

func find(hash, wordlist, typeS string) {
	if typeS == "file" {
		defer wg.Done()
	}

	if hash == "" || wordlist == "" {
		l.Fatal("\033[1;31m[✘]\033[m Hash or wordlist cannot be empty.")

	}

	cmd := &exec.Cmd{}
	if runtime.GOOS != "windows" {
		cmd = exec.Command("grep", hash, wordlist)

	} else {
		cmd = exec.Command("findstr", "/i", hash, wordlist)

	}

	stdout, err := cmd.Output()
	if err != nil {
		if fmt.Sprintf("%v", err) != "exit status 1" {
			l.Fatal("\033[1;31m[✘]\033[m Wordlist not exists in wordlists folder.")

		}
	}

	if string(stdout) != "" {
		fmt.Printf("\033[1;32m[🗸]\033[m %v", string(stdout))

	}
}
