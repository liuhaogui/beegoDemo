package utils

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

const (
	SUCCESS        = 200
	ACCEPTED       = 202
	ERROR          = 501
	FAILED         = 502
	PARAM_ERROR    = 400
	FORBIDDEN      = 403
	NO_AUTH        = 405
	LOGIN_TIME_OUT = 401

	ADDRESS_VALID = 4001
)

var ReturnMessage = map[int]string{
	SUCCESS:        "success",
	ERROR:          "error",
	FAILED:         "failed",
	PARAM_ERROR:    "parameter error",
	FORBIDDEN:      "forbidden",
	LOGIN_TIME_OUT: "login time out",
	NO_AUTH:        "Insufficient authority",
	ADDRESS_VALID:  "Reset password address is no longer valid",
}

func GetMsgByCode(code int) string {
	if msg, ok := ReturnMessage[code]; ok {
		return msg
	} else {
		log.Error("undefined return message code %d ", code)
		return fmt.Sprintf("error, code : %d ", code)
	}
}

func JsonMsg(msg interface{}, code int, data interface{}) (returnData map[string]interface{}) {
	errorMsg := fmt.Sprintf("%s", msg)
	if len(errorMsg) < 1 {
		errorMsg = GetMsgByCode(code)
	}
	returnData = map[string]interface{}{"code": code, "msg": errorMsg, "data": data}
	return returnData
}
