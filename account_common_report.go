package odoo

import (
	"fmt"
)

// AccountCommonReport represents account.common.report model.
type AccountCommonReport struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateFrom    *Time      `xmlrpc:"date_from,omptempty"`
	DateTo      *Time      `xmlrpc:"date_to,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	JournalIds  *Relation  `xmlrpc:"journal_ids,omptempty"`
	TargetMove  *Selection `xmlrpc:"target_move,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountCommonReports represents array of account.common.report model.
type AccountCommonReports []AccountCommonReport

// AccountCommonReportModel is the odoo model name.
const AccountCommonReportModel = "account.common.report"

// Many2One convert AccountCommonReport to *Many2One.
func (acr *AccountCommonReport) Many2One() *Many2One {
	return NewMany2One(acr.Id.Get(), "")
}

// CreateAccountCommonReport creates a new account.common.report model and returns its id.
func (c *Client) CreateAccountCommonReport(acr *AccountCommonReport) (int64, error) {
	return c.Create(AccountCommonReportModel, acr)
}

// UpdateAccountCommonReport updates an existing account.common.report record.
func (c *Client) UpdateAccountCommonReport(acr *AccountCommonReport) error {
	return c.UpdateAccountCommonReports([]int64{acr.Id.Get()}, acr)
}

// UpdateAccountCommonReports updates existing account.common.report records.
// All records (represented by ids) will be updated by acr values.
func (c *Client) UpdateAccountCommonReports(ids []int64, acr *AccountCommonReport) error {
	return c.Update(AccountCommonReportModel, ids, acr)
}

// DeleteAccountCommonReport deletes an existing account.common.report record.
func (c *Client) DeleteAccountCommonReport(id int64) error {
	return c.DeleteAccountCommonReports([]int64{id})
}

// DeleteAccountCommonReports deletes existing account.common.report records.
func (c *Client) DeleteAccountCommonReports(ids []int64) error {
	return c.Delete(AccountCommonReportModel, ids)
}

// GetAccountCommonReport gets account.common.report existing record.
func (c *Client) GetAccountCommonReport(id int64) (*AccountCommonReport, error) {
	acrs, err := c.GetAccountCommonReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if acrs != nil && len(*acrs) > 0 {
		return &((*acrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.common.report not found", id)
}

// GetAccountCommonReports gets account.common.report existing records.
func (c *Client) GetAccountCommonReports(ids []int64) (*AccountCommonReports, error) {
	acrs := &AccountCommonReports{}
	if err := c.Read(AccountCommonReportModel, ids, nil, acrs); err != nil {
		return nil, err
	}
	return acrs, nil
}

// FindAccountCommonReport finds account.common.report record by querying it with criteria.
func (c *Client) FindAccountCommonReport(criteria *Criteria) (*AccountCommonReport, error) {
	acrs := &AccountCommonReports{}
	if err := c.SearchRead(AccountCommonReportModel, criteria, NewOptions().Limit(1), acrs); err != nil {
		return nil, err
	}
	if acrs != nil && len(*acrs) > 0 {
		return &((*acrs)[0]), nil
	}
	return nil, fmt.Errorf("account.common.report was not found with criteria %v", criteria)
}

// FindAccountCommonReports finds account.common.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCommonReports(criteria *Criteria, options *Options) (*AccountCommonReports, error) {
	acrs := &AccountCommonReports{}
	if err := c.SearchRead(AccountCommonReportModel, criteria, options, acrs); err != nil {
		return nil, err
	}
	return acrs, nil
}

// FindAccountCommonReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCommonReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountCommonReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountCommonReportId finds record id by querying it with criteria.
func (c *Client) FindAccountCommonReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCommonReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.common.report was not found with criteria %v and options %v", criteria, options)
}
