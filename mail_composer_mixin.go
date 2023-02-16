package odoo

import (
	"fmt"
)

// MailComposerMixin represents mail.composer.mixin model.
type MailComposerMixin struct {
	Body                 *String   `xmlrpc:"body,omptempty"`
	CanEditBody          *Bool     `xmlrpc:"can_edit_body,omptempty"`
	IsMailTemplateEditor *Bool     `xmlrpc:"is_mail_template_editor,omptempty"`
	Lang                 *String   `xmlrpc:"lang,omptempty"`
	RenderModel          *String   `xmlrpc:"render_model,omptempty"`
	Subject              *String   `xmlrpc:"subject,omptempty"`
	TemplateId           *Many2One `xmlrpc:"template_id,omptempty"`
}

// MailComposerMixins represents array of mail.composer.mixin model.
type MailComposerMixins []MailComposerMixin

// MailComposerMixinModel is the odoo model name.
const MailComposerMixinModel = "mail.composer.mixin"

// Many2One convert MailComposerMixin to *Many2One.
func (mcm *MailComposerMixin) Many2One() *Many2One {
	return NewMany2One(mcm.Id.Get(), "")
}

// CreateMailComposerMixin creates a new mail.composer.mixin model and returns its id.
func (c *Client) CreateMailComposerMixin(mcm *MailComposerMixin) (int64, error) {
	return c.Create(MailComposerMixinModel, mcm)
}

// UpdateMailComposerMixin updates an existing mail.composer.mixin record.
func (c *Client) UpdateMailComposerMixin(mcm *MailComposerMixin) error {
	return c.UpdateMailComposerMixins([]int64{mcm.Id.Get()}, mcm)
}

// UpdateMailComposerMixins updates existing mail.composer.mixin records.
// All records (represented by ids) will be updated by mcm values.
func (c *Client) UpdateMailComposerMixins(ids []int64, mcm *MailComposerMixin) error {
	return c.Update(MailComposerMixinModel, ids, mcm)
}

// DeleteMailComposerMixin deletes an existing mail.composer.mixin record.
func (c *Client) DeleteMailComposerMixin(id int64) error {
	return c.DeleteMailComposerMixins([]int64{id})
}

// DeleteMailComposerMixins deletes existing mail.composer.mixin records.
func (c *Client) DeleteMailComposerMixins(ids []int64) error {
	return c.Delete(MailComposerMixinModel, ids)
}

// GetMailComposerMixin gets mail.composer.mixin existing record.
func (c *Client) GetMailComposerMixin(id int64) (*MailComposerMixin, error) {
	mcms, err := c.GetMailComposerMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcms != nil && len(*mcms) > 0 {
		return &((*mcms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.composer.mixin not found", id)
}

// GetMailComposerMixins gets mail.composer.mixin existing records.
func (c *Client) GetMailComposerMixins(ids []int64) (*MailComposerMixins, error) {
	mcms := &MailComposerMixins{}
	if err := c.Read(MailComposerMixinModel, ids, nil, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposerMixin finds mail.composer.mixin record by querying it with criteria.
func (c *Client) FindMailComposerMixin(criteria *Criteria) (*MailComposerMixin, error) {
	mcms := &MailComposerMixins{}
	if err := c.SearchRead(MailComposerMixinModel, criteria, NewOptions().Limit(1), mcms); err != nil {
		return nil, err
	}
	if mcms != nil && len(*mcms) > 0 {
		return &((*mcms)[0]), nil
	}
	return nil, fmt.Errorf("mail.composer.mixin was not found with criteria %v", criteria)
}

// FindMailComposerMixins finds mail.composer.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposerMixins(criteria *Criteria, options *Options) (*MailComposerMixins, error) {
	mcms := &MailComposerMixins{}
	if err := c.SearchRead(MailComposerMixinModel, criteria, options, mcms); err != nil {
		return nil, err
	}
	return mcms, nil
}

// FindMailComposerMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailComposerMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailComposerMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailComposerMixinId finds record id by querying it with criteria.
func (c *Client) FindMailComposerMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailComposerMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.composer.mixin was not found with criteria %v and options %v", criteria, options)
}
