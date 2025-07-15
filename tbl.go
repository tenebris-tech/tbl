// Copyright (c) 2021-2025 Tenebris Technologies Inc. (https://www.tenebris.com)
// Use of this source code is governed by the MIT license.
// Please see the LICENSE for details.

package tbl

import (
	"encoding/json"
	"fmt"
	"os"
)

// TBL is the main TBL struct
type TBL struct {
	User  string // TBL username
	Pass  string // TBL password
	Key   string // TBL API KEY
	Debug bool   // Debug mode
}

// Credentials is the struct for the OneTrust/TubBoat Logic credentials file
type Credentials struct {
	User string `json:"Username"`
	Pass string `json:"Password"`
	Key  string `json:"X-API-KEY"`
}

type Option func(*TBL) error

func New(options ...Option) (TBL, error) {
	tbl := TBL{}
	for _, opt := range options {
		if err := opt(&tbl); err != nil {
			return TBL{}, err
		}
	}
	return tbl, nil
}

func WithCredentialFile(fileName string) Option {
	return func(t *TBL) error {
		var creds Credentials
		data, err := os.ReadFile(fileName)
		if err != nil {
			return fmt.Errorf("failed to read credentials file: %w", err)
		}
		if err := json.Unmarshal(data, &creds); err != nil {
			return fmt.Errorf("failed to parse credentials file: %w", err)
		}
		t.User = creds.User
		t.Pass = creds.Pass
		t.Key = creds.Key

		if t.User == "" {
			return fmt.Errorf("username is empty")
		}
		if t.Pass == "" {
			return fmt.Errorf("password is empty")
		}
		if t.Key == "" {
			return fmt.Errorf("API key is empty")
		}

		return nil
	}
}

func WithDebug(debug bool) Option {
	return func(t *TBL) error {
		t.Debug = debug
		return nil
	}
}
