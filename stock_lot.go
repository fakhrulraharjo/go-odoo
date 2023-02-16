package odoo

import (
	"fmt"
)

// StockLot represents stock.lot model.
type StockLot struct {
	LastUpdate                  *Time      `xmlrpc:"__last_update,omptempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omptempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omptempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omptempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DeliveryCount               *Int       `xmlrpc:"delivery_count,omptempty"`
	DeliveryIds                 *Relation  `xmlrpc:"delivery_ids,omptempty"`
	DisplayComplete             *Bool      `xmlrpc:"display_complete,omptempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omptempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omptempty"`
	Id                          *Int       `xmlrpc:"id,omptempty"`
	LastDeliveryPartnerId       *Many2One  `xmlrpc:"last_delivery_partner_id,omptempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omptempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omptempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omptempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omptempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageMainAttachmentId     *Many2One  `xmlrpc:"message_main_attachment_id,omptempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omptempty"`
	Name                        *String    `xmlrpc:"name,omptempty"`
	Note                        *String    `xmlrpc:"note,omptempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductQty                  *Float     `xmlrpc:"product_qty,omptempty"`
	ProductUomId                *Many2One  `xmlrpc:"product_uom_id,omptempty"`
	QuantIds                    *Relation  `xmlrpc:"quant_ids,omptempty"`
	Ref                         *String    `xmlrpc:"ref,omptempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockLots represents array of stock.lot model.
type StockLots []StockLot

// StockLotModel is the odoo model name.
const StockLotModel = "stock.lot"

// Many2One convert StockLot to *Many2One.
func (sl *StockLot) Many2One() *Many2One {
	return NewMany2One(sl.Id.Get(), "")
}

// CreateStockLot creates a new stock.lot model and returns its id.
func (c *Client) CreateStockLot(sl *StockLot) (int64, error) {
	return c.Create(StockLotModel, sl)
}

// UpdateStockLot updates an existing stock.lot record.
func (c *Client) UpdateStockLot(sl *StockLot) error {
	return c.UpdateStockLots([]int64{sl.Id.Get()}, sl)
}

// UpdateStockLots updates existing stock.lot records.
// All records (represented by ids) will be updated by sl values.
func (c *Client) UpdateStockLots(ids []int64, sl *StockLot) error {
	return c.Update(StockLotModel, ids, sl)
}

// DeleteStockLot deletes an existing stock.lot record.
func (c *Client) DeleteStockLot(id int64) error {
	return c.DeleteStockLots([]int64{id})
}

// DeleteStockLots deletes existing stock.lot records.
func (c *Client) DeleteStockLots(ids []int64) error {
	return c.Delete(StockLotModel, ids)
}

// GetStockLot gets stock.lot existing record.
func (c *Client) GetStockLot(id int64) (*StockLot, error) {
	sls, err := c.GetStockLots([]int64{id})
	if err != nil {
		return nil, err
	}
	if sls != nil && len(*sls) > 0 {
		return &((*sls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.lot not found", id)
}

// GetStockLots gets stock.lot existing records.
func (c *Client) GetStockLots(ids []int64) (*StockLots, error) {
	sls := &StockLots{}
	if err := c.Read(StockLotModel, ids, nil, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindStockLot finds stock.lot record by querying it with criteria.
func (c *Client) FindStockLot(criteria *Criteria) (*StockLot, error) {
	sls := &StockLots{}
	if err := c.SearchRead(StockLotModel, criteria, NewOptions().Limit(1), sls); err != nil {
		return nil, err
	}
	if sls != nil && len(*sls) > 0 {
		return &((*sls)[0]), nil
	}
	return nil, fmt.Errorf("stock.lot was not found with criteria %v", criteria)
}

// FindStockLots finds stock.lot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLots(criteria *Criteria, options *Options) (*StockLots, error) {
	sls := &StockLots{}
	if err := c.SearchRead(StockLotModel, criteria, options, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindStockLotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockLotIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockLotModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockLotId finds record id by querying it with criteria.
func (c *Client) FindStockLotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockLotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.lot was not found with criteria %v and options %v", criteria, options)
}
