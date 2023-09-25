package unionpay

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

// GenerateSignature 生成签名
func GenerateSignature(params interface{}, key string) (string, error) {
	// 使用反射获取参数的字段名和字段值
	v := reflect.ValueOf(params)
	if v.Kind() != reflect.Struct {
		return "", fmt.Errorf("参数必须是结构体")
	}

	var keys []string
	var paramStr strings.Builder

	for i := 0; i < v.NumField(); i++ {
		fieldName := v.Type().Field(i).Name
		paramStr.WriteString(fieldName)
		paramStr.WriteString("=")
		paramStr.WriteString(fmt.Sprintf("%v", v.Field(i).Interface()))
		paramStr.WriteString("&")
		keys = append(keys, fieldName)
	}

	// 将参数名按ASCII字典序排序
	sort.Strings(keys)

	// 构建待签名字符串
	var signStr strings.Builder
	for _, k := range keys {
		signStr.WriteString(k)
		signStr.WriteString("=")
		signStr.WriteString(fmt.Sprintf("%v", v.FieldByName(k).Interface()))
	}
	signStr.WriteString(key)

	// 计算SHA256签名
	hash := sha256.Sum256([]byte(signStr.String()))
	signature := hex.EncodeToString(hash[:])

	return strings.ToUpper(signature), nil
}
