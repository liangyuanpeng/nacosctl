module github.com/liangyuanpeng/nacosctl

go 1.13

require (
	github.com/goreleaser/goreleaser v0.149.0 // indirect
	github.com/mitchellh/go-homedir v1.1.0
	github.com/nacos-group/nacos-sdk-go v1.0.3
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
)

replace github.com/liangyuanpeng/nacosctl/cmd/configuration => ../cmd/configuration
