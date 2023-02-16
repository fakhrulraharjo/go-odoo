package odoo

import (
	"fmt"
)

// AccountAccruedOrdersWizard represents account.accrued.orders.wizard model.
type AccountAccruedOrdersWizard struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	AccountId     *Many2One `xmlrpc:"account_id,omptempty"`
	Amount        *Float    `xmlrpc:"amount,omptempty"`
	CompanyId     *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId    *Many2One `xmlrpc:"currency_id,omptempty"`
	Date          *Time     `xmlrpc:"date,omptempty"`
	DisplayAmount *Bool     `xmlrpc:"display_amount,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	JournalId     *Many2One `xmlrpc:"journal_id,omptempty"`
	PreviewData   *String   `xmlrpc:"preview_data,omptempty"`
	ReversalDate  *Time     `xmlrpc:"reversal_date,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountAccruedOrdersWizards represents array of account.accrued.orders.wizard model.
type AccountAccruedOrdersWizards []AccountAccruedOrdersWizard

// AccountAccruedOrdersWizardModel is the odoo model name.
const AccountAccruedOrdersWizardModel = "account.accrued.orders.wizard"

// Many2One convert AccountAccruedOrdersWizard to *Many2One.
func (aaow *AccountAccruedOrdersWizard) Many2One() *Many2One {
	return NewMany2One(aaow.Id.Get(), "")
}

// CreateAccountAccruedOrdersWizard creates a new account.accrued.orders.wizard model and returns its id.
func (c *Client) CreateAccountAccruedOrdersWizard(aaow *AccountAccruedOrdersWizard) (int64, error) {
	return c.Create(AccountAccruedOrdersWizardModel, aaow)
}

// UpdateAccountAccruedOrdersWizard updates an existing account.accrued.orders.wizard record.
func (c *Client) UpdateAccountAccruedOrdersWizard(aaow *AccountAccruedOrdersWizard) error {
	return c.UpdateAccountAccruedOrdersWizards([]int64{aaow.Id.Get()}, aaow)
}

// UpdateAccountAccruedOrdersWizards updates existing account.accrued.orders.wizard records.
// All records (represented by ids) will be updated by aaow values.
func (c *Client) UpdateAccountAccruedOrdersWizards(ids []int64, aaow *AccountAccruedOrdersWizard) error {
	return c.Update(AccountAccruedOrdersWizardModel, ids, aaow)
}

// DeleteAccountAccruedOrdersWizard deletes an existing account.accrued.orders.wizard record.
func (c *Client) DeleteAccountAccruedOrdersWizard(id int64) error {
	return c.DeleteAccountAccruedOrdersWizards([]int64{id})
}

// DeleteAccountAccruedOrdersWizards deletes existing account.accrued.orders.wizard records.
func (c *Client) DeleteAccountAccruedOrdersWizards(ids []int64) error {
	return c.Delete(AccountAccruedOrdersWizardModel, ids)
}

// GetAccountAccruedOrdersWizard gets account.accrued.orders.wizard existing record.
func (c *Client) GetAccountAccruedOrdersWizard(id int64) (*AccountAccruedOrdersWizard, error) {
	aaows, err := c.GetAccountAccruedOrdersWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if aaows != nil && len(*aaows) > 0 {
		return &((*aaows)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.accrued.orders.wizard not found", id)
}

// GetAccountAccruedOrdersWizards gets account.accrued.orders.wizard existing records.
func (c *Client) GetAccountAccruedOrdersWizards(ids []int64) (*AccountAccruedOrdersWizards, error) {
	aaows := &AccountAccruedOrdersWizards{}
	if err := c.Read(AccountAccruedOrdersWizardModel, ids, nil, aaows); err != nil {
		return nil, err
	}
	return aaows, nil
}

// FindAccountAccruedOrdersWizard finds account.accrued.orders.wizard record by querying it with criteria.
func (c *Client) FindAccountAccruedOrdersWizard(criteria *Criteria) (*AccountAccruedOrdersWizard, error) {
	aaows := &AccountAccruedOrdersWizards{}
	if err := c.SearchRead(AccountAccruedOrdersWizardModel, criteria, NewOptions().Limit(1), aaows); err != nil {
		return nil, err
	}
	if aaows != nil && len(*aaows) > 0 {
		return &((*aaows)[0]), nil
	}
	return nil, fmt.Errorf("account.accrued.orders.wizard was not found with criteria %v", criteria)
}

// FindAccountAccruedOrdersWizards finds account.accrued.orders.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccruedOrdersWizards(criteria *Criteria, options *Options) (*AccountAccruedOrdersWizards, error) {
	aaows := &AccountAccruedOrdersWizards{}
	if err := c.SearchRead(AccountAccruedOrdersWizardModel, criteria, options, aaows); err != nil {
		return nil, err
	}
	return aaows, nil
}

// FindAccountAccruedOrdersWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccruedOrdersWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAccruedOrdersWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAccruedOrdersWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountAccruedOrdersWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccruedOrdersWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.accrued.orders.wizard was not found with criteria %v and options %v", criteria, options)
}
