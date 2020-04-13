# dockerhub-go
CLI tool for searching images on Docker Hub

### Building
```
go build -o dockerhub src/main.go
```

### Sample Commands

#### Help
```
$ dockerhub help
```

#### Search
```
$ dockerhub search --help

$ dockerhub search mysql
```

#### Info
```
$ dockerhub info --help

$ dockerhub info mysql
```

#### Tags
```
$ dockerhub tags --help

$ dockerhub tags mysql
```
