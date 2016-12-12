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
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/api/drive/v3"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func getClient(ctx context.Context, path string, user string) *http.Client {

	jwtFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Cannot fint JWT file: %v", err)
	}

	config, err := google.JWTConfigFromJSON(jwtFile, drive.DriveScope)
	if err != nil {
		log.Fatalf("Cannot parse JWT file: %v", err)
	}

	config.Subject = user

	ts := config.TokenSource(ctx)

	return oauth2.NewClient(ctx, ts)
}
