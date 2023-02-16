package odoo

import (
	"fmt"
)

// MultiLivechatProviderContext represents multi_livechat.provider_context model.
type MultiLivechatProviderContext struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	MessageId       *Many2One  `xmlrpc:"message_id,omptempty"`
	ProviderContext *Selection `xmlrpc:"provider_context,omptempty"`
	Status          *Selection `xmlrpc:"status,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MultiLivechatProviderContexts represents array of multi_livechat.provider_context model.
type MultiLivechatProviderContexts []MultiLivechatProviderContext

// MultiLivechatProviderContextModel is the odoo model name.
const MultiLivechatProviderContextModel = "multi_livechat.provider_context"

// Many2One convert MultiLivechatProviderContext to *Many2One.
func (mp *MultiLivechatProviderContext) Many2One() *Many2One {
	return NewMany2One(mp.Id.Get(), "")
}

// CreateMultiLivechatProviderContext creates a new multi_livechat.provider_context model and returns its id.
func (c *Client) CreateMultiLivechatProviderContext(mp *MultiLivechatProviderContext) (int64, error) {
	return c.Create(MultiLivechatProviderContextModel, mp)
}

// UpdateMultiLivechatProviderContext updates an existing multi_livechat.provider_context record.
func (c *Client) UpdateMultiLivechatProviderContext(mp *MultiLivechatProviderContext) error {
	return c.UpdateMultiLivechatProviderContexts([]int64{mp.Id.Get()}, mp)
}

// UpdateMultiLivechatProviderContexts updates existing multi_livechat.provider_context records.
// All records (represented by ids) will be updated by mp values.
func (c *Client) UpdateMultiLivechatProviderContexts(ids []int64, mp *MultiLivechatProviderContext) error {
	return c.Update(MultiLivechatProviderContextModel, ids, mp)
}

// DeleteMultiLivechatProviderContext deletes an existing multi_livechat.provider_context record.
func (c *Client) DeleteMultiLivechatProviderContext(id int64) error {
	return c.DeleteMultiLivechatProviderContexts([]int64{id})
}

// DeleteMultiLivechatProviderContexts deletes existing multi_livechat.provider_context records.
func (c *Client) DeleteMultiLivechatProviderContexts(ids []int64) error {
	return c.Delete(MultiLivechatProviderContextModel, ids)
}

// GetMultiLivechatProviderContext gets multi_livechat.provider_context existing record.
func (c *Client) GetMultiLivechatProviderContext(id int64) (*MultiLivechatProviderContext, error) {
	mps, err := c.GetMultiLivechatProviderContexts([]int64{id})
	if err != nil {
		return nil, err
	}
	if mps != nil && len(*mps) > 0 {
		return &((*mps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of multi_livechat.provider_context not found", id)
}

// GetMultiLivechatProviderContexts gets multi_livechat.provider_context existing records.
func (c *Client) GetMultiLivechatProviderContexts(ids []int64) (*MultiLivechatProviderContexts, error) {
	mps := &MultiLivechatProviderContexts{}
	if err := c.Read(MultiLivechatProviderContextModel, ids, nil, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMultiLivechatProviderContext finds multi_livechat.provider_context record by querying it with criteria.
func (c *Client) FindMultiLivechatProviderContext(criteria *Criteria) (*MultiLivechatProviderContext, error) {
	mps := &MultiLivechatProviderContexts{}
	if err := c.SearchRead(MultiLivechatProviderContextModel, criteria, NewOptions().Limit(1), mps); err != nil {
		return nil, err
	}
	if mps != nil && len(*mps) > 0 {
		return &((*mps)[0]), nil
	}
	return nil, fmt.Errorf("multi_livechat.provider_context was not found with criteria %v", criteria)
}

// FindMultiLivechatProviderContexts finds multi_livechat.provider_context records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMultiLivechatProviderContexts(criteria *Criteria, options *Options) (*MultiLivechatProviderContexts, error) {
	mps := &MultiLivechatProviderContexts{}
	if err := c.SearchRead(MultiLivechatProviderContextModel, criteria, options, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMultiLivechatProviderContextIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMultiLivechatProviderContextIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MultiLivechatProviderContextModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMultiLivechatProviderContextId finds record id by querying it with criteria.
func (c *Client) FindMultiLivechatProviderContextId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MultiLivechatProviderContextModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("multi_livechat.provider_context was not found with criteria %v and options %v", criteria, options)
}
