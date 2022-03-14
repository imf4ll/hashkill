install:
	echo "Installing and setup hashkill..."
	go install
	echo "Creating hashkill folder in .config"
	mkdir ${HOME}/.config/hashkill
	echo "Setup finished, enter 'make download' to download stantard hashlists (md5 and sha1)."

download:
	echo "About 700MB+ download..."
	wget "https://cdn-103.anonfiles.com/18JbdeO2xd/1c5d555e-1647231058/wordlists.tar.gz"
	echo "Decompressing..."
	tar -xvf wordlists.tar.gz -C ${HOME}/.config/hashkill/
