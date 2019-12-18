package payment

import (
	"fmt"
	"strings"

	"github.com/satori/go.uuid"
)

func GetTransferOutBizNo() (string, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	newUid := ""
	for _, v := range strings.Split(fmt.Sprintf("%v", uid), "-") {
		newUid = fmt.Sprintf("%v%v", newUid, v)
	}

	return newUid, nil
}
