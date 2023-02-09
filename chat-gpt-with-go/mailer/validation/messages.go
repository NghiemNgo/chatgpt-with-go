package validation

import (
    "reflect"
)

func TransTag(tag string) string {
    tags := map[string]string{
        "eq":       "phải bằng",
        "gt":       "phải lớn hơn",
        "gte":      "phải lớn hơn hoặc bằng",
        "lt":       "phải nhỏ hơn",
        "lte":      "phải nhỏ hơn hoặc bằng",
        "ne":       "không bằng",
        "file":     "không đúng đường dẫn file",  
        "max":      "tối đa",
        "min":      "tối thiểu",
        "required": "là bắt buộc",
        "unique":   "đã tồn tại",
        "email":    "sai định dạng email",
    }

    if val, ok := tags[tag]; ok {
        return val
    }

    return tag

}

func Messages(field, tag, val string, param any) string {
    if (reflect.ValueOf(param).Kind() == reflect.String) {
        val += " ký tự"
    }
    return (field + " " + TransTag(tag) + " " + val)
}