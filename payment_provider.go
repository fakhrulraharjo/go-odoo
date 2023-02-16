package odoo

import (
	"fmt"
)

// PaymentProvider represents payment.provider model.
type PaymentProvider struct {
	LastUpdate                *Time      `xmlrpc:"__last_update,omptempty"`
	AllowExpressCheckout      *Bool      `xmlrpc:"allow_express_checkout,omptempty"`
	AllowTokenization         *Bool      `xmlrpc:"allow_tokenization,omptempty"`
	AuthMsg                   *String    `xmlrpc:"auth_msg,omptempty"`
	AvailableCountryIds       *Relation  `xmlrpc:"available_country_ids,omptempty"`
	CancelMsg                 *String    `xmlrpc:"cancel_msg,omptempty"`
	CaptureManually           *Bool      `xmlrpc:"capture_manually,omptempty"`
	Code                      *Selection `xmlrpc:"code,omptempty"`
	Color                     *Int       `xmlrpc:"color,omptempty"`
	CompanyId                 *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayAs                 *String    `xmlrpc:"display_as,omptempty"`
	DisplayName               *String    `xmlrpc:"display_name,omptempty"`
	DoneMsg                   *String    `xmlrpc:"done_msg,omptempty"`
	ExpressCheckoutFormViewId *Many2One  `xmlrpc:"express_checkout_form_view_id,omptempty"`
	FeesActive                *Bool      `xmlrpc:"fees_active,omptempty"`
	FeesDomFixed              *Float     `xmlrpc:"fees_dom_fixed,omptempty"`
	FeesDomVar                *Float     `xmlrpc:"fees_dom_var,omptempty"`
	FeesIntFixed              *Float     `xmlrpc:"fees_int_fixed,omptempty"`
	FeesIntVar                *Float     `xmlrpc:"fees_int_var,omptempty"`
	Id                        *Int       `xmlrpc:"id,omptempty"`
	Image128                  *String    `xmlrpc:"image_128,omptempty"`
	InlineFormViewId          *Many2One  `xmlrpc:"inline_form_view_id,omptempty"`
	IsPublished               *Bool      `xmlrpc:"is_published,omptempty"`
	JournalId                 *Many2One  `xmlrpc:"journal_id,omptempty"`
	MainCurrencyId            *Many2One  `xmlrpc:"main_currency_id,omptempty"`
	MaximumAmount             *Float     `xmlrpc:"maximum_amount,omptempty"`
	ModuleId                  *Many2One  `xmlrpc:"module_id,omptempty"`
	ModuleState               *Selection `xmlrpc:"module_state,omptempty"`
	ModuleToBuy               *Bool      `xmlrpc:"module_to_buy,omptempty"`
	Name                      *String    `xmlrpc:"name,omptempty"`
	PaymentIconIds            *Relation  `xmlrpc:"payment_icon_ids,omptempty"`
	PendingMsg                *String    `xmlrpc:"pending_msg,omptempty"`
	PreMsg                    *String    `xmlrpc:"pre_msg,omptempty"`
	RedirectFormViewId        *Many2One  `xmlrpc:"redirect_form_view_id,omptempty"`
	Sequence                  *Int       `xmlrpc:"sequence,omptempty"`
	ShowAllowExpressCheckout  *Bool      `xmlrpc:"show_allow_express_checkout,omptempty"`
	ShowAllowTokenization     *Bool      `xmlrpc:"show_allow_tokenization,omptempty"`
	ShowAuthMsg               *Bool      `xmlrpc:"show_auth_msg,omptempty"`
	ShowCancelMsg             *Bool      `xmlrpc:"show_cancel_msg,omptempty"`
	ShowCredentialsPage       *Bool      `xmlrpc:"show_credentials_page,omptempty"`
	ShowDoneMsg               *Bool      `xmlrpc:"show_done_msg,omptempty"`
	ShowPaymentIconIds        *Bool      `xmlrpc:"show_payment_icon_ids,omptempty"`
	ShowPendingMsg            *Bool      `xmlrpc:"show_pending_msg,omptempty"`
	ShowPreMsg                *Bool      `xmlrpc:"show_pre_msg,omptempty"`
	State                     *Selection `xmlrpc:"state,omptempty"`
	SupportExpressCheckout    *Bool      `xmlrpc:"support_express_checkout,omptempty"`
	SupportFees               *Bool      `xmlrpc:"support_fees,omptempty"`
	SupportManualCapture      *Bool      `xmlrpc:"support_manual_capture,omptempty"`
	SupportRefund             *Selection `xmlrpc:"support_refund,omptempty"`
	SupportTokenization       *Bool      `xmlrpc:"support_tokenization,omptempty"`
	TokenInlineFormViewId     *Many2One  `xmlrpc:"token_inline_form_view_id,omptempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PaymentProviders represents array of payment.provider model.
type PaymentProviders []PaymentProvider

// PaymentProviderModel is the odoo model name.
const PaymentProviderModel = "payment.provider"

// Many2One convert PaymentProvider to *Many2One.
func (pp *PaymentProvider) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreatePaymentProvider creates a new payment.provider model and returns its id.
func (c *Client) CreatePaymentProvider(pp *PaymentProvider) (int64, error) {
	return c.Create(PaymentProviderModel, pp)
}

// UpdatePaymentProvider updates an existing payment.provider record.
func (c *Client) UpdatePaymentProvider(pp *PaymentProvider) error {
	return c.UpdatePaymentProviders([]int64{pp.Id.Get()}, pp)
}

// UpdatePaymentProviders updates existing payment.provider records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdatePaymentProviders(ids []int64, pp *PaymentProvider) error {
	return c.Update(PaymentProviderModel, ids, pp)
}

// DeletePaymentProvider deletes an existing payment.provider record.
func (c *Client) DeletePaymentProvider(id int64) error {
	return c.DeletePaymentProviders([]int64{id})
}

// DeletePaymentProviders deletes existing payment.provider records.
func (c *Client) DeletePaymentProviders(ids []int64) error {
	return c.Delete(PaymentProviderModel, ids)
}

// GetPaymentProvider gets payment.provider existing record.
func (c *Client) GetPaymentProvider(id int64) (*PaymentProvider, error) {
	pps, err := c.GetPaymentProviders([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of payment.provider not found", id)
}

// GetPaymentProviders gets payment.provider existing records.
func (c *Client) GetPaymentProviders(ids []int64) (*PaymentProviders, error) {
	pps := &PaymentProviders{}
	if err := c.Read(PaymentProviderModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPaymentProvider finds payment.provider record by querying it with criteria.
func (c *Client) FindPaymentProvider(criteria *Criteria) (*PaymentProvider, error) {
	pps := &PaymentProviders{}
	if err := c.SearchRead(PaymentProviderModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("payment.provider was not found with criteria %v", criteria)
}

// FindPaymentProviders finds payment.provider records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentProviders(criteria *Criteria, options *Options) (*PaymentProviders, error) {
	pps := &PaymentProviders{}
	if err := c.SearchRead(PaymentProviderModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindPaymentProviderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentProviderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PaymentProviderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPaymentProviderId finds record id by querying it with criteria.
func (c *Client) FindPaymentProviderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentProviderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("payment.provider was not found with criteria %v and options %v", criteria, options)
}
