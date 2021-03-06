//
// Copyright (c) 2015 Snowplow Analytics Ltd. All rights reserved.
//
// This program is licensed to you under the Apache License Version 2.0,
// and you may not use this file except in compliance with the Apache License Version 2.0.
// You may obtain a copy of the Apache License Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0.
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the Apache License Version 2.0 is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the Apache License Version 2.0 for the specific language governing permissions and limitations there under.
//
package main

import (
	"flag"
	"fmt"
)

type Options struct {
	help     bool
	version  bool
	playbook string
	sqlroot  string
}

func (o *Options) GetFlagSet() *flag.FlagSet {
	var fs = flag.NewFlagSet("Options", flag.ContinueOnError)

	fs.BoolVar(&(o.help), "help", false, "Shows this message")
	fs.BoolVar(&(o.version), "version", false, "Shows the program version")
	fs.StringVar(&(o.playbook), "playbook", "", "Playbook of SQL scripts to execute")
	fs.StringVar(&(o.sqlroot), "sqlroot", SQLROOT_PLAYBOOK, fmt.Sprintf("Absolute path to SQL scripts. Use %s and %s for those respective paths", SQLROOT_PLAYBOOK, SQLROOT_BINARY))
	// TODO: add format flag if/when we support TOML

	return fs
}
