package odoo

import (
	"fmt"
)

// AccountReport represents account.report model.
type AccountReport struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AvailabilityCondition    *Selection `xmlrpc:"availability_condition,omptempty"`
	ChartTemplateId          *Many2One  `xmlrpc:"chart_template_id,omptempty"`
	ColumnIds                *Relation  `xmlrpc:"column_ids,omptempty"`
	CountryId                *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DefaultOpeningDateFilter *Selection `xmlrpc:"default_opening_date_filter,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	FilterAccountType        *Bool      `xmlrpc:"filter_account_type,omptempty"`
	FilterAnalytic           *Bool      `xmlrpc:"filter_analytic,omptempty"`
	FilterDateRange          *Bool      `xmlrpc:"filter_date_range,omptempty"`
	FilterFiscalPosition     *Bool      `xmlrpc:"filter_fiscal_position,omptempty"`
	FilterGrowthComparison   *Bool      `xmlrpc:"filter_growth_comparison,omptempty"`
	FilterHierarchy          *Selection `xmlrpc:"filter_hierarchy,omptempty"`
	FilterJournals           *Bool      `xmlrpc:"filter_journals,omptempty"`
	FilterMultiCompany       *Selection `xmlrpc:"filter_multi_company,omptempty"`
	FilterPartner            *Bool      `xmlrpc:"filter_partner,omptempty"`
	FilterPeriodComparison   *Bool      `xmlrpc:"filter_period_comparison,omptempty"`
	FilterShowDraft          *Bool      `xmlrpc:"filter_show_draft,omptempty"`
	FilterUnfoldAll          *Bool      `xmlrpc:"filter_unfold_all,omptempty"`
	FilterUnreconciled       *Bool      `xmlrpc:"filter_unreconciled,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	LineIds                  *Relation  `xmlrpc:"line_ids,omptempty"`
	LoadMoreLimit            *Int       `xmlrpc:"load_more_limit,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	OnlyTaxExigible          *Bool      `xmlrpc:"only_tax_exigible,omptempty"`
	RootReportId             *Many2One  `xmlrpc:"root_report_id,omptempty"`
	SearchBar                *Bool      `xmlrpc:"search_bar,omptempty"`
	VariantReportIds         *Relation  `xmlrpc:"variant_report_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountReports represents array of account.report model.
type AccountReports []AccountReport

// AccountReportModel is the odoo model name.
const AccountReportModel = "account.report"

// Many2One convert AccountReport to *Many2One.
func (ar *AccountReport) Many2One() *Many2One {
	return NewMany2One(ar.Id.Get(), "")
}

// CreateAccountReport creates a new account.report model and returns its id.
func (c *Client) CreateAccountReport(ar *AccountReport) (int64, error) {
	return c.Create(AccountReportModel, ar)
}

// UpdateAccountReport updates an existing account.report record.
func (c *Client) UpdateAccountReport(ar *AccountReport) error {
	return c.UpdateAccountReports([]int64{ar.Id.Get()}, ar)
}

// UpdateAccountReports updates existing account.report records.
// All records (represented by ids) will be updated by ar values.
func (c *Client) UpdateAccountReports(ids []int64, ar *AccountReport) error {
	return c.Update(AccountReportModel, ids, ar)
}

// DeleteAccountReport deletes an existing account.report record.
func (c *Client) DeleteAccountReport(id int64) error {
	return c.DeleteAccountReports([]int64{id})
}

// DeleteAccountReports deletes existing account.report records.
func (c *Client) DeleteAccountReports(ids []int64) error {
	return c.Delete(AccountReportModel, ids)
}

// GetAccountReport gets account.report existing record.
func (c *Client) GetAccountReport(id int64) (*AccountReport, error) {
	ars, err := c.GetAccountReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if ars != nil && len(*ars) > 0 {
		return &((*ars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.report not found", id)
}

// GetAccountReports gets account.report existing records.
func (c *Client) GetAccountReports(ids []int64) (*AccountReports, error) {
	ars := &AccountReports{}
	if err := c.Read(AccountReportModel, ids, nil, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAccountReport finds account.report record by querying it with criteria.
func (c *Client) FindAccountReport(criteria *Criteria) (*AccountReport, error) {
	ars := &AccountReports{}
	if err := c.SearchRead(AccountReportModel, criteria, NewOptions().Limit(1), ars); err != nil {
		return nil, err
	}
	if ars != nil && len(*ars) > 0 {
		return &((*ars)[0]), nil
	}
	return nil, fmt.Errorf("account.report was not found with criteria %v", criteria)
}

// FindAccountReports finds account.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReports(criteria *Criteria, options *Options) (*AccountReports, error) {
	ars := &AccountReports{}
	if err := c.SearchRead(AccountReportModel, criteria, options, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAccountReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportId finds record id by querying it with criteria.
func (c *Client) FindAccountReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.report was not found with criteria %v and options %v", criteria, options)
}
