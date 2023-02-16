package odoo

import (
	"fmt"
)

// MultiLivechatProviderDictionary represents multi_livechat.provider_dictionary model.
type MultiLivechatProviderDictionary struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omptempty"`
	ChannelType     *String   `xmlrpc:"channel_type,omptempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName     *String   `xmlrpc:"display_name,omptempty"`
	Id              *Int      `xmlrpc:"id,omptempty"`
	ProviderContext *String   `xmlrpc:"provider_context,omptempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MultiLivechatProviderDictionarys represents array of multi_livechat.provider_dictionary model.
type MultiLivechatProviderDictionarys []MultiLivechatProviderDictionary

// MultiLivechatProviderDictionaryModel is the odoo model name.
const MultiLivechatProviderDictionaryModel = "multi_livechat.provider_dictionary"

// Many2One convert MultiLivechatProviderDictionary to *Many2One.
func (mp *MultiLivechatProviderDictionary) Many2One() *Many2One {
	return NewMany2One(mp.Id.Get(), "")
}

// CreateMultiLivechatProviderDictionary creates a new multi_livechat.provider_dictionary model and returns its id.
func (c *Client) CreateMultiLivechatProviderDictionary(mp *MultiLivechatProviderDictionary) (int64, error) {
	return c.Create(MultiLivechatProviderDictionaryModel, mp)
}

// UpdateMultiLivechatProviderDictionary updates an existing multi_livechat.provider_dictionary record.
func (c *Client) UpdateMultiLivechatProviderDictionary(mp *MultiLivechatProviderDictionary) error {
	return c.UpdateMultiLivechatProviderDictionarys([]int64{mp.Id.Get()}, mp)
}

// UpdateMultiLivechatProviderDictionarys updates existing multi_livechat.provider_dictionary records.
// All records (represented by ids) will be updated by mp values.
func (c *Client) UpdateMultiLivechatProviderDictionarys(ids []int64, mp *MultiLivechatProviderDictionary) error {
	return c.Update(MultiLivechatProviderDictionaryModel, ids, mp)
}

// DeleteMultiLivechatProviderDictionary deletes an existing multi_livechat.provider_dictionary record.
func (c *Client) DeleteMultiLivechatProviderDictionary(id int64) error {
	return c.DeleteMultiLivechatProviderDictionarys([]int64{id})
}

// DeleteMultiLivechatProviderDictionarys deletes existing multi_livechat.provider_dictionary records.
func (c *Client) DeleteMultiLivechatProviderDictionarys(ids []int64) error {
	return c.Delete(MultiLivechatProviderDictionaryModel, ids)
}

// GetMultiLivechatProviderDictionary gets multi_livechat.provider_dictionary existing record.
func (c *Client) GetMultiLivechatProviderDictionary(id int64) (*MultiLivechatProviderDictionary, error) {
	mps, err := c.GetMultiLivechatProviderDictionarys([]int64{id})
	if err != nil {
		return nil, err
	}
	if mps != nil && len(*mps) > 0 {
		return &((*mps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of multi_livechat.provider_dictionary not found", id)
}

// GetMultiLivechatProviderDictionarys gets multi_livechat.provider_dictionary existing records.
func (c *Client) GetMultiLivechatProviderDictionarys(ids []int64) (*MultiLivechatProviderDictionarys, error) {
	mps := &MultiLivechatProviderDictionarys{}
	if err := c.Read(MultiLivechatProviderDictionaryModel, ids, nil, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMultiLivechatProviderDictionary finds multi_livechat.provider_dictionary record by querying it with criteria.
func (c *Client) FindMultiLivechatProviderDictionary(criteria *Criteria) (*MultiLivechatProviderDictionary, error) {
	mps := &MultiLivechatProviderDictionarys{}
	if err := c.SearchRead(MultiLivechatProviderDictionaryModel, criteria, NewOptions().Limit(1), mps); err != nil {
		return nil, err
	}
	if mps != nil && len(*mps) > 0 {
		return &((*mps)[0]), nil
	}
	return nil, fmt.Errorf("multi_livechat.provider_dictionary was not found with criteria %v", criteria)
}

// FindMultiLivechatProviderDictionarys finds multi_livechat.provider_dictionary records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMultiLivechatProviderDictionarys(criteria *Criteria, options *Options) (*MultiLivechatProviderDictionarys, error) {
	mps := &MultiLivechatProviderDictionarys{}
	if err := c.SearchRead(MultiLivechatProviderDictionaryModel, criteria, options, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMultiLivechatProviderDictionaryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMultiLivechatProviderDictionaryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MultiLivechatProviderDictionaryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMultiLivechatProviderDictionaryId finds record id by querying it with criteria.
func (c *Client) FindMultiLivechatProviderDictionaryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MultiLivechatProviderDictionaryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("multi_livechat.provider_dictionary was not found with criteria %v and options %v", criteria, options)
}
