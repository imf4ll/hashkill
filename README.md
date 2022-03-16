```
    ______               ______ ______ ________________
    ___  /_______ __________  /____  /____(_)__  /__  /
    __  __ \  __ `/_  ___/_  __ \_  //_/_  /__  /__  / 
    _  / / / /_/ /_(__  )_  / / /  ,<  _  / _  / _  /  
    /_/ /_/\__,_/ /____/ /_/ /_//_/|_| /_/  /_/  /_/
```

<br><br>
## ♻️ Changelog v0.1.2
- Performance update (Reduced average 42s to 27s on creating, about +65% speed, remembering, we use only CPU to generated the hashs, for now)
- Added command 'version'
- Some fixes

<br><br>
## ❓ What is this?
- A simple golang script to generate pre-hashed files to speed up password cracking, using standard methods to speed up hash locating.

<br><br>
## 🤖 Instructions:
- Download done hashlists:
`make download` (Up to 700MB+ download)

- Create custom pre-hashed files:
`hashkill -c -w <wordlist> -a <algorithm> -n <file name>`

- Crack hashs (unique):
`hashkill -C -a <algorithm> -H <hash>`

- Crack hashs (multiple/files):
`hashkill -C -a <algorithm> -f <file path>`

<br><br>
## ✅ Available hashtypes:
- md5
- sha1
- sha224
- sha256
- sha384
- sha512
- sha3_224
- sha3_256
- sha3_384
- sha3_512
- sha512_224
- sha512_256
- blake2s
- blake2b_256
- blake2b_384
- blake2b_512

<br><br>
## 🔧 Setup:
### Clone this repository:<br>
`git clone https://github.com/z3oxs/hashkill` or Download ZIP and unzip;<br><br>
### Move to repository:<br>
`cd hashkill`<br><br>
### Config:<br>
`make install`
### Run:<br>
`hashkill --help`
