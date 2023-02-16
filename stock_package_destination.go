package odoo

import (
	"fmt"
)

// StockPackageDestination represents stock.package.destination model.
type StockPackageDestination struct {
	LastUpdate       *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName      *String   `xmlrpc:"display_name,omptempty"`
	FilteredLocation *Relation `xmlrpc:"filtered_location,omptempty"`
	Id               *Int      `xmlrpc:"id,omptempty"`
	LocationDestId   *Many2One `xmlrpc:"location_dest_id,omptempty"`
	MoveLineIds      *Relation `xmlrpc:"move_line_ids,omptempty"`
	PickingId        *Many2One `xmlrpc:"picking_id,omptempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omptempty"`
}

// StockPackageDestinations represents array of stock.package.destination model.
type StockPackageDestinations []StockPackageDestination

// StockPackageDestinationModel is the odoo model name.
const StockPackageDestinationModel = "stock.package.destination"

// Many2One convert StockPackageDestination to *Many2One.
func (spd *StockPackageDestination) Many2One() *Many2One {
	return NewMany2One(spd.Id.Get(), "")
}

// CreateStockPackageDestination creates a new stock.package.destination model and returns its id.
func (c *Client) CreateStockPackageDestination(spd *StockPackageDestination) (int64, error) {
	return c.Create(StockPackageDestinationModel, spd)
}

// UpdateStockPackageDestination updates an existing stock.package.destination record.
func (c *Client) UpdateStockPackageDestination(spd *StockPackageDestination) error {
	return c.UpdateStockPackageDestinations([]int64{spd.Id.Get()}, spd)
}

// UpdateStockPackageDestinations updates existing stock.package.destination records.
// All records (represented by ids) will be updated by spd values.
func (c *Client) UpdateStockPackageDestinations(ids []int64, spd *StockPackageDestination) error {
	return c.Update(StockPackageDestinationModel, ids, spd)
}

// DeleteStockPackageDestination deletes an existing stock.package.destination record.
func (c *Client) DeleteStockPackageDestination(id int64) error {
	return c.DeleteStockPackageDestinations([]int64{id})
}

// DeleteStockPackageDestinations deletes existing stock.package.destination records.
func (c *Client) DeleteStockPackageDestinations(ids []int64) error {
	return c.Delete(StockPackageDestinationModel, ids)
}

// GetStockPackageDestination gets stock.package.destination existing record.
func (c *Client) GetStockPackageDestination(id int64) (*StockPackageDestination, error) {
	spds, err := c.GetStockPackageDestinations([]int64{id})
	if err != nil {
		return nil, err
	}
	if spds != nil && len(*spds) > 0 {
		return &((*spds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.package.destination not found", id)
}

// GetStockPackageDestinations gets stock.package.destination existing records.
func (c *Client) GetStockPackageDestinations(ids []int64) (*StockPackageDestinations, error) {
	spds := &StockPackageDestinations{}
	if err := c.Read(StockPackageDestinationModel, ids, nil, spds); err != nil {
		return nil, err
	}
	return spds, nil
}

// FindStockPackageDestination finds stock.package.destination record by querying it with criteria.
func (c *Client) FindStockPackageDestination(criteria *Criteria) (*StockPackageDestination, error) {
	spds := &StockPackageDestinations{}
	if err := c.SearchRead(StockPackageDestinationModel, criteria, NewOptions().Limit(1), spds); err != nil {
		return nil, err
	}
	if spds != nil && len(*spds) > 0 {
		return &((*spds)[0]), nil
	}
	return nil, fmt.Errorf("stock.package.destination was not found with criteria %v", criteria)
}

// FindStockPackageDestinations finds stock.package.destination records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageDestinations(criteria *Criteria, options *Options) (*StockPackageDestinations, error) {
	spds := &StockPackageDestinations{}
	if err := c.SearchRead(StockPackageDestinationModel, criteria, options, spds); err != nil {
		return nil, err
	}
	return spds, nil
}

// FindStockPackageDestinationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPackageDestinationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockPackageDestinationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockPackageDestinationId finds record id by querying it with criteria.
func (c *Client) FindStockPackageDestinationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPackageDestinationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.package.destination was not found with criteria %v and options %v", criteria, options)
}
