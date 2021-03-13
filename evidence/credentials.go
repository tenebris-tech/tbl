//
// Copyright (c) 2021 Tenebris Technologies Inc. (https://www.tenebris.com)
// Use of this source code is governed by the MIT license.
// Please see the LICENSE for details.
//
package evidence

import (
	"encoding/json"
	"io/ioutil"
)

type TBLCredentials struct {
	User string `json:"Username"`
	Pass string `json:"Password"`
	Key  string `json:"X-API-KEY"`
}

func (e *Evidence) Credentials(fileName string) error {
	var creds TBLCredentials

	// Load from fileName
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &creds)
	if err != nil {
		return err
	}

	e.User = creds.User
	e.Pass = creds.Pass
	e.Key = creds.Key
	return nil
}
