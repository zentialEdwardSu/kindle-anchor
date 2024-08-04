# kindle-anchor

Kindle-anchor is a KUAL-based extension used to export local floder to LAN via WebDav protocal, which enables you to manage the file in the Kindle with your favorite WebDav Client. Server address wiil stay on the left-top of screen after staing up.

# Install

KUAL are request, bacause kindle-anchor is a KUAL-based extension :D .

Copy the kindle_anchor floder in the repo root to the extension directory.

# Configuration

WebDav will request username and password to access, default username and password can be found at ./kindle_anchor/start.sh

boot args can be edited in ./kindle_anchor/start.sh. All startup parameters can be obtained by reading main.go

# Build from Source

Make sure you have `go` tool-chain installed.

```shell
make linux/arm/7 
```

# Note

We are not responsible for any losses you incur as a result of using this extension.

So please follow the advice below to protect the data:
- Since the webdav service does not currently support the use of Https, please do not use it in public networks. 
- Please modify the initial username and password to ensure data security.

# Citation

The extension script(.sh) comes from [guo-yong-zhi/kindle-filebrowser](https://github.com/guo-yong-zhi/kindle-filebrowser/tree/main), filebrowser is a much-more powerful tool for file management.

