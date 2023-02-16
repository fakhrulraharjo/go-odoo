package odoo

import (
	"fmt"
)

// StockMoveLine represents stock.move.line model.
type StockMoveLine struct {
	LastUpdate                 *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omptempty"`
	ConsumeLineIds             *Relation  `xmlrpc:"consume_line_ids,omptempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date                       *Time      `xmlrpc:"date,omptempty"`
	DescriptionPicking         *String    `xmlrpc:"description_picking,omptempty"`
	DisplayName                *String    `xmlrpc:"display_name,omptempty"`
	Id                         *Int       `xmlrpc:"id,omptempty"`
	IsInitialDemandEditable    *Bool      `xmlrpc:"is_initial_demand_editable,omptempty"`
	IsInventory                *Bool      `xmlrpc:"is_inventory,omptempty"`
	IsLocked                   *Bool      `xmlrpc:"is_locked,omptempty"`
	LocationDestId             *Many2One  `xmlrpc:"location_dest_id,omptempty"`
	LocationDestUsage          *Selection `xmlrpc:"location_dest_usage,omptempty"`
	LocationId                 *Many2One  `xmlrpc:"location_id,omptempty"`
	LocationUsage              *Selection `xmlrpc:"location_usage,omptempty"`
	LotId                      *Many2One  `xmlrpc:"lot_id,omptempty"`
	LotName                    *String    `xmlrpc:"lot_name,omptempty"`
	LotsVisible                *Bool      `xmlrpc:"lots_visible,omptempty"`
	MoveId                     *Many2One  `xmlrpc:"move_id,omptempty"`
	Origin                     *String    `xmlrpc:"origin,omptempty"`
	OwnerId                    *Many2One  `xmlrpc:"owner_id,omptempty"`
	PackageId                  *Many2One  `xmlrpc:"package_id,omptempty"`
	PackageLevelId             *Many2One  `xmlrpc:"package_level_id,omptempty"`
	PickingCode                *Selection `xmlrpc:"picking_code,omptempty"`
	PickingId                  *Many2One  `xmlrpc:"picking_id,omptempty"`
	PickingPartnerId           *Many2One  `xmlrpc:"picking_partner_id,omptempty"`
	PickingTypeEntirePacks     *Bool      `xmlrpc:"picking_type_entire_packs,omptempty"`
	PickingTypeId              *Many2One  `xmlrpc:"picking_type_id,omptempty"`
	PickingTypeUseCreateLots   *Bool      `xmlrpc:"picking_type_use_create_lots,omptempty"`
	PickingTypeUseExistingLots *Bool      `xmlrpc:"picking_type_use_existing_lots,omptempty"`
	ProduceLineIds             *Relation  `xmlrpc:"produce_line_ids,omptempty"`
	ProductCategoryName        *String    `xmlrpc:"product_category_name,omptempty"`
	ProductId                  *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductUomCategoryId       *Many2One  `xmlrpc:"product_uom_category_id,omptempty"`
	ProductUomId               *Many2One  `xmlrpc:"product_uom_id,omptempty"`
	QtyDone                    *Float     `xmlrpc:"qty_done,omptempty"`
	Reference                  *String    `xmlrpc:"reference,omptempty"`
	ReservedQty                *Float     `xmlrpc:"reserved_qty,omptempty"`
	ReservedUomQty             *Float     `xmlrpc:"reserved_uom_qty,omptempty"`
	ResultPackageId            *Many2One  `xmlrpc:"result_package_id,omptempty"`
	State                      *Selection `xmlrpc:"state,omptempty"`
	Tracking                   *Selection `xmlrpc:"tracking,omptempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockMoveLines represents array of stock.move.line model.
type StockMoveLines []StockMoveLine

// StockMoveLineModel is the odoo model name.
const StockMoveLineModel = "stock.move.line"

// Many2One convert StockMoveLine to *Many2One.
func (sml *StockMoveLine) Many2One() *Many2One {
	return NewMany2One(sml.Id.Get(), "")
}

// CreateStockMoveLine creates a new stock.move.line model and returns its id.
func (c *Client) CreateStockMoveLine(sml *StockMoveLine) (int64, error) {
	return c.Create(StockMoveLineModel, sml)
}

// UpdateStockMoveLine updates an existing stock.move.line record.
func (c *Client) UpdateStockMoveLine(sml *StockMoveLine) error {
	return c.UpdateStockMoveLines([]int64{sml.Id.Get()}, sml)
}

// UpdateStockMoveLines updates existing stock.move.line records.
// All records (represented by ids) will be updated by sml values.
func (c *Client) UpdateStockMoveLines(ids []int64, sml *StockMoveLine) error {
	return c.Update(StockMoveLineModel, ids, sml)
}

// DeleteStockMoveLine deletes an existing stock.move.line record.
func (c *Client) DeleteStockMoveLine(id int64) error {
	return c.DeleteStockMoveLines([]int64{id})
}

// DeleteStockMoveLines deletes existing stock.move.line records.
func (c *Client) DeleteStockMoveLines(ids []int64) error {
	return c.Delete(StockMoveLineModel, ids)
}

// GetStockMoveLine gets stock.move.line existing record.
func (c *Client) GetStockMoveLine(id int64) (*StockMoveLine, error) {
	smls, err := c.GetStockMoveLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if smls != nil && len(*smls) > 0 {
		return &((*smls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.move.line not found", id)
}

// GetStockMoveLines gets stock.move.line existing records.
func (c *Client) GetStockMoveLines(ids []int64) (*StockMoveLines, error) {
	smls := &StockMoveLines{}
	if err := c.Read(StockMoveLineModel, ids, nil, smls); err != nil {
		return nil, err
	}
	return smls, nil
}

// FindStockMoveLine finds stock.move.line record by querying it with criteria.
func (c *Client) FindStockMoveLine(criteria *Criteria) (*StockMoveLine, error) {
	smls := &StockMoveLines{}
	if err := c.SearchRead(StockMoveLineModel, criteria, NewOptions().Limit(1), smls); err != nil {
		return nil, err
	}
	if smls != nil && len(*smls) > 0 {
		return &((*smls)[0]), nil
	}
	return nil, fmt.Errorf("stock.move.line was not found with criteria %v", criteria)
}

// FindStockMoveLines finds stock.move.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockMoveLines(criteria *Criteria, options *Options) (*StockMoveLines, error) {
	smls := &StockMoveLines{}
	if err := c.SearchRead(StockMoveLineModel, criteria, options, smls); err != nil {
		return nil, err
	}
	return smls, nil
}

// FindStockMoveLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockMoveLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockMoveLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockMoveLineId finds record id by querying it with criteria.
func (c *Client) FindStockMoveLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockMoveLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.move.line was not found with criteria %v and options %v", criteria, options)
}
