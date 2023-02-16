package odoo

import (
	"fmt"
)

// MultiLivechatConfigParameter represents multi_livechat.config_parameter model.
type MultiLivechatConfigParameter struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Key         *String   `xmlrpc:"key,omptempty"`
	Value       *String   `xmlrpc:"value,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MultiLivechatConfigParameters represents array of multi_livechat.config_parameter model.
type MultiLivechatConfigParameters []MultiLivechatConfigParameter

// MultiLivechatConfigParameterModel is the odoo model name.
const MultiLivechatConfigParameterModel = "multi_livechat.config_parameter"

// Many2One convert MultiLivechatConfigParameter to *Many2One.
func (mc *MultiLivechatConfigParameter) Many2One() *Many2One {
	return NewMany2One(mc.Id.Get(), "")
}

// CreateMultiLivechatConfigParameter creates a new multi_livechat.config_parameter model and returns its id.
func (c *Client) CreateMultiLivechatConfigParameter(mc *MultiLivechatConfigParameter) (int64, error) {
	return c.Create(MultiLivechatConfigParameterModel, mc)
}

// UpdateMultiLivechatConfigParameter updates an existing multi_livechat.config_parameter record.
func (c *Client) UpdateMultiLivechatConfigParameter(mc *MultiLivechatConfigParameter) error {
	return c.UpdateMultiLivechatConfigParameters([]int64{mc.Id.Get()}, mc)
}

// UpdateMultiLivechatConfigParameters updates existing multi_livechat.config_parameter records.
// All records (represented by ids) will be updated by mc values.
func (c *Client) UpdateMultiLivechatConfigParameters(ids []int64, mc *MultiLivechatConfigParameter) error {
	return c.Update(MultiLivechatConfigParameterModel, ids, mc)
}

// DeleteMultiLivechatConfigParameter deletes an existing multi_livechat.config_parameter record.
func (c *Client) DeleteMultiLivechatConfigParameter(id int64) error {
	return c.DeleteMultiLivechatConfigParameters([]int64{id})
}

// DeleteMultiLivechatConfigParameters deletes existing multi_livechat.config_parameter records.
func (c *Client) DeleteMultiLivechatConfigParameters(ids []int64) error {
	return c.Delete(MultiLivechatConfigParameterModel, ids)
}

// GetMultiLivechatConfigParameter gets multi_livechat.config_parameter existing record.
func (c *Client) GetMultiLivechatConfigParameter(id int64) (*MultiLivechatConfigParameter, error) {
	mcs, err := c.GetMultiLivechatConfigParameters([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of multi_livechat.config_parameter not found", id)
}

// GetMultiLivechatConfigParameters gets multi_livechat.config_parameter existing records.
func (c *Client) GetMultiLivechatConfigParameters(ids []int64) (*MultiLivechatConfigParameters, error) {
	mcs := &MultiLivechatConfigParameters{}
	if err := c.Read(MultiLivechatConfigParameterModel, ids, nil, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMultiLivechatConfigParameter finds multi_livechat.config_parameter record by querying it with criteria.
func (c *Client) FindMultiLivechatConfigParameter(criteria *Criteria) (*MultiLivechatConfigParameter, error) {
	mcs := &MultiLivechatConfigParameters{}
	if err := c.SearchRead(MultiLivechatConfigParameterModel, criteria, NewOptions().Limit(1), mcs); err != nil {
		return nil, err
	}
	if mcs != nil && len(*mcs) > 0 {
		return &((*mcs)[0]), nil
	}
	return nil, fmt.Errorf("multi_livechat.config_parameter was not found with criteria %v", criteria)
}

// FindMultiLivechatConfigParameters finds multi_livechat.config_parameter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMultiLivechatConfigParameters(criteria *Criteria, options *Options) (*MultiLivechatConfigParameters, error) {
	mcs := &MultiLivechatConfigParameters{}
	if err := c.SearchRead(MultiLivechatConfigParameterModel, criteria, options, mcs); err != nil {
		return nil, err
	}
	return mcs, nil
}

// FindMultiLivechatConfigParameterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMultiLivechatConfigParameterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MultiLivechatConfigParameterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMultiLivechatConfigParameterId finds record id by querying it with criteria.
func (c *Client) FindMultiLivechatConfigParameterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MultiLivechatConfigParameterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("multi_livechat.config_parameter was not found with criteria %v and options %v", criteria, options)
}
