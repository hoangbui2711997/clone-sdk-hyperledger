package helps

import (
	"strings"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

func IsError(err error) bool {
	return err != nil && !strings.Contains(err.Error(), generic.TErrorCode_EGood.String())
}

func IsNotExisted(err error) bool {
	return err != nil && strings.Contains(err.Error(), generic.TErrorCode_EItemNotExisted.String())
}
