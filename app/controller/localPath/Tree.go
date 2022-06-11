package localPath

import (
	"io/ioutil"
	"note_go_api/app/controller"
	"note_go_api/app/res"
)

type ElementTree struct {
	ID uint `json:"id"`
	//Key      uint           `json:"key"`
	Label    string `json:"label"`
	FullPath string `json:"full_path"`
	//Title    string         `json:"title"`
	//Parent   uint           `json:"parent"`
	Type     string         `json:"type"`
	Children *[]ElementTree `json:"children"`
}

func Tree(c *controller.BaseController) {
	treeList := filesToTree(_path, 0)
	c.RJson(res.Data(treeList))
}
func filesToTree(path string, startId uint) *[]ElementTree {
	trees := []ElementTree{}
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		return &trees
	}
	for _, fInfo := range dir {
		startId += 1
		tree := ElementTree{}
		nowPath := path + "/" + fInfo.Name()
		tree.Label = fInfo.Name()
		tree.FullPath = nowPath
		tree.ID = startId
		tree.Type = localPathTypeFile
		if fInfo.IsDir() {
			tree.Type = localPathTypeDir
			tree.Children = filesToTree(nowPath, startId)
		}
		trees = append(trees, tree)
	}
	return &trees
}
