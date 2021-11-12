//go:build generate
// +build generate

package main

import (
	"io/ioutil"

	"github.com/kpetku/syndie-gui/plugin"
	. "i2pgit.org/idk/clientapp"
)

func main() {
	ioutil.WriteFile(SyndieClientApp+".java", []byte(GenerateJava(&plugin.SyndieClientApp{})), 0644)
}
