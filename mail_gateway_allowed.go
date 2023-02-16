package odoo

import (
	"fmt"
)

// MailGatewayAllowed represents mail.gateway.allowed model.
type MailGatewayAllowed struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Email           *String   `xmlrpc:"email,omptempty"`
	EmailNormalized *String   `xmlrpc:"email_normalized,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailGatewayAlloweds represents array of mail.gateway.allowed model.
type MailGatewayAlloweds []MailGatewayAllowed

// MailGatewayAllowedModel is the odoo model name.
const MailGatewayAllowedModel = "mail.gateway.allowed"

// Many2One convert MailGatewayAllowed to *Many2One.
func (mga *MailGatewayAllowed) Many2One() *Many2One {
	return NewMany2One(mga.Id.Get(), "")
}

// CreateMailGatewayAllowed creates a new mail.gateway.allowed model and returns its id.
func (c *Client) CreateMailGatewayAllowed(mga *MailGatewayAllowed) (int64, error) {
	return c.Create(MailGatewayAllowedModel, mga)
}

// UpdateMailGatewayAllowed updates an existing mail.gateway.allowed record.
func (c *Client) UpdateMailGatewayAllowed(mga *MailGatewayAllowed) error {
	return c.UpdateMailGatewayAlloweds([]int64{mga.Id.Get()}, mga)
}

// UpdateMailGatewayAlloweds updates existing mail.gateway.allowed records.
// All records (represented by ids) will be updated by mga values.
func (c *Client) UpdateMailGatewayAlloweds(ids []int64, mga *MailGatewayAllowed) error {
	return c.Update(MailGatewayAllowedModel, ids, mga)
}

// DeleteMailGatewayAllowed deletes an existing mail.gateway.allowed record.
func (c *Client) DeleteMailGatewayAllowed(id int64) error {
	return c.DeleteMailGatewayAlloweds([]int64{id})
}

// DeleteMailGatewayAlloweds deletes existing mail.gateway.allowed records.
func (c *Client) DeleteMailGatewayAlloweds(ids []int64) error {
	return c.Delete(MailGatewayAllowedModel, ids)
}

// GetMailGatewayAllowed gets mail.gateway.allowed existing record.
func (c *Client) GetMailGatewayAllowed(id int64) (*MailGatewayAllowed, error) {
	mgas, err := c.GetMailGatewayAlloweds([]int64{id})
	if err != nil {
		return nil, err
	}
	if mgas != nil && len(*mgas) > 0 {
		return &((*mgas)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.gateway.allowed not found", id)
}

// GetMailGatewayAlloweds gets mail.gateway.allowed existing records.
func (c *Client) GetMailGatewayAlloweds(ids []int64) (*MailGatewayAlloweds, error) {
	mgas := &MailGatewayAlloweds{}
	if err := c.Read(MailGatewayAllowedModel, ids, nil, mgas); err != nil {
		return nil, err
	}
	return mgas, nil
}

// FindMailGatewayAllowed finds mail.gateway.allowed record by querying it with criteria.
func (c *Client) FindMailGatewayAllowed(criteria *Criteria) (*MailGatewayAllowed, error) {
	mgas := &MailGatewayAlloweds{}
	if err := c.SearchRead(MailGatewayAllowedModel, criteria, NewOptions().Limit(1), mgas); err != nil {
		return nil, err
	}
	if mgas != nil && len(*mgas) > 0 {
		return &((*mgas)[0]), nil
	}
	return nil, fmt.Errorf("mail.gateway.allowed was not found with criteria %v", criteria)
}

// FindMailGatewayAlloweds finds mail.gateway.allowed records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailGatewayAlloweds(criteria *Criteria, options *Options) (*MailGatewayAlloweds, error) {
	mgas := &MailGatewayAlloweds{}
	if err := c.SearchRead(MailGatewayAllowedModel, criteria, options, mgas); err != nil {
		return nil, err
	}
	return mgas, nil
}

// FindMailGatewayAllowedIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailGatewayAllowedIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailGatewayAllowedModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailGatewayAllowedId finds record id by querying it with criteria.
func (c *Client) FindMailGatewayAllowedId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailGatewayAllowedModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.gateway.allowed was not found with criteria %v and options %v", criteria, options)
}
