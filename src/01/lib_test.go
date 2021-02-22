
package main

import "testing"

// TestOpen tests the functionality of the open library function.
func TestOpen(t *testing.T) {
	library := initLibrary()

	if library.open() != nil {
		t.Errorf("open failed.")
	}
}

// TestClose tests the functionality of the close library function.
func TestClose(t *testing.T) {
	library := initLibrary()

	if library.close() != nil {
		t.Errorf("close failed.")
	}
}

// TestMkDir tests the functionality of the mkDir library function.
func TestMkDir(t *testing.T) {
	library := initLibrary()

	if library.mkDir() != nil {
		t.Errorf("mkDir failed")
	}
}

// TestRemoveFile tests the functionality of the removeDir library function. 
func TestRemoveFile(t *testing.T) {
	library := initLibrary()

	if library.removeFile() != nil {
		t.Errorf("removeFile failed.")
	}
}

// TestRemoveDir tests the functionality of the removeDir library function. 
func TestRemoveDir(t *testing.T) {
	library := initLibrary()

	if library.removeDir() != nil {
		t.Errorf("removeDir failed.")
	}
}

// listDir tests the functionality of the removeDir library function. 
func listDir(t *testing.T) {
	library := initLibrary()

	if library.listDir() != nil {
		t.Errorf("removeDir failed.")
	}
}