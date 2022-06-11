package localPath

import (
	"note_go_api/app/controller"
	model "note_go_api/app/model"
	"note_go_api/app/res"
)

type ElementTree struct {
	ID uint `json:"id"`
	//Key      uint           `json:"key"`
	Label string `json:"label"`
	//Title    string         `json:"title"`
	Parent   uint           `json:"parent"`
	Children *[]ElementTree `json:"children"`
}

func treeJson(eData []model.Topic, parentID uint) *[]ElementTree {
	elTreeData := []ElementTree{}
	childrenList := ElementTree{}

	for _, v := range eData {
		if v.Parent == parentID {
			childrenList.Label = v.Title
			//childrenList.Title = v.Title
			//childrenList.Key = v.BaseModel.ID
			childrenList.ID = v.BaseModel.ID
			childrenList.Parent = v.Parent
			jsonList := treeJson(eData, v.BaseModel.ID)
			childrenList.Children = jsonList
			elTreeData = append(elTreeData, childrenList)
		}
	}
	return &elTreeData
}

func Tree(c *controller.BaseController) {
	topicList := []model.Topic{}
	model.DB.Model(model.Topic{}).Find(&topicList)

	tree := c.Query("noTree")
	if tree != "" {
		list := []model.TopicJson{}
		for _, v := range topicList {
			v.TopicJson.ID = v.BaseModel.ID
			list = append(list, v.TopicJson)
		}
		c.RJson(res.Data(list))
		return
	}
	json := treeJson(topicList, 0)
	c.RJson(res.Data(json))
}
