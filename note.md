# note for go mod, go work, git submodule

## 初始化 go workspace

- `go work ini` 初始化go workspace

## 初始化 go module

- `mkdir example`
- `cd example`
- `go mod init example`
- `vi main.go`
- `go get https://github.com/xmh19936688/go-dispatcher`
- `cd ..`

`go get` 需要在module下执行。
`go mod tidy` 需要在module下执行。

## 将module添加到workspace

- `go work use ./example` 添加module

## 添加git子模块

- `git submodule add https://github.com/xmh19936688/go-dispatcher ./go-dispatcher` 添加子模块

## go-dispatcher git子模块配置文件

`.git/modules/go-dispatcher/config`

## 使用本地依赖 go-dispatcher

- `go work use ./go-dispatcher/` 使用本地依赖

## 进入git子模块编辑并提交

- `cd go-dispatcher`
- `git checkout -b test`
- `vi dispatcher/dispatcher.go`
- `git commit -m "add func for test"`
- `git push origin test`

## 在外层记录所用的git子模块分支

git只能记录所用的submodule的commit号，
也是根据submodule的head指向的commit号来判断是否clean。
`.gitmodules`中的branch是human-face的，git并不识别。

- `config -f .gitmodules submodule.go-dispatcher.branch test`
- `git add .gitmodules`
- `git commit -m "change submodule branch"`
