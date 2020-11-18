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
const byteCodeSuffix = ".pipe"

var scriptStoragePath = getStoragePath()
var funcDefMapping = make(map[string]*CompactFunction)

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
		if strings.HasSuffix(fileInfo.Name(), byteCodeSuffix) {
			byteCode, err := ioutil.ReadFile(filepath.Join(scriptStoragePath, fileInfo.Name()))
			if err != nil {
				return err
			}
			funcDef := deserializeFunction(byteCode)
			// register function to registry
			registry.RegisterFunction(DefLibFunc(funcDef.Name, funcDef.Callable.Exec, funcDef.Params))
			// add function to mapping for convenience to query function md5
			funcDefMapping[funcDef.Name] = funcDef
		}
	}

	return nil
}

func serializeFunction(funcDef *CompactFunction) {

}

func deserializeFunction(bytecode []byte) *CompactFunction {
	return nil
}

func Install(filename string) error {
	pipeScript := parser.ParseScript(filename)
	for idx := range pipeScript.Funcs {
		funcDef := pipeScript.Funcs[idx]
		old, ok := getOldVersionFuncDef(funcDef)
		if ok {
			if funcDef.Md5 == old.Md5 {
				fmt.Println("no update on function", funcDef.Name)
			}
		}
		// TODO:
		// 如果func已加载，则比对md5，若md5不相等则序列化funcDef写入storagePath
		// 若没加载直接序列化写入
	}

	return nil
}

func getOldVersionFuncDef(function CompactFunction) (*CompactFunction, bool) {
	_, err := registry.GetFunction(function.Name)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return nil, true
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
