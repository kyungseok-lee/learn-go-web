# Go로 만드는 웹
- 원문: Tucker의 GoLang 프로그래밍
  - github: https://github.com/tuckersGo/goWeb
  - youtube: https://www.youtube.com/channel/UCZp_ftx6UB_32VfVmlS3o_A

## go mod
```
go mod init github.com/kyungseok-lee/learn-go-web
go mod tidy
go list -m all
```

## Test
- [GoConvey](https://github.com/smartystreets/goconvey)
```
go get github.com/smartystreets/goconvey
$GOPATH/bin/goconvey
```

- [testify](https://github.com/stretchr/testify)
```
go get github.com/stretchr/testify/assert
```
```
github.com/stretchr/testify/assert
github.com/stretchr/testify/require
github.com/stretchr/testify/mock
github.com/stretchr/testify/suite
github.com/stretchr/testify/http (deprecated)
```