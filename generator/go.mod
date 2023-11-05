module github.com/per1234/inoplatforms/generator

go 1.21

replace github.com/per1234/inoplatforms/registry/assets/go-registry => ../registry/assets/go-registry/

require (
	github.com/arduino/go-paths-helper v1.9.2
	github.com/per1234/inoplatforms/registry/assets/go-registry v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.7.0
	go.bug.st/relaxed-semver v0.11.0
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
