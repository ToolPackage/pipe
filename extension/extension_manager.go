package extension

import (
	"fmt"
	"github.com/ToolPackage/pipe/parser"
	. "github.com/ToolPackage/pipe/parser/definition"
	"github.com/ToolPackage/pipe/registry"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

const storageDirName = ".pipe"
const funcBinarySuffix = ".pipe"

var scriptStoragePath = getStoragePath()
var funcDefMapping = make(map[string]*CompactFunction)

// pipe node type
const (
	variableNode = iota
	streamNode
	pipeNodeArray
	functionNode
)

func init() {
	if err := loadLibraries(); err != nil {
		panic(err)
	}
}

func loadLibraries() error {
	files, err := ioutil.ReadDir(scriptStoragePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, fileInfo := range files {
		if strings.HasSuffix(fileInfo.Name(), funcBinarySuffix) {
			bytes, err := ioutil.ReadFile(filepath.Join(scriptStoragePath, fileInfo.Name()))
			if err != nil {
				return err
			}
			funcDef := deserialize(bytes)
			// register function to registry
			registry.RegisterFunction(DefLibFunc(funcDef.Name, funcDef.Callable.Exec, funcDef.Params))
			// add function to mapping for convenience to query function md5
			funcDefMapping[funcDef.Name] = funcDef
		}
	}

	return nil
}

func Install(filename string) error {
	pipeScript := parser.ParseScript(filename)
	// TODO: loading animation
	for idx := range pipeScript.Funcs {
		funcDef := pipeScript.Funcs[idx]
		old, ok := funcDefMapping[funcDef.Name]
		// save function when:
		// - if old version definition is found and their md5 are not equal
		// - no old version definition is found
		if ok && funcDef.Md5 == old.Md5 {
			fmt.Println("no update on function", funcDef.Name)
			continue
		}
		// convert to byte code and save to file
		bytes := serialize(&funcDef)
		if err := ioutil.WriteFile(filepath.Join(storageDirName, funcDef.Name+funcBinarySuffix),
			bytes, 0666); err != nil {
			return err
		}
	}
	return nil
}

func getStoragePath() string {
	return filepath.Join(getUserHomeDir(), storageDirName)
}

func getUserHomeDir() string {
	home := os.Getenv("HOME")
	if home != "" {
		return home
	}

	if runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
		if home != "" {
			return home
		}

		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home != "" {
			return home
		}
	}

	panic("could not detect home directory")
}
