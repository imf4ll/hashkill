package cmd

import (
    "os"
    "log"
    "bufio"
    "fmt"
)

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
    writer := bufio.NewWriter(newFile)
    for scanner.Scan() {
        writer.WriteString(fmt.Sprintf("%v:%v\n", Crypto(algo, scanner.Text()), scanner.Text()))
    
    }
}
