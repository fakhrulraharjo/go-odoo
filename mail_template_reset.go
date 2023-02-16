package odoo

import (
	"fmt"
)

// MailTemplateReset represents mail.template.reset model.
type MailTemplateReset struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	TemplateIds *Relation `xmlrpc:"template_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailTemplateResets represents array of mail.template.reset model.
type MailTemplateResets []MailTemplateReset

// MailTemplateResetModel is the odoo model name.
const MailTemplateResetModel = "mail.template.reset"

// Many2One convert MailTemplateReset to *Many2One.
func (mtr *MailTemplateReset) Many2One() *Many2One {
	return NewMany2One(mtr.Id.Get(), "")
}

// CreateMailTemplateReset creates a new mail.template.reset model and returns its id.
func (c *Client) CreateMailTemplateReset(mtr *MailTemplateReset) (int64, error) {
	return c.Create(MailTemplateResetModel, mtr)
}

// UpdateMailTemplateReset updates an existing mail.template.reset record.
func (c *Client) UpdateMailTemplateReset(mtr *MailTemplateReset) error {
	return c.UpdateMailTemplateResets([]int64{mtr.Id.Get()}, mtr)
}

// UpdateMailTemplateResets updates existing mail.template.reset records.
// All records (represented by ids) will be updated by mtr values.
func (c *Client) UpdateMailTemplateResets(ids []int64, mtr *MailTemplateReset) error {
	return c.Update(MailTemplateResetModel, ids, mtr)
}

// DeleteMailTemplateReset deletes an existing mail.template.reset record.
func (c *Client) DeleteMailTemplateReset(id int64) error {
	return c.DeleteMailTemplateResets([]int64{id})
}

// DeleteMailTemplateResets deletes existing mail.template.reset records.
func (c *Client) DeleteMailTemplateResets(ids []int64) error {
	return c.Delete(MailTemplateResetModel, ids)
}

// GetMailTemplateReset gets mail.template.reset existing record.
func (c *Client) GetMailTemplateReset(id int64) (*MailTemplateReset, error) {
	mtrs, err := c.GetMailTemplateResets([]int64{id})
	if err != nil {
		return nil, err
	}
	if mtrs != nil && len(*mtrs) > 0 {
		return &((*mtrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.template.reset not found", id)
}

// GetMailTemplateResets gets mail.template.reset existing records.
func (c *Client) GetMailTemplateResets(ids []int64) (*MailTemplateResets, error) {
	mtrs := &MailTemplateResets{}
	if err := c.Read(MailTemplateResetModel, ids, nil, mtrs); err != nil {
		return nil, err
	}
	return mtrs, nil
}

// FindMailTemplateReset finds mail.template.reset record by querying it with criteria.
func (c *Client) FindMailTemplateReset(criteria *Criteria) (*MailTemplateReset, error) {
	mtrs := &MailTemplateResets{}
	if err := c.SearchRead(MailTemplateResetModel, criteria, NewOptions().Limit(1), mtrs); err != nil {
		return nil, err
	}
	if mtrs != nil && len(*mtrs) > 0 {
		return &((*mtrs)[0]), nil
	}
	return nil, fmt.Errorf("mail.template.reset was not found with criteria %v", criteria)
}

// FindMailTemplateResets finds mail.template.reset records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTemplateResets(criteria *Criteria, options *Options) (*MailTemplateResets, error) {
	mtrs := &MailTemplateResets{}
	if err := c.SearchRead(MailTemplateResetModel, criteria, options, mtrs); err != nil {
		return nil, err
	}
	return mtrs, nil
}

// FindMailTemplateResetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailTemplateResetIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailTemplateResetModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailTemplateResetId finds record id by querying it with criteria.
func (c *Client) FindMailTemplateResetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailTemplateResetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.template.reset was not found with criteria %v and options %v", criteria, options)
}
