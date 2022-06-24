// Copyright 2022 InfraCloud Technologies
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

// Greeter is a webserver which reads value from GREETER_NAME
// environment variable and replies with a greeting message to the
// HTTP requests.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	nameEnv = "GREETER_NAME"
)

func main() {
	var (
		listenAddress = flag.String("web.listen-address", ":8080",
			"Address on which to expose the web interface.")
		cliOnly = flag.Bool("cli-only", false,
			"Print the greeting on CLI and exit.")
	)
	flag.Parse()

	greeting := "Hello from %s!"
	greeter := "Anonymous"
	if n, ok := os.LookupEnv(nameEnv); ok {
		greeter = n
	} else {
		log.Printf("environment variable %s is empty, falling back to %q", nameEnv, greeter)
	}
	message := fmt.Sprintf(greeting, greeter)

	if *cliOnly {
		fmt.Println(message)
		os.Exit(0)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("sending greetings to %s", r.URL)
		fmt.Fprintln(w, message)
	})
	log.Printf("listening on %s", *listenAddress)
	log.Fatal(http.ListenAndServe(*listenAddress, nil))
}
