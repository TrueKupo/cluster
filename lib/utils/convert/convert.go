package convert

import (
	"sync"

	pr "github.com/truekupo/cluster/common/interfaces/messages/response"
	pt "github.com/truekupo/cluster/common/interfaces/messages/types"
	"github.com/truekupo/cluster/lib/blockchain/chain"
	ee "github.com/truekupo/cluster/lib/errors"
)

var (
	lock sync.Mutex

	errorToProtoStatusMap map[error]pr.StatusCode = map[error]pr.StatusCode{
		ee.ErrNotFound:   pr.StatusCode_NotFound,
		ee.ErrBadRequest: pr.StatusCode_BadRequest,
		ee.ErrInternal:   pr.StatusCode_InternalServerError,
		nil:              pr.StatusCode_OK,
	}

	directionToStringMap map[pt.TxDirection]string = map[pt.TxDirection]string{
		pt.TxDirection_INPUT:  chain.TransactionDirectionIn,
		pt.TxDirection_OUTPUT: chain.TransactionDirectionOut,
		pt.TxDirection_INOUT:  chain.TransactionDirectionInOut,
	}

	directionToProtoMap map[string]pt.TxDirection = map[string]pt.TxDirection{
		chain.TransactionDirectionIn:    pt.TxDirection_INPUT,
		chain.TransactionDirectionOut:   pt.TxDirection_OUTPUT,
		chain.TransactionDirectionInOut: pt.TxDirection_INOUT,
	}

	statusToStringMap map[pt.TxStatus]string = map[pt.TxStatus]string{
		pt.TxStatus_NEW:        chain.TransactionStatusNew,
		pt.TxStatus_PROCESSING: chain.TransactionStatusProcessing,
		pt.TxStatus_FINALIZED:  chain.TransactionStatusFinalized,
		pt.TxStatus_FAILED:     chain.TransactionStatusFailed,
	}

	statusToProtoMap map[string]pt.TxStatus = map[string]pt.TxStatus{
		chain.TransactionStatusNew:        pt.TxStatus_NEW,
		chain.TransactionStatusProcessing: pt.TxStatus_PROCESSING,
		chain.TransactionStatusFinalized:  pt.TxStatus_FINALIZED,
		chain.TransactionStatusFailed:     pt.TxStatus_FAILED,
	}
)

func ToProtoError(err error) pr.StatusCode {
	lock.Lock()
	defer lock.Unlock()

	return toProtoError(err)
}

func ToProtoTxDirection(direction string) pt.TxDirection {
	lock.Lock()
	defer lock.Unlock()

	return toProtoTxDirection(direction)
}

func ToStringTxDirection(direction pt.TxDirection) string {
	lock.Lock()
	defer lock.Unlock()

	return toStringTxDirection(direction)
}

func ToProtoTxStatus(status string) pt.TxStatus {
	lock.Lock()
	defer lock.Unlock()

	return toProtoTxStatus(status)
}

func ToStringTxStatus(status pt.TxStatus) string {
	lock.Lock()
	defer lock.Unlock()

	return toStringTxStatus(status)
}

func toProtoError(err error) pr.StatusCode {
	s, ok := errorToProtoStatusMap[err]
	if !ok {
		s = pr.StatusCode_InternalServerError
	}

	return s
}

func toProtoTxDirection(direction string) pt.TxDirection {
	d, ok := directionToProtoMap[direction]
	if !ok {
		d = pt.TxDirection_INPUT
	}

	return d
}

func toStringTxDirection(direction pt.TxDirection) string {
	d, ok := directionToStringMap[direction]
	if !ok {
		d = chain.TransactionDirectionIn
	}

	return d
}

func toProtoTxStatus(status string) pt.TxStatus {
	d, ok := statusToProtoMap[status]
	if !ok {
		d = pt.TxStatus_NEW
	}

	return d
}

func toStringTxStatus(status pt.TxStatus) string {
	d, ok := statusToStringMap[status]
	if !ok {
		d = chain.TransactionStatusNew
	}

	return d
}
