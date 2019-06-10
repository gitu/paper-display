# paper-display


# Dependencies
## GO
```
wget https://storage.googleapis.com/golang/go1.12.5.linux-armv6l.tar.gz
sudo tar -C /usr/local -xvf go1.12.5.linux-armv6l.tar.gz
cat >> ~/.bashrc << 'EOF'
export GOPATH=$HOME/go
export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin
EOF
source ~/.bashrc
```
   
### Check Version 
```
go version
```



## Install Paper Display
```
go get -u github.com/gitu/paper-display
sudo ln -s /home/pi/go/src/github.com/gitu/paper-display/paper-display.service /lib/systemd/system/paper-display.service
sudo systemctl daemon-reload
sudo systemctl enable paper-display.service
```

