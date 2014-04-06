package fix50

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type RequestForPositionsAck struct {
	quickfix.Message
}

func (m *RequestForPositionsAck) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SettleOnOpenFlag() (*field.SettleOnOpenFlag, error) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) NTPositionLimit() (*field.NTPositionLimit, error) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) TotalNumPosReports() (*field.TotalNumPosReports, error) {
	f := new(field.TotalNumPosReports)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) UnsolicitedIndicator() (*field.UnsolicitedIndicator, error) {
	f := new(field.UnsolicitedIndicator)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) PosReqResult() (*field.PosReqResult, error) {
	f := new(field.PosReqResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) PosMaintRptID() (*field.PosMaintRptID, error) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) PosReqID() (*field.PosReqID, error) {
	f := new(field.PosReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) NoInstrumentParties() (*field.NoInstrumentParties, error) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) MinPriceIncrement() (*field.MinPriceIncrement, error) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) ResponseDestination() (*field.ResponseDestination, error) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) UnitOfMeasure() (*field.UnitOfMeasure, error) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) ResponseTransportType() (*field.ResponseTransportType, error) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, error) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SecurityStatus() (*field.SecurityStatus, error) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) PositionLimit() (*field.PositionLimit, error) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) MaturityTime() (*field.MaturityTime, error) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SettlSessSubID() (*field.SettlSessSubID, error) {
	f := new(field.SettlSessSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) TimeUnit() (*field.TimeUnit, error) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) StrikeMultiplier() (*field.StrikeMultiplier, error) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) PosReqStatus() (*field.PosReqStatus, error) {
	f := new(field.PosReqStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) StrikeValue() (*field.StrikeValue, error) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *RequestForPositionsAck) PosReqType() (*field.PosReqType, error) {
	f := new(field.PosReqType)
	err := m.Body.Get(f)
	return f, err
}
