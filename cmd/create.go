package cmd

import (
    "os"
    "bufio"
    "fmt"
)

func create(algo, wordlist, hashedwl string) {
    file, err := os.Open(wordlist)
    if err != nil {
        l.Fatal("\033[1;31m[-] Undefined wordlist, use '--wordlist' or '-w' to define.\033[m")

    }

    defer file.Close()

    newFile, err := os.Create(hashedwl)
    if err != nil {
        l.Fatal("\033[1;31m[-] Unable to create the hashed wordlist file.\033[m")

    }

    defer newFile.Close()

    fmt.Println("\033[1;32m[+] Creating hashed wordlist named\033[1;33m", hashedwl, "\033[m")

    scanner := bufio.NewScanner(file)
    writer := bufio.NewWriter(newFile)
    for scanner.Scan() {
        writer.WriteString(fmt.Sprintf("%v:%v\n", Crypto(algo, scanner.Text()), scanner.Text()))
    
    }
}
