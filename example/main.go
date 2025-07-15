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
// Note that each TBL Custom Integration can have one or more "Custom TBL Integration",
// each of which are associated with a single evidence task and have a unique Collector URL.
//
// For further information, please see the TBL TBL API documentation.
package main

import (
	"fmt"
	"log"

	"github.com/tenebris-tech/tbl"
)

func main() {
	var err error

	// Instantiate new evidence uploader and read the credentials from the JSON file
	t, err := tbl.New(
		tbl.WithCredentialFile("http-headers-example.json"), // Point to file created by OneTrust/TugBoat Logic
		tbl.WithDebug(true), // Enable debug output
	)
	checkError(err)

	/**************************************************************************
	The following must be set:

		t.URL  - TBL endpoint for the evidence task
		t.File - File to upload as evidence
		t.Type - Type of file (text/plain, text/csv, etc.)
	****************************************************************************/

	// Set TBL endpoint for the specific evidence task (*** MUST BE UPDATED ***)\
	// This can be a comma-delimited list of URLs if you want to upload to multiple tasks
	url := "https://openapi.tugboatlogic.com/api/v0/evidence/collector/xxxx/"

	// Upload a plain text file using Evidence struct with empty date
	err = t.Upload(tbl.Evidence{
		FileName:  "sample.txt",
		FileType:  "text/plain",
		URL:       url,
		Collected: "", // Empty date will use current date
	})
	checkError(err)

	// Upload a CSV file using Evidence struct with empty date
	err = t.Upload(tbl.Evidence{
		FileName:  "sample.csv",
		FileType:  "text/csv",
		URL:       url,
		Collected: "", // Empty date will use current date
	})
	checkError(err)

	// Upload an image file using Evidence struct with empty date
	err = t.Upload(tbl.Evidence{
		FileName:  "sample.jpg",
		FileType:  "image/jpeg",
		URL:       url,
		Collected: "", // Empty date will use current date
	})
	checkError(err)

	fmt.Println("TBL uploaded successfully.")
}

// Check for error
func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
