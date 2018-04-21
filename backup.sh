brew list > brew/deps.txt
brew cask list > brew/cask.txt
rm -rf fish/*
cat ~/.config/fish/config.fish > fish/config.fish
cp -r ~/.config/fish/functions fish/functions
cat ~/.gitconfig > git/gitconfig
cat ~/.gitignore > git/gitignore
cat ~/.hyper.js > hyper/hyper.js
