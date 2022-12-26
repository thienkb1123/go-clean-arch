package converter

import (
	"bytes"
	"encoding/json"
)

// Convert bytes to buffer helper
func AnyToBytesBuffer(i interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(i)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

// Convert bytes to any
func BytesToAny(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

// Convert any to bytes
func AnyToBytes(v any) ([]byte, error) {
	return json.Marshal(v)
}

// ...
func MapStringToSlice(maps map[string]any) []any {
	var results []any
	for key, value := range maps {
		results = append(results, key, value)
	}
	return results
}

// ...
// func ContextToSlice(ctx context.Context) []any {
// 	for key, value := range ctx {
// 		fmt.Println(key, value)
// 	}
// }
