package odoo

import (
	"fmt"
)

// PaymentTransaction represents payment.transaction model.
type PaymentTransaction struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	Amount              *Float     `xmlrpc:"amount,omptempty"`
	CallbackHash        *String    `xmlrpc:"callback_hash,omptempty"`
	CallbackIsDone      *Bool      `xmlrpc:"callback_is_done,omptempty"`
	CallbackMethod      *String    `xmlrpc:"callback_method,omptempty"`
	CallbackModelId     *Many2One  `xmlrpc:"callback_model_id,omptempty"`
	CallbackResId       *Int       `xmlrpc:"callback_res_id,omptempty"`
	ChildTransactionIds *Relation  `xmlrpc:"child_transaction_ids,omptempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId          *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	Fees                *Float     `xmlrpc:"fees,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	InvoiceIds          *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoicesCount       *Int       `xmlrpc:"invoices_count,omptempty"`
	IsPostProcessed     *Bool      `xmlrpc:"is_post_processed,omptempty"`
	LandingRoute        *String    `xmlrpc:"landing_route,omptempty"`
	LastStateChange     *Time      `xmlrpc:"last_state_change,omptempty"`
	Operation           *Selection `xmlrpc:"operation,omptempty"`
	PartnerAddress      *String    `xmlrpc:"partner_address,omptempty"`
	PartnerCity         *String    `xmlrpc:"partner_city,omptempty"`
	PartnerCountryId    *Many2One  `xmlrpc:"partner_country_id,omptempty"`
	PartnerEmail        *String    `xmlrpc:"partner_email,omptempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerLang         *Selection `xmlrpc:"partner_lang,omptempty"`
	PartnerName         *String    `xmlrpc:"partner_name,omptempty"`
	PartnerPhone        *String    `xmlrpc:"partner_phone,omptempty"`
	PartnerStateId      *Many2One  `xmlrpc:"partner_state_id,omptempty"`
	PartnerZip          *String    `xmlrpc:"partner_zip,omptempty"`
	PaymentId           *Many2One  `xmlrpc:"payment_id,omptempty"`
	ProviderCode        *Selection `xmlrpc:"provider_code,omptempty"`
	ProviderId          *Many2One  `xmlrpc:"provider_id,omptempty"`
	ProviderReference   *String    `xmlrpc:"provider_reference,omptempty"`
	Reference           *String    `xmlrpc:"reference,omptempty"`
	RefundsCount        *Int       `xmlrpc:"refunds_count,omptempty"`
	SourceTransactionId *Many2One  `xmlrpc:"source_transaction_id,omptempty"`
	State               *Selection `xmlrpc:"state,omptempty"`
	StateMessage        *String    `xmlrpc:"state_message,omptempty"`
	TokenId             *Many2One  `xmlrpc:"token_id,omptempty"`
	Tokenize            *Bool      `xmlrpc:"tokenize,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PaymentTransactions represents array of payment.transaction model.
type PaymentTransactions []PaymentTransaction

// PaymentTransactionModel is the odoo model name.
const PaymentTransactionModel = "payment.transaction"

// Many2One convert PaymentTransaction to *Many2One.
func (pt *PaymentTransaction) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreatePaymentTransaction creates a new payment.transaction model and returns its id.
func (c *Client) CreatePaymentTransaction(pt *PaymentTransaction) (int64, error) {
	return c.Create(PaymentTransactionModel, pt)
}

// UpdatePaymentTransaction updates an existing payment.transaction record.
func (c *Client) UpdatePaymentTransaction(pt *PaymentTransaction) error {
	return c.UpdatePaymentTransactions([]int64{pt.Id.Get()}, pt)
}

// UpdatePaymentTransactions updates existing payment.transaction records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdatePaymentTransactions(ids []int64, pt *PaymentTransaction) error {
	return c.Update(PaymentTransactionModel, ids, pt)
}

// DeletePaymentTransaction deletes an existing payment.transaction record.
func (c *Client) DeletePaymentTransaction(id int64) error {
	return c.DeletePaymentTransactions([]int64{id})
}

// DeletePaymentTransactions deletes existing payment.transaction records.
func (c *Client) DeletePaymentTransactions(ids []int64) error {
	return c.Delete(PaymentTransactionModel, ids)
}

// GetPaymentTransaction gets payment.transaction existing record.
func (c *Client) GetPaymentTransaction(id int64) (*PaymentTransaction, error) {
	pts, err := c.GetPaymentTransactions([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of payment.transaction not found", id)
}

// GetPaymentTransactions gets payment.transaction existing records.
func (c *Client) GetPaymentTransactions(ids []int64) (*PaymentTransactions, error) {
	pts := &PaymentTransactions{}
	if err := c.Read(PaymentTransactionModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindPaymentTransaction finds payment.transaction record by querying it with criteria.
func (c *Client) FindPaymentTransaction(criteria *Criteria) (*PaymentTransaction, error) {
	pts := &PaymentTransactions{}
	if err := c.SearchRead(PaymentTransactionModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("payment.transaction was not found with criteria %v", criteria)
}

// FindPaymentTransactions finds payment.transaction records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentTransactions(criteria *Criteria, options *Options) (*PaymentTransactions, error) {
	pts := &PaymentTransactions{}
	if err := c.SearchRead(PaymentTransactionModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindPaymentTransactionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentTransactionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PaymentTransactionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPaymentTransactionId finds record id by querying it with criteria.
func (c *Client) FindPaymentTransactionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentTransactionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("payment.transaction was not found with criteria %v and options %v", criteria, options)
}
