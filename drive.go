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
