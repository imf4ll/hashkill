package cmd

import (
    "os"
    "log"
    "bufio"
    "fmt"
)

func crack(algo, wordlist, hash, ftype string) {
    if ftype == "file" {
        hashFile, err := os.Open(hash)
        if err != nil {
            log.Fatal("Unable to read the hashs file.")

        }

        defer hashFile.Close()
    
        scannerHashs := bufio.NewScanner(hashFile)
        for scannerHashs.Scan() {
            wg.Add(1)
            
            go grep(scannerHashs.Text(), fmt.Sprintf("%v/.config/hashkill/%v.txt", os.Getenv("HOME"), algo), "file")
        }

        wg.Wait()
    
    } else {
        grep(hash, fmt.Sprintf("%v/.config/hashkill/%v.txt", os.Getenv("HOME"), algo), "hash")

    }
}
