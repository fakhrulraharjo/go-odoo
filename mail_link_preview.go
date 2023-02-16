package odoo

import (
	"fmt"
)

// MailLinkPreview represents mail.link.preview model.
type MailLinkPreview struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	ImageMimetype *String   `xmlrpc:"image_mimetype,omptempty"`
	MessageId     *Many2One `xmlrpc:"message_id,omptempty"`
	OgDescription *String   `xmlrpc:"og_description,omptempty"`
	OgImage       *String   `xmlrpc:"og_image,omptempty"`
	OgMimetype    *String   `xmlrpc:"og_mimetype,omptempty"`
	OgTitle       *String   `xmlrpc:"og_title,omptempty"`
	OgType        *String   `xmlrpc:"og_type,omptempty"`
	SourceUrl     *String   `xmlrpc:"source_url,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailLinkPreviews represents array of mail.link.preview model.
type MailLinkPreviews []MailLinkPreview

// MailLinkPreviewModel is the odoo model name.
const MailLinkPreviewModel = "mail.link.preview"

// Many2One convert MailLinkPreview to *Many2One.
func (mlp *MailLinkPreview) Many2One() *Many2One {
	return NewMany2One(mlp.Id.Get(), "")
}

// CreateMailLinkPreview creates a new mail.link.preview model and returns its id.
func (c *Client) CreateMailLinkPreview(mlp *MailLinkPreview) (int64, error) {
	return c.Create(MailLinkPreviewModel, mlp)
}

// UpdateMailLinkPreview updates an existing mail.link.preview record.
func (c *Client) UpdateMailLinkPreview(mlp *MailLinkPreview) error {
	return c.UpdateMailLinkPreviews([]int64{mlp.Id.Get()}, mlp)
}

// UpdateMailLinkPreviews updates existing mail.link.preview records.
// All records (represented by ids) will be updated by mlp values.
func (c *Client) UpdateMailLinkPreviews(ids []int64, mlp *MailLinkPreview) error {
	return c.Update(MailLinkPreviewModel, ids, mlp)
}

// DeleteMailLinkPreview deletes an existing mail.link.preview record.
func (c *Client) DeleteMailLinkPreview(id int64) error {
	return c.DeleteMailLinkPreviews([]int64{id})
}

// DeleteMailLinkPreviews deletes existing mail.link.preview records.
func (c *Client) DeleteMailLinkPreviews(ids []int64) error {
	return c.Delete(MailLinkPreviewModel, ids)
}

// GetMailLinkPreview gets mail.link.preview existing record.
func (c *Client) GetMailLinkPreview(id int64) (*MailLinkPreview, error) {
	mlps, err := c.GetMailLinkPreviews([]int64{id})
	if err != nil {
		return nil, err
	}
	if mlps != nil && len(*mlps) > 0 {
		return &((*mlps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.link.preview not found", id)
}

// GetMailLinkPreviews gets mail.link.preview existing records.
func (c *Client) GetMailLinkPreviews(ids []int64) (*MailLinkPreviews, error) {
	mlps := &MailLinkPreviews{}
	if err := c.Read(MailLinkPreviewModel, ids, nil, mlps); err != nil {
		return nil, err
	}
	return mlps, nil
}

// FindMailLinkPreview finds mail.link.preview record by querying it with criteria.
func (c *Client) FindMailLinkPreview(criteria *Criteria) (*MailLinkPreview, error) {
	mlps := &MailLinkPreviews{}
	if err := c.SearchRead(MailLinkPreviewModel, criteria, NewOptions().Limit(1), mlps); err != nil {
		return nil, err
	}
	if mlps != nil && len(*mlps) > 0 {
		return &((*mlps)[0]), nil
	}
	return nil, fmt.Errorf("mail.link.preview was not found with criteria %v", criteria)
}

// FindMailLinkPreviews finds mail.link.preview records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailLinkPreviews(criteria *Criteria, options *Options) (*MailLinkPreviews, error) {
	mlps := &MailLinkPreviews{}
	if err := c.SearchRead(MailLinkPreviewModel, criteria, options, mlps); err != nil {
		return nil, err
	}
	return mlps, nil
}

// FindMailLinkPreviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailLinkPreviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailLinkPreviewModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailLinkPreviewId finds record id by querying it with criteria.
func (c *Client) FindMailLinkPreviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailLinkPreviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.link.preview was not found with criteria %v and options %v", criteria, options)
}
