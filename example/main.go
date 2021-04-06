//
// Copyright (c) 2021 Tenebris Technologies Inc. (https://www.tenebris.com)
// Use of this source code is governed by the MIT license.
// Please see the LICENSE for details.
//
// This is an example of how to upload a file as evidence to TBL. It is intended as
// a starting point for automatic evidence collection from a source that is not
// natively supported by TBL.
//
// A TBL Custom Integration must be created and the username, password, and API key
// must be set. A convenience function can be used to import them from the JSON
// file offered by TBL when creating the custom integration.
//
// Note that each TBL Custom Integration can have one or more "Custom Evidence Integration",
// each of which are associated with a single evidence task and have a unique Collector URL.
//
// For further information, please see the TBL Evidence API documentation.
//
package main

import (
	"fmt"
	"log"

	"github.com/tenebris-tech/tbl/evidence"
)

func main() {
	var err error

	// Instantiate new evidence uploader
	e := evidence.New()

	/**************************************************************************
	The following must be set:

		e.User - TLB username
		e.Pass - TBL password
		e.Key  - TBL API key
		e.URL  - TBL endpoint for the evidence task
		e.File - File to upload as evidence
		e.Type - Type of file (text/plain, text/csv, etc.)
	****************************************************************************/

	// Read TBL credentials from TBL http-headers.json file
	err = e.Credentials("http-headers.json")
	checkError(err)

	// Set TBL endpoint for the specific evidence task (*** MUST BE UPDATED ***)
	e.URL = "https://openapi.tugboatlogic.com/api/v0/evidence/collector/xxxx/"

	// Set file (full or relative path) and MIME type
	e.File = "sample.txt"
	e.Type = "text/plain"

	// Upload
	err = e.Upload()
	checkError(err)

	// Test .csv
	e.File = "sample.csv"
	e.Type = "text/csv"
	err = e.Upload()
	checkError(err)

	// Test .jpg
	e.File = "sample.jpg"
	e.Type = "image/jpeg"
	err = e.Upload()
	checkError(err)

	fmt.Println("Evidence uploaded successfully.")
}

// Check for error
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
