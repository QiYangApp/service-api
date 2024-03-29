// Copyright 2022 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package timeutil

import (
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	executablModTime     = time.Now()
	executablModTimeOnce sync.Once
)

// GetExecutableModTime get executable file modified time of current process.
func GetExecutableModTime() time.Time {
	executablModTimeOnce.Do(func() {
		exePath, err := os.Executable()
		if err != nil {
			log.Fatalf("os.Executable: %v", err)
			return
		}

		exePath, err = filepath.Abs(exePath)
		if err != nil {
			log.Fatalf("filepath.Abs: %v", err)
			return
		}

		exePath, err = filepath.EvalSymlinks(exePath)
		if err != nil {
			log.Fatalf("filepath.EvalSymlinks: %v", err)
			return
		}

		st, err := os.Stat(exePath)
		if err != nil {
			log.Fatalf("os.Stat: %v", err)
			return
		}

		executablModTime = st.ModTime()
	})
	return executablModTime
}
