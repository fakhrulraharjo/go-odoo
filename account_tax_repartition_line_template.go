package odoo

import (
	"fmt"
)

// AccountTaxRepartitionLineTemplate represents account.tax.repartition.line.template model.
type AccountTaxRepartitionLineTemplate struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AccountId                *Many2One  `xmlrpc:"account_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	FactorPercent            *Float     `xmlrpc:"factor_percent,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	InvoiceTaxId             *Many2One  `xmlrpc:"invoice_tax_id,omptempty"`
	MinusReportExpressionIds *Relation  `xmlrpc:"minus_report_expression_ids,omptempty"`
	PlusReportExpressionIds  *Relation  `xmlrpc:"plus_report_expression_ids,omptempty"`
	RefundTaxId              *Many2One  `xmlrpc:"refund_tax_id,omptempty"`
	RepartitionType          *Selection `xmlrpc:"repartition_type,omptempty"`
	TagIds                   *Relation  `xmlrpc:"tag_ids,omptempty"`
	UseInTaxClosing          *Bool      `xmlrpc:"use_in_tax_closing,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountTaxRepartitionLineTemplates represents array of account.tax.repartition.line.template model.
type AccountTaxRepartitionLineTemplates []AccountTaxRepartitionLineTemplate

// AccountTaxRepartitionLineTemplateModel is the odoo model name.
const AccountTaxRepartitionLineTemplateModel = "account.tax.repartition.line.template"

// Many2One convert AccountTaxRepartitionLineTemplate to *Many2One.
func (atrlt *AccountTaxRepartitionLineTemplate) Many2One() *Many2One {
	return NewMany2One(atrlt.Id.Get(), "")
}

// CreateAccountTaxRepartitionLineTemplate creates a new account.tax.repartition.line.template model and returns its id.
func (c *Client) CreateAccountTaxRepartitionLineTemplate(atrlt *AccountTaxRepartitionLineTemplate) (int64, error) {
	return c.Create(AccountTaxRepartitionLineTemplateModel, atrlt)
}

// UpdateAccountTaxRepartitionLineTemplate updates an existing account.tax.repartition.line.template record.
func (c *Client) UpdateAccountTaxRepartitionLineTemplate(atrlt *AccountTaxRepartitionLineTemplate) error {
	return c.UpdateAccountTaxRepartitionLineTemplates([]int64{atrlt.Id.Get()}, atrlt)
}

// UpdateAccountTaxRepartitionLineTemplates updates existing account.tax.repartition.line.template records.
// All records (represented by ids) will be updated by atrlt values.
func (c *Client) UpdateAccountTaxRepartitionLineTemplates(ids []int64, atrlt *AccountTaxRepartitionLineTemplate) error {
	return c.Update(AccountTaxRepartitionLineTemplateModel, ids, atrlt)
}

// DeleteAccountTaxRepartitionLineTemplate deletes an existing account.tax.repartition.line.template record.
func (c *Client) DeleteAccountTaxRepartitionLineTemplate(id int64) error {
	return c.DeleteAccountTaxRepartitionLineTemplates([]int64{id})
}

// DeleteAccountTaxRepartitionLineTemplates deletes existing account.tax.repartition.line.template records.
func (c *Client) DeleteAccountTaxRepartitionLineTemplates(ids []int64) error {
	return c.Delete(AccountTaxRepartitionLineTemplateModel, ids)
}

// GetAccountTaxRepartitionLineTemplate gets account.tax.repartition.line.template existing record.
func (c *Client) GetAccountTaxRepartitionLineTemplate(id int64) (*AccountTaxRepartitionLineTemplate, error) {
	atrlts, err := c.GetAccountTaxRepartitionLineTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if atrlts != nil && len(*atrlts) > 0 {
		return &((*atrlts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.tax.repartition.line.template not found", id)
}

// GetAccountTaxRepartitionLineTemplates gets account.tax.repartition.line.template existing records.
func (c *Client) GetAccountTaxRepartitionLineTemplates(ids []int64) (*AccountTaxRepartitionLineTemplates, error) {
	atrlts := &AccountTaxRepartitionLineTemplates{}
	if err := c.Read(AccountTaxRepartitionLineTemplateModel, ids, nil, atrlts); err != nil {
		return nil, err
	}
	return atrlts, nil
}

// FindAccountTaxRepartitionLineTemplate finds account.tax.repartition.line.template record by querying it with criteria.
func (c *Client) FindAccountTaxRepartitionLineTemplate(criteria *Criteria) (*AccountTaxRepartitionLineTemplate, error) {
	atrlts := &AccountTaxRepartitionLineTemplates{}
	if err := c.SearchRead(AccountTaxRepartitionLineTemplateModel, criteria, NewOptions().Limit(1), atrlts); err != nil {
		return nil, err
	}
	if atrlts != nil && len(*atrlts) > 0 {
		return &((*atrlts)[0]), nil
	}
	return nil, fmt.Errorf("account.tax.repartition.line.template was not found with criteria %v", criteria)
}

// FindAccountTaxRepartitionLineTemplates finds account.tax.repartition.line.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxRepartitionLineTemplates(criteria *Criteria, options *Options) (*AccountTaxRepartitionLineTemplates, error) {
	atrlts := &AccountTaxRepartitionLineTemplates{}
	if err := c.SearchRead(AccountTaxRepartitionLineTemplateModel, criteria, options, atrlts); err != nil {
		return nil, err
	}
	return atrlts, nil
}

// FindAccountTaxRepartitionLineTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountTaxRepartitionLineTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountTaxRepartitionLineTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountTaxRepartitionLineTemplateId finds record id by querying it with criteria.
func (c *Client) FindAccountTaxRepartitionLineTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountTaxRepartitionLineTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.tax.repartition.line.template was not found with criteria %v and options %v", criteria, options)
}
