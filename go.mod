module github.com/go-app-blazar/example-plugin

go 1.26.2

tool (
	github.com/go-app-blazar/blazar/tool/generate-interface-package
	github.com/go-app-blazar/blazar/tool/generate-mixin-files
)

require (
	github.com/go-app-blazar/blazar v0.5.3
	github.com/maxence-charriere/go-app/v11 v11.0.4
)

require (
	github.com/alexflint/go-arg v1.6.1 // indirect
	github.com/alexflint/go-scalar v1.2.0 // indirect
	github.com/go-app-blazar/router v0.1.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
)
