package cmd

import (
    "os"
    "bufio"
    "fmt"
)

func crack(algo, wordlist, hash, ftype string) {
    if ftype == "file" {
        hashFile, err := os.Open(hash)
        if err != nil {
            l.Fatal("\033[1;31m[✘]\033[m Unable to read the hashs file.")

        }

        defer hashFile.Close()
    
        scannerHashs := bufio.NewScanner(hashFile)
        for scannerHashs.Scan() {
            wg.Add(1)
            
            go find(scannerHashs.Text(), fmt.Sprintf("%v/.config/hashkill/%v.txt", os.Getenv("HOME"), algo), "file")
        }

        wg.Wait()
    
    } else {
        find(hash, fmt.Sprintf("%v/.config/hashkill/%v.txt", os.Getenv("HOME"), algo), "hash")

    }
}
