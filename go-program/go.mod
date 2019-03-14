module angenalZZZ/go-program // 新建项目时,选择go-module: set GO111MODULE=on && go mod init

require (
	github.com/chai2010/winsvc latest
	github.com/gobuffalo/envy v1.6.15
	github.com/kr/pty v1.1.3 // indirect
	github.com/rogpeppe/go-internal v1.2.2 // indirect
	github.com/stretchr/objx v0.1.1 // indirect
)

replace (
	golang.org/x/sys/windows latest => github.com/golang/sys/windows latest
)
