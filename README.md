# paper-display





--> Install go on Raspian

wget https://storage.googleapis.com/golang/go1.10.1.linux-armv6l.tar.gz
sudo tar -C /usr/local -xvf go1.10.1.linux-armv6l.tar.gz
cat >> ~/.bashrc << 'EOF'
export GOPATH=$HOME/go
export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
EOF
source ~/.bashrc