//go:build darwin || freebsd
// +build darwin freebsd

/* Top Disk Usage.
 * Copyright (C) 2019 Joseph Paul <joseph.paul1@gmx.com>
 * https://github.com/CorrectRoadH/tdu
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 */

package tdu

import (
	"syscall"
)

func tcgets() uintptr {
	return uintptr(syscall.TIOCGETA)
}
