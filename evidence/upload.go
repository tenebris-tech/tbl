//
// Copyright (c) 2021 Tenebris Technologies Inc. (https://www.tenebris.com)
// Use of this source code is governed by the MIT license.
// Please see the LICENSE for details.
//
package evidence

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"time"
)

func (e *Evidence) Upload() error {

	// Ensure username, password, and API key are set
	if e.User == "" {
		return errors.New("username must be set")
	}

	if e.Pass == "" {
		return errors.New("password must be set")
	}

	if e.Key == "" {
		return errors.New("API key must be set")
	}

	if e.URL == "" {
		return errors.New("TBL evidence task endpoint (URL) must be set")
	}

	if e.File == "" {
		return errors.New("evidence filename must be set")
	}

	if e.Type == "" {
		return errors.New("file type must be set")
	}

	// Get current date in YYYY-MM-DD
	date := time.Now().Format("2006-01-02")

	// Open file for read
	file, err := os.Open(e.File)
	if err != nil {
		return err
	}

	//goland:noinspection GoUnhandledErrorResult
	defer file.Close()

	// Set up buffer for body and instantiate multipart writer
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create new form-data header with specified Content-Type
	// This is used instead of the CreateFormFile wrapper because it
	// always sets application/octet-stream, which is not what we want
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="%s"; filename="%s"`,
			escapeQuotes("file"), escapeQuotes(filepath.Base(file.Name()))))
	h.Set("Content-Type", e.Type)

	// Create the MIME part
	part, err := writer.CreatePart(h)
	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return err
	}

	// Add collected field
	err = writer.WriteField("collected", date)
	if err != nil {
		return err
	}

	err = writer.Close()
	if err != nil {
		return err
	}

	// Set up HTTP request
	req, _ := http.NewRequest("POST", e.URL, body)
	req.SetBasicAuth(e.User, e.Pass)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("X-API-KEY", e.Key)

	// Create HTTP client and perform request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// TBL will return HTTP 201 Created for success
	if resp.StatusCode != 201 {
		txt := fmt.Sprintf("HTTP client returned %s", resp.Status)
		return errors.New(txt)
	}

	return nil
}
