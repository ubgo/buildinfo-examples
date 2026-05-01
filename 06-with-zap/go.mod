module github.com/ubgo/buildinfo-examples/06-with-zap

go 1.24

require (
	github.com/ubgo/buildinfo/contrib/buildinfo-zap v0.0.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/ubgo/buildinfo v0.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
)

replace (
	github.com/ubgo/buildinfo => ../../buildinfo
	github.com/ubgo/buildinfo/contrib/buildinfo-zap => ../../buildinfo/contrib/buildinfo-zap
)
