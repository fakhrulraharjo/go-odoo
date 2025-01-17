package odoo

import (
	"fmt"
)

// SmsTemplate represents sms.template model.
type SmsTemplate struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	Body            *String   `xmlrpc:"body,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	Lang            *String   `xmlrpc:"lang,omptempty"`
	Model           *String   `xmlrpc:"model,omptempty"`
	ModelId         *Many2One `xmlrpc:"model_id,omptempty"`
	Name            *String   `xmlrpc:"name,omptempty"`
	RenderModel     *String   `xmlrpc:"render_model,omptempty"`
	SidebarActionId *Many2One `xmlrpc:"sidebar_action_id,omptempty"`
	TemplateFs      *String   `xmlrpc:"template_fs,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// SmsTemplates represents array of sms.template model.
type SmsTemplates []SmsTemplate

// SmsTemplateModel is the odoo model name.
const SmsTemplateModel = "sms.template"

// Many2One convert SmsTemplate to *Many2One.
func (st *SmsTemplate) Many2One() *Many2One {
	return NewMany2One(st.Id.Get(), "")
}

// CreateSmsTemplate creates a new sms.template model and returns its id.
func (c *Client) CreateSmsTemplate(st *SmsTemplate) (int64, error) {
	return c.Create(SmsTemplateModel, st)
}

// UpdateSmsTemplate updates an existing sms.template record.
func (c *Client) UpdateSmsTemplate(st *SmsTemplate) error {
	return c.UpdateSmsTemplates([]int64{st.Id.Get()}, st)
}

// UpdateSmsTemplates updates existing sms.template records.
// All records (represented by ids) will be updated by st values.
func (c *Client) UpdateSmsTemplates(ids []int64, st *SmsTemplate) error {
	return c.Update(SmsTemplateModel, ids, st)
}

// DeleteSmsTemplate deletes an existing sms.template record.
func (c *Client) DeleteSmsTemplate(id int64) error {
	return c.DeleteSmsTemplates([]int64{id})
}

// DeleteSmsTemplates deletes existing sms.template records.
func (c *Client) DeleteSmsTemplates(ids []int64) error {
	return c.Delete(SmsTemplateModel, ids)
}

// GetSmsTemplate gets sms.template existing record.
func (c *Client) GetSmsTemplate(id int64) (*SmsTemplate, error) {
	sts, err := c.GetSmsTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if sts != nil && len(*sts) > 0 {
		return &((*sts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sms.template not found", id)
}

// GetSmsTemplates gets sms.template existing records.
func (c *Client) GetSmsTemplates(ids []int64) (*SmsTemplates, error) {
	sts := &SmsTemplates{}
	if err := c.Read(SmsTemplateModel, ids, nil, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSmsTemplate finds sms.template record by querying it with criteria.
func (c *Client) FindSmsTemplate(criteria *Criteria) (*SmsTemplate, error) {
	sts := &SmsTemplates{}
	if err := c.SearchRead(SmsTemplateModel, criteria, NewOptions().Limit(1), sts); err != nil {
		return nil, err
	}
	if sts != nil && len(*sts) > 0 {
		return &((*sts)[0]), nil
	}
	return nil, fmt.Errorf("sms.template was not found with criteria %v", criteria)
}

// FindSmsTemplates finds sms.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTemplates(criteria *Criteria, options *Options) (*SmsTemplates, error) {
	sts := &SmsTemplates{}
	if err := c.SearchRead(SmsTemplateModel, criteria, options, sts); err != nil {
		return nil, err
	}
	return sts, nil
}

// FindSmsTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SmsTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSmsTemplateId finds record id by querying it with criteria.
func (c *Client) FindSmsTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sms.template was not found with criteria %v and options %v", criteria, options)
}
