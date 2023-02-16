package odoo

import (
	"fmt"
)

// MailMail represents mail.mail model.
type MailMail struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omptempty"`
	AttachmentIds              *Relation  `xmlrpc:"attachment_ids,omptempty"`
	AuthorAvatar               *String    `xmlrpc:"author_avatar,omptempty"`
	AuthorGuestId              *Many2One  `xmlrpc:"author_guest_id,omptempty"`
	AuthorId                   *Many2One  `xmlrpc:"author_id,omptempty"`
	AutoDelete                 *Bool      `xmlrpc:"auto_delete,omptempty"`
	Body                       *String    `xmlrpc:"body,omptempty"`
	BodyHtml                   *String    `xmlrpc:"body_html,omptempty"`
	CannedResponseIds          *Relation  `xmlrpc:"canned_response_ids,omptempty"`
	ChildIds                   *Relation  `xmlrpc:"child_ids,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date                       *Time      `xmlrpc:"date,omptempty"`
	Description                *String    `xmlrpc:"description,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	EmailAddSignature          *Bool      `xmlrpc:"email_add_signature,omptempty"`
	EmailCc                    *String    `xmlrpc:"email_cc,omptempty"`
	EmailFrom                  *String    `xmlrpc:"email_from,omptempty"`
	EmailLayoutXmlid           *String    `xmlrpc:"email_layout_xmlid,omptempty"`
	EmailTo                    *String    `xmlrpc:"email_to,omptempty"`
	FailureReason              *String    `xmlrpc:"failure_reason,omptempty"`
	FailureType                *Selection `xmlrpc:"failure_type,omptempty"`
	FetchmailServerId          *Many2One  `xmlrpc:"fetchmail_server_id,omptempty"`
	HasError                   *Bool      `xmlrpc:"has_error,omptempty"`
	HasSmsError                *Bool      `xmlrpc:"has_sms_error,omptempty"`
	Headers                    *String    `xmlrpc:"headers,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	IsCurrentUserOrGuestAuthor *Bool      `xmlrpc:"is_current_user_or_guest_author,omptempty"`
	IsInternal                 *Bool      `xmlrpc:"is_internal,omptempty"`
	IsNotification             *Bool      `xmlrpc:"is_notification,omptempty"`
	LetterIds                  *Relation  `xmlrpc:"letter_ids,omptempty"`
	LinkPreviewIds             *Relation  `xmlrpc:"link_preview_ids,omptempty"`
	MailActivityTypeId         *Many2One  `xmlrpc:"mail_activity_type_id,omptempty"`
	MailIds                    *Relation  `xmlrpc:"mail_ids,omptempty"`
	MailMessageId              *Many2One  `xmlrpc:"mail_message_id,omptempty"`
	MailMessageIdInt           *Int       `xmlrpc:"mail_message_id_int,omptempty"`
	MailServerId               *Many2One  `xmlrpc:"mail_server_id,omptempty"`
	MessageId                  *String    `xmlrpc:"message_id,omptempty"`
	MessageType                *Selection `xmlrpc:"message_type,omptempty"`
	Model                      *String    `xmlrpc:"model,omptempty"`
	Needaction                 *Bool      `xmlrpc:"needaction,omptempty"`
	NotificationIds            *Relation  `xmlrpc:"notification_ids,omptempty"`
	NotifiedPartnerIds         *Relation  `xmlrpc:"notified_partner_ids,omptempty"`
	ParentId                   *Many2One  `xmlrpc:"parent_id,omptempty"`
	PartnerIds                 *Relation  `xmlrpc:"partner_ids,omptempty"`
	Preview                    *String    `xmlrpc:"preview,omptempty"`
	ProviderContextId          *Many2One  `xmlrpc:"provider_context_id,omptempty"`
	ProviderContextIds         *Relation  `xmlrpc:"provider_context_ids,omptempty"`
	ReactionIds                *Relation  `xmlrpc:"reaction_ids,omptempty"`
	RecipientIds               *Relation  `xmlrpc:"recipient_ids,omptempty"`
	RecordName                 *String    `xmlrpc:"record_name,omptempty"`
	References                 *String    `xmlrpc:"references,omptempty"`
	ReplyTo                    *String    `xmlrpc:"reply_to,omptempty"`
	ReplyToForceNew            *Bool      `xmlrpc:"reply_to_force_new,omptempty"`
	ResId                      *Many2One  `xmlrpc:"res_id,omptempty"`
	RestrictedAttachmentCount  *Int       `xmlrpc:"restricted_attachment_count,omptempty"`
	ScheduledDate              *Time      `xmlrpc:"scheduled_date,omptempty"`
	SnailmailError             *Bool      `xmlrpc:"snailmail_error,omptempty"`
	Starred                    *Bool      `xmlrpc:"starred,omptempty"`
	StarredPartnerIds          *Relation  `xmlrpc:"starred_partner_ids,omptempty"`
	State                      *Selection `xmlrpc:"state,omptempty"`
	Subject                    *String    `xmlrpc:"subject,omptempty"`
	SubtypeId                  *Many2One  `xmlrpc:"subtype_id,omptempty"`
	ToDelete                   *Bool      `xmlrpc:"to_delete,omptempty"`
	TrackingValueIds           *Relation  `xmlrpc:"tracking_value_ids,omptempty"`
	UnrestrictedAttachmentIds  *Relation  `xmlrpc:"unrestricted_attachment_ids,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailMails represents array of mail.mail model.
