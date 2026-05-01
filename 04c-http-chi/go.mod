module github.com/ubgo/buildinfo-examples/04c-http-chi

go 1.24

require (
	github.com/go-chi/chi/v5 v5.1.0
	github.com/ubgo/buildinfo/contrib/buildinfo-chi v0.0.0
)

require github.com/ubgo/buildinfo v0.0.0 // indirect

replace (
	github.com/ubgo/buildinfo => ../../buildinfo
	github.com/ubgo/buildinfo/contrib/buildinfo-chi => ../../buildinfo/contrib/buildinfo-chi
)
