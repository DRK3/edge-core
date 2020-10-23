/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package metadata

import (
	"sync"
)

// nolint:gochecknoglobals // package-private globals
var (
	rwmutex     = &sync.RWMutex{}
	levels      = newModuledLevels()
	callerInfos = newCallerInfo()
)

// SetLevel - setting log level for given module.
func SetLevel(module string, level Level) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	levels.SetLevel(module, level)
}

// GetLevel - getting log level for given module.
func GetLevel(module string) Level {
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return levels.GetLevel(module)
}

// GetAllLevels - getting all set log levels.
func GetAllLevels() map[string]Level {
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return levels.GetAllLevels()
}

// IsEnabledFor - Check if given log level is enabled for given module.
func IsEnabledFor(module string, level Level) bool {
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return levels.IsEnabledFor(module, level)
}

// ShowCallerInfo - Show caller info in log lines for given log level and module.
func ShowCallerInfo(module string, level Level) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	callerInfos.ShowCallerInfo(module, level)
}

// HideCallerInfo - Do not show caller info in log lines for given log level and module.
func HideCallerInfo(module string, level Level) {
	rwmutex.Lock()
	defer rwmutex.Unlock()
	callerInfos.HideCallerInfo(module, level)
}

// IsCallerInfoEnabled - returns if caller info enabled for given log level and module.
func IsCallerInfoEnabled(module string, level Level) bool {
	rwmutex.RLock()
	defer rwmutex.RUnlock()

	return callerInfos.IsCallerInfoEnabled(module, level)
}
