package odoo

import (
	"fmt"
)

// MailTemplatePreview represents mail.template.preview model.
type MailTemplatePreview struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	AttachmentIds  *Relation  `xmlrpc:"attachment_ids,omptempty"`
	BodyHtml       *String    `xmlrpc:"body_html,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	EmailCc        *String    `xmlrpc:"email_cc,omptempty"`
	EmailFrom      *String    `xmlrpc:"email_from,omptempty"`
	EmailTo        *String    `xmlrpc:"email_to,omptempty"`
	ErrorMsg       *String    `xmlrpc:"error_msg,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	Lang           *Selection `xmlrpc:"lang,omptempty"`
	MailTemplateId *Many2One  `xmlrpc:"mail_template_id,omptempty"`
	ModelId        *Many2One  `xmlrpc:"model_id,omptempty"`
	NoRecord       *Bool      `xmlrpc:"no_record,omptempty"`
	PartnerIds     *Relation  `xmlrpc:"partner_ids,omptempty"`
	ReplyTo        *String    `xmlrpc:"reply_to,omptempty"`
	ResourceRef    *String    `xmlrpc:"resource_ref,omptempty"`
	ScheduledDate  *String    `xmlrpc:"scheduled_date,omptempty"`
	Subject        *String    `xmlrpc:"subject,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailTemplatePreviews represents array of mail.template.preview model.
type MailTemplatePreviews []MailTemplatePreview

// MailTemplatePreviewModel is the odoo model name.
const MailTemplatePreviewModel = "mail.template.preview"

// Many2One convert MailTemplatePreview to *Many2One.
func (mtp *MailTemplatePreview) Many2One() *Many2One {
	return NewMany2One(mtp.Id.Get(), "")
}

// CreateMailTemplatePreview creates a new mail.template.preview model and returns its id.
func (c *Client) CreateMailTemplatePreview(mtp *MailTemplatePreview) (int64, error) {
	return c.Create(MailTemplatePreviewModel, mtp)
}

// UpdateMailTemplatePreview updates an existing mail.template.preview record.
func (c *Client) UpdateMailTemplatePreview(mtp *MailTemplatePreview) error {
	return c.UpdateMailTemplatePreviews([]int64{mtp.Id.Get()}, mtp)
}

// UpdateMailTemplatePreviews updates existing mail.template.preview records.
// All records (represented by ids) will be updated by mtp values.
func (c *Client) UpdateMailTemplatePreviews(ids []int64, mtp *MailTemplatePreview) error {
	return c.Update(MailTemplatePreviewModel, ids, mtp)
}

// DeleteMailTemplatePreview deletes an existing mail.template.preview record.
func (c *Client) DeleteMailTemplatePreview(id int64) error {
	return c.DeleteMailTemplatePreviews([]int64{id})
}

// DeleteMailTemplatePreviews deletes existing mail.template.preview records.
func (c *Client) DeleteMailTemplatePreviews(ids []int64) error {
	return c.Delete(MailTemplatePreviewModel, ids)
}

// GetMailTemplatePreview gets mail.template.preview existing record.
func (c *Client) GetMailTemplatePreview(id int64) (*MailTemplatePreview, error) {
	mtps, err := c.GetMailTemplatePreviews([]int64{id})
	if err != nil {
		return nil, err
	}
	if mtps != nil && len(*mtps) > 0 {
		return &((*mtps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.template.preview not found", id)
}

// GetMailTemplatePreviews gets mail.template.preview existing records.
func (c *Client) GetMailTemplatePreviews(ids []int64) (*MailTemplatePreviews, error) {
	mtps := &MailTemplatePreviews{}
	if err := c.Read(MailTemplatePreviewModel, ids, nil, mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

// FindMailTemplatePreview finds mail.template.preview record by querying it with criteria.
func (c *Client) FindMailTemplatePreview(criteria *Criteria) (*MailTemplatePreview, error) {
	mtps := &MailTemplatePreviews{}
	if err := c.SearchRead(MailTemplatePreviewModel, criteria, NewOptions().Limit(1), mtps); err != nil {
		return nil, err
	}
	if mtps != nil && len(*mtps) > 0 {
		return &((*mtps)[0]), nil
	}
	return nil, fmt.Errorf("mail.template.preview was not found with criteria %v", criteria)
}

// FindMailTemplatePreviews finds mail.template.preview records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTemplatePreviews(criteria *Criteria, options *Options) (*MailTemplatePreviews, error) {
	mtps := &MailTemplatePreviews{}
	if err := c.SearchRead(MailTemplatePreviewModel, criteria, options, mtps); err != nil {
		return nil, err
	}
	return mtps, nil
}

// FindMailTemplatePreviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTemplatePreviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailTemplatePreviewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailTemplatePreviewId finds record id by querying it with criteria.
func (c *Client) FindMailTemplatePreviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTemplatePreviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.template.preview was not found with criteria %v and options %v", criteria, options)
}
