package odoo

import (
	"fmt"
)

// MailMessageSchedule represents mail.message.schedule model.
type MailMessageSchedule struct {
	LastUpdate             *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate             *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName            *String   `xmlrpc:"display_name,omptempty"`
	Id                     *Int      `xmlrpc:"id,omptempty"`
	MailMessageId          *Many2One `xmlrpc:"mail_message_id,omptempty"`
	NotificationParameters *String   `xmlrpc:"notification_parameters,omptempty"`
	ScheduledDatetime      *Time     `xmlrpc:"scheduled_datetime,omptempty"`
	WriteDate              *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailMessageSchedules represents array of mail.message.schedule model.
type MailMessageSchedules []MailMessageSchedule

// MailMessageScheduleModel is the odoo model name.
const MailMessageScheduleModel = "mail.message.schedule"

// Many2One convert MailMessageSchedule to *Many2One.
func (mms *MailMessageSchedule) Many2One() *Many2One {
	return NewMany2One(mms.Id.Get(), "")
}

// CreateMailMessageSchedule creates a new mail.message.schedule model and returns its id.
func (c *Client) CreateMailMessageSchedule(mms *MailMessageSchedule) (int64, error) {
	return c.Create(MailMessageScheduleModel, mms)
}

// UpdateMailMessageSchedule updates an existing mail.message.schedule record.
func (c *Client) UpdateMailMessageSchedule(mms *MailMessageSchedule) error {
	return c.UpdateMailMessageSchedules([]int64{mms.Id.Get()}, mms)
}

// UpdateMailMessageSchedules updates existing mail.message.schedule records.
// All records (represented by ids) will be updated by mms values.
func (c *Client) UpdateMailMessageSchedules(ids []int64, mms *MailMessageSchedule) error {
	return c.Update(MailMessageScheduleModel, ids, mms)
}

// DeleteMailMessageSchedule deletes an existing mail.message.schedule record.
func (c *Client) DeleteMailMessageSchedule(id int64) error {
	return c.DeleteMailMessageSchedules([]int64{id})
}

// DeleteMailMessageSchedules deletes existing mail.message.schedule records.
func (c *Client) DeleteMailMessageSchedules(ids []int64) error {
	return c.Delete(MailMessageScheduleModel, ids)
}

// GetMailMessageSchedule gets mail.message.schedule existing record.
func (c *Client) GetMailMessageSchedule(id int64) (*MailMessageSchedule, error) {
	mmss, err := c.GetMailMessageSchedules([]int64{id})
	if err != nil {
		return nil, err
	}
	if mmss != nil && len(*mmss) > 0 {
		return &((*mmss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.message.schedule not found", id)
}

// GetMailMessageSchedules gets mail.message.schedule existing records.
func (c *Client) GetMailMessageSchedules(ids []int64) (*MailMessageSchedules, error) {
	mmss := &MailMessageSchedules{}
	if err := c.Read(MailMessageScheduleModel, ids, nil, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMessageSchedule finds mail.message.schedule record by querying it with criteria.
func (c *Client) FindMailMessageSchedule(criteria *Criteria) (*MailMessageSchedule, error) {
	mmss := &MailMessageSchedules{}
	if err := c.SearchRead(MailMessageScheduleModel, criteria, NewOptions().Limit(1), mmss); err != nil {
		return nil, err
	}
	if mmss != nil && len(*mmss) > 0 {
		return &((*mmss)[0]), nil
	}
	return nil, fmt.Errorf("mail.message.schedule was not found with criteria %v", criteria)
}

// FindMailMessageSchedules finds mail.message.schedule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageSchedules(criteria *Criteria, options *Options) (*MailMessageSchedules, error) {
	mmss := &MailMessageSchedules{}
	if err := c.SearchRead(MailMessageScheduleModel, criteria, options, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMessageScheduleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageScheduleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailMessageScheduleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMessageScheduleId finds record id by querying it with criteria.
func (c *Client) FindMailMessageScheduleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMessageScheduleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.message.schedule was not found with criteria %v and options %v", criteria, options)
}
