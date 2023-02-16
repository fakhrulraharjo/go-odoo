package odoo

import (
	"fmt"
)

// AccountAnalyticDistributionModel represents account.analytic.distribution.model model.
type AccountAnalyticDistributionModel struct {
	LastUpdate           *Time       `xmlrpc:"__last_update,omptempty"`
	AccountPrefix        *String     `xmlrpc:"account_prefix,omptempty"`
	AnalyticDistribution interface{} `xmlrpc:"analytic_distribution,omptempty"`
	AnalyticPrecision    *Int        `xmlrpc:"analytic_precision,omptempty"`
	CompanyId            *Many2One   `xmlrpc:"company_id,omptempty"`
	CreateDate           *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid            *Many2One   `xmlrpc:"create_uid,omptempty"`
	DisplayName          *String     `xmlrpc:"display_name,omptempty"`
	Id                   *Int        `xmlrpc:"id,omptempty"`
	PartnerCategoryId    *Many2One   `xmlrpc:"partner_category_id,omptempty"`
	PartnerId            *Many2One   `xmlrpc:"partner_id,omptempty"`
	ProductCategId       *Many2One   `xmlrpc:"product_categ_id,omptempty"`
	ProductId            *Many2One   `xmlrpc:"product_id,omptempty"`
	WriteDate            *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid             *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// AccountAnalyticDistributionModels represents array of account.analytic.distribution.model model.
type AccountAnalyticDistributionModels []AccountAnalyticDistributionModel

// AccountAnalyticDistributionModelModel is the odoo model name.
const AccountAnalyticDistributionModelModel = "account.analytic.distribution.model"

// Many2One convert AccountAnalyticDistributionModel to *Many2One.
func (aadm *AccountAnalyticDistributionModel) Many2One() *Many2One {
	return NewMany2One(aadm.Id.Get(), "")
}

// CreateAccountAnalyticDistributionModel creates a new account.analytic.distribution.model model and returns its id.
func (c *Client) CreateAccountAnalyticDistributionModel(aadm *AccountAnalyticDistributionModel) (int64, error) {
	return c.Create(AccountAnalyticDistributionModelModel, aadm)
}

// UpdateAccountAnalyticDistributionModel updates an existing account.analytic.distribution.model record.
func (c *Client) UpdateAccountAnalyticDistributionModel(aadm *AccountAnalyticDistributionModel) error {
	return c.UpdateAccountAnalyticDistributionModels([]int64{aadm.Id.Get()}, aadm)
}

// UpdateAccountAnalyticDistributionModels updates existing account.analytic.distribution.model records.
// All records (represented by ids) will be updated by aadm values.
func (c *Client) UpdateAccountAnalyticDistributionModels(ids []int64, aadm *AccountAnalyticDistributionModel) error {
	return c.Update(AccountAnalyticDistributionModelModel, ids, aadm)
}

// DeleteAccountAnalyticDistributionModel deletes an existing account.analytic.distribution.model record.
func (c *Client) DeleteAccountAnalyticDistributionModel(id int64) error {
	return c.DeleteAccountAnalyticDistributionModels([]int64{id})
}

// DeleteAccountAnalyticDistributionModels deletes existing account.analytic.distribution.model records.
func (c *Client) DeleteAccountAnalyticDistributionModels(ids []int64) error {
	return c.Delete(AccountAnalyticDistributionModelModel, ids)
}

// GetAccountAnalyticDistributionModel gets account.analytic.distribution.model existing record.
func (c *Client) GetAccountAnalyticDistributionModel(id int64) (*AccountAnalyticDistributionModel, error) {
	aadms, err := c.GetAccountAnalyticDistributionModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if aadms != nil && len(*aadms) > 0 {
		return &((*aadms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.analytic.distribution.model not found", id)
}

// GetAccountAnalyticDistributionModels gets account.analytic.distribution.model existing records.
func (c *Client) GetAccountAnalyticDistributionModels(ids []int64) (*AccountAnalyticDistributionModels, error) {
	aadms := &AccountAnalyticDistributionModels{}
	if err := c.Read(AccountAnalyticDistributionModelModel, ids, nil, aadms); err != nil {
		return nil, err
	}
	return aadms, nil
}

// FindAccountAnalyticDistributionModel finds account.analytic.distribution.model record by querying it with criteria.
func (c *Client) FindAccountAnalyticDistributionModel(criteria *Criteria) (*AccountAnalyticDistributionModel, error) {
	aadms := &AccountAnalyticDistributionModels{}
	if err := c.SearchRead(AccountAnalyticDistributionModelModel, criteria, NewOptions().Limit(1), aadms); err != nil {
		return nil, err
	}
	if aadms != nil && len(*aadms) > 0 {
		return &((*aadms)[0]), nil
	}
	return nil, fmt.Errorf("account.analytic.distribution.model was not found with criteria %v", criteria)
}

// FindAccountAnalyticDistributionModels finds account.analytic.distribution.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticDistributionModels(criteria *Criteria, options *Options) (*AccountAnalyticDistributionModels, error) {
	aadms := &AccountAnalyticDistributionModels{}
	if err := c.SearchRead(AccountAnalyticDistributionModelModel, criteria, options, aadms); err != nil {
		return nil, err
	}
	return aadms, nil
}

// FindAccountAnalyticDistributionModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAnalyticDistributionModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAnalyticDistributionModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAnalyticDistributionModelId finds record id by querying it with criteria.
func (c *Client) FindAccountAnalyticDistributionModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAnalyticDistributionModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.analytic.distribution.model was not found with criteria %v and options %v", criteria, options)
}
