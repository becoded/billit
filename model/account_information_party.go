// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountInformationParty account information party
//
// swagger:model AccountInformationParty
type AccountInformationParty struct {

	// API allowed
	APIAllowed bool `json:"APIAllowed,omitempty"`

	// API enabled features
	APIEnabledFeatures *APIEnabledFeatures `json:"APIEnabledFeatures,omitempty"`

	// accountant code
	AccountantCode int32 `json:"AccountantCode,omitempty"`

	// accounting cost type name
	AccountingCostTypeName string `json:"AccountingCostTypeName,omitempty"`

	// addresses
	Addresses []*PartyAddress `json:"Addresses"`

	// auto send invoice reminders
	AutoSendInvoiceReminders bool `json:"AutoSendInvoiceReminders,omitempty"`

	// b i c
	BIC string `json:"BIC,omitempty"`

	// bank accounts
	BankAccounts []*BankAccount `json:"BankAccounts"`

	// book category t c
	BookCategoryTC string `json:"BookCategoryTC,omitempty"`

	// book year start
	// Format: date-time
	BookYearStart *strfmt.DateTime `json:"BookYearStart,omitempty"`

	// box
	Box string `json:"Box,omitempty"`

	// city
	City string `json:"City,omitempty"`

	// commercial name
	CommercialName string `json:"CommercialName,omitempty"`

	// contact
	Contact string `json:"Contact,omitempty"`

	// contact first name
	ContactFirstName string `json:"ContactFirstName,omitempty"`

	// contact last name
	ContactLastName string `json:"ContactLastName,omitempty"`

	// cost category code
	CostCategoryCode int32 `json:"CostCategoryCode,omitempty"`

	// country code
	CountryCode string `json:"CountryCode,omitempty"`

	// created
	// Format: date-time
	Created *strfmt.DateTime `json:"Created,omitempty"`

	// custom fields
	CustomFields []*MutableKeyValuePairStringString `json:"CustomFields"`

	// default currency t c
	DefaultCurrencyTC string `json:"DefaultCurrencyTC,omitempty"`

	// default expiry offset
	DefaultExpiryOffset int32 `json:"DefaultExpiryOffset,omitempty"`

	// default paid
	DefaultPaid bool `json:"DefaultPaid,omitempty"`

	// default payment method t c
	DefaultPaymentMethodTC string `json:"DefaultPaymentMethodTC,omitempty"`

	// default reduction percentage
	DefaultReductionPercentage float64 `json:"DefaultReductionPercentage,omitempty"`

	// default transport type t c
	DefaultTransportTypeTC string `json:"DefaultTransportTypeTC,omitempty"`

	// display name
	DisplayName string `json:"DisplayName,omitempty"`

	// email
	Email string `json:"Email,omitempty"`

	// external provider ID
	ExternalProviderID string `json:"ExternalProviderID,omitempty"`

	// external provider t c
	ExternalProviderTC string `json:"ExternalProviderTC,omitempty"`

	// fax
	Fax string `json:"Fax,omitempty"`

	// financial reduction days default
	FinancialReductionDaysDefault int32 `json:"FinancialReductionDaysDefault,omitempty"`

	// financial reduction percentage default
	FinancialReductionPercentageDefault float64 `json:"FinancialReductionPercentageDefault,omitempty"`

	// g l account code
	GLAccountCode int64 `json:"GLAccountCode,omitempty"`

	// g l default expiry offset
	GLDefaultExpiryOffset int64 `json:"GLDefaultExpiryOffset,omitempty"`

	// i b a n
	IBAN string `json:"IBAN,omitempty"`

	// identifiers
	Identifiers []*PartyIdentifier `json:"Identifiers"`

	// internal info
	InternalInfo string `json:"InternalInfo,omitempty"`

	// is high risk
	IsHighRisk bool `json:"IsHighRisk,omitempty"`

	// language
	Language string `json:"Language,omitempty"`

	// last modified
	// Format: date-time
	LastModified *strfmt.DateTime `json:"LastModified,omitempty"`

	// logo file ID
	// Example: 00000000-0000-0000-0000-000000000000
	// Format: uuid
	LogoFileID strfmt.UUID `json:"LogoFileID,omitempty"`

	// mobile
	Mobile string `json:"Mobile,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// nr
	Nr string `json:"Nr,omitempty"`

	// party ID
	PartyID int32 `json:"PartyID,omitempty"`

	// party type
	// Enum: ["Customer","Supplier"]
	PartyType string `json:"PartyType,omitempty"`

	// peppol identifier
	PeppolIdentifier string `json:"PeppolIdentifier,omitempty"`

	// phone
	Phone string `json:"Phone,omitempty"`

	// r p r
	RPR string `json:"RPR,omitempty"`

	// regime
	Regime string `json:"Regime,omitempty"`

	// role
	Role string `json:"Role,omitempty"`

	// send p d f
	SendPDF bool `json:"SendPDF,omitempty"`

	// send u b l
	SendUBL bool `json:"SendUBL,omitempty"`

	// small enterprise
	SmallEnterprise bool `json:"SmallEnterprise,omitempty"`

	// street
	Street string `json:"Street,omitempty"`

	// street number
	StreetNumber string `json:"StreetNumber,omitempty"`

	// sub regime
	SubRegime string `json:"SubRegime,omitempty"`

	// tab name alias
	TabNameAlias string `json:"TabNameAlias,omitempty"`

	// user language
	UserLanguage string `json:"UserLanguage,omitempty"`

	// users
	Users []*UserCoreInfo `json:"Users"`

	// v a t deductable
	VATDeductable bool `json:"VATDeductable,omitempty"`

	// v a t liable
	VATLiable bool `json:"VATLiable,omitempty"`

	// v a t number
	VATNumber string `json:"VATNumber,omitempty"`

	// ventilation code
	VentilationCode string `json:"VentilationCode,omitempty"`

	// website
	Website string `json:"Website,omitempty"`

	// zipcode
	Zipcode string `json:"Zipcode,omitempty"`
}

// Validate validates this account information party
func (m *AccountInformationParty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIEnabledFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBankAccounts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBookYearStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomFields(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogoFileID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountInformationParty) validateAPIEnabledFeatures(formats strfmt.Registry) error {
	if swag.IsZero(m.APIEnabledFeatures) { // not required
		return nil
	}

	if m.APIEnabledFeatures != nil {
		if err := m.APIEnabledFeatures.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("APIEnabledFeatures")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("APIEnabledFeatures")
			}
			return err
		}
	}

	return nil
}

func (m *AccountInformationParty) validateAddresses(formats strfmt.Registry) error {
	if swag.IsZero(m.Addresses) { // not required
		return nil
	}

	for i := 0; i < len(m.Addresses); i++ {
		if swag.IsZero(m.Addresses[i]) { // not required
			continue
		}

		if m.Addresses[i] != nil {
			if err := m.Addresses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountInformationParty) validateBankAccounts(formats strfmt.Registry) error {
	if swag.IsZero(m.BankAccounts) { // not required
		return nil
	}

	for i := 0; i < len(m.BankAccounts); i++ {
		if swag.IsZero(m.BankAccounts[i]) { // not required
			continue
		}

		if m.BankAccounts[i] != nil {
			if err := m.BankAccounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BankAccounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BankAccounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountInformationParty) validateBookYearStart(formats strfmt.Registry) error {
	if swag.IsZero(m.BookYearStart) { // not required
		return nil
	}

	if err := validate.FormatOf("BookYearStart", "body", "date-time", m.BookYearStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountInformationParty) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("Created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountInformationParty) validateCustomFields(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomFields) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomFields); i++ {
		if swag.IsZero(m.CustomFields[i]) { // not required
			continue
		}

		if m.CustomFields[i] != nil {
			if err := m.CustomFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CustomFields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CustomFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountInformationParty) validateIdentifiers(formats strfmt.Registry) error {
	if swag.IsZero(m.Identifiers) { // not required
		return nil
	}

	for i := 0; i < len(m.Identifiers); i++ {
		if swag.IsZero(m.Identifiers[i]) { // not required
			continue
		}

		if m.Identifiers[i] != nil {
			if err := m.Identifiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Identifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Identifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountInformationParty) validateLastModified(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AccountInformationParty) validateLogoFileID(formats strfmt.Registry) error {
	if swag.IsZero(m.LogoFileID) { // not required
		return nil
	}

	if err := validate.FormatOf("LogoFileID", "body", "uuid", m.LogoFileID.String(), formats); err != nil {
		return err
	}

	return nil
}

var accountInformationPartyTypePartyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Customer","Supplier"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountInformationPartyTypePartyTypePropEnum = append(accountInformationPartyTypePartyTypePropEnum, v)
	}
}

const (

	// AccountInformationPartyPartyTypeCustomer captures enum value "Customer"
	AccountInformationPartyPartyTypeCustomer string = "Customer"

	// AccountInformationPartyPartyTypeSupplier captures enum value "Supplier"
	AccountInformationPartyPartyTypeSupplier string = "Supplier"
)

// prop value enum
func (m *AccountInformationParty) validatePartyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, accountInformationPartyTypePartyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AccountInformationParty) validatePartyType(formats strfmt.Registry) error {
	if swag.IsZero(m.PartyType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePartyTypeEnum("PartyType", "body", m.PartyType); err != nil {
		return err
	}

	return nil
}

func (m *AccountInformationParty) validateUsers(formats strfmt.Registry) error {
	if swag.IsZero(m.Users) { // not required
		return nil
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this account information party based on the context it is used
func (m *AccountInformationParty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAPIEnabledFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAddresses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBankAccounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomFields(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIdentifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountInformationParty) contextValidateAPIEnabledFeatures(ctx context.Context, formats strfmt.Registry) error {

	if m.APIEnabledFeatures != nil {

		if swag.IsZero(m.APIEnabledFeatures) { // not required
			return nil
		}

		if err := m.APIEnabledFeatures.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("APIEnabledFeatures")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("APIEnabledFeatures")
			}
			return err
		}
	}

	return nil
}

func (m *AccountInformationParty) contextValidateAddresses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Addresses); i++ {

		if m.Addresses[i] != nil {

			if swag.IsZero(m.Addresses[i]) { // not required
				return nil
			}

			if err := m.Addresses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Addresses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Addresses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountInformationParty) contextValidateBankAccounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BankAccounts); i++ {

		if m.BankAccounts[i] != nil {

			if swag.IsZero(m.BankAccounts[i]) { // not required
				return nil
			}

			if err := m.BankAccounts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BankAccounts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BankAccounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountInformationParty) contextValidateCustomFields(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomFields); i++ {

		if m.CustomFields[i] != nil {

			if swag.IsZero(m.CustomFields[i]) { // not required
				return nil
			}

			if err := m.CustomFields[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("CustomFields" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("CustomFields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountInformationParty) contextValidateIdentifiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Identifiers); i++ {

		if m.Identifiers[i] != nil {

			if swag.IsZero(m.Identifiers[i]) { // not required
				return nil
			}

			if err := m.Identifiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Identifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Identifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AccountInformationParty) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Users); i++ {

		if m.Users[i] != nil {

			if swag.IsZero(m.Users[i]) { // not required
				return nil
			}

			if err := m.Users[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountInformationParty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountInformationParty) UnmarshalBinary(b []byte) error {
	var res AccountInformationParty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
