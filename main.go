// Copyright 2017 XUEQIU.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"

	"github.com/urfave/cli"

	"fmt"

	"github.com/xueqiu/rdr/dump"
)

//go:generate go-bindata -prefix "static/" -o=static/static.go -pkg=static -ignore static.go static/...
//go:generate go-bindata -prefix "views/" -o=views/views.go -pkg=views -ignore views.go views/...

func main() {
	app := cli.NewApp()
	app.Name = "rdr"
	app.Usage = "a tool to parse redis rdbfile"
	app.Version = "v0.0.1"
	app.Writer = os.Stdout
	app.ErrWriter = os.Stderr
	app.Commands = []cli.Command{
		cli.Command{
			Name:      "dump",
			Usage:     "dump statistical information of rdbfile to STDOUT",
			ArgsUsage: "FILE1 [FILE2] [FILE3]...",
			Action:    dump.ToCliWriter,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-expire, x",
					Usage: "remove expiry from all keys (default: false)",
				},
			},
		},
		cli.Command{
			Name:      "show",
			Usage:     "show statistical information of rdbfile by webpage",
			ArgsUsage: "DIR1 [DIR2] [DIR3] or FILE1 [FILE2] [FILE3]...",
			Flags: []cli.Flag{
				cli.UintFlag{
					Name:  "port, p",
					Value: 8080,
					Usage: "Port for rdr to listen",
				},
				cli.BoolFlag{
					Name:  "no-expire, x",
					Usage: "remove expiry from all keys (default: false)",
				},
			},
			Action: dump.Show,
		},
		cli.Command{
			Name:      "keys",
			Usage:     "get all keys from rdbfile",
			ArgsUsage: "FILE1 [FILE2] [FILE3]...",
			Action:    dump.Keys,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-expire, x",
					Usage: "remove expiry from all keys (default: false)",
				},
			},
		},
		cli.Command{
			Name:      "hash-fields",
			Usage:     "get hash keys all field from rdbfile",
			ArgsUsage: "hash-key regexp pattern  FILE1 [FILE2] [FILE3]...",
			Action:    dump.HashFields,
		},
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		_, _ = fmt.Fprintf(c.App.ErrWriter, "command %q can not be found.\n", command)
		_ = cli.ShowAppHelp(c)
	}
	_ = app.Run(os.Args)
}
