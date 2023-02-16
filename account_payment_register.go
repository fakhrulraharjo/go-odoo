package odoo

import (
	"fmt"
)

// AccountPaymentRegister represents account.payment.register model.
type AccountPaymentRegister struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	Amount                        *Float     `xmlrpc:"amount,omptempty"`
	AvailableJournalIds           *Relation  `xmlrpc:"available_journal_ids,omptempty"`
	AvailablePartnerBankIds       *Relation  `xmlrpc:"available_partner_bank_ids,omptempty"`
	AvailablePaymentMethodLineIds *Relation  `xmlrpc:"available_payment_method_line_ids,omptempty"`
	CanEditWizard                 *Bool      `xmlrpc:"can_edit_wizard,omptempty"`
	CanGroupPayments              *Bool      `xmlrpc:"can_group_payments,omptempty"`
	Communication                 *String    `xmlrpc:"communication,omptempty"`
	CompanyCurrencyId             *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryCode                   *String    `xmlrpc:"country_code,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	EarlyPaymentDiscountMode      *Bool      `xmlrpc:"early_payment_discount_mode,omptempty"`
	GroupPayment                  *Bool      `xmlrpc:"group_payment,omptempty"`
	HideWriteoffSection           *Bool      `xmlrpc:"hide_writeoff_section,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	JournalId                     *Many2One  `xmlrpc:"journal_id,omptempty"`
	LineIds                       *Relation  `xmlrpc:"line_ids,omptempty"`
	PartnerBankId                 *Many2One  `xmlrpc:"partner_bank_id,omptempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerType                   *Selection `xmlrpc:"partner_type,omptempty"`
	PaymentDate                   *Time      `xmlrpc:"payment_date,omptempty"`
	PaymentDifference             *Float     `xmlrpc:"payment_difference,omptempty"`
	PaymentDifferenceHandling     *Selection `xmlrpc:"payment_difference_handling,omptempty"`
	PaymentMethodCode             *String    `xmlrpc:"payment_method_code,omptempty"`
	PaymentMethodLineId           *Many2One  `xmlrpc:"payment_method_line_id,omptempty"`
	PaymentTokenId                *Many2One  `xmlrpc:"payment_token_id,omptempty"`
	PaymentType                   *Selection `xmlrpc:"payment_type,omptempty"`
	RequirePartnerBankAccount     *Bool      `xmlrpc:"require_partner_bank_account,omptempty"`
	ShowPartnerBankAccount        *Bool      `xmlrpc:"show_partner_bank_account,omptempty"`
	SourceAmount                  *Float     `xmlrpc:"source_amount,omptempty"`
	SourceAmountCurrency          *Float     `xmlrpc:"source_amount_currency,omptempty"`
	SourceCurrencyId              *Many2One  `xmlrpc:"source_currency_id,omptempty"`
	SuitablePaymentTokenIds       *Relation  `xmlrpc:"suitable_payment_token_ids,omptempty"`
	UseElectronicPaymentMethod    *Bool      `xmlrpc:"use_electronic_payment_method,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
	WriteoffAccountId             *Many2One  `xmlrpc:"writeoff_account_id,omptempty"`
	WriteoffLabel                 *String    `xmlrpc:"writeoff_label,omptempty"`
}

// AccountPaymentRegisters represents array of account.payment.register model.
type AccountPaymentRegisters []AccountPaymentRegister

// AccountPaymentRegisterModel is the odoo model name.
const AccountPaymentRegisterModel = "account.payment.register"

// Many2One convert AccountPaymentRegister to *Many2One.
func (apr *AccountPaymentRegister) Many2One() *Many2One {
	return NewMany2One(apr.Id.Get(), "")
}

// CreateAccountPaymentRegister creates a new account.payment.register model and returns its id.
func (c *Client) CreateAccountPaymentRegister(apr *AccountPaymentRegister) (int64, error) {
	return c.Create(AccountPaymentRegisterModel, apr)
}

// UpdateAccountPaymentRegister updates an existing account.payment.register record.
func (c *Client) UpdateAccountPaymentRegister(apr *AccountPaymentRegister) error {
	return c.UpdateAccountPaymentRegisters([]int64{apr.Id.Get()}, apr)
}

// UpdateAccountPaymentRegisters updates existing account.payment.register records.
// All records (represented by ids) will be updated by apr values.
func (c *Client) UpdateAccountPaymentRegisters(ids []int64, apr *AccountPaymentRegister) error {
	return c.Update(AccountPaymentRegisterModel, ids, apr)
}

// DeleteAccountPaymentRegister deletes an existing account.payment.register record.
func (c *Client) DeleteAccountPaymentRegister(id int64) error {
	return c.DeleteAccountPaymentRegisters([]int64{id})
}

// DeleteAccountPaymentRegisters deletes existing account.payment.register records.
func (c *Client) DeleteAccountPaymentRegisters(ids []int64) error {
	return c.Delete(AccountPaymentRegisterModel, ids)
}

// GetAccountPaymentRegister gets account.payment.register existing record.
func (c *Client) GetAccountPaymentRegister(id int64) (*AccountPaymentRegister, error) {
	aprs, err := c.GetAccountPaymentRegisters([]int64{id})
	if err != nil {
		return nil, err
	}
	if aprs != nil && len(*aprs) > 0 {
		return &((*aprs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.payment.register not found", id)
}

// GetAccountPaymentRegisters gets account.payment.register existing records.
func (c *Client) GetAccountPaymentRegisters(ids []int64) (*AccountPaymentRegisters, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.Read(AccountPaymentRegisterModel, ids, nil, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPaymentRegister finds account.payment.register record by querying it with criteria.
func (c *Client) FindAccountPaymentRegister(criteria *Criteria) (*AccountPaymentRegister, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.SearchRead(AccountPaymentRegisterModel, criteria, NewOptions().Limit(1), aprs); err != nil {
		return nil, err
	}
	if aprs != nil && len(*aprs) > 0 {
		return &((*aprs)[0]), nil
	}
	return nil, fmt.Errorf("account.payment.register was not found with criteria %v", criteria)
}

// FindAccountPaymentRegisters finds account.payment.register records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentRegisters(criteria *Criteria, options *Options) (*AccountPaymentRegisters, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.SearchRead(AccountPaymentRegisterModel, criteria, options, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPaymentRegisterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentRegisterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountPaymentRegisterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountPaymentRegisterId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentRegisterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentRegisterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.payment.register was not found with criteria %v and options %v", criteria, options)
}
