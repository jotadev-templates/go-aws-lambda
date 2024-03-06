package dmncompany

import "encoding/json"

func ConvertDatabaseToDomain(data any) (company, error) {
	var result company

	dataByte, err := json.Marshal(data.(map[string]any))
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(dataByte, &result)
	return result, err
}

func ConvertRequestBodyToDomain(data []byte) (company, error) {
	var result company
	err := json.Unmarshal(data, &result)
	return result, err
}
