package odoo

import (
	"fmt"
)

// PaymentProviderOnboardingWizard represents payment.provider.onboarding.wizard model.
type PaymentProviderOnboardingWizard struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	DataFetched         *Bool      `xmlrpc:"_data_fetched,omptempty"`
	AccNumber           *String    `xmlrpc:"acc_number,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	JournalName         *String    `xmlrpc:"journal_name,omptempty"`
	ManualName          *String    `xmlrpc:"manual_name,omptempty"`
	ManualPostMsg       *String    `xmlrpc:"manual_post_msg,omptempty"`
	PaymentMethod       *Selection `xmlrpc:"payment_method,omptempty"`
	PaypalEmailAccount  *String    `xmlrpc:"paypal_email_account,omptempty"`
	PaypalPdtToken      *String    `xmlrpc:"paypal_pdt_token,omptempty"`
	PaypalSellerAccount *String    `xmlrpc:"paypal_seller_account,omptempty"`
	PaypalUserType      *Selection `xmlrpc:"paypal_user_type,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PaymentProviderOnboardingWizards represents array of payment.provider.onboarding.wizard model.
type PaymentProviderOnboardingWizards []PaymentProviderOnboardingWizard

// PaymentProviderOnboardingWizardModel is the odoo model name.
const PaymentProviderOnboardingWizardModel = "payment.provider.onboarding.wizard"

// Many2One convert PaymentProviderOnboardingWizard to *Many2One.
func (ppow *PaymentProviderOnboardingWizard) Many2One() *Many2One {
	return NewMany2One(ppow.Id.Get(), "")
}

// CreatePaymentProviderOnboardingWizard creates a new payment.provider.onboarding.wizard model and returns its id.
func (c *Client) CreatePaymentProviderOnboardingWizard(ppow *PaymentProviderOnboardingWizard) (int64, error) {
	return c.Create(PaymentProviderOnboardingWizardModel, ppow)
}

// UpdatePaymentProviderOnboardingWizard updates an existing payment.provider.onboarding.wizard record.
func (c *Client) UpdatePaymentProviderOnboardingWizard(ppow *PaymentProviderOnboardingWizard) error {
	return c.UpdatePaymentProviderOnboardingWizards([]int64{ppow.Id.Get()}, ppow)
}

// UpdatePaymentProviderOnboardingWizards updates existing payment.provider.onboarding.wizard records.
// All records (represented by ids) will be updated by ppow values.
func (c *Client) UpdatePaymentProviderOnboardingWizards(ids []int64, ppow *PaymentProviderOnboardingWizard) error {
	return c.Update(PaymentProviderOnboardingWizardModel, ids, ppow)
}

// DeletePaymentProviderOnboardingWizard deletes an existing payment.provider.onboarding.wizard record.
func (c *Client) DeletePaymentProviderOnboardingWizard(id int64) error {
	return c.DeletePaymentProviderOnboardingWizards([]int64{id})
}

// DeletePaymentProviderOnboardingWizards deletes existing payment.provider.onboarding.wizard records.
func (c *Client) DeletePaymentProviderOnboardingWizards(ids []int64) error {
	return c.Delete(PaymentProviderOnboardingWizardModel, ids)
}

// GetPaymentProviderOnboardingWizard gets payment.provider.onboarding.wizard existing record.
func (c *Client) GetPaymentProviderOnboardingWizard(id int64) (*PaymentProviderOnboardingWizard, error) {
	ppows, err := c.GetPaymentProviderOnboardingWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if ppows != nil && len(*ppows) > 0 {
		return &((*ppows)[0]), nil
	}
	return nil, fmt.Errorf("id %v of payment.provider.onboarding.wizard not found", id)
}

// GetPaymentProviderOnboardingWizards gets payment.provider.onboarding.wizard existing records.
func (c *Client) GetPaymentProviderOnboardingWizards(ids []int64) (*PaymentProviderOnboardingWizards, error) {
	ppows := &PaymentProviderOnboardingWizards{}
	if err := c.Read(PaymentProviderOnboardingWizardModel, ids, nil, ppows); err != nil {
		return nil, err
	}
	return ppows, nil
}

// FindPaymentProviderOnboardingWizard finds payment.provider.onboarding.wizard record by querying it with criteria.
func (c *Client) FindPaymentProviderOnboardingWizard(criteria *Criteria) (*PaymentProviderOnboardingWizard, error) {
	ppows := &PaymentProviderOnboardingWizards{}
	if err := c.SearchRead(PaymentProviderOnboardingWizardModel, criteria, NewOptions().Limit(1), ppows); err != nil {
		return nil, err
	}
	if ppows != nil && len(*ppows) > 0 {
		return &((*ppows)[0]), nil
	}
	return nil, fmt.Errorf("payment.provider.onboarding.wizard was not found with criteria %v", criteria)
}

// FindPaymentProviderOnboardingWizards finds payment.provider.onboarding.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentProviderOnboardingWizards(criteria *Criteria, options *Options) (*PaymentProviderOnboardingWizards, error) {
	ppows := &PaymentProviderOnboardingWizards{}
	if err := c.SearchRead(PaymentProviderOnboardingWizardModel, criteria, options, ppows); err != nil {
		return nil, err
	}
	return ppows, nil
}

// FindPaymentProviderOnboardingWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentProviderOnboardingWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PaymentProviderOnboardingWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPaymentProviderOnboardingWizardId finds record id by querying it with criteria.
func (c *Client) FindPaymentProviderOnboardingWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentProviderOnboardingWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("payment.provider.onboarding.wizard was not found with criteria %v and options %v", criteria, options)
}
