package odoo

import (
	"fmt"
)

// PosSession represents pos.session model.
type PosSession struct {
	LastUpdate                     *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityDateDeadline           *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration    *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon          *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                    *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                  *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary                *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon               *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId                 *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                 *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	BankPaymentIds                 *Relation  `xmlrpc:"bank_payment_ids,omptempty"`
	CashControl                    *Bool      `xmlrpc:"cash_control,omptempty"`
	CashJournalId                  *Many2One  `xmlrpc:"cash_journal_id,omptempty"`
	CashRealTransaction            *Float     `xmlrpc:"cash_real_transaction,omptempty"`
	CashRegisterBalanceEnd         *Float     `xmlrpc:"cash_register_balance_end,omptempty"`
	CashRegisterBalanceEndReal     *Float     `xmlrpc:"cash_register_balance_end_real,omptempty"`
	CashRegisterBalanceStart       *Float     `xmlrpc:"cash_register_balance_start,omptempty"`
	CashRegisterDifference         *Float     `xmlrpc:"cash_register_difference,omptempty"`
	CashRegisterTotalEntryEncoding *Float     `xmlrpc:"cash_register_total_entry_encoding,omptempty"`
	CompanyId                      *Many2One  `xmlrpc:"company_id,omptempty"`
	ConfigId                       *Many2One  `xmlrpc:"config_id,omptempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                     *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omptempty"`
	FailedPickings                 *Bool      `xmlrpc:"failed_pickings,omptempty"`
	HasMessage                     *Bool      `xmlrpc:"has_message,omptempty"`
	Id                             *Int       `xmlrpc:"id,omptempty"`
	IsInCompanyCurrency            *Bool      `xmlrpc:"is_in_company_currency,omptempty"`
	LoginNumber                    *Int       `xmlrpc:"login_number,omptempty"`
	MessageAttachmentCount         *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds             *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError                *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter         *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError             *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                     *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower              *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId        *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction              *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter       *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds              *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MoveId                         *Many2One  `xmlrpc:"move_id,omptempty"`
	MyActivityDateDeadline         *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                           *String    `xmlrpc:"name,omptempty"`
	OpeningNotes                   *String    `xmlrpc:"opening_notes,omptempty"`
	OrderCount                     *Int       `xmlrpc:"order_count,omptempty"`
	OrderIds                       *Relation  `xmlrpc:"order_ids,omptempty"`
	PaymentMethodIds               *Relation  `xmlrpc:"payment_method_ids,omptempty"`
	PickingCount                   *Int       `xmlrpc:"picking_count,omptempty"`
	PickingIds                     *Relation  `xmlrpc:"picking_ids,omptempty"`
	Rescue                         *Bool      `xmlrpc:"rescue,omptempty"`
	SequenceNumber                 *Int       `xmlrpc:"sequence_number,omptempty"`
	StartAt                        *Time      `xmlrpc:"start_at,omptempty"`
	State                          *Selection `xmlrpc:"state,omptempty"`
	StatementLineIds               *Relation  `xmlrpc:"statement_line_ids,omptempty"`
	StopAt                         *Time      `xmlrpc:"stop_at,omptempty"`
	TotalPaymentsAmount            *Float     `xmlrpc:"total_payments_amount,omptempty"`
	UpdateStockAtClosing           *Bool      `xmlrpc:"update_stock_at_closing,omptempty"`
	UserId                         *Many2One  `xmlrpc:"user_id,omptempty"`
	WebsiteMessageIds              *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PosSessions represents array of pos.session model.
type PosSessions []PosSession

// PosSessionModel is the odoo model name.
const PosSessionModel = "pos.session"

// Many2One convert PosSession to *Many2One.
func (ps *PosSession) Many2One() *Many2One {
	return NewMany2One(ps.Id.Get(), "")
}

// CreatePosSession creates a new pos.session model and returns its id.
func (c *Client) CreatePosSession(ps *PosSession) (int64, error) {
	return c.Create(PosSessionModel, ps)
}

// UpdatePosSession updates an existing pos.session record.
func (c *Client) UpdatePosSession(ps *PosSession) error {
	return c.UpdatePosSessions([]int64{ps.Id.Get()}, ps)
}

// UpdatePosSessions updates existing pos.session records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdatePosSessions(ids []int64, ps *PosSession) error {
	return c.Update(PosSessionModel, ids, ps)
}

// DeletePosSession deletes an existing pos.session record.
func (c *Client) DeletePosSession(id int64) error {
	return c.DeletePosSessions([]int64{id})
}

// DeletePosSessions deletes existing pos.session records.
func (c *Client) DeletePosSessions(ids []int64) error {
	return c.Delete(PosSessionModel, ids)
}

// GetPosSession gets pos.session existing record.
func (c *Client) GetPosSession(id int64) (*PosSession, error) {
	pss, err := c.GetPosSessions([]int64{id})
	if err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.session not found", id)
}

// GetPosSessions gets pos.session existing records.
func (c *Client) GetPosSessions(ids []int64) (*PosSessions, error) {
	pss := &PosSessions{}
	if err := c.Read(PosSessionModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPosSession finds pos.session record by querying it with criteria.
func (c *Client) FindPosSession(criteria *Criteria) (*PosSession, error) {
	pss := &PosSessions{}
	if err := c.SearchRead(PosSessionModel, criteria, NewOptions().Limit(1), pss); err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("pos.session was not found with criteria %v", criteria)
}

// FindPosSessions finds pos.session records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosSessions(criteria *Criteria, options *Options) (*PosSessions, error) {
	pss := &PosSessions{}
	if err := c.SearchRead(PosSessionModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPosSessionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosSessionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosSessionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosSessionId finds record id by querying it with criteria.
func (c *Client) FindPosSessionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosSessionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.session was not found with criteria %v and options %v", criteria, options)
}
