package fls

import (
//	"fmt"
	"os"
	"path"
)

func CreateDirForFile(rootdir string, file string) {

	dir := path.Dir(file)

	dirtocreate := rootdir + dir
	if _, err := os.Stat(dirtocreate); os.IsNotExist(err) {

		os.MkdirAll(dirtocreate, 0777)
	
	}
	
	
}
func CreateDirForDir(rootdir string, dir string) {
	
		dirtocreate := rootdir+dir		
		if _, err := os.Stat(dirtocreate ); os.IsNotExist(err) {

		os.MkdirAll(dirtocreate, 0777)
	
	}	
	
}
