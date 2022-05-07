package htmlr

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

var includeExpressionReg = regexp.MustCompile(`(?mi)\{% *include *["'](.*)["'] *%\}`)

func Resolve(input, output string) {
	er := ioutil.WriteFile(output, resolve(input), 0765)
	if er != nil {
		panic(fmt.Sprintf("couldn't write resolved template to '%s'", output))
	}
}

func resolve(input string) []byte {
	return resolveIncludes(loadFile(input), filepath.Dir(input))
}

func resolveIncludes(src []byte, currentDir string) []byte {
	return includeExpressionReg.ReplaceAllFunc(src, func(b []byte) []byte {
		fpath := resolveFilePath(currentDir, extractPath(b))
		fmt.Println("repl func", "'"+extractPath(b)+"'")
		return resolveIncludes(loadFile(fpath), filepath.Dir(fpath))
	})
}

func resolveFilePath(dir, fpath string) string {
	np := fpath
	var err error
	if _, err = os.Stat(np); errors.Is(err, os.ErrNotExist) {
		np = filepath.Join(dir, np)
		if _, err = os.Stat(np); errors.Is(err, os.ErrNotExist) {
			np, err = filepath.Abs(fpath)
			if err != nil {
				np = filepath.Join(dir, fpath)
				np, err = filepath.Abs(np)
				if err != nil {
					panic(fmt.Sprintf("couldn't find template 1 '%s'", fpath))
				}
			}
			if _, err = os.Stat(np); errors.Is(err, os.ErrNotExist) {
				panic(fmt.Sprintf("couldn't find template 2 '%s' '%s'", fpath, np))
			}
		}
	}
	return np
}

func loadFile(fpath string) []byte {
	if fpath == "" {
		panic(fmt.Sprintf("template file path cannot be empty '%s'", fpath))
	}
	src, er := fileToBytes(fpath)
	if er != nil {
		panic(fmt.Sprintf("couldn't find template '%s'", fpath))
	}
	return src
}

func extractPath(inclExpr []byte) string {
	for _, match := range includeExpressionReg.FindAllSubmatch(inclExpr, -1) {
		if len(match) >= 2 {
			return string(match[1])
		}
	}
	return ""
}
