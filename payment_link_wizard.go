package odoo

import (
	"fmt"
)

// PaymentLinkWizard represents payment.link.wizard model.
type PaymentLinkWizard struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	Amount                   *Float     `xmlrpc:"amount,omptempty"`
	AmountMax                *Float     `xmlrpc:"amount_max,omptempty"`
	AvailableProviderIds     *Relation  `xmlrpc:"available_provider_ids,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omptempty"`
	Description              *String    `xmlrpc:"description,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	HasMultipleProviders     *Bool      `xmlrpc:"has_multiple_providers,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	Link                     *String    `xmlrpc:"link,omptempty"`
	PartnerEmail             *String    `xmlrpc:"partner_email,omptempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omptempty"`
	PaymentProviderSelection *Selection `xmlrpc:"payment_provider_selection,omptempty"`
	ResId                    *Int       `xmlrpc:"res_id,omptempty"`
	ResModel                 *String    `xmlrpc:"res_model,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PaymentLinkWizards represents array of payment.link.wizard model.
type PaymentLinkWizards []PaymentLinkWizard

// PaymentLinkWizardModel is the odoo model name.
const PaymentLinkWizardModel = "payment.link.wizard"

// Many2One convert PaymentLinkWizard to *Many2One.
func (plw *PaymentLinkWizard) Many2One() *Many2One {
	return NewMany2One(plw.Id.Get(), "")
}

// CreatePaymentLinkWizard creates a new payment.link.wizard model and returns its id.
func (c *Client) CreatePaymentLinkWizard(plw *PaymentLinkWizard) (int64, error) {
	return c.Create(PaymentLinkWizardModel, plw)
}

// UpdatePaymentLinkWizard updates an existing payment.link.wizard record.
func (c *Client) UpdatePaymentLinkWizard(plw *PaymentLinkWizard) error {
	return c.UpdatePaymentLinkWizards([]int64{plw.Id.Get()}, plw)
}

// UpdatePaymentLinkWizards updates existing payment.link.wizard records.
// All records (represented by ids) will be updated by plw values.
func (c *Client) UpdatePaymentLinkWizards(ids []int64, plw *PaymentLinkWizard) error {
	return c.Update(PaymentLinkWizardModel, ids, plw)
}

// DeletePaymentLinkWizard deletes an existing payment.link.wizard record.
func (c *Client) DeletePaymentLinkWizard(id int64) error {
	return c.DeletePaymentLinkWizards([]int64{id})
}

// DeletePaymentLinkWizards deletes existing payment.link.wizard records.
func (c *Client) DeletePaymentLinkWizards(ids []int64) error {
	return c.Delete(PaymentLinkWizardModel, ids)
}

// GetPaymentLinkWizard gets payment.link.wizard existing record.
func (c *Client) GetPaymentLinkWizard(id int64) (*PaymentLinkWizard, error) {
	plws, err := c.GetPaymentLinkWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if plws != nil && len(*plws) > 0 {
		return &((*plws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of payment.link.wizard not found", id)
}

// GetPaymentLinkWizards gets payment.link.wizard existing records.
func (c *Client) GetPaymentLinkWizards(ids []int64) (*PaymentLinkWizards, error) {
	plws := &PaymentLinkWizards{}
	if err := c.Read(PaymentLinkWizardModel, ids, nil, plws); err != nil {
		return nil, err
	}
	return plws, nil
}

// FindPaymentLinkWizard finds payment.link.wizard record by querying it with criteria.
func (c *Client) FindPaymentLinkWizard(criteria *Criteria) (*PaymentLinkWizard, error) {
	plws := &PaymentLinkWizards{}
	if err := c.SearchRead(PaymentLinkWizardModel, criteria, NewOptions().Limit(1), plws); err != nil {
		return nil, err
	}
	if plws != nil && len(*plws) > 0 {
		return &((*plws)[0]), nil
	}
	return nil, fmt.Errorf("payment.link.wizard was not found with criteria %v", criteria)
}

// FindPaymentLinkWizards finds payment.link.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentLinkWizards(criteria *Criteria, options *Options) (*PaymentLinkWizards, error) {
	plws := &PaymentLinkWizards{}
	if err := c.SearchRead(PaymentLinkWizardModel, criteria, options, plws); err != nil {
		return nil, err
	}
	return plws, nil
}

// FindPaymentLinkWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentLinkWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PaymentLinkWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPaymentLinkWizardId finds record id by querying it with criteria.
func (c *Client) FindPaymentLinkWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentLinkWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("payment.link.wizard was not found with criteria %v and options %v", criteria, options)
}
