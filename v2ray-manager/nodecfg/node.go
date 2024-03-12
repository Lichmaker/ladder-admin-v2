package nodecfg

import (
	"encoding/json"
	"fmt"
	"v2ray-manager/pkg/logger"
)

type NodeConfigJson map[string]*NodeInfo

func Init() {
	if err := loadNodeFromJSON(); err != nil {
		logger.Logger.Warn(fmt.Sprintf("加载节点配置json失败:%s", err))
	} else {
		for _, v := range GetAllNode() {
			if err := loadUserFromJSON(v.ID); err != nil {
				logger.Logger.Warn(fmt.Sprintf("加载节点用户json失败:%s", err))
			}
		}
	}
}

func SetNode(arr []*NodeInfo) error {
	nodeCache.Clear()
	for _, v := range arr {
		nodeCache.Store(v.ID, v)
		userCache[v.ID] = &Cache{}
	}
	return writeNodeFile()
}

func DeleteNode(ID string) error {
	nodeCache.Delete(ID)
	delete(userCache, ID)
	return writeNodeFile()
}

func AddNode(info *NodeInfo) error {
	nodeCache.Store(info.ID, info)
	userCache[info.ID] = &Cache{}
	return writeNodeFile()
}

func GetAllNode() []*NodeInfo {
	data := nodeCache.GetAll()
	var result []*NodeInfo = make([]*NodeInfo, 0)
	for _, v := range data {
		result = append(result, v.(*NodeInfo))
	}
	return result
}

func GetByID(nodeID string) *NodeInfo {
	v, exists := nodeCache.Get(nodeID)
	if !exists {
		return nil
	}
	return v.(*NodeInfo)
}

func loadNodeFromJSON() error {
	fileData, err := readFile(NODE_JSON_FILE_PATH)
	if err != nil {
		return err
	}

	expectData := make(NodeConfigJson)
	if err := json.Unmarshal(fileData, &expectData); err != nil {
		return err
	}
	for _, v := range expectData {
		nodeCache.Store(v.ID, v)
	}
	return nil
}

func writeNodeFile() error {
	cacheData := nodeCache.GetAll()
	jsonData := make(NodeConfigJson)
	for k, v := range cacheData {
		newKey := k.(string)
		newValue := v.(*NodeInfo)
		jsonData[newKey] = newValue
	}
	return writeFile(NODE_JSON_FILE_PATH, jsonData)
}
