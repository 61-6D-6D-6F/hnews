# HNews

HNews is a cli app for reading stories and replies on the terminal from [Hacker News](https://news.ycombinator.com/news) using its official [API](https://github.com/HackerNews/API).

## Run:
```
go run cmd/hnews/hnews.go
```

## Test:
```
go test internal/*.go -v
```

## Build:
### Linux x86-64:
```
GOOS=linux GOARCH=amd64 go build -o hnews-linux-amd64 cmd/hnews/hnews.go
```

### macOS arm64:
```
GOOS=darwin GOARCH=arm64 go build -o hnews-macos-arm64 cmd/hnews/hnews.go
```

### Windows x86-64:
```
GOOS=windows GOARCH=amd64 go build -o hnews-windows-amd64.exe cmd/hnews/hnews.go
```

## Screens:

### List of stories:
<img width="656" alt="list" src="https://github.com/user-attachments/assets/4a479df7-8ffd-47b2-807f-e3bfe533e8e4" />

Commands:
- x - exit app
- n - next story page
- p - previous story page
- 1-9 - display story details

### Story details:
<img width="656" alt="details" src="https://github.com/user-attachments/assets/e48ec3a4-846c-4aa5-99a8-62e93e64849c" />

Commands:
- x - exit app
- b - go back to stories
- c - display comment

### Comments:
<img width="656" alt="comments" src="https://github.com/user-attachments/assets/6f8dfe49-8094-4cba-a100-00c649cd22f0" />

Commands:
- x - exit app
- n - next story page
- p - previous story page
- b - go back one layer
- r - display reply

## Acknowledgement:
The state of the app is WIP, in need of refactoring.
