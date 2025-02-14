package to

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
)

// string to bool
func ToBool(s string) (bool, error) {
	return strconv.ParseBool(s)
}

// any to []byte
func ToBytes(value any) []byte {
	var err error
	defer func() {
		if err != nil {
			log.Println("run ToBytes err ", err)
		}
	}()
	v := reflect.ValueOf(value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		number := v.Int()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err = binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes()
	case uint, uint8, uint16, uint32, uint64:
		number := v.Uint()
		buf := bytes.NewBuffer([]byte{})
		buf.Reset()
		err = binary.Write(buf, binary.BigEndian, number)
		return buf.Bytes()
	case float32:
		number := float32(v.Float())
		bits := math.Float32bits(number)
		b := make([]byte, 4)
		binary.BigEndian.PutUint32(b, bits)
		return b
	case float64:
		number := v.Float()
		bits := math.Float64bits(number)
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, bits)
		return b
	case bool:
		return strconv.AppendBool([]byte{}, v.Bool())
	case string:
		return []byte(v.String())
	case []byte:
		return v.Bytes()
	default:
		fmt.Append([]byte{}, value)
	}
	return nil

}

// any to string
func ToString(value any) string {
	if value == nil {
		return ""
	}

	switch val := value.(type) {
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64)
	case int:
		return strconv.FormatInt(int64(val), 10)
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(val, 10)
	case string:
		return val
	case []byte:
		return string(val)
	default:
		b, err := json.Marshal(val)
		if err != nil {
			return ""
		}
		return string(b)

	}
}

// any to json
func ToJson(value any) (string, error) {
	result, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// any to float64
func ToFloat(value any) (float64, error) {
	v := reflect.ValueOf(value)

	result := 0.0
	err := fmt.Errorf("ToInt: unvalid interface type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = float64(v.Int())
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = float64(v.Uint())
		return result, nil
	case float32, float64:
		result = v.Float()
		return result, nil
	case string:
		result, err = strconv.ParseFloat(v.String(), 64)
		if err != nil {
			result = 0.0
		}
		return result, err
	default:
		return result, err
	}
}

// any to int64
func ToInt(value any) (int64, error) {
	v := reflect.ValueOf(value)

	var result int64
	err := fmt.Errorf("ToInt: invalid value type %T", value)
	switch value.(type) {
	case int, int8, int16, int32, int64:
		result = v.Int()
		return result, nil
	case uint, uint8, uint16, uint32, uint64:
		result = int64(v.Uint())
		return result, nil
	case float32, float64:
		result = int64(v.Float())
		return result, nil
	case string:
		result, err = strconv.ParseInt(v.String(), 0, 64)
		if err != nil {
			result = 0
		}
		return result, err
	default:
		return result, err
	}
}
