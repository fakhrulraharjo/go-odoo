package odoo

import (
	"fmt"
)

// PublisherWarrantyContract represents publisher_warranty.contract model.
type PublisherWarrantyContract struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// PublisherWarrantyContracts represents array of publisher_warranty.contract model.
type PublisherWarrantyContracts []PublisherWarrantyContract

// PublisherWarrantyContractModel is the odoo model name.
const PublisherWarrantyContractModel = "publisher_warranty.contract"

// Many2One convert PublisherWarrantyContract to *Many2One.
func (pc *PublisherWarrantyContract) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreatePublisherWarrantyContract creates a new publisher_warranty.contract model and returns its id.
func (c *Client) CreatePublisherWarrantyContract(pc *PublisherWarrantyContract) (int64, error) {
	return c.Create(PublisherWarrantyContractModel, pc)
}

// UpdatePublisherWarrantyContract updates an existing publisher_warranty.contract record.
func (c *Client) UpdatePublisherWarrantyContract(pc *PublisherWarrantyContract) error {
	return c.UpdatePublisherWarrantyContracts([]int64{pc.Id.Get()}, pc)
}

// UpdatePublisherWarrantyContracts updates existing publisher_warranty.contract records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdatePublisherWarrantyContracts(ids []int64, pc *PublisherWarrantyContract) error {
	return c.Update(PublisherWarrantyContractModel, ids, pc)
}

// DeletePublisherWarrantyContract deletes an existing publisher_warranty.contract record.
func (c *Client) DeletePublisherWarrantyContract(id int64) error {
	return c.DeletePublisherWarrantyContracts([]int64{id})
}

// DeletePublisherWarrantyContracts deletes existing publisher_warranty.contract records.
func (c *Client) DeletePublisherWarrantyContracts(ids []int64) error {
	return c.Delete(PublisherWarrantyContractModel, ids)
}

// GetPublisherWarrantyContract gets publisher_warranty.contract existing record.
func (c *Client) GetPublisherWarrantyContract(id int64) (*PublisherWarrantyContract, error) {
	pcs, err := c.GetPublisherWarrantyContracts([]int64{id})
	if err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of publisher_warranty.contract not found", id)
}

// GetPublisherWarrantyContracts gets publisher_warranty.contract existing records.
func (c *Client) GetPublisherWarrantyContracts(ids []int64) (*PublisherWarrantyContracts, error) {
	pcs := &PublisherWarrantyContracts{}
	if err := c.Read(PublisherWarrantyContractModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPublisherWarrantyContract finds publisher_warranty.contract record by querying it with criteria.
func (c *Client) FindPublisherWarrantyContract(criteria *Criteria) (*PublisherWarrantyContract, error) {
	pcs := &PublisherWarrantyContracts{}
	if err := c.SearchRead(PublisherWarrantyContractModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("publisher_warranty.contract was not found with criteria %v", criteria)
}

// FindPublisherWarrantyContracts finds publisher_warranty.contract records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPublisherWarrantyContracts(criteria *Criteria, options *Options) (*PublisherWarrantyContracts, error) {
	pcs := &PublisherWarrantyContracts{}
	if err := c.SearchRead(PublisherWarrantyContractModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPublisherWarrantyContractIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPublisherWarrantyContractIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PublisherWarrantyContractModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPublisherWarrantyContractId finds record id by querying it with criteria.
func (c *Client) FindPublisherWarrantyContractId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PublisherWarrantyContractModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("publisher_warranty.contract was not found with criteria %v and options %v", criteria, options)
}
