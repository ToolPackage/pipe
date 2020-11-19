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

const libraryDirName = ".pipe-lib"
const libScriptSuffix = ".pipe"

var centralLibraryPath = getCentralLibraryPath()
var funcDefMapping = make(map[string]*CompactFunction)

func LoadLibraries() {
	cwd, err := os.Getwd()
	exitWhenError(err)
	err = loadLibrary(filepath.Join(cwd, libraryDirName))
	exitWhenError(err)
	err = loadLibrary(centralLibraryPath)
	exitWhenError(err)
}

func exitWhenError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func loadLibrary(libPath string) error {
	files, err := ioutil.ReadDir(libPath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	for _, fileInfo := range files {
		if strings.HasSuffix(fileInfo.Name(), libScriptSuffix) {
			bytes, err := ioutil.ReadFile(filepath.Join(libPath, fileInfo.Name()))
			if err != nil {
				return err
			}
			funcDef := Deserialize(bytes)
			// register function to registry
			registry.RegisterFunction(DefLibFunc(funcDef.Name, funcDef.Exec, funcDef.Params))
			// add function to mapping for convenience to query function md5
			funcDefMapping[funcDef.Name] = funcDef
		}
	}

	return nil
}

func Install(filename string, global bool) error {
	libPath, err := getLibraryPath(global)
	if err != nil {
		return err
	}

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
		bytes := Serialize(&funcDef)
		if err := ioutil.WriteFile(filepath.Join(libPath, funcDef.Name+libScriptSuffix),
			bytes, 0666); err != nil {
			return err
		}
	}
	return nil
}

func getLibraryPath(central bool) (libPath string, err error) {
	if central {
		libPath = centralLibraryPath
	} else {
		libPath, err = os.Getwd()
		if err != nil {
			return
		}
		libPath = filepath.Join(libPath, libraryDirName)
	}
	err = os.MkdirAll(libPath, 0666)
	return
}

func getCentralLibraryPath() string {
	return filepath.Join(getUserHomeDir(), libraryDirName)
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
