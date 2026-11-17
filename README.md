# example-plugin
This showcases how to create a simple Blazar plugin.

# Setup

Include the Blazar tools.

```
go get -tool github.com/go-app-blazar/blazar/tool/generate-mixin-files
go get -tool github.com/go-app-blazar/blazar/tool/generate-interface-package
```

Determine your package name (this one is `example`).

Put your actual components in an `internal` directory (this one is `internal/base/example`).

Create a top-level `generate.go` file that calls the Blazar tools.

```
//go:generate go tool github.com/go-app-blazar/blazar/tool/generate-mixin-files --source-directory=./internal/base/example

//go:generate go tool github.com/go-app-blazar/blazar/tool/generate-interface-package --source-directory=./internal/base/example --target-directory=./example --private-prefix=example --public-prefix=Example
```

Create a Blazar plugin package under `plugin/example`.

That package should contain the CSS file in an `embed.FS`.

It should also define a new `blazarplugin.Plugin`.

# Usage

Once done, you can add your plugin to your Blazar app.

```
blazarApp.AddPlugin(example.NewPlugin(example.Config{
    Location: "/web/example/",
}))
```

And you can use your component normally.

```
app.Div().
    Body(
        example.Banner().
            HeadlineText("This is my banner!"),
    )
```