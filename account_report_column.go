package odoo

import (
	"fmt"
)

// AccountReportColumn represents account.report.column model.
type AccountReportColumn struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	BlankIfZero         *Bool      `xmlrpc:"blank_if_zero,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	CustomAuditActionId *Many2One  `xmlrpc:"custom_audit_action_id,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	ExpressionLabel     *String    `xmlrpc:"expression_label,omptempty"`
	FigureType          *Selection `xmlrpc:"figure_type,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	Name                *String    `xmlrpc:"name,omptempty"`
	ReportId            *Many2One  `xmlrpc:"report_id,omptempty"`
	Sequence            *Int       `xmlrpc:"sequence,omptempty"`
	Sortable            *Bool      `xmlrpc:"sortable,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountReportColumns represents array of account.report.column model.
type AccountReportColumns []AccountReportColumn

// AccountReportColumnModel is the odoo model name.
const AccountReportColumnModel = "account.report.column"

// Many2One convert AccountReportColumn to *Many2One.
func (arc *AccountReportColumn) Many2One() *Many2One {
	return NewMany2One(arc.Id.Get(), "")
}

// CreateAccountReportColumn creates a new account.report.column model and returns its id.
func (c *Client) CreateAccountReportColumn(arc *AccountReportColumn) (int64, error) {
	return c.Create(AccountReportColumnModel, arc)
}

// UpdateAccountReportColumn updates an existing account.report.column record.
func (c *Client) UpdateAccountReportColumn(arc *AccountReportColumn) error {
	return c.UpdateAccountReportColumns([]int64{arc.Id.Get()}, arc)
}

// UpdateAccountReportColumns updates existing account.report.column records.
// All records (represented by ids) will be updated by arc values.
func (c *Client) UpdateAccountReportColumns(ids []int64, arc *AccountReportColumn) error {
	return c.Update(AccountReportColumnModel, ids, arc)
}

// DeleteAccountReportColumn deletes an existing account.report.column record.
func (c *Client) DeleteAccountReportColumn(id int64) error {
	return c.DeleteAccountReportColumns([]int64{id})
}

// DeleteAccountReportColumns deletes existing account.report.column records.
func (c *Client) DeleteAccountReportColumns(ids []int64) error {
	return c.Delete(AccountReportColumnModel, ids)
}

// GetAccountReportColumn gets account.report.column existing record.
func (c *Client) GetAccountReportColumn(id int64) (*AccountReportColumn, error) {
	arcs, err := c.GetAccountReportColumns([]int64{id})
	if err != nil {
		return nil, err
	}
	if arcs != nil && len(*arcs) > 0 {
		return &((*arcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.report.column not found", id)
}

// GetAccountReportColumns gets account.report.column existing records.
func (c *Client) GetAccountReportColumns(ids []int64) (*AccountReportColumns, error) {
	arcs := &AccountReportColumns{}
	if err := c.Read(AccountReportColumnModel, ids, nil, arcs); err != nil {
		return nil, err
	}
	return arcs, nil
}

// FindAccountReportColumn finds account.report.column record by querying it with criteria.
func (c *Client) FindAccountReportColumn(criteria *Criteria) (*AccountReportColumn, error) {
	arcs := &AccountReportColumns{}
	if err := c.SearchRead(AccountReportColumnModel, criteria, NewOptions().Limit(1), arcs); err != nil {
		return nil, err
	}
	if arcs != nil && len(*arcs) > 0 {
		return &((*arcs)[0]), nil
	}
	return nil, fmt.Errorf("account.report.column was not found with criteria %v", criteria)
}

// FindAccountReportColumns finds account.report.column records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportColumns(criteria *Criteria, options *Options) (*AccountReportColumns, error) {
	arcs := &AccountReportColumns{}
	if err := c.SearchRead(AccountReportColumnModel, criteria, options, arcs); err != nil {
		return nil, err
	}
	return arcs, nil
}

// FindAccountReportColumnIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportColumnIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportColumnModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportColumnId finds record id by querying it with criteria.
func (c *Client) FindAccountReportColumnId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportColumnModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.report.column was not found with criteria %v and options %v", criteria, options)
}
