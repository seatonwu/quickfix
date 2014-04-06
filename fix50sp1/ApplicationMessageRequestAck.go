package fix50sp1

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type ApplicationMessageRequestAck struct {
	quickfix.Message
}

func (m *ApplicationMessageRequestAck) ApplResponseID() (*field.ApplResponseID, error) {
	f := new(field.ApplResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequestAck) ApplResponseType() (*field.ApplResponseType, error) {
	f := new(field.ApplResponseType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequestAck) ApplTotalMessageCount() (*field.ApplTotalMessageCount, error) {
	f := new(field.ApplTotalMessageCount)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequestAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequestAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequestAck) ApplReqID() (*field.ApplReqID, error) {
	f := new(field.ApplReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequestAck) ApplReqType() (*field.ApplReqType, error) {
	f := new(field.ApplReqType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequestAck) NoApplIDs() (*field.NoApplIDs, error) {
	f := new(field.NoApplIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *ApplicationMessageRequestAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
