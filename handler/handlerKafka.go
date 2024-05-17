package handler

import (
	"api/util/time_relate"
	"encoding/json"
	"sync"
	"time"
)

var ListenKafkaTestMappingMap = make(map[string][]byte, 0)

type DeleteTimeOutData struct {
	CreatedTime time.Time `json:"createdTime"`
}

func ListenKafkaTest(dataByte []byte) {
	type receivedData struct {
		Uuid    string `json:"uuid"`
		Status  bool   `json:"status"`
		Message string `json:"message"`
	}

	var data receivedData
	var mutex sync.Mutex
	if err := json.Unmarshal(dataByte, &data); err == nil {
		mutex.Lock()
		ListenKafkaTestMappingMap[data.Uuid] = dataByte
		if len(ListenKafkaTestMappingMap) > 1000 {
			var delData DeleteTimeOutData
			for key, value := range ListenKafkaTestMappingMap {
				if err := json.Unmarshal(value, &delData); err == nil {
					if time_relate.TimeNow().Sub(delData.CreatedTime).Minutes() >= 5 {
						delete(ListenKafkaTestMappingMap, key)
					}
				}
			}
		}
		mutex.Unlock()
	}
}
