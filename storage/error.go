// Copyright (c) 2018 ContentBox Authors.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package storage

import "errors"

//error
var (
	ErrKeyNotExists = errors.New("specified key does not exists")
	ErrKeyNotFound  = errors.New("key not found")
)
