package fls

import (
    "testing"
)

func TestCreateDirIfNotExist(t *testing.T) {

	CreateDirForFile("/tmp/","test/test2/test.html")
	CreateDirForDir("/tmp/","test3/test4")	
//	CreateDirIfNotExist("/tmp/","test3/test4")	
		
}

