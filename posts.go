package main

import (
	"fmt"
	"github.com/writeas/writeas-cli/utils"
	"os"
	"path/filepath"
)

const (
	POSTS_FILE = "posts.psv"
	SEPARATOR  = `|`
)

func userDataDir() string {
	return filepath.Join(parentDataDir(), DATA_DIR_NAME)
}

func dataDirExists() bool {
	return fileutils.Exists(userDataDir())
}

func createDataDir() {
	err := os.Mkdir(userDataDir(), 0700)
	if err != nil && DEBUG {
		panic(err)
	}
}

func addPost(id, token string) {
	f, err := os.OpenFile(filepath.Join(userDataDir(), POSTS_FILE), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		if DEBUG {
			panic(err)
		} else {
			return
		}
	}
	defer f.Close()

	l := fmt.Sprintf("%s%s%s\n", id, SEPARATOR, token)

	if _, err = f.WriteString(l); err != nil && DEBUG {
		panic(err)
	}
}