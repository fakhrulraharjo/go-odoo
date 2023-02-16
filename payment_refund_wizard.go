package odoo

import (
	"fmt"
)

// PaymentRefundWizard represents payment.refund.wizard model.
type PaymentRefundWizard struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AmountAvailableForRefund *Float     `xmlrpc:"amount_available_for_refund,omptempty"`
	AmountToRefund           *Float     `xmlrpc:"amount_to_refund,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	HasPendingRefund         *Bool      `xmlrpc:"has_pending_refund,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	PaymentAmount            *Float     `xmlrpc:"payment_amount,omptempty"`
	PaymentId                *Many2One  `xmlrpc:"payment_id,omptempty"`
	RefundedAmount           *Float     `xmlrpc:"refunded_amount,omptempty"`
	SupportRefund            *Selection `xmlrpc:"support_refund,omptempty"`
	TransactionId            *Many2One  `xmlrpc:"transaction_id,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PaymentRefundWizards represents array of payment.refund.wizard model.
type PaymentRefundWizards []PaymentRefundWizard

// PaymentRefundWizardModel is the odoo model name.
const PaymentRefundWizardModel = "payment.refund.wizard"

// Many2One convert PaymentRefundWizard to *Many2One.
func (prw *PaymentRefundWizard) Many2One() *Many2One {
	return NewMany2One(prw.Id.Get(), "")
}

// CreatePaymentRefundWizard creates a new payment.refund.wizard model and returns its id.
func (c *Client) CreatePaymentRefundWizard(prw *PaymentRefundWizard) (int64, error) {
	return c.Create(PaymentRefundWizardModel, prw)
}

// UpdatePaymentRefundWizard updates an existing payment.refund.wizard record.
func (c *Client) UpdatePaymentRefundWizard(prw *PaymentRefundWizard) error {
	return c.UpdatePaymentRefundWizards([]int64{prw.Id.Get()}, prw)
}

// UpdatePaymentRefundWizards updates existing payment.refund.wizard records.
// All records (represented by ids) will be updated by prw values.
func (c *Client) UpdatePaymentRefundWizards(ids []int64, prw *PaymentRefundWizard) error {
	return c.Update(PaymentRefundWizardModel, ids, prw)
}

// DeletePaymentRefundWizard deletes an existing payment.refund.wizard record.
func (c *Client) DeletePaymentRefundWizard(id int64) error {
	return c.DeletePaymentRefundWizards([]int64{id})
}

// DeletePaymentRefundWizards deletes existing payment.refund.wizard records.
func (c *Client) DeletePaymentRefundWizards(ids []int64) error {
	return c.Delete(PaymentRefundWizardModel, ids)
}

// GetPaymentRefundWizard gets payment.refund.wizard existing record.
func (c *Client) GetPaymentRefundWizard(id int64) (*PaymentRefundWizard, error) {
	prws, err := c.GetPaymentRefundWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if prws != nil && len(*prws) > 0 {
		return &((*prws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of payment.refund.wizard not found", id)
}

// GetPaymentRefundWizards gets payment.refund.wizard existing records.
func (c *Client) GetPaymentRefundWizards(ids []int64) (*PaymentRefundWizards, error) {
	prws := &PaymentRefundWizards{}
	if err := c.Read(PaymentRefundWizardModel, ids, nil, prws); err != nil {
		return nil, err
	}
	return prws, nil
}

// FindPaymentRefundWizard finds payment.refund.wizard record by querying it with criteria.
func (c *Client) FindPaymentRefundWizard(criteria *Criteria) (*PaymentRefundWizard, error) {
	prws := &PaymentRefundWizards{}
	if err := c.SearchRead(PaymentRefundWizardModel, criteria, NewOptions().Limit(1), prws); err != nil {
		return nil, err
	}
	if prws != nil && len(*prws) > 0 {
		return &((*prws)[0]), nil
	}
	return nil, fmt.Errorf("payment.refund.wizard was not found with criteria %v", criteria)
}

// FindPaymentRefundWizards finds payment.refund.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentRefundWizards(criteria *Criteria, options *Options) (*PaymentRefundWizards, error) {
	prws := &PaymentRefundWizards{}
	if err := c.SearchRead(PaymentRefundWizardModel, criteria, options, prws); err != nil {
		return nil, err
	}
	return prws, nil
}

// FindPaymentRefundWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentRefundWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PaymentRefundWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPaymentRefundWizardId finds record id by querying it with criteria.
func (c *Client) FindPaymentRefundWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentRefundWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("payment.refund.wizard was not found with criteria %v and options %v", criteria, options)
}
