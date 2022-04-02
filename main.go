package main

import (
	cmd "github.com/caddyserver/caddy/v2/cmd"
	_ "github.com/caddyserver/caddy/v2/modules/standard"

	_ "github.com/rabbice/toml-adapter"
	_ "github.com/rabbice/auth-prefix"
)

func main() {
	cmd.Main()
}
