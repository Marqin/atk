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
