/*
   Copyright 2021 MediaExchange.io

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	Hostname string
	Username string
	Password string
)

func main() {
	flag.StringVar(&Hostname, "hostname", "", "NNTP server hostname")
	flag.StringVar(&Username, "username", "", "NNTP account username")
	flag.StringVar(&Password, "password", "", "NNTP account password")

	if len(os.Args) == 1 {
		help()
		return
	}
}

// help displays the program's sub-commands and arguments.
func help() {
	fmt.Println("Usage:")
	fmt.Print('\n')
	fmt.Println("    nntp [arguments]")
	fmt.Print('\n')
	flag.PrintDefaults()
}

func download() {
}