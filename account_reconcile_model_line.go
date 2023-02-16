package odoo

import (
	"fmt"
)

// AccountReconcileModelLine represents account.reconcile.model.line model.
type AccountReconcileModelLine struct {
	LastUpdate            *Time       `xmlrpc:"__last_update,omptempty"`
	AccountId             *Many2One   `xmlrpc:"account_id,omptempty"`
	AllowPaymentTolerance *Bool       `xmlrpc:"allow_payment_tolerance,omptempty"`
	Amount                *Float      `xmlrpc:"amount,omptempty"`
	AmountString          *String     `xmlrpc:"amount_string,omptempty"`
	AmountType            *Selection  `xmlrpc:"amount_type,omptempty"`
	AnalyticDistribution  interface{} `xmlrpc:"analytic_distribution,omptempty"`
	AnalyticPrecision     *Int        `xmlrpc:"analytic_precision,omptempty"`
	CompanyId             *Many2One   `xmlrpc:"company_id,omptempty"`
	CreateDate            *Time       `xmlrpc:"create_date,omptempty"`
	CreateUid             *Many2One   `xmlrpc:"create_uid,omptempty"`
	DisplayName           *String     `xmlrpc:"display_name,omptempty"`
	ForceTaxIncluded      *Bool       `xmlrpc:"force_tax_included,omptempty"`
	Id                    *Int        `xmlrpc:"id,omptempty"`
	JournalId             *Many2One   `xmlrpc:"journal_id,omptempty"`
	Label                 *String     `xmlrpc:"label,omptempty"`
	ModelId               *Many2One   `xmlrpc:"model_id,omptempty"`
	PaymentToleranceParam *Float      `xmlrpc:"payment_tolerance_param,omptempty"`
	RuleType              *Selection  `xmlrpc:"rule_type,omptempty"`
	Sequence              *Int        `xmlrpc:"sequence,omptempty"`
	ShowForceTaxIncluded  *Bool       `xmlrpc:"show_force_tax_included,omptempty"`
	TaxIds                *Relation   `xmlrpc:"tax_ids,omptempty"`
	WriteDate             *Time       `xmlrpc:"write_date,omptempty"`
	WriteUid              *Many2One   `xmlrpc:"write_uid,omptempty"`
}

// AccountReconcileModelLines represents array of account.reconcile.model.line model.
type AccountReconcileModelLines []AccountReconcileModelLine

// AccountReconcileModelLineModel is the odoo model name.
const AccountReconcileModelLineModel = "account.reconcile.model.line"

// Many2One convert AccountReconcileModelLine to *Many2One.
func (arml *AccountReconcileModelLine) Many2One() *Many2One {
	return NewMany2One(arml.Id.Get(), "")
}

// CreateAccountReconcileModelLine creates a new account.reconcile.model.line model and returns its id.
func (c *Client) CreateAccountReconcileModelLine(arml *AccountReconcileModelLine) (int64, error) {
	return c.Create(AccountReconcileModelLineModel, arml)
}

// UpdateAccountReconcileModelLine updates an existing account.reconcile.model.line record.
func (c *Client) UpdateAccountReconcileModelLine(arml *AccountReconcileModelLine) error {
	return c.UpdateAccountReconcileModelLines([]int64{arml.Id.Get()}, arml)
}

// UpdateAccountReconcileModelLines updates existing account.reconcile.model.line records.
// All records (represented by ids) will be updated by arml values.
func (c *Client) UpdateAccountReconcileModelLines(ids []int64, arml *AccountReconcileModelLine) error {
	return c.Update(AccountReconcileModelLineModel, ids, arml)
}

// DeleteAccountReconcileModelLine deletes an existing account.reconcile.model.line record.
func (c *Client) DeleteAccountReconcileModelLine(id int64) error {
	return c.DeleteAccountReconcileModelLines([]int64{id})
}

// DeleteAccountReconcileModelLines deletes existing account.reconcile.model.line records.
func (c *Client) DeleteAccountReconcileModelLines(ids []int64) error {
	return c.Delete(AccountReconcileModelLineModel, ids)
}

// GetAccountReconcileModelLine gets account.reconcile.model.line existing record.
func (c *Client) GetAccountReconcileModelLine(id int64) (*AccountReconcileModelLine, error) {
	armls, err := c.GetAccountReconcileModelLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if armls != nil && len(*armls) > 0 {
		return &((*armls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.reconcile.model.line not found", id)
}

// GetAccountReconcileModelLines gets account.reconcile.model.line existing records.
func (c *Client) GetAccountReconcileModelLines(ids []int64) (*AccountReconcileModelLines, error) {
	armls := &AccountReconcileModelLines{}
	if err := c.Read(AccountReconcileModelLineModel, ids, nil, armls); err != nil {
		return nil, err
	}
	return armls, nil
}

// FindAccountReconcileModelLine finds account.reconcile.model.line record by querying it with criteria.
func (c *Client) FindAccountReconcileModelLine(criteria *Criteria) (*AccountReconcileModelLine, error) {
	armls := &AccountReconcileModelLines{}
	if err := c.SearchRead(AccountReconcileModelLineModel, criteria, NewOptions().Limit(1), armls); err != nil {
		return nil, err
	}
	if armls != nil && len(*armls) > 0 {
		return &((*armls)[0]), nil
	}
	return nil, fmt.Errorf("account.reconcile.model.line was not found with criteria %v", criteria)
}

// FindAccountReconcileModelLines finds account.reconcile.model.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelLines(criteria *Criteria, options *Options) (*AccountReconcileModelLines, error) {
	armls := &AccountReconcileModelLines{}
	if err := c.SearchRead(AccountReconcileModelLineModel, criteria, options, armls); err != nil {
		return nil, err
	}
	return armls, nil
}

// FindAccountReconcileModelLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReconcileModelLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReconcileModelLineId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.reconcile.model.line was not found with criteria %v and options %v", criteria, options)
}
