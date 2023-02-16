package odoo

import (
	"fmt"
)

// MailChannelRtcSession represents mail.channel.rtc.session model.
type MailChannelRtcSession struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	ChannelId         *Many2One `xmlrpc:"channel_id,omptempty"`
	ChannelMemberId   *Many2One `xmlrpc:"channel_member_id,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	GuestId           *Many2One `xmlrpc:"guest_id,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	IsCameraOn        *Bool     `xmlrpc:"is_camera_on,omptempty"`
	IsDeaf            *Bool     `xmlrpc:"is_deaf,omptempty"`
	IsMuted           *Bool     `xmlrpc:"is_muted,omptempty"`
	IsScreenSharingOn *Bool     `xmlrpc:"is_screen_sharing_on,omptempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailChannelRtcSessions represents array of mail.channel.rtc.session model.
type MailChannelRtcSessions []MailChannelRtcSession

// MailChannelRtcSessionModel is the odoo model name.
const MailChannelRtcSessionModel = "mail.channel.rtc.session"

// Many2One convert MailChannelRtcSession to *Many2One.
func (mcrs *MailChannelRtcSession) Many2One() *Many2One {
	return NewMany2One(mcrs.Id.Get(), "")
}

// CreateMailChannelRtcSession creates a new mail.channel.rtc.session model and returns its id.
func (c *Client) CreateMailChannelRtcSession(mcrs *MailChannelRtcSession) (int64, error) {
	return c.Create(MailChannelRtcSessionModel, mcrs)
}

// UpdateMailChannelRtcSession updates an existing mail.channel.rtc.session record.
func (c *Client) UpdateMailChannelRtcSession(mcrs *MailChannelRtcSession) error {
	return c.UpdateMailChannelRtcSessions([]int64{mcrs.Id.Get()}, mcrs)
}

// UpdateMailChannelRtcSessions updates existing mail.channel.rtc.session records.
// All records (represented by ids) will be updated by mcrs values.
func (c *Client) UpdateMailChannelRtcSessions(ids []int64, mcrs *MailChannelRtcSession) error {
	return c.Update(MailChannelRtcSessionModel, ids, mcrs)
}

// DeleteMailChannelRtcSession deletes an existing mail.channel.rtc.session record.
func (c *Client) DeleteMailChannelRtcSession(id int64) error {
	return c.DeleteMailChannelRtcSessions([]int64{id})
}

// DeleteMailChannelRtcSessions deletes existing mail.channel.rtc.session records.
func (c *Client) DeleteMailChannelRtcSessions(ids []int64) error {
	return c.Delete(MailChannelRtcSessionModel, ids)
}

// GetMailChannelRtcSession gets mail.channel.rtc.session existing record.
func (c *Client) GetMailChannelRtcSession(id int64) (*MailChannelRtcSession, error) {
	mcrss, err := c.GetMailChannelRtcSessions([]int64{id})
	if err != nil {
		return nil, err
	}
	if mcrss != nil && len(*mcrss) > 0 {
		return &((*mcrss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.channel.rtc.session not found", id)
}

// GetMailChannelRtcSessions gets mail.channel.rtc.session existing records.
func (c *Client) GetMailChannelRtcSessions(ids []int64) (*MailChannelRtcSessions, error) {
	mcrss := &MailChannelRtcSessions{}
	if err := c.Read(MailChannelRtcSessionModel, ids, nil, mcrss); err != nil {
		return nil, err
	}
	return mcrss, nil
}

// FindMailChannelRtcSession finds mail.channel.rtc.session record by querying it with criteria.
func (c *Client) FindMailChannelRtcSession(criteria *Criteria) (*MailChannelRtcSession, error) {
	mcrss := &MailChannelRtcSessions{}
	if err := c.SearchRead(MailChannelRtcSessionModel, criteria, NewOptions().Limit(1), mcrss); err != nil {
		return nil, err
	}
	if mcrss != nil && len(*mcrss) > 0 {
		return &((*mcrss)[0]), nil
	}
	return nil, fmt.Errorf("mail.channel.rtc.session was not found with criteria %v", criteria)
}

// FindMailChannelRtcSessions finds mail.channel.rtc.session records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelRtcSessions(criteria *Criteria, options *Options) (*MailChannelRtcSessions, error) {
	mcrss := &MailChannelRtcSessions{}
	if err := c.SearchRead(MailChannelRtcSessionModel, criteria, options, mcrss); err != nil {
		return nil, err
	}
	return mcrss, nil
}

// FindMailChannelRtcSessionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelRtcSessionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailChannelRtcSessionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailChannelRtcSessionId finds record id by querying it with criteria.
func (c *Client) FindMailChannelRtcSessionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailChannelRtcSessionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.channel.rtc.session was not found with criteria %v and options %v", criteria, options)
}
