// Copyright (c) 2021 Tenebris Technologies Inc. (https://www.tenebris.com)
// Use of this source code is governed by the MIT license.
// Please see the LICENSE for details.

package tbl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Upload uploads a file to one or more TBL evidence task endpoints.
// The url parameter can contain one or more comma-delimited URLs.
func (e *TBL) Upload(fileName string, fileType string, url string) error {

	if url == "" {
		return errors.New("TBL evidence task endpoint (URL) must be specified")
	}

	// Split the URL string into individual URLs
	urls := splitURLs(url)

	if e.Debug {
		log.Printf("[DEBUG] Upload request for file: %s, type: %s", fileName, fileType)
		log.Printf("[DEBUG] Found %d URLs to upload to: %v", len(urls), urls)
	}

	if fileName == "" {
		return errors.New("evidence filename must be specified")
	}

	if fileType == "" {
		return errors.New("file type must be specified")
	}

	// Get current date in YYYY-MM-DD
	date := time.Now().Format("2006-01-02")

	// Open file for read
	file, err := os.Open(fileName)
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
	h.Set("Content-Type", fileType)

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

	// If there's only one URL, perform a simple upload
	if len(urls) == 1 {
		if e.Debug {
			log.Printf("[DEBUG] Uploading to single URL: %s", urls[0])
		}
		return uploadToSingleURL(e, urls[0], body, writer.FormDataContentType())
	}

	// For multiple URLs, attempt to upload to all of them
	var lastErr error
	successCount := 0

	if e.Debug {
		log.Printf("[DEBUG] Attempting to upload to %d URLs", len(urls))
	}

	for _, singleURL := range urls {
		// Create a new buffer with the same content for each request
		// This is necessary because the body buffer can only be read once
		bodyClone := bytes.NewBuffer(body.Bytes())

		if e.Debug {
			log.Printf("[DEBUG] Uploading to URL: %s", singleURL)
		}

		err := uploadToSingleURL(e, singleURL, bodyClone, writer.FormDataContentType())
		if err != nil {
			lastErr = err
			if e.Debug {
				log.Printf("[DEBUG] Upload to %s failed: %v", singleURL, err)
			}
		} else {
			successCount++
			if e.Debug {
				log.Printf("[DEBUG] Upload to %s succeeded", singleURL)
			}
		}
	}

	// If no uploads succeeded, return the last error
	if successCount == 0 && lastErr != nil {
		if e.Debug {
			log.Printf("[DEBUG] All uploads failed, last error: %v", lastErr)
		}
		return lastErr
	}

	// If some but not all uploads succeeded, return a partial success error
	if successCount < len(urls) {
		partialMsg := fmt.Sprintf("uploaded to %d of %d URLs, last error: %v", successCount, len(urls), lastErr)
		if e.Debug {
			log.Printf("[DEBUG] Partial success: %s", partialMsg)
		}
		return fmt.Errorf(partialMsg)
	}

	if e.Debug {
		log.Printf("[DEBUG] All %d uploads completed successfully", len(urls))
	}
	return nil
}

// splitURLs splits a comma-delimited string of URLs into a slice of individual URLs
func splitURLs(urlString string) []string {
	// Split the URL string by commas
	urls := strings.Split(urlString, ",")

	// Trim whitespace from each URL
	for i, url := range urls {
		urls[i] = strings.TrimSpace(url)
	}

	// Filter out empty URLs
	var validURLs []string
	for _, url := range urls {
		if url != "" {
			validURLs = append(validURLs, url)
		}
	}

	return validURLs
}

// uploadToSingleURL uploads the file to a single URL
func uploadToSingleURL(e *TBL, url string, body io.Reader, contentType string) error {
	if e.Debug {
		log.Printf("[DEBUG] Setting up HTTP request to %s", url)
	}
	// Set up HTTP request
	req, _ := http.NewRequest("POST", url, body)
	req.SetBasicAuth(e.User, e.Pass)
	req.Header.Add("Content-Type", contentType)
	req.Header.Add("X-API-KEY", e.Key)

	// Create HTTP client and perform request
	client := &http.Client{}

	if e.Debug {
		log.Printf("[DEBUG] Sending HTTP request to %s", url)
	}

	resp, err := client.Do(req)
	if err != nil {
		if e.Debug {
			log.Printf("[DEBUG] HTTP request failed: %v", err)
		}
		return err
	}

	// TBL will return HTTP 201 Created for success
	if resp.StatusCode != 201 {
		errMsg := fmt.Sprintf("HTTP client returned %s", resp.Status)
		if e.Debug {
			log.Printf("[DEBUG] Upload failed: %s", errMsg)
		}
		return fmt.Errorf(errMsg)
	}

	if e.Debug {
		log.Printf("[DEBUG] Upload to %s successful (HTTP 201)", url)
	}

	return nil
}
