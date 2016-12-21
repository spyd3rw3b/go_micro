package main

import (
  "fmt"
  "os"

  "github.com/urfave/cli"
)

var (
  Revision = "1900a"
)

func main() {
  cli.VersionPrinter = func(c *cli.Context) {
    fmt.Printf("version=%s revision=%s\n", c.App.Version, Revision)
  }

  app := cli.NewApp()
  app.Name = "cli_wrap"
  app.Version = "1.1.0"
  app.Run(os.Args)
}