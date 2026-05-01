module github.com/ubgo/buildinfo-examples/07-with-slog

go 1.24

require github.com/ubgo/buildinfo/contrib/buildinfo-slog v0.0.0

require github.com/ubgo/buildinfo v0.0.0 // indirect

replace (
	github.com/ubgo/buildinfo => ../../buildinfo
	github.com/ubgo/buildinfo/contrib/buildinfo-slog => ../../buildinfo/contrib/buildinfo-slog
)
