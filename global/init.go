package global

import (
	"os"
	"path/filepath"
	"sync"
)

const (
	SecKey               string = "pb5gm867BDsS2WUIjALaYRrNIYzduUo1vm247qNBnUX9i5QMa3pxHj7lw96jMS2Hg8Xe2KXRkj7n90c0JZMuIymODDSmtlnpdZzlqiB36COoc5UNbOfW2WwJHgKWP1WM"
	RChatAdminChanPrefix string = "admin-*"
	RChatUserChanPrefix  string = "user-*"

	RChatAdminSet string = "admins"
	RChatUserSet  string = "users"

	RoleAdmin string = "admin"
	RoleUser  string = "user"
)

func init() {
	Init()
}

var RootDir string

var once = new(sync.Once)

func Init() {
	once.Do(func() {
		inferRootDir()
		initConfig()
	})
}

func inferRootDir() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	var infer func(d string) string
	infer = func(d string) string {
		if exists(d + "/template") {
			return d
		}

		return infer(filepath.Dir(d))
	}
	RootDir = infer(cwd)
}

func exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
