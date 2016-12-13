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
	"fmt"
	"log"

	"google.golang.org/api/drive/v3"
)

func _printFiles(fileList *drive.FileList) error {
	for _, f := range fileList.Files {

		fmt.Printf("%s\n", f.Name)

	}

	return nil
}

func printFiles(srv *drive.Service, ctx context.Context, user string) {

	query := fmt.Sprintf("\"%s\" in owners", user)

	err := srv.Files.List().Q(query).Pages(ctx, _printFiles)
	if err != nil {
		log.Fatalf("Error while querying for files: %v", err)
	}

}
