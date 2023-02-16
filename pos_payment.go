package odoo

import (
	"fmt"
)

// PosPayment represents pos.payment model.
type PosPayment struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	AccountMoveId   *Many2One `xmlrpc:"account_move_id,omptempty"`
	Amount          *Float    `xmlrpc:"amount,omptempty"`
	CardType        *String   `xmlrpc:"card_type,omptempty"`
	CardholderName  *String   `xmlrpc:"cardholder_name,omptempty"`
	CompanyId       *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	CurrencyId      *Many2One `xmlrpc:"currency_id,omptempty"`
	CurrencyRate    *Float    `xmlrpc:"currency_rate,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	IsChange        *Bool     `xmlrpc:"is_change,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	PartnerId       *Many2One `xmlrpc:"partner_id,omptempty"`
	PaymentDate     *Time     `xmlrpc:"payment_date,omptempty"`
	PaymentMethodId *Many2One `xmlrpc:"payment_method_id,omptempty"`
	PaymentStatus   *String   `xmlrpc:"payment_status,omptempty"`
	PosOrderId      *Many2One `xmlrpc:"pos_order_id,omptempty"`
	SessionId       *Many2One `xmlrpc:"session_id,omptempty"`
	Ticket          *String   `xmlrpc:"ticket,omptempty"`
	TransactionId   *String   `xmlrpc:"transaction_id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosPayments represents array of pos.payment model.
type PosPayments []PosPayment

// PosPaymentModel is the odoo model name.
const PosPaymentModel = "pos.payment"

// Many2One convert PosPayment to *Many2One.
func (pp *PosPayment) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreatePosPayment creates a new pos.payment model and returns its id.
func (c *Client) CreatePosPayment(pp *PosPayment) (int64, error) {
	return c.Create(PosPaymentModel, pp)
}

// UpdatePosPayment updates an existing pos.payment record.
func (c *Client) UpdatePosPayment(pp *PosPayment) error {
	return c.UpdatePosPayments([]int64{pp.Id.Get()}, pp)
}

// UpdatePosPayments updates existing pos.payment records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdatePosPayments(ids []int64, pp *PosPayment) error {
	return c.Update(PosPaymentModel, ids, pp)
}

// DeletePosPayment deletes an existing pos.payment record.
func (c *Client) DeletePosPayment(id int64) error {
	return c.DeletePosPayments([]int64{id})
}

// DeletePosPayments deletes existing pos.payment records.
func (c *Client) DeletePosPayments(ids []int64) error {
	return c.Delete(PosPaymentModel, ids)
}

// GetPosPayment gets pos.payment existing record.
func (c *Client) GetPosPayment(id int64) (*PosPayment, error) {
	pps, err := c.GetPosPayments([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.payment not found", id)
}

// GetPosPayments gets pos.payment existing records.
func (c *Client) GetPosPayments(ids []int64) (*PosPayments, error) {
	pps := &PosPayments{}
	if err := c.Read(PosPaymentModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPosPayment finds pos.payment record by querying it with criteria.
func (c *Client) FindPosPayment(criteria *Criteria) (*PosPayment, error) {
	pps := &PosPayments{}
	if err := c.SearchRead(PosPaymentModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("pos.payment was not found with criteria %v", criteria)
}

// FindPosPayments finds pos.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPayments(criteria *Criteria, options *Options) (*PosPayments, error) {
	pps := &PosPayments{}
	if err := c.SearchRead(PosPaymentModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPosPaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosPaymentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosPaymentId finds record id by querying it with criteria.
func (c *Client) FindPosPaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.payment was not found with criteria %v and options %v", criteria, options)
}
