package odoo

import (
	"fmt"
)

// AccountAccount represents account.account model.
type AccountAccount struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AccountType              *Selection `xmlrpc:"account_type,omptempty"`
	AllowedJournalIds        *Relation  `xmlrpc:"allowed_journal_ids,omptempty"`
	Code                     *String    `xmlrpc:"code,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId               *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrentBalance           *Float     `xmlrpc:"current_balance,omptempty"`
	Deprecated               *Bool      `xmlrpc:"deprecated,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	GroupId                  *Many2One  `xmlrpc:"group_id,omptempty"`
	HasMessage               *Bool      `xmlrpc:"has_message,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	IncludeInitialBalance    *Bool      `xmlrpc:"include_initial_balance,omptempty"`
	InternalGroup            *Selection `xmlrpc:"internal_group,omptempty"`
	IsOffBalance             *Bool      `xmlrpc:"is_off_balance,omptempty"`
	MessageAttachmentCount   *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds       *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError          *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter   *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError       *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds               *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower        *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId  *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction        *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds        *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	NonTrade                 *Bool      `xmlrpc:"non_trade,omptempty"`
	Note                     *String    `xmlrpc:"note,omptempty"`
	OpeningBalance           *Float     `xmlrpc:"opening_balance,omptempty"`
	OpeningCredit            *Float     `xmlrpc:"opening_credit,omptempty"`
	OpeningDebit             *Float     `xmlrpc:"opening_debit,omptempty"`
	Reconcile                *Bool      `xmlrpc:"reconcile,omptempty"`
	RelatedTaxesAmount       *Int       `xmlrpc:"related_taxes_amount,omptempty"`
	RootId                   *Many2One  `xmlrpc:"root_id,omptempty"`
	TagIds                   *Relation  `xmlrpc:"tag_ids,omptempty"`
	TaxIds                   *Relation  `xmlrpc:"tax_ids,omptempty"`
	Used                     *Bool      `xmlrpc:"used,omptempty"`
	WebsiteMessageIds        *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountAccounts represents array of account.account model.
type AccountAccounts []AccountAccount

// AccountAccountModel is the odoo model name.
const AccountAccountModel = "account.account"

// Many2One convert AccountAccount to *Many2One.
func (aa *AccountAccount) Many2One() *Many2One {
	return NewMany2One(aa.Id.Get(), "")
}

// CreateAccountAccount creates a new account.account model and returns its id.
func (c *Client) CreateAccountAccount(aa *AccountAccount) (int64, error) {
	return c.Create(AccountAccountModel, aa)
}

// UpdateAccountAccount updates an existing account.account record.
func (c *Client) UpdateAccountAccount(aa *AccountAccount) error {
	return c.UpdateAccountAccounts([]int64{aa.Id.Get()}, aa)
}

// UpdateAccountAccounts updates existing account.account records.
// All records (represented by ids) will be updated by aa values.
func (c *Client) UpdateAccountAccounts(ids []int64, aa *AccountAccount) error {
	return c.Update(AccountAccountModel, ids, aa)
}

// DeleteAccountAccount deletes an existing account.account record.
func (c *Client) DeleteAccountAccount(id int64) error {
	return c.DeleteAccountAccounts([]int64{id})
}

// DeleteAccountAccounts deletes existing account.account records.
func (c *Client) DeleteAccountAccounts(ids []int64) error {
	return c.Delete(AccountAccountModel, ids)
}

// GetAccountAccount gets account.account existing record.
func (c *Client) GetAccountAccount(id int64) (*AccountAccount, error) {
	aas, err := c.GetAccountAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.account not found", id)
}

// GetAccountAccounts gets account.account existing records.
func (c *Client) GetAccountAccounts(ids []int64) (*AccountAccounts, error) {
	aas := &AccountAccounts{}
	if err := c.Read(AccountAccountModel, ids, nil, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAccount finds account.account record by querying it with criteria.
func (c *Client) FindAccountAccount(criteria *Criteria) (*AccountAccount, error) {
	aas := &AccountAccounts{}
	if err := c.SearchRead(AccountAccountModel, criteria, NewOptions().Limit(1), aas); err != nil {
		return nil, err
	}
	if aas != nil && len(*aas) > 0 {
		return &((*aas)[0]), nil
	}
	return nil, fmt.Errorf("account.account was not found with criteria %v", criteria)
}

// FindAccountAccounts finds account.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccounts(criteria *Criteria, options *Options) (*AccountAccounts, error) {
	aas := &AccountAccounts{}
	if err := c.SearchRead(AccountAccountModel, criteria, options, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAccountModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAccountId finds record id by querying it with criteria.
func (c *Client) FindAccountAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.account was not found with criteria %v and options %v", criteria, options)
}
