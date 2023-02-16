package odoo

import (
	"fmt"
)

// AccountReportExternalValue represents account.report.external.value model.
type AccountReportExternalValue struct {
	LastUpdate                     *Time     `xmlrpc:"__last_update,omptempty"`
	CarryoverOriginExpressionLabel *String   `xmlrpc:"carryover_origin_expression_label,omptempty"`
	CarryoverOriginReportLineId    *Many2One `xmlrpc:"carryover_origin_report_line_id,omptempty"`
	CompanyId                      *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate                     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                      *Many2One `xmlrpc:"create_uid,omptempty"`
	Date                           *Time     `xmlrpc:"date,omptempty"`
	DisplayName                    *String   `xmlrpc:"display_name,omptempty"`
	ForeignVatFiscalPositionId     *Many2One `xmlrpc:"foreign_vat_fiscal_position_id,omptempty"`
	Id                             *Int      `xmlrpc:"id,omptempty"`
	Name                           *String   `xmlrpc:"name,omptempty"`
	ReportCountryId                *Many2One `xmlrpc:"report_country_id,omptempty"`
	TargetReportExpressionId       *Many2One `xmlrpc:"target_report_expression_id,omptempty"`
	TargetReportExpressionLabel    *String   `xmlrpc:"target_report_expression_label,omptempty"`
	TargetReportLineId             *Many2One `xmlrpc:"target_report_line_id,omptempty"`
	Value                          *Float    `xmlrpc:"value,omptempty"`
	WriteDate                      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountReportExternalValues represents array of account.report.external.value model.
type AccountReportExternalValues []AccountReportExternalValue

// AccountReportExternalValueModel is the odoo model name.
const AccountReportExternalValueModel = "account.report.external.value"

// Many2One convert AccountReportExternalValue to *Many2One.
func (arev *AccountReportExternalValue) Many2One() *Many2One {
	return NewMany2One(arev.Id.Get(), "")
}

// CreateAccountReportExternalValue creates a new account.report.external.value model and returns its id.
func (c *Client) CreateAccountReportExternalValue(arev *AccountReportExternalValue) (int64, error) {
	return c.Create(AccountReportExternalValueModel, arev)
}

// UpdateAccountReportExternalValue updates an existing account.report.external.value record.
func (c *Client) UpdateAccountReportExternalValue(arev *AccountReportExternalValue) error {
	return c.UpdateAccountReportExternalValues([]int64{arev.Id.Get()}, arev)
}

// UpdateAccountReportExternalValues updates existing account.report.external.value records.
// All records (represented by ids) will be updated by arev values.
func (c *Client) UpdateAccountReportExternalValues(ids []int64, arev *AccountReportExternalValue) error {
	return c.Update(AccountReportExternalValueModel, ids, arev)
}

// DeleteAccountReportExternalValue deletes an existing account.report.external.value record.
func (c *Client) DeleteAccountReportExternalValue(id int64) error {
	return c.DeleteAccountReportExternalValues([]int64{id})
}

// DeleteAccountReportExternalValues deletes existing account.report.external.value records.
func (c *Client) DeleteAccountReportExternalValues(ids []int64) error {
	return c.Delete(AccountReportExternalValueModel, ids)
}

// GetAccountReportExternalValue gets account.report.external.value existing record.
func (c *Client) GetAccountReportExternalValue(id int64) (*AccountReportExternalValue, error) {
	arevs, err := c.GetAccountReportExternalValues([]int64{id})
	if err != nil {
		return nil, err
	}
	if arevs != nil && len(*arevs) > 0 {
		return &((*arevs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.report.external.value not found", id)
}

// GetAccountReportExternalValues gets account.report.external.value existing records.
func (c *Client) GetAccountReportExternalValues(ids []int64) (*AccountReportExternalValues, error) {
	arevs := &AccountReportExternalValues{}
	if err := c.Read(AccountReportExternalValueModel, ids, nil, arevs); err != nil {
		return nil, err
	}
	return arevs, nil
}

// FindAccountReportExternalValue finds account.report.external.value record by querying it with criteria.
func (c *Client) FindAccountReportExternalValue(criteria *Criteria) (*AccountReportExternalValue, error) {
	arevs := &AccountReportExternalValues{}
	if err := c.SearchRead(AccountReportExternalValueModel, criteria, NewOptions().Limit(1), arevs); err != nil {
		return nil, err
	}
	if arevs != nil && len(*arevs) > 0 {
		return &((*arevs)[0]), nil
	}
	return nil, fmt.Errorf("account.report.external.value was not found with criteria %v", criteria)
}

// FindAccountReportExternalValues finds account.report.external.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportExternalValues(criteria *Criteria, options *Options) (*AccountReportExternalValues, error) {
	arevs := &AccountReportExternalValues{}
	if err := c.SearchRead(AccountReportExternalValueModel, criteria, options, arevs); err != nil {
		return nil, err
	}
	return arevs, nil
}

// FindAccountReportExternalValueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportExternalValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportExternalValueModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportExternalValueId finds record id by querying it with criteria.
func (c *Client) FindAccountReportExternalValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportExternalValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.report.external.value was not found with criteria %v and options %v", criteria, options)
}
