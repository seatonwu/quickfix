package fix50sp1

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type BusinessMessageReject struct {
	quickfix.Message
}

func (m *BusinessMessageReject) RefSeqNum() (*field.RefSeqNum, error) {
	f := new(field.RefSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) BusinessRejectReason() (*field.BusinessRejectReason, error) {
	f := new(field.BusinessRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) RefApplVerID() (*field.RefApplVerID, error) {
	f := new(field.RefApplVerID)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) RefApplExtID() (*field.RefApplExtID, error) {
	f := new(field.RefApplExtID)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) RefMsgType() (*field.RefMsgType, error) {
	f := new(field.RefMsgType)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) BusinessRejectRefID() (*field.BusinessRejectRefID, error) {
	f := new(field.BusinessRejectRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *BusinessMessageReject) RefCstmApplVerID() (*field.RefCstmApplVerID, error) {
	f := new(field.RefCstmApplVerID)
	err := m.Body.Get(f)
	return f, err
}
