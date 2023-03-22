package nmid

import (
	"douyin-grab/grab"
	"douyin-grab/pkg/logger"
	"encoding/json"
	"fmt"
	"os"

	"github.com/HughNian/nmid/pkg/model"
	wor "github.com/HughNian/nmid/pkg/worker"
	"github.com/vmihailenco/msgpack"
)

const WorkerName = "DOUYIN-GRAB"

func RunWorker() {
	nmidSerAddr := os.Getenv("NMID_SERVER_HOST") + ":" + os.Getenv("NMID_SERVER_PORT")
	worker := wor.NewWorker().SetWorkerName(WorkerName)
	err := worker.AddServer("tcp", nmidSerAddr)
	if err != nil {
		logger.Error("worker init error %s", err)
		worker.WorkerClose()
		return
	}

	worker.AddFunction("GetLiveRoomInfo", GetLiveRoomInfo)

	if err = worker.WorkerReady(); err != nil {
		logger.Error("worker not ready error %s", err)
		worker.WorkerClose()
		return
	}

	worker.WorkerDo()
}

func GetLiveRoomInfo(job wor.Job) ([]byte, error) {
	paramKey := "room_url"

	resp := job.GetResponse()
	if nil == resp {
		return []byte(``), fmt.Errorf("response data error")
	}

	if len(resp.ParamsMap) > 0 {
		roomUrl := resp.ParamsMap[paramKey].(string)

		roomInfo, ttwid := grab.FetchLiveRoomInfo(roomUrl)
		retData := make(map[string]interface{})
		retData["info"] = roomInfo
		retData["ttwid"] = ttwid
		rdata, err := json.Marshal(retData)
		if err != nil {
			return []byte(``), err
		}

		retStruct := model.GetRetStruct()
		retStruct.Msg = "ok"
		retStruct.Data = rdata
		ret, err := msgpack.Marshal(retStruct)
		if nil != err {
			return []byte(``), err
		}

		resp.RetLen = uint32(len(ret))
		resp.Ret = ret

		return ret, nil
	}

	return nil, fmt.Errorf("response data error")
}
