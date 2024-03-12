package nodecfg

import (
	"fmt"
	"os"
	"sync"
	"v2ray-manager/pkg/helpers"
)

const NODE_JSON_FILE_PATH = "etc/v2ray-node-cfg.json"
const USER_JSON_FILE_PATH = "etc/v2ray-user-cfg.json"

var fileLocker sync.RWMutex

func readFile(path string) ([]byte, error) {
	fileLocker.RLock()
	defer fileLocker.RUnlock()
	return os.ReadFile(path)
}

func writeFile(path string, value any) error {
	fileLocker.Lock()
	defer fileLocker.Unlock()
	str := helpers.ToJson(value, true)
	return os.WriteFile(path, []byte(str), os.ModePerm)
}

func getNodeUserFileName(nodeID string) string {
	return fmt.Sprintf("etc/v2ray-user-%s.json", nodeID)
}
