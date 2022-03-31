package cmd

import (
    "fmt"
    "log"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "golang.org/x/crypto/sha3"
    "golang.org/x/crypto/blake2s"
    "golang.org/x/crypto/blake2b"
)

func Crypto(algo, word string) string {
    result := ""

    switch algo {
        case "md5": result = fmt.Sprintf("%x", md5.Sum([]byte(word)))
        case "sha1": result = fmt.Sprintf("%x", sha1.Sum([]byte(word)))
        case "sha224": result = fmt.Sprintf("%x", sha256.Sum224([]byte(word)))
        case "sha256": result = fmt.Sprintf("%x", sha256.Sum256([]byte(word)))
        case "sha384": result = fmt.Sprintf("%x", sha512.Sum384([]byte(word)))
        case "sha512": result = fmt.Sprintf("%x", sha512.Sum512([]byte(word)))
        case "sha3_224": result = fmt.Sprintf("%x", sha3.Sum224([]byte(word)))
        case "sha3_256": result = fmt.Sprintf("%x", sha3.Sum256([]byte(word)))
        case "sha3_384": result = fmt.Sprintf("%x", sha3.Sum384([]byte(word)))
        case "sha3_512": result = fmt.Sprintf("%x", sha3.Sum512([]byte(word)))
        case "sha512_224": result = fmt.Sprintf("%x", sha512.Sum512_224([]byte(word)))
        case "sha512_256": result = fmt.Sprintf("%x", sha512.Sum512_256([]byte(word)))
        case "blake2s": result = fmt.Sprintf("%x", blake2s.Sum256([]byte(word)))
        case "blake2b_256": result = fmt.Sprintf("%x", blake2b.Sum256([]byte(word)))
        case "blake2b_384": result = fmt.Sprintf("%x", blake2b.Sum384([]byte(word)))
        case "blake2b_512": result = fmt.Sprintf("%x", blake2b.Sum512([]byte(word)))
        default: log.Fatal("\033[1;31m[✘]\033[m Invalid hashtype.")
    }
    
    return result
}
