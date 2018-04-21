cat brew/deps.txt | xargs brew install
cat brew/cask.txt | xargs brew cask install
mkdir -p ~/.config/fish
cat fish/config.fish > ~/.config/fish/config.fish
cp -r fish/functions ~/.config/fish/functions
cat git/gitconfig > ~/.gitconfig
cat git/gitignore > ~/.gitignore
cat hyper/hyper.js > ~/.hyper.js
