package nodecfg

import (
	"encoding/json"
	"os"
	"v2ray-manager/pkg/logger"
	"v2ray-manager/protobuf/manager"

	"go.uber.org/zap"
)

type UserConfigJson map[string]*manager.UserInfo // key-email

func SetUsers(nodeID string, arr []*manager.UserInfo) error {
	userCache[nodeID] = &Cache{}
	for _, v := range arr {
		userCache[nodeID].Store(v.Email, v)
	}
	return writeUserFile(nodeID)
}

func DeleteUser(nodeID string, email string) error {
	userCache[nodeID].Delete(email)
	return writeUserFile(nodeID)
}

func AddUser(nodeID string, data *manager.UserInfo) error {
	if data == nil {
		return InvalidDataError
	}
	// 把相同uuid的数据删掉
	if oldValue := FindUserByUUID(nodeID, data.ID); oldValue != nil {
		userCache[nodeID].Delete(oldValue.Email)
	}
	userCache[nodeID].Store(data.Email, data)
	return writeUserFile(nodeID)
}

func GetUsers(nodeID string) []*manager.UserInfo {
	result := make([]*manager.UserInfo, 0)
	data := userCache[nodeID].GetAll()
	for _, v := range data {
		result = append(result, v.(*manager.UserInfo))
	}
	return result
}

func FindUserByUUID(nodeID, uuid string) *manager.UserInfo {
	data := userCache[nodeID].GetAll()
	for _, v := range data {
		value := v.(*manager.UserInfo)
		if value.ID == uuid {
			return value
		}
	}
	return nil
}

func loadUserFromJSON(nodeID string) error {
	userCache[nodeID] = &Cache{}
	jsonFilePath := getNodeUserFileName(nodeID)

	// 文件存在才去读取
	if _, err := os.Stat(jsonFilePath); err != nil {
		logger.Logger.Warn("文件不存在", zap.String("path", jsonFilePath))
		return nil
	}

	fileData, err := readFile(jsonFilePath)
	if err != nil {
		return err
	}

	expectData := make(UserConfigJson)
	if err := json.Unmarshal(fileData, &expectData); err != nil {
		return err
	}
	for _, v := range expectData {
		userCache[nodeID].Store(v.Email, v)
	}
	return nil
}

func writeUserFile(nodeID string) error {
	v := userCache[nodeID]
	jsonData := make(UserConfigJson)
	for emailKey, userValue := range v.GetAll() {
		email := emailKey.(string)
		user := userValue.(*manager.UserInfo)
		jsonData[email] = user
	}
	return writeFile(getNodeUserFileName(nodeID), jsonData)
}
