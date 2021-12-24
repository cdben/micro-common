package common

import "encoding/json"

func SwapTo(source, model interface{}) error {
	dataByte, err := json.Marshal(source)
	if err != nil {
		return err
	}
	err = json.Unmarshal(dataByte, model)
	return err
}
