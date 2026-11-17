package exampleplugin

//go:generate go tool github.com/go-app-blazar/blazar/tool/generate-mixin-files --source-directory=./internal/base/example

//go:generate go tool github.com/go-app-blazar/blazar/tool/generate-interface-package --source-directory=./internal/base/example --target-directory=./example --private-prefix=example --public-prefix=Example
