/*
   Copyright 2016 Hubert Jarosz

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
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
)

func main() {

	user := flag.String("list-files", "", "`e-mail` of user whose files you want to list")
	jwtFile := flag.String("jwt", "./jwt.json", "`path` to JWT file with secret key")
	flag.Parse()

	if *user == "" {
		flag.Usage()
		os.Exit(2)
	}

	ctx := context.Background()

	client := getClient(ctx, *jwtFile, *user)

	srv, err := drive.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve drive Client %v", err)
	}

	fmt.Printf("Files that belong to %s:\n", *user)
	printFiles(srv, ctx, *user)
}
