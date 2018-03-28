function cleanup_mac_os
	brew update
brew upgrade
brew cleanup -s
brew cask cleanup
brew doctor
brew missing
npm update -g
end
