package odoo

import (
	"fmt"
)

// SmsResend represents sms.resend model.
type SmsResend struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	CanCancel              *Bool     `xmlrpc:"can_cancel,omptempty"`
	CanResend              *Bool     `xmlrpc:"can_resend,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	HasInsufficientCredit  *Bool     `xmlrpc:"has_insufficient_credit,omptempty"`
	HasUnregisteredAccount *Bool     `xmlrpc:"has_unregistered_account,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	MailMessageId          *Many2One `xmlrpc:"mail_message_id,omptempty"`
	RecipientIds           *Relation `xmlrpc:"recipient_ids,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SmsResends represents array of sms.resend model.
type SmsResends []SmsResend

// SmsResendModel is the odoo model name.
const SmsResendModel = "sms.resend"

// Many2One convert SmsResend to *Many2One.
func (sr *SmsResend) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateSmsResend creates a new sms.resend model and returns its id.
func (c *Client) CreateSmsResend(sr *SmsResend) (int64, error) {
	return c.Create(SmsResendModel, sr)
}

// UpdateSmsResend updates an existing sms.resend record.
func (c *Client) UpdateSmsResend(sr *SmsResend) error {
	return c.UpdateSmsResends([]int64{sr.Id.Get()}, sr)
}

// UpdateSmsResends updates existing sms.resend records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateSmsResends(ids []int64, sr *SmsResend) error {
	return c.Update(SmsResendModel, ids, sr)
}

// DeleteSmsResend deletes an existing sms.resend record.
func (c *Client) DeleteSmsResend(id int64) error {
	return c.DeleteSmsResends([]int64{id})
}

// DeleteSmsResends deletes existing sms.resend records.
func (c *Client) DeleteSmsResends(ids []int64) error {
	return c.Delete(SmsResendModel, ids)
}

// GetSmsResend gets sms.resend existing record.
func (c *Client) GetSmsResend(id int64) (*SmsResend, error) {
	srs, err := c.GetSmsResends([]int64{id})
	if err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sms.resend not found", id)
}

// GetSmsResends gets sms.resend existing records.
func (c *Client) GetSmsResends(ids []int64) (*SmsResends, error) {
	srs := &SmsResends{}
	if err := c.Read(SmsResendModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSmsResend finds sms.resend record by querying it with criteria.
func (c *Client) FindSmsResend(criteria *Criteria) (*SmsResend, error) {
	srs := &SmsResends{}
	if err := c.SearchRead(SmsResendModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	if srs != nil && len(*srs) > 0 {
		return &((*srs)[0]), nil
	}
	return nil, fmt.Errorf("sms.resend was not found with criteria %v", criteria)
}

// FindSmsResends finds sms.resend records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsResends(criteria *Criteria, options *Options) (*SmsResends, error) {
	srs := &SmsResends{}
	if err := c.SearchRead(SmsResendModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindSmsResendIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsResendIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SmsResendModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSmsResendId finds record id by querying it with criteria.
func (c *Client) FindSmsResendId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsResendModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sms.resend was not found with criteria %v and options %v", criteria, options)
}
