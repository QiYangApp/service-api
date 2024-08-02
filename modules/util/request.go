// Copyright 2020 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package util

import (
	"net/http"
	"strings"
)

// IsAPIPath returns true if the specified URL is an API path
func IsAPIPath(req *http.Request) bool {
	return strings.HasPrefix(req.URL.Path, "/api/")
}

// IsAttachmentDownload check if request is a file download (GET) with URL to an attachment
func IsAttachmentDownload(req *http.Request) bool {
	return strings.HasPrefix(req.URL.Path, "/attachments/") && req.Method == "GET"
}

// IsContainerPath checks if the request targets the container endpoint
func IsContainerPath(req *http.Request) bool {
	return strings.HasPrefix(req.URL.Path, "/v2/")
}
