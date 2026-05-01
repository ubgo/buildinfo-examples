module github.com/ubgo/buildinfo-examples/04-http-nethttp

go 1.24

require github.com/ubgo/buildinfo/contrib/buildinfo-nethttp v0.0.0

require github.com/ubgo/buildinfo v0.0.0 // indirect

replace (
	github.com/ubgo/buildinfo => ../../buildinfo
	github.com/ubgo/buildinfo/contrib/buildinfo-nethttp => ../../buildinfo/contrib/buildinfo-nethttp
)
