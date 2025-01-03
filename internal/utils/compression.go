package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
)

func CompressJSON(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&buffer)
	_, err = gzipWriter.Write(jsonData)
	if err != nil {
		return nil, err
	}
	gzipWriter.Close()

	return buffer.Bytes(), nil
}
