package odoo

import (
	"fmt"
)

// AccountAnalyticPlan represents account.analytic.plan model.
type AccountAnalyticPlan struct {
	LastUpdate           *Time      `xmlrpc:"__last_update,omptempty"`
	AccountCount         *Int       `xmlrpc:"account_count,omptempty"`
	AccountIds           *Relation  `xmlrpc:"account_ids,omptempty"`
	AllAccountCount      *Int       `xmlrpc:"all_account_count,omptempty"`
	ApplicabilityIds     *Relation  `xmlrpc:"applicability_ids,omptempty"`
	ChildrenCount        *Int       `xmlrpc:"children_count,omptempty"`
	ChildrenIds          *Relation  `xmlrpc:"children_ids,omptempty"`
	Color                *Int       `xmlrpc:"color,omptempty"`
	CompanyId            *Many2One  `xmlrpc:"company_id,omptempty"`
	CompleteName         *String    `xmlrpc:"complete_name,omptempty"`
	CreateDate           *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One  `xmlrpc:"create_uid,omptempty"`
	DefaultApplicability *Selection `xmlrpc:"default_applicability,omptempty"`
	Description          *String    `xmlrpc:"description,omptempty"`
	DisplayName          *String    `xmlrpc:"display_name,omptempty"`
	Id                   *Int       `xmlrpc:"id,omptempty"`
	Name                 *String    `xmlrpc:"name,omptempty"`
	ParentId             *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentPath           *String    `xmlrpc:"parent_path,omptempty"`
	WriteDate            *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountAnalyticPlans represents array of account.analytic.plan model.
type AccountAnalyticPlans []AccountAnalyticPlan

// AccountAnalyticPlanModel is the odoo model name.
const AccountAnalyticPlanModel = "account.analytic.plan"

// Many2One convert AccountAnalyticPlan to *Many2One.
func (aap *AccountAnalyticPlan) Many2One() *Many2One {
	return NewMany2One(aap.Id.Get(), "")
}

// CreateAccountAnalyticPlan creates a new account.analytic.plan model and returns its id.
func (c *Client) CreateAccountAnalyticPlan(aap *AccountAnalyticPlan) (int64, error) {
	return c.Create(AccountAnalyticPlanModel, aap)
}

// UpdateAccountAnalyticPlan updates an existing account.analytic.plan record.
func (c *Client) UpdateAccountAnalyticPlan(aap *AccountAnalyticPlan) error {
	return c.UpdateAccountAnalyticPlans([]int64{aap.Id.Get()}, aap)
}

// UpdateAccountAnalyticPlans updates existing account.analytic.plan records.
// All records (represented by ids) will be updated by aap values.
func (c *Client) UpdateAccountAnalyticPlans(ids []int64, aap *AccountAnalyticPlan) error {
	return c.Update(AccountAnalyticPlanModel, ids, aap)
}

// DeleteAccountAnalyticPlan deletes an existing account.analytic.plan record.
func (c *Client) DeleteAccountAnalyticPlan(id int64) error {
	return c.DeleteAccountAnalyticPlans([]int64{id})
}

// DeleteAccountAnalyticPlans deletes existing account.analytic.plan records.
func (c *Client) DeleteAccountAnalyticPlans(ids []int64) error {
	return c.Delete(AccountAnalyticPlanModel, ids)
}

// GetAccountAnalyticPlan gets account.analytic.plan existing record.
func (c *Client) GetAccountAnalyticPlan(id int64) (*AccountAnalyticPlan, error) {
	aaps, err := c.GetAccountAnalyticPlans([]int64{id})
	if err != nil {
		return nil, err
	}
	if aaps != nil && len(*aaps) > 0 {
		return &((*aaps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.analytic.plan not found", id)
}

// GetAccountAnalyticPlans gets account.analytic.plan existing records.
func (c *Client) GetAccountAnalyticPlans(ids []int64) (*AccountAnalyticPlans, error) {
	aaps := &AccountAnalyticPlans{}
	if err := c.Read(AccountAnalyticPlanModel, ids, nil, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAnalyticPlan finds account.analytic.plan record by querying it with criteria.
func (c *Client) FindAccountAnalyticPlan(criteria *Criteria) (*AccountAnalyticPlan, error) {
	aaps := &AccountAnalyticPlans{}
	if err := c.SearchRead(AccountAnalyticPlanModel, criteria, NewOptions().Limit(1), aaps); err != nil {
		return nil, err
	}
	if aaps != nil && len(*aaps) > 0 {
		return &((*aaps)[0]), nil
	}
	return nil, fmt.Errorf("account.analytic.plan was not found with criteria %v", criteria)
}

// FindAccountAnalyticPlans finds account.analytic.plan records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticPlans(criteria *Criteria, options *Options) (*AccountAnalyticPlans, error) {
	aaps := &AccountAnalyticPlans{}
	if err := c.SearchRead(AccountAnalyticPlanModel, criteria, options, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAnalyticPlanIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticPlanIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAnalyticPlanModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAnalyticPlanId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticPlanId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticPlanModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.analytic.plan was not found with criteria %v and options %v", criteria, options)
}
