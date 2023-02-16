package odoo

import (
	"fmt"
)

// AccountJournal represents account.journal model.
type AccountJournal struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	AccountControlIds            *Relation  `xmlrpc:"account_control_ids,omptempty"`
	Active                       *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline         *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration  *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon        *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                  *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary              *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon             *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId               *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId               *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AliasDomain                  *String    `xmlrpc:"alias_domain,omptempty"`
	AliasId                      *Many2One  `xmlrpc:"alias_id,omptempty"`
	AliasName                    *String    `xmlrpc:"alias_name,omptempty"`
	AvailablePaymentMethodIds    *Relation  `xmlrpc:"available_payment_method_ids,omptempty"`
	BankAccNumber                *String    `xmlrpc:"bank_acc_number,omptempty"`
	BankAccountId                *Many2One  `xmlrpc:"bank_account_id,omptempty"`
	BankId                       *Many2One  `xmlrpc:"bank_id,omptempty"`
	BankStatementsSource         *Selection `xmlrpc:"bank_statements_source,omptempty"`
	Code                         *String    `xmlrpc:"code,omptempty"`
	Color                        *Int       `xmlrpc:"color,omptempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyPartnerId             *Many2One  `xmlrpc:"company_partner_id,omptempty"`
	CompatibleEdiIds             *Relation  `xmlrpc:"compatible_edi_ids,omptempty"`
	CountryCode                  *String    `xmlrpc:"country_code,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omptempty"`
	DefaultAccountId             *Many2One  `xmlrpc:"default_account_id,omptempty"`
	DefaultAccountType           *String    `xmlrpc:"default_account_type,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	EdiFormatIds                 *Relation  `xmlrpc:"edi_format_ids,omptempty"`
	EntriesCount                 *Int       `xmlrpc:"entries_count,omptempty"`
	HasMessage                   *Bool      `xmlrpc:"has_message,omptempty"`
	HasSequenceHoles             *Bool      `xmlrpc:"has_sequence_holes,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	InboundPaymentMethodLineIds  *Relation  `xmlrpc:"inbound_payment_method_line_ids,omptempty"`
	InvoiceReferenceModel        *Selection `xmlrpc:"invoice_reference_model,omptempty"`
	InvoiceReferenceType         *Selection `xmlrpc:"invoice_reference_type,omptempty"`
	JournalGroupIds              *Relation  `xmlrpc:"journal_group_ids,omptempty"`
	JsonActivityData             *String    `xmlrpc:"json_activity_data,omptempty"`
	KanbanDashboard              *String    `xmlrpc:"kanban_dashboard,omptempty"`
	KanbanDashboardGraph         *String    `xmlrpc:"kanban_dashboard_graph,omptempty"`
	LossAccountId                *Many2One  `xmlrpc:"loss_account_id,omptempty"`
	MessageAttachmentCount       *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds           *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError              *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter       *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError           *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                   *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower            *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId      *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction            *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter     *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds            *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline       *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	OutboundPaymentMethodLineIds *Relation  `xmlrpc:"outbound_payment_method_line_ids,omptempty"`
	PaymentSequence              *Bool      `xmlrpc:"payment_sequence,omptempty"`
	PosPaymentMethodIds          *Relation  `xmlrpc:"pos_payment_method_ids,omptempty"`
	ProfitAccountId              *Many2One  `xmlrpc:"profit_account_id,omptempty"`
	RefundSequence               *Bool      `xmlrpc:"refund_sequence,omptempty"`
	RestrictModeHashTable        *Bool      `xmlrpc:"restrict_mode_hash_table,omptempty"`
	SaleActivityNote             *String    `xmlrpc:"sale_activity_note,omptempty"`
	SaleActivityTypeId           *Many2One  `xmlrpc:"sale_activity_type_id,omptempty"`
	SaleActivityUserId           *Many2One  `xmlrpc:"sale_activity_user_id,omptempty"`
	SecureSequenceId             *Many2One  `xmlrpc:"secure_sequence_id,omptempty"`
	SelectedPaymentMethodCodes   *String    `xmlrpc:"selected_payment_method_codes,omptempty"`
	Sequence                     *Int       `xmlrpc:"sequence,omptempty"`
	SequenceOverrideRegex        *String    `xmlrpc:"sequence_override_regex,omptempty"`
	ShowOnDashboard              *Bool      `xmlrpc:"show_on_dashboard,omptempty"`
	SuspenseAccountId            *Many2One  `xmlrpc:"suspense_account_id,omptempty"`
	Type                         *Selection `xmlrpc:"type,omptempty"`
	WebsiteMessageIds            *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountJournals represents array of account.journal model.
type AccountJournals []AccountJournal

// AccountJournalModel is the odoo model name.
const AccountJournalModel = "account.journal"

// Many2One convert AccountJournal to *Many2One.
func (aj *AccountJournal) Many2One() *Many2One {
	return NewMany2One(aj.Id.Get(), "")
}

// CreateAccountJournal creates a new account.journal model and returns its id.
func (c *Client) CreateAccountJournal(aj *AccountJournal) (int64, error) {
	return c.Create(AccountJournalModel, aj)
}

// UpdateAccountJournal updates an existing account.journal record.
func (c *Client) UpdateAccountJournal(aj *AccountJournal) error {
	return c.UpdateAccountJournals([]int64{aj.Id.Get()}, aj)
}

// UpdateAccountJournals updates existing account.journal records.
// All records (represented by ids) will be updated by aj values.
func (c *Client) UpdateAccountJournals(ids []int64, aj *AccountJournal) error {
	return c.Update(AccountJournalModel, ids, aj)
}

// DeleteAccountJournal deletes an existing account.journal record.
func (c *Client) DeleteAccountJournal(id int64) error {
	return c.DeleteAccountJournals([]int64{id})
}

// DeleteAccountJournals deletes existing account.journal records.
func (c *Client) DeleteAccountJournals(ids []int64) error {
	return c.Delete(AccountJournalModel, ids)
}

// GetAccountJournal gets account.journal existing record.
func (c *Client) GetAccountJournal(id int64) (*AccountJournal, error) {
	ajs, err := c.GetAccountJournals([]int64{id})
	if err != nil {
		return nil, err
	}
	if ajs != nil && len(*ajs) > 0 {
		return &((*ajs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.journal not found", id)
}

// GetAccountJournals gets account.journal existing records.
func (c *Client) GetAccountJournals(ids []int64) (*AccountJournals, error) {
	ajs := &AccountJournals{}
	if err := c.Read(AccountJournalModel, ids, nil, ajs); err != nil {
		return nil, err
	}
	return ajs, nil
}

// FindAccountJournal finds account.journal record by querying it with criteria.
func (c *Client) FindAccountJournal(criteria *Criteria) (*AccountJournal, error) {
	ajs := &AccountJournals{}
	if err := c.SearchRead(AccountJournalModel, criteria, NewOptions().Limit(1), ajs); err != nil {
		return nil, err
	}
	if ajs != nil && len(*ajs) > 0 {
		return &((*ajs)[0]), nil
	}
	return nil, fmt.Errorf("account.journal was not found with criteria %v", criteria)
}

// FindAccountJournals finds account.journal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountJournals(criteria *Criteria, options *Options) (*AccountJournals, error) {
	ajs := &AccountJournals{}
	if err := c.SearchRead(AccountJournalModel, criteria, options, ajs); err != nil {
		return nil, err
	}
	return ajs, nil
}

// FindAccountJournalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountJournalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountJournalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountJournalId finds record id by querying it with criteria.
func (c *Client) FindAccountJournalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountJournalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.journal was not found with criteria %v and options %v", criteria, options)
}
