package odoo

import (
	"fmt"
)

// MailGuest represents mail.guest model.
type MailGuest struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	AccessToken *String    `xmlrpc:"access_token,omptempty"`
	Avatar1024  *String    `xmlrpc:"avatar_1024,omptempty"`
	Avatar128   *String    `xmlrpc:"avatar_128,omptempty"`
	Avatar1920  *String    `xmlrpc:"avatar_1920,omptempty"`
	Avatar256   *String    `xmlrpc:"avatar_256,omptempty"`
	Avatar512   *String    `xmlrpc:"avatar_512,omptempty"`
	ChannelIds  *Relation  `xmlrpc:"channel_ids,omptempty"`
	CountryId   *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	ImStatus    *String    `xmlrpc:"im_status,omptempty"`
	Image1024   *String    `xmlrpc:"image_1024,omptempty"`
	Image128    *String    `xmlrpc:"image_128,omptempty"`
	Image1920   *String    `xmlrpc:"image_1920,omptempty"`
	Image256    *String    `xmlrpc:"image_256,omptempty"`
	Image512    *String    `xmlrpc:"image_512,omptempty"`
	Lang        *Selection `xmlrpc:"lang,omptempty"`
	Name        *String    `xmlrpc:"name,omptempty"`
	Timezone    *Selection `xmlrpc:"timezone,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailGuests represents array of mail.guest model.
type MailGuests []MailGuest

// MailGuestModel is the odoo model name.
const MailGuestModel = "mail.guest"

// Many2One convert MailGuest to *Many2One.
func (mg *MailGuest) Many2One() *Many2One {
	return NewMany2One(mg.Id.Get(), "")
}

// CreateMailGuest creates a new mail.guest model and returns its id.
func (c *Client) CreateMailGuest(mg *MailGuest) (int64, error) {
	return c.Create(MailGuestModel, mg)
}

// UpdateMailGuest updates an existing mail.guest record.
func (c *Client) UpdateMailGuest(mg *MailGuest) error {
	return c.UpdateMailGuests([]int64{mg.Id.Get()}, mg)
}

// UpdateMailGuests updates existing mail.guest records.
// All records (represented by ids) will be updated by mg values.
func (c *Client) UpdateMailGuests(ids []int64, mg *MailGuest) error {
	return c.Update(MailGuestModel, ids, mg)
}

// DeleteMailGuest deletes an existing mail.guest record.
func (c *Client) DeleteMailGuest(id int64) error {
	return c.DeleteMailGuests([]int64{id})
}

// DeleteMailGuests deletes existing mail.guest records.
func (c *Client) DeleteMailGuests(ids []int64) error {
	return c.Delete(MailGuestModel, ids)
}

// GetMailGuest gets mail.guest existing record.
func (c *Client) GetMailGuest(id int64) (*MailGuest, error) {
	mgs, err := c.GetMailGuests([]int64{id})
	if err != nil {
		return nil, err
	}
	if mgs != nil && len(*mgs) > 0 {
		return &((*mgs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.guest not found", id)
}

// GetMailGuests gets mail.guest existing records.
func (c *Client) GetMailGuests(ids []int64) (*MailGuests, error) {
	mgs := &MailGuests{}
	if err := c.Read(MailGuestModel, ids, nil, mgs); err != nil {
		return nil, err
	}
	return mgs, nil
}

// FindMailGuest finds mail.guest record by querying it with criteria.
func (c *Client) FindMailGuest(criteria *Criteria) (*MailGuest, error) {
	mgs := &MailGuests{}
	if err := c.SearchRead(MailGuestModel, criteria, NewOptions().Limit(1), mgs); err != nil {
		return nil, err
	}
	if mgs != nil && len(*mgs) > 0 {
		return &((*mgs)[0]), nil
	}
	return nil, fmt.Errorf("mail.guest was not found with criteria %v", criteria)
}

// FindMailGuests finds mail.guest records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailGuests(criteria *Criteria, options *Options) (*MailGuests, error) {
	mgs := &MailGuests{}
	if err := c.SearchRead(MailGuestModel, criteria, options, mgs); err != nil {
		return nil, err
	}
	return mgs, nil
}

// FindMailGuestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailGuestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailGuestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailGuestId finds record id by querying it with criteria.
func (c *Client) FindMailGuestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailGuestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.guest was not found with criteria %v and options %v", criteria, options)
}
