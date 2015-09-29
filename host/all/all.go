// Package all conviniently loads all the inbuilt/supported host drivers.
package all

import (
	_ "github.com/gavincabbage/embd/host/bbb"
	_ "github.com/gavincabbage/embd/host/rpi"
)
