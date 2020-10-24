# jtl
Useful command line tool for developers

# Build & Install
```bash
go get github.com/parkjangwon/jtl
go build jtl.go

cp ./jtl /usr/local/bin; chmod a+x /usr/local/bin/jtl
```

# If you want to build on a different platform
```bash
go get github.com/parkjangwon/jtl

# for linux x86_64
cd $GOPATH/src/github.com/parkjangwon/jtl
env GOOS=linux GOARCH=amd64 go build -v jtl.go

cp ./jtl /usr/local/bin; chmod a+x /usr/local/bin/jtl
```

# How to Usage
```bash
# jtl help

jtl <command> [arguments]
```

# Functions Supported
```bash
# Base64 Encoder : Outputs the Base64 encoded result of the arguments.
jtl b64e foo
Zm9v

# Base64 Decoder: Outputs the Base64 decoded result of the arguments.
jtl b64d Zm9v
foo

# Telnet Client : This is a telnet client. argument must in the form "host:port".
jtl tel aspmx.l.google.com:25
220 mx.google.com ESMTP bj14si3555710pjb.77 - gsmtp
helo aa
250 mx.google.com at your service
........
........
quit
```
