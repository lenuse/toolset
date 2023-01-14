package pdf

import (
	"bufio"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
	"github.com/pkg/errors"
	"os"
	"path"
	"strconv"
	"strings"
)

type PathConf struct {
	InfilePath    string
	OutfilePath   string
	BookmarksPath string
}

func AddBookmarks(conf PathConf) error {
	bms, err := parasBookmark(conf)
	if err != nil {
		return err
	}
	fileNameWithSuffix := path.Base(conf.InfilePath)
	fileType := path.Ext(fileNameWithSuffix)
	if fileType != ".pdf" {
		return errors.New("file type err")
	}
	fileNameOnly := strings.TrimSuffix(fileNameWithSuffix, fileType)
	tmpFile := "./" + fileNameOnly + ".bak.pdf"
	if conf.OutfilePath != "" && conf.InfilePath != conf.OutfilePath {
		tmpFile = conf.OutfilePath
	}
	pdfCxt, err := pdfcpu.ReadFile(conf.InfilePath, nil)
	if err != nil {
		return err
	}
	rootDict, err := pdfCxt.XRefTable.Catalog()
	if err != nil {
		return err
	}
	_, ok := rootDict.Find("Outlines")
	if !ok {
		rootDict.Delete("Outlines")
	}
	err = pdfCxt.AddBookmarks(bms)
	if err != nil {
		return err
	}
	return api.CreatePDFFile(pdfCxt.XRefTable, tmpFile, nil)
}

func parasBookmark(conf PathConf) ([]pdfcpu.Bookmark, error) {
	bmsf, err := os.Open(conf.BookmarksPath)
	if err != nil {
		return nil, err
	}
	defer bmsf.Close()
	var bmsList []pdfcpu.Bookmark
	buf := bufio.NewScanner(bmsf)
	offset := 0
	for buf.Scan() {
		line := buf.Text()
		if line == "" {
			continue
		}
		line = strings.TrimSpace(line)
		// 读取偏移量
		if strings.Index(line, "offset:") == 0 {
			v := strings.Split(line, ":")
			offset, _ = strconv.Atoi(v[1])
			continue
		}
		// 读取配置
		line = strings.TrimRight(line, " ")
		info := strings.Split(line, " ")
		level := info[0]
		levelNum := getLevel(level)
		pageNum, _ := strconv.Atoi(info[2])
		pageNum = pageNum + offset
		title := strings.Join(info[0:2], " ")
		bold := false
		if levelNum == 0 {
			bold = true
		}

		currBm := pdfcpu.Bookmark{
			Title:    title,
			PageFrom: pageNum,
			PageThru: pageNum,
			Bold:     bold,
			Italic:   false,
			Color:    &pdfcpu.Black,
			Children: nil,
			Parent:   nil,
		}
		if levelNum == 0 {
			bmsList = append(bmsList, currBm)
		} else {
			tmpBm := &bmsList[len(bmsList)-1]
			for i := 0; i < levelNum-1; i++ {
				tmpBm, err = getChildrenLastBm(tmpBm)
				if err != nil {
					return nil, err
				}
			}
			currBm.Parent = tmpBm
			tmpBm.Children = append(tmpBm.Children, currBm)
		}

	}
	return bmsList, nil
}

func getLevel(level string) int {
	return strings.Count(level, ".")
}

func getChildrenLastBm(bm *pdfcpu.Bookmark) (*pdfcpu.Bookmark, error) {
	if len(bm.Children)-1 < 0 {
		return nil, errors.New("目录层级错误")
	}
	tmp := bm.Children[len(bm.Children)-1]
	return &tmp, nil
}
