package api

import (
	"reflect"
	"strings"
)

func ReplaceTagsInKeyField(key string, field string, paramTable map[string]interface{}) (string, string) {
	for tag, value := range paramTable {
		if vstr, ok := value.(string); ok {
			if strings.Contains(key, "@"+tag) {
				key = strings.Replace(key, "@"+tag, vstr, 1)
			}
			if strings.Contains(field, "@"+tag) {
				field = strings.Replace(field, "@"+tag, vstr, 1)
			}
		}
	}
	return key, field
}

func MixinParamEnhancer[i any, o any](f func(InParameter i) (ret o, err error), fixParam func(paramMap map[string]interface{}, param i) (out i, err error)) {
	var (
		_api   ApiInterface
		_apiIO *Api[i, o]
		ok     bool
	)
	if _api, ok = GetApiByFunc(reflect.ValueOf(f).Pointer()); !ok {
		return
	}
	if _apiIO, ok = _api.(*Api[i, o]); !ok {
		return
	}
	_apiIO.ParamEnhancer = fixParam
}

// MixinResultSaver is a mixin function to save result to db , and response the value to the web client.
// The resultFinalizer is a function used to save the result excuted by the service. & response the value to the web client. (hide the fields if you need)

func MixinResultSaver[i any, o any](f func(InParameter i) (ret o, err error), resultFinalizer func(param i, ret o, paramMap map[string]interface{}) (valueToWebclient interface{}, err error)) {
	var (
		_api   ApiInterface
		_apiIO *Api[i, o]
		ok     bool
	)
	if _api, ok = GetApiByFunc(reflect.ValueOf(f).Pointer()); !ok {
		return
	}
	if _apiIO, ok = _api.(*Api[i, o]); !ok {
		return
	}
	_apiIO.ResultFinalizer = resultFinalizer
}
