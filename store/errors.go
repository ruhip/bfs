package main

import (
	"errors"
)

var (
	// block
	ErrSuperBlockMagic   = errors.New("super block magic error")
	ErrSuperBlockVer     = errors.New("super block ver error")
	ErrSuperBlockPadding = errors.New("super block padding error")
	// needle
	ErrNeedleExists      = errors.New("needle already exists")
	ErrNoNeedle          = errors.New("needle not exists")
	ErrNeedleChecksum    = errors.New("needle checksum error")
	ErrNeedleFlag        = errors.New("needle flag error")
	ErrNeedleSize        = errors.New("needle size error")
	ErrNeedleHeaderMagic = errors.New("needle header magic number error")
	ErrNeedleFooterMagic = errors.New("needle footer magic number error")
	ErrNeedleKey         = errors.New("needle key not match")
	ErrNeedlePadding     = errors.New("needle padding error")
	ErrNeedleCookie      = errors.New("needle cookie error")
	ErrNeedleDeleted     = errors.New("needle deleted")
	// ring
	ErrRingEmpty = errors.New("ring buffer empty")
	ErrRingFull  = errors.New("ring buffer full")
	// store
	ErrStoreVolumeIndex = errors.New("store volume index error")
	// volume
	ErrVolumeNotExist = errors.New("volume not exist")
)
