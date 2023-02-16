package odoo

import (
	"fmt"
)

// AccountBankStatementLine represents account.bank.statement.line model.
type AccountBankStatementLine struct {
	LastUpdate                            *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken                           *String    `xmlrpc:"access_token,omptempty"`
	AccessUrl                             *String    `xmlrpc:"access_url,omptempty"`
	AccessWarning                         *String    `xmlrpc:"access_warning,omptempty"`
	AccountNumber                         *String    `xmlrpc:"account_number,omptempty"`
	ActivityDateDeadline                  *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration           *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon                 *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                           *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                         *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                       *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon                      *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                        *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                        *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AlwaysTaxExigible                     *Bool      `xmlrpc:"always_tax_exigible,omptempty"`
	Amount                                *Float     `xmlrpc:"amount,omptempty"`
	AmountCurrency                        *Float     `xmlrpc:"amount_currency,omptempty"`
	AmountPaid                            *Float     `xmlrpc:"amount_paid,omptempty"`
	AmountResidual                        *Float     `xmlrpc:"amount_residual,omptempty"`
	AmountResidualSigned                  *Float     `xmlrpc:"amount_residual_signed,omptempty"`
	AmountTax                             *Float     `xmlrpc:"amount_tax,omptempty"`
	AmountTaxSigned                       *Float     `xmlrpc:"amount_tax_signed,omptempty"`
	AmountTotal                           *Float     `xmlrpc:"amount_total,omptempty"`
	AmountTotalInCurrencySigned           *Float     `xmlrpc:"amount_total_in_currency_signed,omptempty"`
	AmountTotalSigned                     *Float     `xmlrpc:"amount_total_signed,omptempty"`
	AmountUntaxed                         *Float     `xmlrpc:"amount_untaxed,omptempty"`
	AmountUntaxedSigned                   *Float     `xmlrpc:"amount_untaxed_signed,omptempty"`
	AttachmentIds                         *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AuthorizedTransactionIds              *Relation  `xmlrpc:"authorized_transaction_ids,omptempty"`
	AutoPost                              *Selection `xmlrpc:"auto_post,omptempty"`
	AutoPostOriginId                      *Many2One  `xmlrpc:"auto_post_origin_id,omptempty"`
	AutoPostUntil                         *Time      `xmlrpc:"auto_post_until,omptempty"`
	BankPartnerId                         *Many2One  `xmlrpc:"bank_partner_id,omptempty"`
	CommercialPartnerId                   *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompanyCurrencyId                     *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryCode                           *String    `xmlrpc:"country_code,omptempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                            *Many2One  `xmlrpc:"currency_id,omptempty"`
	Date                                  *Time      `xmlrpc:"date,omptempty"`
	DirectionSign                         *Int       `xmlrpc:"direction_sign,omptempty"`
	DisplayInactiveCurrencyWarning        *Bool      `xmlrpc:"display_inactive_currency_warning,omptempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omptempty"`
	DisplayQrCode                         *Bool      `xmlrpc:"display_qr_code,omptempty"`
	DuplicatedRefIds                      *Relation  `xmlrpc:"duplicated_ref_ids,omptempty"`
	EdiBlockingLevel                      *Selection `xmlrpc:"edi_blocking_level,omptempty"`
	EdiDocumentIds                        *Relation  `xmlrpc:"edi_document_ids,omptempty"`
	EdiErrorCount                         *Int       `xmlrpc:"edi_error_count,omptempty"`
	EdiErrorMessage                       *String    `xmlrpc:"edi_error_message,omptempty"`
	EdiShowAbandonCancelButton            *Bool      `xmlrpc:"edi_show_abandon_cancel_button,omptempty"`
	EdiShowCancelButton                   *Bool      `xmlrpc:"edi_show_cancel_button,omptempty"`
	EdiState                              *Selection `xmlrpc:"edi_state,omptempty"`
	EdiWebServicesToProcess               *String    `xmlrpc:"edi_web_services_to_process,omptempty"`
	FiscalPositionId                      *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	ForeignCurrencyId                     *Many2One  `xmlrpc:"foreign_currency_id,omptempty"`
	HasMessage                            *Bool      `xmlrpc:"has_message,omptempty"`
	HasReconciledEntries                  *Bool      `xmlrpc:"has_reconciled_entries,omptempty"`
	HidePostButton                        *Bool      `xmlrpc:"hide_post_button,omptempty"`
	HighestName                           *String    `xmlrpc:"highest_name,omptempty"`
	Id                                    *Int       `xmlrpc:"id,omptempty"`
	InalterableHash                       *String    `xmlrpc:"inalterable_hash,omptempty"`
	InternalIndex                         *String    `xmlrpc:"internal_index,omptempty"`
	InvoiceCashRoundingId                 *Many2One  `xmlrpc:"invoice_cash_rounding_id,omptempty"`
	InvoiceDate                           *Time      `xmlrpc:"invoice_date,omptempty"`
	InvoiceDateDue                        *Time      `xmlrpc:"invoice_date_due,omptempty"`
	InvoiceFilterTypeDomain               *String    `xmlrpc:"invoice_filter_type_domain,omptempty"`
	InvoiceHasOutstanding                 *Bool      `xmlrpc:"invoice_has_outstanding,omptempty"`
	InvoiceIncotermId                     *Many2One  `xmlrpc:"invoice_incoterm_id,omptempty"`
	InvoiceLineIds                        *Relation  `xmlrpc:"invoice_line_ids,omptempty"`
	InvoiceOrigin                         *String    `xmlrpc:"invoice_origin,omptempty"`
	InvoiceOutstandingCreditsDebitsWidget *String    `xmlrpc:"invoice_outstanding_credits_debits_widget,omptempty"`
	InvoicePartnerDisplayName             *String    `xmlrpc:"invoice_partner_display_name,omptempty"`
	InvoicePaymentTermId                  *Many2One  `xmlrpc:"invoice_payment_term_id,omptempty"`
	InvoicePaymentsWidget                 *String    `xmlrpc:"invoice_payments_widget,omptempty"`
	InvoiceSourceEmail                    *String    `xmlrpc:"invoice_source_email,omptempty"`
	InvoiceUserId                         *Many2One  `xmlrpc:"invoice_user_id,omptempty"`
	InvoiceVendorBillId                   *Many2One  `xmlrpc:"invoice_vendor_bill_id,omptempty"`
	IsMoveSent                            *Bool      `xmlrpc:"is_move_sent,omptempty"`
	IsReconciled                          *Bool      `xmlrpc:"is_reconciled,omptempty"`
	IsStorno                              *Bool      `xmlrpc:"is_storno,omptempty"`
	JournalId                             *Many2One  `xmlrpc:"journal_id,omptempty"`
	L10NIdAttachmentId                    *Many2One  `xmlrpc:"l10n_id_attachment_id,omptempty"`
	L10NIdCsvCreated                      *Bool      `xmlrpc:"l10n_id_csv_created,omptempty"`
	L10NIdKodeTransaksi                   *Selection `xmlrpc:"l10n_id_kode_transaksi,omptempty"`
	L10NIdNeedKodeTransaksi               *Bool      `xmlrpc:"l10n_id_need_kode_transaksi,omptempty"`
	L10NIdReplaceInvoiceId                *Many2One  `xmlrpc:"l10n_id_replace_invoice_id,omptempty"`
	L10NIdTaxNumber                       *String    `xmlrpc:"l10n_id_tax_number,omptempty"`
	LineIds                               *Relation  `xmlrpc:"line_ids,omptempty"`
	MadeSequenceHole                      *Bool      `xmlrpc:"made_sequence_hole,omptempty"`
	MessageAttachmentCount                *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds                    *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                       *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter                *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError                    *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                            *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower                     *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId               *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction                     *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter              *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds                     *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MoveId                                *Many2One  `xmlrpc:"move_id,omptempty"`
	MoveType                              *Selection `xmlrpc:"move_type,omptempty"`
	MyActivityDateDeadline                *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                                  *String    `xmlrpc:"name,omptempty"`
	Narration                             *String    `xmlrpc:"narration,omptempty"`
	NeededTerms                           *String    `xmlrpc:"needed_terms,omptempty"`
	NeededTermsDirty                      *Bool      `xmlrpc:"needed_terms_dirty,omptempty"`
	PartnerBankId                         *Many2One  `xmlrpc:"partner_bank_id,omptempty"`
	PartnerCreditWarning                  *String    `xmlrpc:"partner_credit_warning,omptempty"`
	PartnerId                             *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerName                           *String    `xmlrpc:"partner_name,omptempty"`
	PartnerShippingId                     *Many2One  `xmlrpc:"partner_shipping_id,omptempty"`
	PaymentId                             *Many2One  `xmlrpc:"payment_id,omptempty"`
	PaymentIds                            *Relation  `xmlrpc:"payment_ids,omptempty"`
	PaymentRef                            *String    `xmlrpc:"payment_ref,omptempty"`
	PaymentReference                      *String    `xmlrpc:"payment_reference,omptempty"`
	PaymentState                          *Selection `xmlrpc:"payment_state,omptempty"`
	PaymentTermDetails                    *String    `xmlrpc:"payment_term_details,omptempty"`
	PosOrderIds                           *Relation  `xmlrpc:"pos_order_ids,omptempty"`
	PosPaymentIds                         *Relation  `xmlrpc:"pos_payment_ids,omptempty"`
	PosSessionId                          *Many2One  `xmlrpc:"pos_session_id,omptempty"`
	PostedBefore                          *Bool      `xmlrpc:"posted_before,omptempty"`
	QrCodeMethod                          *Selection `xmlrpc:"qr_code_method,omptempty"`
	QuickEditMode                         *Bool      `xmlrpc:"quick_edit_mode,omptempty"`
	QuickEditTotalAmount                  *Float     `xmlrpc:"quick_edit_total_amount,omptempty"`
	QuickEncodingVals                     *String    `xmlrpc:"quick_encoding_vals,omptempty"`
	Ref                                   *String    `xmlrpc:"ref,omptempty"`
	RestrictModeHashTable                 *Bool      `xmlrpc:"restrict_mode_hash_table,omptempty"`
	ReversalMoveId                        *Relation  `xmlrpc:"reversal_move_id,omptempty"`
	ReversedEntryId                       *Many2One  `xmlrpc:"reversed_entry_id,omptempty"`
	RunningBalance                        *Float     `xmlrpc:"running_balance,omptempty"`
	SecureSequenceNumber                  *Int       `xmlrpc:"secure_sequence_number,omptempty"`
	Sequence                              *Int       `xmlrpc:"sequence,omptempty"`
	SequenceNumber                        *Int       `xmlrpc:"sequence_number,omptempty"`
	SequencePrefix                        *String    `xmlrpc:"sequence_prefix,omptempty"`
	ShowDiscountDetails                   *Bool      `xmlrpc:"show_discount_details,omptempty"`
	ShowNameWarning                       *Bool      `xmlrpc:"show_name_warning,omptempty"`
	ShowPaymentTermDetails                *Bool      `xmlrpc:"show_payment_term_details,omptempty"`
	ShowResetToDraftButton                *Bool      `xmlrpc:"show_reset_to_draft_button,omptempty"`
	State                                 *Selection `xmlrpc:"state,omptempty"`
	StatementComplete                     *Bool      `xmlrpc:"statement_complete,omptempty"`
	StatementId                           *Many2One  `xmlrpc:"statement_id,omptempty"`
	StatementLineId                       *Many2One  `xmlrpc:"statement_line_id,omptempty"`
	StatementLineIds                      *Relation  `xmlrpc:"statement_line_ids,omptempty"`
	StatementValid                        *Bool      `xmlrpc:"statement_valid,omptempty"`
	StockMoveId                           *Many2One  `xmlrpc:"stock_move_id,omptempty"`
	StockValuationLayerIds                *Relation  `xmlrpc:"stock_valuation_layer_ids,omptempty"`
	StringToHash                          *String    `xmlrpc:"string_to_hash,omptempty"`
	SuitableJournalIds                    *Relation  `xmlrpc:"suitable_journal_ids,omptempty"`
	TaxCashBasisCreatedMoveIds            *Relation  `xmlrpc:"tax_cash_basis_created_move_ids,omptempty"`
	TaxCashBasisOriginMoveId              *Many2One  `xmlrpc:"tax_cash_basis_origin_move_id,omptempty"`
	TaxCashBasisRecId                     *Many2One  `xmlrpc:"tax_cash_basis_rec_id,omptempty"`
	TaxCountryCode                        *String    `xmlrpc:"tax_country_code,omptempty"`
	TaxCountryId                          *Many2One  `xmlrpc:"tax_country_id,omptempty"`
	TaxLockDateMessage                    *String    `xmlrpc:"tax_lock_date_message,omptempty"`
	TaxTotals                             *String    `xmlrpc:"tax_totals,omptempty"`
	ToCheck                               *Bool      `xmlrpc:"to_check,omptempty"`
	TransactionIds                        *Relation  `xmlrpc:"transaction_ids,omptempty"`
	TransactionType                       *String    `xmlrpc:"transaction_type,omptempty"`
	TypeName                              *String    `xmlrpc:"type_name,omptempty"`
	UserId                                *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds                     *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountBankStatementLines represents array of account.bank.statement.line model.
type AccountBankStatementLines []AccountBankStatementLine

// AccountBankStatementLineModel is the odoo model name.
const AccountBankStatementLineModel = "account.bank.statement.line"

// Many2One convert AccountBankStatementLine to *Many2One.
func (absl *AccountBankStatementLine) Many2One() *Many2One {
	return NewMany2One(absl.Id.Get(), "")
}

// CreateAccountBankStatementLine creates a new account.bank.statement.line model and returns its id.
func (c *Client) CreateAccountBankStatementLine(absl *AccountBankStatementLine) (int64, error) {
	return c.Create(AccountBankStatementLineModel, absl)
}

// UpdateAccountBankStatementLine updates an existing account.bank.statement.line record.
func (c *Client) UpdateAccountBankStatementLine(absl *AccountBankStatementLine) error {
	return c.UpdateAccountBankStatementLines([]int64{absl.Id.Get()}, absl)
}

// UpdateAccountBankStatementLines updates existing account.bank.statement.line records.
// All records (represented by ids) will be updated by absl values.
func (c *Client) UpdateAccountBankStatementLines(ids []int64, absl *AccountBankStatementLine) error {
	return c.Update(AccountBankStatementLineModel, ids, absl)
}

// DeleteAccountBankStatementLine deletes an existing account.bank.statement.line record.
func (c *Client) DeleteAccountBankStatementLine(id int64) error {
	return c.DeleteAccountBankStatementLines([]int64{id})
}

// DeleteAccountBankStatementLines deletes existing account.bank.statement.line records.
func (c *Client) DeleteAccountBankStatementLines(ids []int64) error {
	return c.Delete(AccountBankStatementLineModel, ids)
}

// GetAccountBankStatementLine gets account.bank.statement.line existing record.
func (c *Client) GetAccountBankStatementLine(id int64) (*AccountBankStatementLine, error) {
	absls, err := c.GetAccountBankStatementLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if absls != nil && len(*absls) > 0 {
		return &((*absls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.statement.line not found", id)
}

// GetAccountBankStatementLines gets account.bank.statement.line existing records.
func (c *Client) GetAccountBankStatementLines(ids []int64) (*AccountBankStatementLines, error) {
	absls := &AccountBankStatementLines{}
	if err := c.Read(AccountBankStatementLineModel, ids, nil, absls); err != nil {
		return nil, err
	}
	return absls, nil
}

// FindAccountBankStatementLine finds account.bank.statement.line record by querying it with criteria.
func (c *Client) FindAccountBankStatementLine(criteria *Criteria) (*AccountBankStatementLine, error) {
	absls := &AccountBankStatementLines{}
	if err := c.SearchRead(AccountBankStatementLineModel, criteria, NewOptions().Limit(1), absls); err != nil {
		return nil, err
	}
	if absls != nil && len(*absls) > 0 {
		return &((*absls)[0]), nil
	}
	return nil, fmt.Errorf("account.bank.statement.line was not found with criteria %v", criteria)
}

// FindAccountBankStatementLines finds account.bank.statement.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementLines(criteria *Criteria, options *Options) (*AccountBankStatementLines, error) {
	absls := &AccountBankStatementLines{}
	if err := c.SearchRead(AccountBankStatementLineModel, criteria, options, absls); err != nil {
		return nil, err
	}
	return absls, nil
}

// FindAccountBankStatementLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankStatementLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankStatementLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankStatementLineId finds record id by querying it with criteria.
func (c *Client) FindAccountBankStatementLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankStatementLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.bank.statement.line was not found with criteria %v and options %v", criteria, options)
}
