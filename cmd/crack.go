package cmd

import (
	"bufio"
	"os"
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
            
            go find(scannerHashs.Text(), wordlist, "file")
        }

        wg.Wait()
    
    } else {
        find(hash, wordlist, "hash")

    }
}
