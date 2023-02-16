package odoo

import (
	"fmt"
)

// AccountPaymentMethodLine represents account.payment.method.line model.
type AccountPaymentMethodLine struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omptempty"`
	AvailablePaymentMethodIds *Relation  `xmlrpc:"available_payment_method_ids,omptempty"`
	Code                      *String    `xmlrpc:"code,omptempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName               *String    `xmlrpc:"display_name,omptempty"`
	Id                        *Int       `xmlrpc:"id,omptempty"`
	JournalId                 *Many2One  `xmlrpc:"journal_id,omptempty"`
	Name                      *String    `xmlrpc:"name,omptempty"`
	PaymentAccountId          *Many2One  `xmlrpc:"payment_account_id,omptempty"`
	PaymentMethodId           *Many2One  `xmlrpc:"payment_method_id,omptempty"`
	PaymentProviderId         *Many2One  `xmlrpc:"payment_provider_id,omptempty"`
	PaymentProviderState      *Selection `xmlrpc:"payment_provider_state,omptempty"`
	PaymentType               *Selection `xmlrpc:"payment_type,omptempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omptempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountPaymentMethodLines represents array of account.payment.method.line model.
type AccountPaymentMethodLines []AccountPaymentMethodLine

// AccountPaymentMethodLineModel is the odoo model name.
const AccountPaymentMethodLineModel = "account.payment.method.line"

// Many2One convert AccountPaymentMethodLine to *Many2One.
func (apml *AccountPaymentMethodLine) Many2One() *Many2One {
	return NewMany2One(apml.Id.Get(), "")
}

// CreateAccountPaymentMethodLine creates a new account.payment.method.line model and returns its id.
func (c *Client) CreateAccountPaymentMethodLine(apml *AccountPaymentMethodLine) (int64, error) {
	return c.Create(AccountPaymentMethodLineModel, apml)
}

// UpdateAccountPaymentMethodLine updates an existing account.payment.method.line record.
func (c *Client) UpdateAccountPaymentMethodLine(apml *AccountPaymentMethodLine) error {
	return c.UpdateAccountPaymentMethodLines([]int64{apml.Id.Get()}, apml)
}

// UpdateAccountPaymentMethodLines updates existing account.payment.method.line records.
// All records (represented by ids) will be updated by apml values.
func (c *Client) UpdateAccountPaymentMethodLines(ids []int64, apml *AccountPaymentMethodLine) error {
	return c.Update(AccountPaymentMethodLineModel, ids, apml)
}

// DeleteAccountPaymentMethodLine deletes an existing account.payment.method.line record.
func (c *Client) DeleteAccountPaymentMethodLine(id int64) error {
	return c.DeleteAccountPaymentMethodLines([]int64{id})
}

// DeleteAccountPaymentMethodLines deletes existing account.payment.method.line records.
func (c *Client) DeleteAccountPaymentMethodLines(ids []int64) error {
	return c.Delete(AccountPaymentMethodLineModel, ids)
}

// GetAccountPaymentMethodLine gets account.payment.method.line existing record.
func (c *Client) GetAccountPaymentMethodLine(id int64) (*AccountPaymentMethodLine, error) {
	apmls, err := c.GetAccountPaymentMethodLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if apmls != nil && len(*apmls) > 0 {
		return &((*apmls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.payment.method.line not found", id)
}

// GetAccountPaymentMethodLines gets account.payment.method.line existing records.
func (c *Client) GetAccountPaymentMethodLines(ids []int64) (*AccountPaymentMethodLines, error) {
	apmls := &AccountPaymentMethodLines{}
	if err := c.Read(AccountPaymentMethodLineModel, ids, nil, apmls); err != nil {
		return nil, err
	}
	return apmls, nil
}

// FindAccountPaymentMethodLine finds account.payment.method.line record by querying it with criteria.
func (c *Client) FindAccountPaymentMethodLine(criteria *Criteria) (*AccountPaymentMethodLine, error) {
	apmls := &AccountPaymentMethodLines{}
	if err := c.SearchRead(AccountPaymentMethodLineModel, criteria, NewOptions().Limit(1), apmls); err != nil {
		return nil, err
	}
	if apmls != nil && len(*apmls) > 0 {
		return &((*apmls)[0]), nil
	}
	return nil, fmt.Errorf("account.payment.method.line was not found with criteria %v", criteria)
}

// FindAccountPaymentMethodLines finds account.payment.method.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentMethodLines(criteria *Criteria, options *Options) (*AccountPaymentMethodLines, error) {
	apmls := &AccountPaymentMethodLines{}
	if err := c.SearchRead(AccountPaymentMethodLineModel, criteria, options, apmls); err != nil {
		return nil, err
	}
	return apmls, nil
}

// FindAccountPaymentMethodLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentMethodLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountPaymentMethodLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountPaymentMethodLineId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentMethodLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentMethodLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.payment.method.line was not found with criteria %v and options %v", criteria, options)
}
