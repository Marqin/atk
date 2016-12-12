/*
   Copyright (C) 2016  Hubert Jarosz

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
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
