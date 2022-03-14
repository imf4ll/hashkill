package cmd

import (
    "bufio"
    "os"
    "log"
    "fmt"
    "os/exec"
    "sync"

    "github.com/z3oxs/hashkill/utils"
    "github.com/spf13/cobra"
)

var wg sync.WaitGroup

func Hashkill(c *cobra.Command, args []string) {
    algo, _ := c.Flags().GetString("algo")
    hashtype, _ := c.Flags().GetBool("hashtype")
    wordlist, _ := c.Flags().GetString("wordlist")
    hashedwl, _ := c.Flags().GetString("name")
    file, _ := c.Flags().GetString("file")
    hash, _ := c.Flags().GetString("hash")
    createM, _ := c.Flags().GetBool("create")
    crackM, _ := c.Flags().GetBool("crack")

    if hashtype {
        fmt.Println(utils.Hashtype())

        return
    }

    if createM {
        utils.Banner("Create", algo, hashedwl, wordlist, file, hash)
        create(algo, wordlist, hashedwl)

    } else if crackM {
        utils.Banner("Crack", algo, "", wordlist, file, hash)
        
        if file != "" {
            crack(algo, wordlist, file, "file")

        } else {
            crack(algo, wordlist, hash, "hash")

        }
    }

    fmt.Println("\n\033[33m[!] \033[mThanks for the preference :).")
}

func create(algo, wordlist, hashedwl string) {
    file, err := os.Open(wordlist)
    if err != nil {
        log.Fatal("Undefined wordlist, use '--wordlist' or '-w' to define.")

    }

    defer file.Close()

    newFile, err := os.Create(hashedwl)
    if err != nil {
        log.Fatal("Unable to create the hashed wordlist file.")

    }

    defer newFile.Close()

    fmt.Println("\033[32m[+] \033[mCreating hashed wordlist named\033[33m", hashedwl, "\033[m")

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        newFile.WriteString(fmt.Sprintf("%v:%v\n", Crypto(algo, scanner.Text()), scanner.Text()))
    
    }
}

func crack(algo, wordlist, hash, ftype string) {
    pwd, _ := os.Executable()
    
    if ftype == "file" {
        hashFile, err := os.Open(hash)
        if err != nil {
            log.Fatal("Unable to read the hashs file.")

        }

        defer hashFile.Close()
    
        scannerHashs := bufio.NewScanner(hashFile)
        for scannerHashs.Scan() {
            wg.Add(1)
            go grep(scannerHashs.Text(), fmt.Sprintf("%v/wordlists/%v.txt", pwd, algo))

        }

    } else {
        wg.Add(1)
        go grep(hash, fmt.Sprintf("%v/.config/hashkill/%v.txt", os.Getenv("HOME"), algo))

    }

    wg.Wait()
}

func grep(hash, wordlist string) {
    if hash  == "" || wordlist == "" {
        log.Fatal("Hash or wordlist cannot be empty.")

    }

    cmd := exec.Command("grep", hash, wordlist)
    stdout, err := cmd.Output()
    if err != nil {
        log.Fatal("Wordlist not exists in .config")
    }

    fmt.Printf("\033[32m[+] \033[m%v", string(stdout))

    wg.Done()
}
