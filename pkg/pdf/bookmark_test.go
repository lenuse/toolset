package pdf

import (
	"fmt"
	"github.com/lenuse/toolset/cmd"
	"testing"
)

func TestParasBookmark(t *testing.T) {
	conf := cmd.PathConf{
		BookmarksPath: "../../test/pdfbookmark.txt",
	}
	bms, err := parasBookmark(conf)
	fmt.Println(bms)
	if err != nil {
		t.Error(err)
	}
}

func TestAddBookmarksFile(t *testing.T) {
	conf := cmd.PathConf{
		BookmarksPath: "../../test/pdfbookmark.txt",
		InfilePath:    "/Users/XXX/Documents/pdf/new.pdf",
	}
	err := addBookmarks(conf)
	if err != nil {
		t.Error(err)
	}
}
