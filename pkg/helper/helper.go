package helper

import (
	"encoding/json"
	"strings"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

func MarshalToStruct(data interface{}, resp interface{}) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, resp)
	if err != nil {
		return err
	}

	return nil
}

func ConvertMapToStruct(inputMap map[string]interface{}) (*structpb.Struct, error) {
	marshledInputMap, err := json.Marshal(inputMap)
	outputStruct := &structpb.Struct{}
	if err != nil {
		return outputStruct, err
	}
	err = protojson.Unmarshal(marshledInputMap, outputStruct)

	return outputStruct, err
}

func GetURLWithTableSlug(c *gin.Context) string {
	url := c.FullPath()
	if strings.Contains(url, ":table_slug") {
		tableSlug := c.Param("table_slug")
		url = strings.Replace(url, ":table_slug", tableSlug, -1)
	}

	return url
}

func GetUniqueValue(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; !ok {
			c = append(c, item)
		}
	}
	return
}
