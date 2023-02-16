package odoo

import (
	"fmt"
)

// AccountReconcileModel represents account.reconcile.model model.
type AccountReconcileModel struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omptempty"`
	Active                     *Bool      `xmlrpc:"active,omptempty"`
	AllowPaymentTolerance      *Bool      `xmlrpc:"allow_payment_tolerance,omptempty"`
	AutoReconcile              *Bool      `xmlrpc:"auto_reconcile,omptempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	DecimalSeparator           *String    `xmlrpc:"decimal_separator,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	HasMessage                 *Bool      `xmlrpc:"has_message,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	LineIds                    *Relation  `xmlrpc:"line_ids,omptempty"`
	MatchAmount                *Selection `xmlrpc:"match_amount,omptempty"`
	MatchAmountMax             *Float     `xmlrpc:"match_amount_max,omptempty"`
	MatchAmountMin             *Float     `xmlrpc:"match_amount_min,omptempty"`
	MatchJournalIds            *Relation  `xmlrpc:"match_journal_ids,omptempty"`
	MatchLabel                 *Selection `xmlrpc:"match_label,omptempty"`
	MatchLabelParam            *String    `xmlrpc:"match_label_param,omptempty"`
	MatchNature                *Selection `xmlrpc:"match_nature,omptempty"`
	MatchNote                  *Selection `xmlrpc:"match_note,omptempty"`
	MatchNoteParam             *String    `xmlrpc:"match_note_param,omptempty"`
	MatchPartner               *Bool      `xmlrpc:"match_partner,omptempty"`
	MatchPartnerCategoryIds    *Relation  `xmlrpc:"match_partner_category_ids,omptempty"`
	MatchPartnerIds            *Relation  `xmlrpc:"match_partner_ids,omptempty"`
	MatchSameCurrency          *Bool      `xmlrpc:"match_same_currency,omptempty"`
	MatchTextLocationLabel     *Bool      `xmlrpc:"match_text_location_label,omptempty"`
	MatchTextLocationNote      *Bool      `xmlrpc:"match_text_location_note,omptempty"`
	MatchTextLocationReference *Bool      `xmlrpc:"match_text_location_reference,omptempty"`
	MatchTransactionType       *Selection `xmlrpc:"match_transaction_type,omptempty"`
	MatchTransactionTypeParam  *String    `xmlrpc:"match_transaction_type_param,omptempty"`
	MatchingOrder              *Selection `xmlrpc:"matching_order,omptempty"`
	MessageAttachmentCount     *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds         *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError            *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter     *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError         *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                 *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower          *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId    *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction          *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter   *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds          *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Name                       *String    `xmlrpc:"name,omptempty"`
	NumberEntries              *Int       `xmlrpc:"number_entries,omptempty"`
	PartnerMappingLineIds      *Relation  `xmlrpc:"partner_mapping_line_ids,omptempty"`
	PastMonthsLimit            *Int       `xmlrpc:"past_months_limit,omptempty"`
	PaymentToleranceParam      *Float     `xmlrpc:"payment_tolerance_param,omptempty"`
	PaymentToleranceType       *Selection `xmlrpc:"payment_tolerance_type,omptempty"`
	RuleType                   *Selection `xmlrpc:"rule_type,omptempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omptempty"`
	ShowDecimalSeparator       *Bool      `xmlrpc:"show_decimal_separator,omptempty"`
	ToCheck                    *Bool      `xmlrpc:"to_check,omptempty"`
	WebsiteMessageIds          *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountReconcileModels represents array of account.reconcile.model model.
type AccountReconcileModels []AccountReconcileModel

// AccountReconcileModelModel is the odoo model name.
const AccountReconcileModelModel = "account.reconcile.model"

// Many2One convert AccountReconcileModel to *Many2One.
func (arm *AccountReconcileModel) Many2One() *Many2One {
	return NewMany2One(arm.Id.Get(), "")
}

// CreateAccountReconcileModel creates a new account.reconcile.model model and returns its id.
func (c *Client) CreateAccountReconcileModel(arm *AccountReconcileModel) (int64, error) {
	return c.Create(AccountReconcileModelModel, arm)
}

// UpdateAccountReconcileModel updates an existing account.reconcile.model record.
func (c *Client) UpdateAccountReconcileModel(arm *AccountReconcileModel) error {
	return c.UpdateAccountReconcileModels([]int64{arm.Id.Get()}, arm)
}

// UpdateAccountReconcileModels updates existing account.reconcile.model records.
// All records (represented by ids) will be updated by arm values.
func (c *Client) UpdateAccountReconcileModels(ids []int64, arm *AccountReconcileModel) error {
	return c.Update(AccountReconcileModelModel, ids, arm)
}

// DeleteAccountReconcileModel deletes an existing account.reconcile.model record.
func (c *Client) DeleteAccountReconcileModel(id int64) error {
	return c.DeleteAccountReconcileModels([]int64{id})
}

// DeleteAccountReconcileModels deletes existing account.reconcile.model records.
func (c *Client) DeleteAccountReconcileModels(ids []int64) error {
	return c.Delete(AccountReconcileModelModel, ids)
}

// GetAccountReconcileModel gets account.reconcile.model existing record.
func (c *Client) GetAccountReconcileModel(id int64) (*AccountReconcileModel, error) {
	arms, err := c.GetAccountReconcileModels([]int64{id})
	if err != nil {
		return nil, err
	}
	if arms != nil && len(*arms) > 0 {
		return &((*arms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.reconcile.model not found", id)
}

// GetAccountReconcileModels gets account.reconcile.model existing records.
func (c *Client) GetAccountReconcileModels(ids []int64) (*AccountReconcileModels, error) {
	arms := &AccountReconcileModels{}
	if err := c.Read(AccountReconcileModelModel, ids, nil, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReconcileModel finds account.reconcile.model record by querying it with criteria.
func (c *Client) FindAccountReconcileModel(criteria *Criteria) (*AccountReconcileModel, error) {
	arms := &AccountReconcileModels{}
	if err := c.SearchRead(AccountReconcileModelModel, criteria, NewOptions().Limit(1), arms); err != nil {
		return nil, err
	}
	if arms != nil && len(*arms) > 0 {
		return &((*arms)[0]), nil
	}
	return nil, fmt.Errorf("account.reconcile.model was not found with criteria %v", criteria)
}

// FindAccountReconcileModels finds account.reconcile.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModels(criteria *Criteria, options *Options) (*AccountReconcileModels, error) {
	arms := &AccountReconcileModels{}
	if err := c.SearchRead(AccountReconcileModelModel, criteria, options, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReconcileModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReconcileModelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReconcileModelId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.reconcile.model was not found with criteria %v and options %v", criteria, options)
}
