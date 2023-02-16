package odoo

import (
	"fmt"
)

// MailIceServer represents mail.ice.server model.
type MailIceServer struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	Credential  *String    `xmlrpc:"credential,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	ServerType  *Selection `xmlrpc:"server_type,omptempty"`
	Uri         *String    `xmlrpc:"uri,omptempty"`
	Username    *String    `xmlrpc:"username,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailIceServers represents array of mail.ice.server model.
type MailIceServers []MailIceServer

// MailIceServerModel is the odoo model name.
const MailIceServerModel = "mail.ice.server"

// Many2One convert MailIceServer to *Many2One.
func (mis *MailIceServer) Many2One() *Many2One {
	return NewMany2One(mis.Id.Get(), "")
}

// CreateMailIceServer creates a new mail.ice.server model and returns its id.
func (c *Client) CreateMailIceServer(mis *MailIceServer) (int64, error) {
	return c.Create(MailIceServerModel, mis)
}

// UpdateMailIceServer updates an existing mail.ice.server record.
func (c *Client) UpdateMailIceServer(mis *MailIceServer) error {
	return c.UpdateMailIceServers([]int64{mis.Id.Get()}, mis)
}

// UpdateMailIceServers updates existing mail.ice.server records.
// All records (represented by ids) will be updated by mis values.
func (c *Client) UpdateMailIceServers(ids []int64, mis *MailIceServer) error {
	return c.Update(MailIceServerModel, ids, mis)
}

// DeleteMailIceServer deletes an existing mail.ice.server record.
func (c *Client) DeleteMailIceServer(id int64) error {
	return c.DeleteMailIceServers([]int64{id})
}

// DeleteMailIceServers deletes existing mail.ice.server records.
func (c *Client) DeleteMailIceServers(ids []int64) error {
	return c.Delete(MailIceServerModel, ids)
}

// GetMailIceServer gets mail.ice.server existing record.
func (c *Client) GetMailIceServer(id int64) (*MailIceServer, error) {
	miss, err := c.GetMailIceServers([]int64{id})
	if err != nil {
		return nil, err
	}
	if miss != nil && len(*miss) > 0 {
		return &((*miss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.ice.server not found", id)
}

// GetMailIceServers gets mail.ice.server existing records.
func (c *Client) GetMailIceServers(ids []int64) (*MailIceServers, error) {
	miss := &MailIceServers{}
	if err := c.Read(MailIceServerModel, ids, nil, miss); err != nil {
		return nil, err
	}
	return miss, nil
}

// FindMailIceServer finds mail.ice.server record by querying it with criteria.
func (c *Client) FindMailIceServer(criteria *Criteria) (*MailIceServer, error) {
	miss := &MailIceServers{}
	if err := c.SearchRead(MailIceServerModel, criteria, NewOptions().Limit(1), miss); err != nil {
		return nil, err
	}
	if miss != nil && len(*miss) > 0 {
		return &((*miss)[0]), nil
	}
	return nil, fmt.Errorf("mail.ice.server was not found with criteria %v", criteria)
}

// FindMailIceServers finds mail.ice.server records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailIceServers(criteria *Criteria, options *Options) (*MailIceServers, error) {
	miss := &MailIceServers{}
	if err := c.SearchRead(MailIceServerModel, criteria, options, miss); err != nil {
		return nil, err
	}
	return miss, nil
}

// FindMailIceServerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailIceServerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailIceServerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailIceServerId finds record id by querying it with criteria.
func (c *Client) FindMailIceServerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailIceServerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.ice.server was not found with criteria %v and options %v", criteria, options)
}