type MailMails []MailMail

// MailMailModel is the odoo model name.
const MailMailModel = "mail.mail"

// Many2One convert MailMail to *Many2One.
func (mm *MailMail) Many2One() *Many2One {
	return NewMany2One(mm.Id.Get(), "")
}

// CreateMailMail creates a new mail.mail model and returns its id.
func (c *Client) CreateMailMail(mm *MailMail) (int64, error) {
	return c.Create(MailMailModel, mm)
}

// UpdateMailMail updates an existing mail.mail record.
func (c *Client) UpdateMailMail(mm *MailMail) error {
	return c.UpdateMailMails([]int64{mm.Id.Get()}, mm)
}

// UpdateMailMails updates existing mail.mail records.
// All records (represented by ids) will be updated by mm values.
func (c *Client) UpdateMailMails(ids []int64, mm *MailMail) error {
	return c.Update(MailMailModel, ids, mm)
}

// DeleteMailMail deletes an existing mail.mail record.
func (c *Client) DeleteMailMail(id int64) error {
	return c.DeleteMailMails([]int64{id})
}

// DeleteMailMails deletes existing mail.mail records.
func (c *Client) DeleteMailMails(ids []int64) error {
	return c.Delete(MailMailModel, ids)
}

// GetMailMail gets mail.mail existing record.
func (c *Client) GetMailMail(id int64) (*MailMail, error) {
	mms, err := c.GetMailMails([]int64{id})
	if err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.mail not found", id)
}

// GetMailMails gets mail.mail existing records.
func (c *Client) GetMailMails(ids []int64) (*MailMails, error) {
	mms := &MailMails{}
	if err := c.Read(MailMailModel, ids, nil, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMail finds mail.mail record by querying it with criteria.
func (c *Client) FindMailMail(criteria *Criteria) (*MailMail, error) {
	mms := &MailMails{}
	if err := c.SearchRead(MailMailModel, criteria, NewOptions().Limit(1), mms); err != nil {
		return nil, err
	}
	if mms != nil && len(*mms) > 0 {
		return &((*mms)[0]), nil
	}
	return nil, fmt.Errorf("mail.mail was not found with criteria %v", criteria)
}

// FindMailMails finds mail.mail records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMails(criteria *Criteria, options *Options) (*MailMails, error) {
	mms := &MailMails{}
	if err := c.SearchRead(MailMailModel, criteria, options, mms); err != nil {
		return nil, err
	}
	return mms, nil
}

// FindMailMailIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMailIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailMailModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMailId finds record id by querying it with criteria.
func (c *Client) FindMailMailId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMailModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.mail was not found with criteria %v and options %v", criteria, options)
}
