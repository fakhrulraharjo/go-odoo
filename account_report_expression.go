package odoo

import (
	"fmt"
)

// AccountReportExpression represents account.report.expression model.
type AccountReportExpression struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	Auditable       *Bool      `xmlrpc:"auditable,omptempty"`
	BlankIfZero     *Bool      `xmlrpc:"blank_if_zero,omptempty"`
	CarryoverTarget *String    `xmlrpc:"carryover_target,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateScope       *Selection `xmlrpc:"date_scope,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Engine          *Selection `xmlrpc:"engine,omptempty"`
	FigureType      *Selection `xmlrpc:"figure_type,omptempty"`
	Formula         *String    `xmlrpc:"formula,omptempty"`
	GreenOnPositive *Bool      `xmlrpc:"green_on_positive,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	Label           *String    `xmlrpc:"label,omptempty"`
	ReportLineId    *Many2One  `xmlrpc:"report_line_id,omptempty"`
	ReportLineName  *String    `xmlrpc:"report_line_name,omptempty"`
	Subformula      *String    `xmlrpc:"subformula,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountReportExpressions represents array of account.report.expression model.
type AccountReportExpressions []AccountReportExpression

// AccountReportExpressionModel is the odoo model name.
const AccountReportExpressionModel = "account.report.expression"

// Many2One convert AccountReportExpression to *Many2One.
func (are *AccountReportExpression) Many2One() *Many2One {
	return NewMany2One(are.Id.Get(), "")
}

// CreateAccountReportExpression creates a new account.report.expression model and returns its id.
func (c *Client) CreateAccountReportExpression(are *AccountReportExpression) (int64, error) {
	return c.Create(AccountReportExpressionModel, are)
}

// UpdateAccountReportExpression updates an existing account.report.expression record.
func (c *Client) UpdateAccountReportExpression(are *AccountReportExpression) error {
	return c.UpdateAccountReportExpressions([]int64{are.Id.Get()}, are)
}

// UpdateAccountReportExpressions updates existing account.report.expression records.
// All records (represented by ids) will be updated by are values.
func (c *Client) UpdateAccountReportExpressions(ids []int64, are *AccountReportExpression) error {
	return c.Update(AccountReportExpressionModel, ids, are)
}

// DeleteAccountReportExpression deletes an existing account.report.expression record.
func (c *Client) DeleteAccountReportExpression(id int64) error {
	return c.DeleteAccountReportExpressions([]int64{id})
}

// DeleteAccountReportExpressions deletes existing account.report.expression records.
func (c *Client) DeleteAccountReportExpressions(ids []int64) error {
	return c.Delete(AccountReportExpressionModel, ids)
}

// GetAccountReportExpression gets account.report.expression existing record.
func (c *Client) GetAccountReportExpression(id int64) (*AccountReportExpression, error) {
	ares, err := c.GetAccountReportExpressions([]int64{id})
	if err != nil {
		return nil, err
	}
	if ares != nil && len(*ares) > 0 {
		return &((*ares)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.report.expression not found", id)
}

// GetAccountReportExpressions gets account.report.expression existing records.
func (c *Client) GetAccountReportExpressions(ids []int64) (*AccountReportExpressions, error) {
	ares := &AccountReportExpressions{}
	if err := c.Read(AccountReportExpressionModel, ids, nil, ares); err != nil {
		return nil, err
	}
	return ares, nil
}

// FindAccountReportExpression finds account.report.expression record by querying it with criteria.
func (c *Client) FindAccountReportExpression(criteria *Criteria) (*AccountReportExpression, error) {
	ares := &AccountReportExpressions{}
	if err := c.SearchRead(AccountReportExpressionModel, criteria, NewOptions().Limit(1), ares); err != nil {
		return nil, err
	}
	if ares != nil && len(*ares) > 0 {
		return &((*ares)[0]), nil
	}
	return nil, fmt.Errorf("account.report.expression was not found with criteria %v", criteria)
}

// FindAccountReportExpressions finds account.report.expression records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportExpressions(criteria *Criteria, options *Options) (*AccountReportExpressions, error) {
	ares := &AccountReportExpressions{}
	if err := c.SearchRead(AccountReportExpressionModel, criteria, options, ares); err != nil {
		return nil, err
	}
	return ares, nil
}

// FindAccountReportExpressionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportExpressionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportExpressionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportExpressionId finds record id by querying it with criteria.
func (c *Client) FindAccountReportExpressionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportExpressionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.report.expression was not found with criteria %v and options %v", criteria, options)
}
