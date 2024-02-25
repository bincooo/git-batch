# git-batch
Automatically synchronize files with git

Run with Binary
```
./git-batch \
 -p ~/project \
 --commit 10s \
 --name bincooo \
 --email bincooo@gmail.com \
 --push 1m 
```

### build
```bash
# Mac
GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o git-batch -trimpath

# Linux
GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o git-batch -trimpath

#Win
GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o git-batch -trimpath
```