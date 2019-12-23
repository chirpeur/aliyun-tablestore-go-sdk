package tablestore

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore/otsprotocol"
)

func TestOtsError_Error(t *testing.T) {
	errCode := "OTSObjectNotExist"
	errMessage := "Requested stream does not exist."
	reqId := "00057eff-2268-8529-1adc-e60b025ba100"
	pbErr := &otsprotocol.Error{
		Code:    &errCode,
		Message: &errMessage,
	}

	oldErrStr := fmt.Errorf("%s %s %s", *pbErr.Code, *pbErr.Message, reqId)

	otsErr := pbErrToOtsError(500, pbErr, reqId)

	if otsErr.Error() != oldErrStr.Error() {
		t.Errorf("error string not equal, old %s new %s", oldErrStr.Error(), otsErr.Error())
	}
}

func TestErrorValue(t *testing.T) {
	badHappen := func() error {
		return &OtsError{
			Message:        "foo is bad",
			RequestId:      "123",
			HttpStatusCode: 403,
			error:          ErrFoo,
		}
	}
	err := badHappen()
	assert.True(t, errors.Is(err, ErrFoo))
	_, ok := err.(*OtsError)
	assert.True(t, ok)
}
