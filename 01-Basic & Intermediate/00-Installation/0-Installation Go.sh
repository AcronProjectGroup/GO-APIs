# Remove any previous Go installation
sudo rm -rf /usr/local/go 

sudo tar -C /usr/local -xzf go1.20.5.linux-amd64.tar.gz

echo $PATH | grep "/usr/local/go/bin"

# for config to all directories:
# echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
# echo 'export PATH=$GOPATH/bin:$GOROOT/bin:$PATH' >> ~/.bashrc
# echo 'export GOPATH=$HOME/go' >> ~/.bashrc
# source ~/.profile


# Or use this for only you want to use Go specific directory:
# export GOROOT=/usr/local/go
# export GOPATH=$HOME/go
# export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

