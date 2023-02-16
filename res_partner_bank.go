package odoo

import (
	"fmt"
)

// ResPartnerBank represents res.partner.bank model.
type ResPartnerBank struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	AccHolderName               *String    `xmlrpc:"acc_holder_name,omptempty"`
	AccNumber                   *String    `xmlrpc:"acc_number,omptempty"`
	AccType                     *Selection `xmlrpc:"acc_type,omptempty"`
	Active                      *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AllowOutPayment             *Bool      `xmlrpc:"allow_out_payment,omptempty"`
	BankBic                     *String    `xmlrpc:"bank_bic,omptempty"`
	BankId                      *Many2One  `xmlrpc:"bank_id,omptempty"`
	BankName                    *String    `xmlrpc:"bank_name,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	JournalId                   *Relation  `xmlrpc:"journal_id,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omptempty"`
	SanitizedAccNumber          *String    `xmlrpc:"sanitized_acc_number,omptempty"`
	Sequence                    *Int       `xmlrpc:"sequence,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ResPartnerBanks represents array of res.partner.bank model.
type ResPartnerBanks []ResPartnerBank

// ResPartnerBankModel is the odoo model name.
const ResPartnerBankModel = "res.partner.bank"

// Many2One convert ResPartnerBank to *Many2One.
func (rpb *ResPartnerBank) Many2One() *Many2One {
	return NewMany2One(rpb.Id.Get(), "")
}

// CreateResPartnerBank creates a new res.partner.bank model and returns its id.
func (c *Client) CreateResPartnerBank(rpb *ResPartnerBank) (int64, error) {
	return c.Create(ResPartnerBankModel, rpb)
}

// UpdateResPartnerBank updates an existing res.partner.bank record.
func (c *Client) UpdateResPartnerBank(rpb *ResPartnerBank) error {
	return c.UpdateResPartnerBanks([]int64{rpb.Id.Get()}, rpb)
}

// UpdateResPartnerBanks updates existing res.partner.bank records.
// All records (represented by ids) will be updated by rpb values.
func (c *Client) UpdateResPartnerBanks(ids []int64, rpb *ResPartnerBank) error {
	return c.Update(ResPartnerBankModel, ids, rpb)
}

// DeleteResPartnerBank deletes an existing res.partner.bank record.
func (c *Client) DeleteResPartnerBank(id int64) error {
	return c.DeleteResPartnerBanks([]int64{id})
}

// DeleteResPartnerBanks deletes existing res.partner.bank records.
func (c *Client) DeleteResPartnerBanks(ids []int64) error {
	return c.Delete(ResPartnerBankModel, ids)
}

// GetResPartnerBank gets res.partner.bank existing record.
func (c *Client) GetResPartnerBank(id int64) (*ResPartnerBank, error) {
	rpbs, err := c.GetResPartnerBanks([]int64{id})
	if err != nil {
		return nil, err
	}
	if rpbs != nil && len(*rpbs) > 0 {
		return &((*rpbs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner.bank not found", id)
}

// GetResPartnerBanks gets res.partner.bank existing records.
func (c *Client) GetResPartnerBanks(ids []int64) (*ResPartnerBanks, error) {
	rpbs := &ResPartnerBanks{}
	if err := c.Read(ResPartnerBankModel, ids, nil, rpbs); err != nil {
		return nil, err
	}
	return rpbs, nil
}

// FindResPartnerBank finds res.partner.bank record by querying it with criteria.
func (c *Client) FindResPartnerBank(criteria *Criteria) (*ResPartnerBank, error) {
	rpbs := &ResPartnerBanks{}
	if err := c.SearchRead(ResPartnerBankModel, criteria, NewOptions().Limit(1), rpbs); err != nil {
		return nil, err
	}
	if rpbs != nil && len(*rpbs) > 0 {
		return &((*rpbs)[0]), nil
	}
	return nil, fmt.Errorf("res.partner.bank was not found with criteria %v", criteria)
}

// FindResPartnerBanks finds res.partner.bank records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerBanks(criteria *Criteria, options *Options) (*ResPartnerBanks, error) {
	rpbs := &ResPartnerBanks{}
	if err := c.SearchRead(ResPartnerBankModel, criteria, options, rpbs); err != nil {
		return nil, err
	}
	return rpbs, nil
}

// FindResPartnerBankIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerBankIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerBankModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerBankId finds record id by querying it with criteria.
func (c *Client) FindResPartnerBankId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerBankModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.partner.bank was not found with criteria %v and options %v", criteria, options)
}
