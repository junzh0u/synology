package synology

import (
	"errors"
	"fmt"
)

func errFromData(data interface{}, knownErrors map[int]string) error {
	success := data.(map[string]interface{})["success"].(bool)
	if success {
		return nil
	}

	errcode := int(data.(map[string]interface{})["error"].(map[string]interface{})["code"].(float64))
	return errFromCode(errcode, knownErrors)
}

func errFromCode(code int, knownErrors map[int]string) error {
	for k, v := range map[int]string{
		100: "Unknown error",
		101: "Invalid parameter",
		102: "The requested API does not exist",
		103: "The requested method does not exist",
		104: "The requested version does not support the functionality",
		105: "The logged in session does not have permission",
		106: "Session timeout",
		107: "Session interrupted by duplicate login",
	} {
		knownErrors[k] = v
	}
	reason, ok := knownErrors[code]
	if ok {
		return errors.New(reason)
	}
	return fmt.Errorf("unknown errocode %d", code)
}
