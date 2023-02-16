package odoo

import (
	"fmt"
)

// AccountAutomaticEntryWizard represents account.automatic.entry.wizard model.
type AccountAutomaticEntryWizard struct {
	LastUpdate            *Time      `xmlrpc:"__last_update,omptempty"`
	AccountType           *Selection `xmlrpc:"account_type,omptempty"`
	Action                *Selection `xmlrpc:"action,omptempty"`
	CompanyCurrencyId     *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date                  *Time      `xmlrpc:"date,omptempty"`
	DestinationAccountId  *Many2One  `xmlrpc:"destination_account_id,omptempty"`
	DisplayCurrencyHelper *Bool      `xmlrpc:"display_currency_helper,omptempty"`
	DisplayName           *String    `xmlrpc:"display_name,omptempty"`
	ExpenseAccrualAccount *Many2One  `xmlrpc:"expense_accrual_account,omptempty"`
	Id                    *Int       `xmlrpc:"id,omptempty"`
	JournalId             *Many2One  `xmlrpc:"journal_id,omptempty"`
	MoveData              *String    `xmlrpc:"move_data,omptempty"`
	MoveLineIds           *Relation  `xmlrpc:"move_line_ids,omptempty"`
	Percentage            *Float     `xmlrpc:"percentage,omptempty"`
	PreviewMoveData       *String    `xmlrpc:"preview_move_data,omptempty"`
	RevenueAccrualAccount *Many2One  `xmlrpc:"revenue_accrual_account,omptempty"`
	TotalAmount           *Float     `xmlrpc:"total_amount,omptempty"`
	WriteDate             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountAutomaticEntryWizards represents array of account.automatic.entry.wizard model.
type AccountAutomaticEntryWizards []AccountAutomaticEntryWizard

// AccountAutomaticEntryWizardModel is the odoo model name.
const AccountAutomaticEntryWizardModel = "account.automatic.entry.wizard"

// Many2One convert AccountAutomaticEntryWizard to *Many2One.
func (aaew *AccountAutomaticEntryWizard) Many2One() *Many2One {
	return NewMany2One(aaew.Id.Get(), "")
}

// CreateAccountAutomaticEntryWizard creates a new account.automatic.entry.wizard model and returns its id.
func (c *Client) CreateAccountAutomaticEntryWizard(aaew *AccountAutomaticEntryWizard) (int64, error) {
	return c.Create(AccountAutomaticEntryWizardModel, aaew)
}

// UpdateAccountAutomaticEntryWizard updates an existing account.automatic.entry.wizard record.
func (c *Client) UpdateAccountAutomaticEntryWizard(aaew *AccountAutomaticEntryWizard) error {
	return c.UpdateAccountAutomaticEntryWizards([]int64{aaew.Id.Get()}, aaew)
}

// UpdateAccountAutomaticEntryWizards updates existing account.automatic.entry.wizard records.
// All records (represented by ids) will be updated by aaew values.
func (c *Client) UpdateAccountAutomaticEntryWizards(ids []int64, aaew *AccountAutomaticEntryWizard) error {
	return c.Update(AccountAutomaticEntryWizardModel, ids, aaew)
}

// DeleteAccountAutomaticEntryWizard deletes an existing account.automatic.entry.wizard record.
func (c *Client) DeleteAccountAutomaticEntryWizard(id int64) error {
	return c.DeleteAccountAutomaticEntryWizards([]int64{id})
}

// DeleteAccountAutomaticEntryWizards deletes existing account.automatic.entry.wizard records.
func (c *Client) DeleteAccountAutomaticEntryWizards(ids []int64) error {
	return c.Delete(AccountAutomaticEntryWizardModel, ids)
}

// GetAccountAutomaticEntryWizard gets account.automatic.entry.wizard existing record.
func (c *Client) GetAccountAutomaticEntryWizard(id int64) (*AccountAutomaticEntryWizard, error) {
	aaews, err := c.GetAccountAutomaticEntryWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if aaews != nil && len(*aaews) > 0 {
		return &((*aaews)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.automatic.entry.wizard not found", id)
}

// GetAccountAutomaticEntryWizards gets account.automatic.entry.wizard existing records.
func (c *Client) GetAccountAutomaticEntryWizards(ids []int64) (*AccountAutomaticEntryWizards, error) {
	aaews := &AccountAutomaticEntryWizards{}
	if err := c.Read(AccountAutomaticEntryWizardModel, ids, nil, aaews); err != nil {
		return nil, err
	}
	return aaews, nil
}

// FindAccountAutomaticEntryWizard finds account.automatic.entry.wizard record by querying it with criteria.
func (c *Client) FindAccountAutomaticEntryWizard(criteria *Criteria) (*AccountAutomaticEntryWizard, error) {
	aaews := &AccountAutomaticEntryWizards{}
	if err := c.SearchRead(AccountAutomaticEntryWizardModel, criteria, NewOptions().Limit(1), aaews); err != nil {
		return nil, err
	}
	if aaews != nil && len(*aaews) > 0 {
		return &((*aaews)[0]), nil
	}
	return nil, fmt.Errorf("account.automatic.entry.wizard was not found with criteria %v", criteria)
}

// FindAccountAutomaticEntryWizards finds account.automatic.entry.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAutomaticEntryWizards(criteria *Criteria, options *Options) (*AccountAutomaticEntryWizards, error) {
	aaews := &AccountAutomaticEntryWizards{}
	if err := c.SearchRead(AccountAutomaticEntryWizardModel, criteria, options, aaews); err != nil {
		return nil, err
	}
	return aaews, nil
}

// FindAccountAutomaticEntryWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAutomaticEntryWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAutomaticEntryWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAutomaticEntryWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountAutomaticEntryWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAutomaticEntryWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.automatic.entry.wizard was not found with criteria %v and options %v", criteria, options)
}
