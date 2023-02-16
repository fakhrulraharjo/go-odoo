package odoo

import (
	"fmt"
)

// StockMove represents stock.move model.
type StockMove struct {
	LastUpdate               *Time      `xmlrpc:"__last_update,omptempty"`
	AccountMoveIds           *Relation  `xmlrpc:"account_move_ids,omptempty"`
	Additional               *Bool      `xmlrpc:"additional,omptempty"`
	AnalyticAccountLineId    *Many2One  `xmlrpc:"analytic_account_line_id,omptempty"`
	Availability             *Float     `xmlrpc:"availability,omptempty"`
	CompanyId                *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate               *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date                     *Time      `xmlrpc:"date,omptempty"`
	DateDeadline             *Time      `xmlrpc:"date_deadline,omptempty"`
	DelayAlertDate           *Time      `xmlrpc:"delay_alert_date,omptempty"`
	DescriptionPicking       *String    `xmlrpc:"description_picking,omptempty"`
	DisplayAssignSerial      *Bool      `xmlrpc:"display_assign_serial,omptempty"`
	DisplayClearSerial       *Bool      `xmlrpc:"display_clear_serial,omptempty"`
	DisplayName              *String    `xmlrpc:"display_name,omptempty"`
	ForecastAvailability     *Float     `xmlrpc:"forecast_availability,omptempty"`
	ForecastExpectedDate     *Time      `xmlrpc:"forecast_expected_date,omptempty"`
	FromImmediateTransfer    *Bool      `xmlrpc:"from_immediate_transfer,omptempty"`
	GroupId                  *Many2One  `xmlrpc:"group_id,omptempty"`
	HasTracking              *Selection `xmlrpc:"has_tracking,omptempty"`
	Id                       *Int       `xmlrpc:"id,omptempty"`
	IsInitialDemandEditable  *Bool      `xmlrpc:"is_initial_demand_editable,omptempty"`
	IsInventory              *Bool      `xmlrpc:"is_inventory,omptempty"`
	IsLocked                 *Bool      `xmlrpc:"is_locked,omptempty"`
	IsQuantityDoneEditable   *Bool      `xmlrpc:"is_quantity_done_editable,omptempty"`
	LocationDestId           *Many2One  `xmlrpc:"location_dest_id,omptempty"`
	LocationDestUsage        *Selection `xmlrpc:"location_dest_usage,omptempty"`
	LocationId               *Many2One  `xmlrpc:"location_id,omptempty"`
	LocationUsage            *Selection `xmlrpc:"location_usage,omptempty"`
	LotIds                   *Relation  `xmlrpc:"lot_ids,omptempty"`
	MoveDestIds              *Relation  `xmlrpc:"move_dest_ids,omptempty"`
	MoveLineIds              *Relation  `xmlrpc:"move_line_ids,omptempty"`
	MoveLineNosuggestIds     *Relation  `xmlrpc:"move_line_nosuggest_ids,omptempty"`
	MoveLinesCount           *Int       `xmlrpc:"move_lines_count,omptempty"`
	MoveOrigIds              *Relation  `xmlrpc:"move_orig_ids,omptempty"`
	Name                     *String    `xmlrpc:"name,omptempty"`
	NextSerial               *String    `xmlrpc:"next_serial,omptempty"`
	NextSerialCount          *Int       `xmlrpc:"next_serial_count,omptempty"`
	OrderpointId             *Many2One  `xmlrpc:"orderpoint_id,omptempty"`
	Origin                   *String    `xmlrpc:"origin,omptempty"`
	OriginReturnedMoveId     *Many2One  `xmlrpc:"origin_returned_move_id,omptempty"`
	PackageLevelId           *Many2One  `xmlrpc:"package_level_id,omptempty"`
	PartnerId                *Many2One  `xmlrpc:"partner_id,omptempty"`
	PickingCode              *Selection `xmlrpc:"picking_code,omptempty"`
	PickingId                *Many2One  `xmlrpc:"picking_id,omptempty"`
	PickingTypeEntirePacks   *Bool      `xmlrpc:"picking_type_entire_packs,omptempty"`
	PickingTypeId            *Many2One  `xmlrpc:"picking_type_id,omptempty"`
	PriceUnit                *Float     `xmlrpc:"price_unit,omptempty"`
	Priority                 *Selection `xmlrpc:"priority,omptempty"`
	ProcureMethod            *Selection `xmlrpc:"procure_method,omptempty"`
	ProductId                *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductPackagingId       *Many2One  `xmlrpc:"product_packaging_id,omptempty"`
	ProductQty               *Float     `xmlrpc:"product_qty,omptempty"`
	ProductTmplId            *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	ProductType              *Selection `xmlrpc:"product_type,omptempty"`
	ProductUom               *Many2One  `xmlrpc:"product_uom,omptempty"`
	ProductUomCategoryId     *Many2One  `xmlrpc:"product_uom_category_id,omptempty"`
	ProductUomQty            *Float     `xmlrpc:"product_uom_qty,omptempty"`
	PropagateCancel          *Bool      `xmlrpc:"propagate_cancel,omptempty"`
	QuantityDone             *Float     `xmlrpc:"quantity_done,omptempty"`
	Reference                *String    `xmlrpc:"reference,omptempty"`
	ReservationDate          *Time      `xmlrpc:"reservation_date,omptempty"`
	ReservedAvailability     *Float     `xmlrpc:"reserved_availability,omptempty"`
	RestrictPartnerId        *Many2One  `xmlrpc:"restrict_partner_id,omptempty"`
	ReturnedMoveIds          *Relation  `xmlrpc:"returned_move_ids,omptempty"`
	RouteIds                 *Relation  `xmlrpc:"route_ids,omptempty"`
	RuleId                   *Many2One  `xmlrpc:"rule_id,omptempty"`
	ScrapIds                 *Relation  `xmlrpc:"scrap_ids,omptempty"`
	Scrapped                 *Bool      `xmlrpc:"scrapped,omptempty"`
	Sequence                 *Int       `xmlrpc:"sequence,omptempty"`
	ShowDetailsVisible       *Bool      `xmlrpc:"show_details_visible,omptempty"`
	ShowOperations           *Bool      `xmlrpc:"show_operations,omptempty"`
	ShowReservedAvailability *Bool      `xmlrpc:"show_reserved_availability,omptempty"`
	State                    *Selection `xmlrpc:"state,omptempty"`
	StockValuationLayerIds   *Relation  `xmlrpc:"stock_valuation_layer_ids,omptempty"`
	ToRefund                 *Bool      `xmlrpc:"to_refund,omptempty"`
	WarehouseId              *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WriteDate                *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                 *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockMoves represents array of stock.move model.
type StockMoves []StockMove

// StockMoveModel is the odoo model name.
const StockMoveModel = "stock.move"

// Many2One convert StockMove to *Many2One.
func (sm *StockMove) Many2One() *Many2One {
	return NewMany2One(sm.Id.Get(), "")
}

// CreateStockMove creates a new stock.move model and returns its id.
func (c *Client) CreateStockMove(sm *StockMove) (int64, error) {
	return c.Create(StockMoveModel, sm)
}

// UpdateStockMove updates an existing stock.move record.
func (c *Client) UpdateStockMove(sm *StockMove) error {
	return c.UpdateStockMoves([]int64{sm.Id.Get()}, sm)
}

// UpdateStockMoves updates existing stock.move records.
// All records (represented by ids) will be updated by sm values.
func (c *Client) UpdateStockMoves(ids []int64, sm *StockMove) error {
	return c.Update(StockMoveModel, ids, sm)
}

// DeleteStockMove deletes an existing stock.move record.
func (c *Client) DeleteStockMove(id int64) error {
	return c.DeleteStockMoves([]int64{id})
}

// DeleteStockMoves deletes existing stock.move records.
func (c *Client) DeleteStockMoves(ids []int64) error {
	return c.Delete(StockMoveModel, ids)
}

// GetStockMove gets stock.move existing record.
func (c *Client) GetStockMove(id int64) (*StockMove, error) {
	sms, err := c.GetStockMoves([]int64{id})
	if err != nil {
		return nil, err
	}
	if sms != nil && len(*sms) > 0 {
		return &((*sms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.move not found", id)
}

// GetStockMoves gets stock.move existing records.
func (c *Client) GetStockMoves(ids []int64) (*StockMoves, error) {
	sms := &StockMoves{}
	if err := c.Read(StockMoveModel, ids, nil, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindStockMove finds stock.move record by querying it with criteria.
func (c *Client) FindStockMove(criteria *Criteria) (*StockMove, error) {
	sms := &StockMoves{}
	if err := c.SearchRead(StockMoveModel, criteria, NewOptions().Limit(1), sms); err != nil {
		return nil, err
	}
	if sms != nil && len(*sms) > 0 {
		return &((*sms)[0]), nil
	}
	return nil, fmt.Errorf("stock.move was not found with criteria %v", criteria)
}

// FindStockMoves finds stock.move records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockMoves(criteria *Criteria, options *Options) (*StockMoves, error) {
	sms := &StockMoves{}
	if err := c.SearchRead(StockMoveModel, criteria, options, sms); err != nil {
		return nil, err
	}
	return sms, nil
}

// FindStockMoveIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockMoveIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockMoveModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockMoveId finds record id by querying it with criteria.
func (c *Client) FindStockMoveId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockMoveModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.move was not found with criteria %v and options %v", criteria, options)
}
