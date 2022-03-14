package cmd

import (
    "log"
    "os/exec"
    "fmt"
)

func grep(hash, wordlist string) {
    if hash  == "" || wordlist == "" {
        log.Fatal("Hash or wordlist cannot be empty.")

    }

    cmd := exec.Command("grep", hash, wordlist)
    stdout, err := cmd.Output()
    if err != nil {
        if fmt.Sprintf("%v", err) != "exit status 1" {
            log.Fatal("Wordlist not exists in wordlists folder.")
        
        }
    }

    if string(stdout) != "" {
        fmt.Printf("\033[32m[+] \033[m%v", string(stdout))
    
    }

    wg.Done()
}
