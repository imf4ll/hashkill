host=

install:
	echo "Installing and setup hashkill..."
	go install
	echo "Creating hashkill folder in .config"
	mkdir ${HOME}/.config/hashkill
	echo "Setup finished, enter 'make download <Algorithm>' to download hashlists or 'make downloadall' to download all hashlists."

downloadall:
	echo "About 20GB download..."
	wget ${host}/md5.txt -O ${HOME}/.config/hashkill/md5.txt
	wget ${host}/sha1.txt -O ${HOME}/.config/hashkill/sha1.txt
	wget ${host}/sha224.txt -O ${HOME}/.config/hashkill/sha224.txt
	wget ${host}/sha256.txt -O ${HOME}/.config/hashkill/sha256.txt
	wget ${host}/sha384.txt -O ${HOME}/.config/hashkill/sha384.txt
	wget ${host}/sha512.txt -O ${HOME}/.config/hashkill/sha512.txt
	wget ${host}/sha3_224.txt -O ${HOME}/.config/hashkill/sha3_224.txt
	wget ${host}/sha3_256.txt -O ${HOME}/.config/hashkill/sha3_256.txt
	wget ${host}/sha3_384.txt -O ${HOME}/.config/hashkill/sha3_384.txt
	wget ${host}/sha3_512.txt -O ${HOME}/.config/hashkill/sha3_512.txt
	wget ${host}/sha512_224.txt -O ${HOME}/.config/hashkill/sha512_224.txt
	wget ${host}/sha512_256.txt -O ${HOME}/.config/hashkill/sha512_256.txt
	wget ${host}/blake2s.txt -O ${HOME}/.config/hashkill/blake2s.txt
	wget ${host}/blake2b_256.txt -O ${HOME}/.config/hashkill/blake2b_256.txt
	wget ${host}/blake2b_384.txt -O ${HOME}/.config/hashkill/blake2b_384.txt
	wget ${host}/blake2b_512.txt -O ${HOME}/.config/hashkill/blake2b_512.txt

downloadmd5:
	wget ${host}/md5.txt -O ${HOME}/.config/hashkill/md5.txt

downloadsha1:
	wget ${host}/sha1.txt -O ${HOME}/.config/hashkill/sha1.txt

downloadsha224:
	wget ${host}/sha224.txt -O ${HOME}/.config/hashkill/sha224.txt

downloadsha256:
	wget ${host}/sha256.txt -O ${HOME}/.config/hashkill/sha256.txt

downloadsha384:
	wget ${host}/sha384.txt -O ${HOME}/.config/hashkill/sha384.txt

downloadsha512:
	wget ${host}/sha512.txt -O ${HOME}/.config/hashkill/sha512.txt

downloadsha3_224:
	wget ${host}/sha3_224.txt -O ${HOME}/.config/hashkill/sha3_224.txt

downloadsha3_256:
	wget ${host}/sha3_256.txt -O ${HOME}/.config/hashkill/sha3_256.txt

downloadsha3_384:
	wget ${host}/sha3_384.txt -O ${HOME}/.config/hashkill/sha3_384.txt

downloadsha3_512:
	wget ${host}/sha3_512.txt -O ${HOME}/.config/hashkill/sha3_512.txt

downloadsha512_224:
	wget ${host}/sha512_224.txt -O ${HOME}/.config/hashkill/sha512_224.txt

download512_256:
	wget ${host}/sha512_256.txt -O ${HOME}/.config/hashkill/sha512_256.txt

downloadblake2s:
	wget ${host}/blake2s.txt -O ${HOME}/.config/hashkill/blake2s.txt

downloadblake2b_256:
	wget ${host}/blake2b_256.txt -O ${HOME}/.config/hashkill/blake2b_256.txt

downloadblake2b_384:
	wget ${host}/blake2b_384.txt -O ${HOME}/.config/hashkill/blake2b_384.txt

downloadblake2b_512:
	wget ${host}/blake2b_512.txt -O ${HOME}/.config/hashkill/blake2b_512.txt
