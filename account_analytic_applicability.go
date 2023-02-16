package odoo

import (
	"fmt"
)

// AccountAnalyticApplicability represents account.analytic.applicability model.
type AccountAnalyticApplicability struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	AccountPrefix  *String    `xmlrpc:"account_prefix,omptempty"`
	AnalyticPlanId *Many2One  `xmlrpc:"analytic_plan_id,omptempty"`
	Applicability  *Selection `xmlrpc:"applicability,omptempty"`
	BusinessDomain *Selection `xmlrpc:"business_domain,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	ProductCategId *Many2One  `xmlrpc:"product_categ_id,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountAnalyticApplicabilitys represents array of account.analytic.applicability model.
type AccountAnalyticApplicabilitys []AccountAnalyticApplicability

// AccountAnalyticApplicabilityModel is the odoo model name.
const AccountAnalyticApplicabilityModel = "account.analytic.applicability"

// Many2One convert AccountAnalyticApplicability to *Many2One.
func (aaa *AccountAnalyticApplicability) Many2One() *Many2One {
	return NewMany2One(aaa.Id.Get(), "")
}

// CreateAccountAnalyticApplicability creates a new account.analytic.applicability model and returns its id.
func (c *Client) CreateAccountAnalyticApplicability(aaa *AccountAnalyticApplicability) (int64, error) {
	return c.Create(AccountAnalyticApplicabilityModel, aaa)
}

// UpdateAccountAnalyticApplicability updates an existing account.analytic.applicability record.
func (c *Client) UpdateAccountAnalyticApplicability(aaa *AccountAnalyticApplicability) error {
	return c.UpdateAccountAnalyticApplicabilitys([]int64{aaa.Id.Get()}, aaa)
}

// UpdateAccountAnalyticApplicabilitys updates existing account.analytic.applicability records.
// All records (represented by ids) will be updated by aaa values.
func (c *Client) UpdateAccountAnalyticApplicabilitys(ids []int64, aaa *AccountAnalyticApplicability) error {
	return c.Update(AccountAnalyticApplicabilityModel, ids, aaa)
}

// DeleteAccountAnalyticApplicability deletes an existing account.analytic.applicability record.
func (c *Client) DeleteAccountAnalyticApplicability(id int64) error {
	return c.DeleteAccountAnalyticApplicabilitys([]int64{id})
}

// DeleteAccountAnalyticApplicabilitys deletes existing account.analytic.applicability records.
func (c *Client) DeleteAccountAnalyticApplicabilitys(ids []int64) error {
	return c.Delete(AccountAnalyticApplicabilityModel, ids)
}

// GetAccountAnalyticApplicability gets account.analytic.applicability existing record.
func (c *Client) GetAccountAnalyticApplicability(id int64) (*AccountAnalyticApplicability, error) {
	aaas, err := c.GetAccountAnalyticApplicabilitys([]int64{id})
	if err != nil {
		return nil, err
	}
	if aaas != nil && len(*aaas) > 0 {
		return &((*aaas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.analytic.applicability not found", id)
}

// GetAccountAnalyticApplicabilitys gets account.analytic.applicability existing records.
func (c *Client) GetAccountAnalyticApplicabilitys(ids []int64) (*AccountAnalyticApplicabilitys, error) {
	aaas := &AccountAnalyticApplicabilitys{}
	if err := c.Read(AccountAnalyticApplicabilityModel, ids, nil, aaas); err != nil {
		return nil, err
	}
	return aaas, nil
}

// FindAccountAnalyticApplicability finds account.analytic.applicability record by querying it with criteria.
func (c *Client) FindAccountAnalyticApplicability(criteria *Criteria) (*AccountAnalyticApplicability, error) {
	aaas := &AccountAnalyticApplicabilitys{}
	if err := c.SearchRead(AccountAnalyticApplicabilityModel, criteria, NewOptions().Limit(1), aaas); err != nil {
		return nil, err
	}
	if aaas != nil && len(*aaas) > 0 {
		return &((*aaas)[0]), nil
	}
	return nil, fmt.Errorf("account.analytic.applicability was not found with criteria %v", criteria)
}

// FindAccountAnalyticApplicabilitys finds account.analytic.applicability records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticApplicabilitys(criteria *Criteria, options *Options) (*AccountAnalyticApplicabilitys, error) {
	aaas := &AccountAnalyticApplicabilitys{}
	if err := c.SearchRead(AccountAnalyticApplicabilityModel, criteria, options, aaas); err != nil {
		return nil, err
	}
	return aaas, nil
}

// FindAccountAnalyticApplicabilityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticApplicabilityIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAnalyticApplicabilityModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAnalyticApplicabilityId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticApplicabilityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticApplicabilityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.analytic.applicability was not found with criteria %v and options %v", criteria, options)
}
