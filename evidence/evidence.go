//
// Copyright (c) 2021 Tenebris Technologies Inc. (https://www.tenebris.com)
// Use of this source code is governed by the MIT license.
// Please see the LICENSE for details.
//
package evidence

type Evidence struct {
	User string // TBL username
	Pass string // TBL password
	Key  string // TBL API KEY
	URL  string // URL for TBL evidence task
	File string // File to upload
	Type string // File type (i.e. text/plain, text/csv)
}

func New() Evidence {
	return Evidence{}
}
