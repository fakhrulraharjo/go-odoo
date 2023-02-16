package odoo

import (
	"fmt"
)

// AccountBankStatement represents account.bank.statement model.
type AccountBankStatement struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	AttachmentIds  *Relation `xmlrpc:"attachment_ids,omptempty"`
	BalanceEnd     *Float    `xmlrpc:"balance_end,omptempty"`
	BalanceEndReal *Float    `xmlrpc:"balance_end_real,omptempty"`
	BalanceStart   *Float    `xmlrpc:"balance_start,omptempty"`
	CompanyId      *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId     *Many2One `xmlrpc:"currency_id,omptempty"`
	Date           *Time     `xmlrpc:"date,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	FirstLineIndex *String   `xmlrpc:"first_line_index,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	IsComplete     *Bool     `xmlrpc:"is_complete,omptempty"`
	IsValid        *Bool     `xmlrpc:"is_valid,omptempty"`
	JournalId      *Many2One `xmlrpc:"journal_id,omptempty"`
	LineIds        *Relation `xmlrpc:"line_ids,omptempty"`
	Name           *String   `xmlrpc:"name,omptempty"`
	Reference      *String   `xmlrpc:"reference,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountBankStatements represents array of account.bank.statement model.
type AccountBankStatements []AccountBankStatement

// AccountBankStatementModel is the odoo model name.
const AccountBankStatementModel = "account.bank.statement"

// Many2One convert AccountBankStatement to *Many2One.
func (abs *AccountBankStatement) Many2One() *Many2One {
	return NewMany2One(abs.Id.Get(), "")
}

// CreateAccountBankStatement creates a new account.bank.statement model and returns its id.
func (c *Client) CreateAccountBankStatement(abs *AccountBankStatement) (int64, error) {
	return c.Create(AccountBankStatementModel, abs)
}

// UpdateAccountBankStatement updates an existing account.bank.statement record.
func (c *Client) UpdateAccountBankStatement(abs *AccountBankStatement) error {
	return c.UpdateAccountBankStatements([]int64{abs.Id.Get()}, abs)
}

// UpdateAccountBankStatements updates existing account.bank.statement records.
// All records (represented by ids) will be updated by abs values.
func (c *Client) UpdateAccountBankStatements(ids []int64, abs *AccountBankStatement) error {
	return c.Update(AccountBankStatementModel, ids, abs)
}

// DeleteAccountBankStatement deletes an existing account.bank.statement record.
func (c *Client) DeleteAccountBankStatement(id int64) error {
	return c.DeleteAccountBankStatements([]int64{id})
}

// DeleteAccountBankStatements deletes existing account.bank.statement records.
func (c *Client) DeleteAccountBankStatements(ids []int64) error {
	return c.Delete(AccountBankStatementModel, ids)
}

// GetAccountBankStatement gets account.bank.statement existing record.
func (c *Client) GetAccountBankStatement(id int64) (*AccountBankStatement, error) {
	abss, err := c.GetAccountBankStatements([]int64{id})
	if err != nil {
		return nil, err
	}
	if abss != nil && len(*abss) > 0 {
		return &((*abss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.statement not found", id)
}

// GetAccountBankStatements gets account.bank.statement existing records.
func (c *Client) GetAccountBankStatements(ids []int64) (*AccountBankStatements, error) {
	abss := &AccountBankStatements{}
	if err := c.Read(AccountBankStatementModel, ids, nil, abss); err != nil {
		return nil, err
	}
	return abss, nil
}

// FindAccountBankStatement finds account.bank.statement record by querying it with criteria.
func (c *Client) FindAccountBankStatement(criteria *Criteria) (*AccountBankStatement, error) {
	abss := &AccountBankStatements{}
	if err := c.SearchRead(AccountBankStatementModel, criteria, NewOptions().Limit(1), abss); err != nil {
		return nil, err
	}
	if abss != nil && len(*abss) > 0 {
		return &((*abss)[0]), nil
	}
	return nil, fmt.Errorf("account.bank.statement was not found with criteria %v", criteria)
}

// FindAccountBankStatements finds account.bank.statement records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatements(criteria *Criteria, options *Options) (*AccountBankStatements, error) {
	abss := &AccountBankStatements{}
	if err := c.SearchRead(AccountBankStatementModel, criteria, options, abss); err != nil {
		return nil, err
	}
	return abss, nil
}

// FindAccountBankStatementIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankStatementModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankStatementId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.bank.statement was not found with criteria %v and options %v", criteria, options)
}
