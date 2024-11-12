# zj-alias

Zellij plugin to show my bash aliases in a separate pane.

## Build instructions

Requires tinygo to compile it. For tinygo, I needed an older go version. Downgraded to `go1.22.9` for this project.

```bash
tinygo build -o plugin.wasm -target=wasi -scheduler=none
```

**TODO**: running the plugin directly in the plugin manager in Zellij seems to work, but I can't make it automatically start when Zellij starts yet. Getting some errors.
