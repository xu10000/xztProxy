
### switchomega 当前文件夹有

### 配置文件的密码数组怎么生成
config/generatePasswd.go 
go test > log.txt

### mac加入开机自启脚本
http://t.zoukankan.com/dongfangzan-p-5976791.html


### 编译时先生成图标 - 要在windows下编译才行
go generate
### 编译linux
set -x CGO_ENABLED 0; set -x  GOOS linux; set -x  GOARCH amd64;  go build -o xztProxyLinux
### 编译windows
set -x CGO_ENABLED 0; set -x  GOOS windows; set -x  GOARCH amd64;  go build -o xztProxyWin.exe
### 编译mac
go build -o xztProxyMac
