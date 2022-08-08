/*
The MIT License (MIT)

# Copyright (c) 2016 Achille

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit
persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
package dl

// #cgo CFLAGS: -W -Wall -Wno-unused-parameter -O3
// #cgo LDFLAGS: -ldl
import "C"
import (
	"os"
	"path/filepath"
	"strings"
)

func find(name string) (path string, err error) {
	if strings.ContainsRune(name, '/') {
		path = name
		return
	}

	if len(filepath.Ext(name)) == 0 {
		name += ".so"
	}

	dirs := make([]string, 0, 100)
	dirs = append(dirs, getPaths("LD_LIBRARY_PATH")...)
	dirs = append(dirs, "/lib", "/usr/lib")

	ok := false

	for _, dir := range dirs {
		filepath.Walk(dir, func(p string, f os.FileInfo, e error) error {
			if !ok && f != nil && !f.IsDir() && (f.Mode()&0600) != 0 && strings.HasPrefix(f.Name(), name) {
				path, ok = p, true
			}
			return nil
		})

		if ok {
			break
		}
	}

	if !ok {
		err = os.ErrNotExist
	}

	return
}

func getPaths(env string) []string {
	paths := strings.Split(os.Getenv(env), ":")
	i := 0

	for _, p := range paths {
		if len(p) != 0 {
			paths[i] = p
			i++
		}
	}

	return paths[:i]
}
