package odoo

import (
	"fmt"
)

// AccountReportLine represents account.report.line model.
type AccountReportLine struct {
	LastUpdate          *Time     `xmlrpc:"__last_update,omptempty"`
	AccountCodesFormula *String   `xmlrpc:"account_codes_formula,omptempty"`
	ActionId            *Many2One `xmlrpc:"action_id,omptempty"`
	AggregationFormula  *String   `xmlrpc:"aggregation_formula,omptempty"`
	ChildrenIds         *Relation `xmlrpc:"children_ids,omptempty"`
	Code                *String   `xmlrpc:"code,omptempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String   `xmlrpc:"display_name,omptempty"`
	DomainFormula       *String   `xmlrpc:"domain_formula,omptempty"`
	ExpressionIds       *Relation `xmlrpc:"expression_ids,omptempty"`
	Foldable            *Bool     `xmlrpc:"foldable,omptempty"`
	Groupby             *String   `xmlrpc:"groupby,omptempty"`
	HideIfZero          *Bool     `xmlrpc:"hide_if_zero,omptempty"`
	HierarchyLevel      *Int      `xmlrpc:"hierarchy_level,omptempty"`
	Id                  *Int      `xmlrpc:"id,omptempty"`
	Name                *String   `xmlrpc:"name,omptempty"`
	ParentId            *Many2One `xmlrpc:"parent_id,omptempty"`
	PrintOnNewPage      *Bool     `xmlrpc:"print_on_new_page,omptempty"`
	ReportId            *Many2One `xmlrpc:"report_id,omptempty"`
	Sequence            *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountReportLines represents array of account.report.line model.
type AccountReportLines []AccountReportLine

// AccountReportLineModel is the odoo model name.
const AccountReportLineModel = "account.report.line"

// Many2One convert AccountReportLine to *Many2One.
func (arl *AccountReportLine) Many2One() *Many2One {
	return NewMany2One(arl.Id.Get(), "")
}

// CreateAccountReportLine creates a new account.report.line model and returns its id.
func (c *Client) CreateAccountReportLine(arl *AccountReportLine) (int64, error) {
	return c.Create(AccountReportLineModel, arl)
}

// UpdateAccountReportLine updates an existing account.report.line record.
func (c *Client) UpdateAccountReportLine(arl *AccountReportLine) error {
	return c.UpdateAccountReportLines([]int64{arl.Id.Get()}, arl)
}

// UpdateAccountReportLines updates existing account.report.line records.
// All records (represented by ids) will be updated by arl values.
func (c *Client) UpdateAccountReportLines(ids []int64, arl *AccountReportLine) error {
	return c.Update(AccountReportLineModel, ids, arl)
}

// DeleteAccountReportLine deletes an existing account.report.line record.
func (c *Client) DeleteAccountReportLine(id int64) error {
	return c.DeleteAccountReportLines([]int64{id})
}

// DeleteAccountReportLines deletes existing account.report.line records.
func (c *Client) DeleteAccountReportLines(ids []int64) error {
	return c.Delete(AccountReportLineModel, ids)
}

// GetAccountReportLine gets account.report.line existing record.
func (c *Client) GetAccountReportLine(id int64) (*AccountReportLine, error) {
	arls, err := c.GetAccountReportLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if arls != nil && len(*arls) > 0 {
		return &((*arls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.report.line not found", id)
}

// GetAccountReportLines gets account.report.line existing records.
func (c *Client) GetAccountReportLines(ids []int64) (*AccountReportLines, error) {
	arls := &AccountReportLines{}
	if err := c.Read(AccountReportLineModel, ids, nil, arls); err != nil {
		return nil, err
	}
	return arls, nil
}

// FindAccountReportLine finds account.report.line record by querying it with criteria.
func (c *Client) FindAccountReportLine(criteria *Criteria) (*AccountReportLine, error) {
	arls := &AccountReportLines{}
	if err := c.SearchRead(AccountReportLineModel, criteria, NewOptions().Limit(1), arls); err != nil {
		return nil, err
	}
	if arls != nil && len(*arls) > 0 {
		return &((*arls)[0]), nil
	}
	return nil, fmt.Errorf("account.report.line was not found with criteria %v", criteria)
}

// FindAccountReportLines finds account.report.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportLines(criteria *Criteria, options *Options) (*AccountReportLines, error) {
	arls := &AccountReportLines{}
	if err := c.SearchRead(AccountReportLineModel, criteria, options, arls); err != nil {
		return nil, err
	}
	return arls, nil
}

// FindAccountReportLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportLineId finds record id by querying it with criteria.
func (c *Client) FindAccountReportLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.report.line was not found with criteria %v and options %v", criteria, options)
}
