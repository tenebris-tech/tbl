// Copyright (c) 2021 Tenebris Technologies Inc. (https://www.tenebris.com)
// Use of this source code is governed by the MIT license.
// Please see the LICENSE for details.

package tbl

import "strings"

// From mime/multipart/writer.go
var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}
