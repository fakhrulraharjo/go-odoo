package odoo

import (
	"fmt"
)

// ResUsersSettings represents res.users.settings model.
type ResUsersSettings struct {
	LastUpdate                           *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate                           *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                            *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName                          *String   `xmlrpc:"display_name,omptempty"`
	Id                                   *Int      `xmlrpc:"id,omptempty"`
	IsDiscussSidebarCategoryChannelOpen  *Bool     `xmlrpc:"is_discuss_sidebar_category_channel_open,omptempty"`
	IsDiscussSidebarCategoryChatOpen     *Bool     `xmlrpc:"is_discuss_sidebar_category_chat_open,omptempty"`
	IsDiscussSidebarCategoryEchoDemoOpen *Bool     `xmlrpc:"is_discuss_sidebar_category_echo_demo_open,omptempty"`
	IsDiscussSidebarCategoryWazapOpen    *Bool     `xmlrpc:"is_discuss_sidebar_category_wazap_open,omptempty"`
	PushToTalkKey                        *String   `xmlrpc:"push_to_talk_key,omptempty"`
	UsePushToTalk                        *Bool     `xmlrpc:"use_push_to_talk,omptempty"`
	UserId                               *Many2One `xmlrpc:"user_id,omptempty"`
	VoiceActiveDuration                  *Int      `xmlrpc:"voice_active_duration,omptempty"`
	VolumeSettingsIds                    *Relation `xmlrpc:"volume_settings_ids,omptempty"`
	WriteDate                            *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                             *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResUsersSettingss represents array of res.users.settings model.
type ResUsersSettingss []ResUsersSettings

// ResUsersSettingsModel is the odoo model name.
const ResUsersSettingsModel = "res.users.settings"

// Many2One convert ResUsersSettings to *Many2One.
func (rus *ResUsersSettings) Many2One() *Many2One {
	return NewMany2One(rus.Id.Get(), "")
}

// CreateResUsersSettings creates a new res.users.settings model and returns its id.
func (c *Client) CreateResUsersSettings(rus *ResUsersSettings) (int64, error) {
	return c.Create(ResUsersSettingsModel, rus)
}

// UpdateResUsersSettings updates an existing res.users.settings record.
func (c *Client) UpdateResUsersSettings(rus *ResUsersSettings) error {
	return c.UpdateResUsersSettingss([]int64{rus.Id.Get()}, rus)
}

// UpdateResUsersSettingss updates existing res.users.settings records.
// All records (represented by ids) will be updated by rus values.
func (c *Client) UpdateResUsersSettingss(ids []int64, rus *ResUsersSettings) error {
	return c.Update(ResUsersSettingsModel, ids, rus)
}

// DeleteResUsersSettings deletes an existing res.users.settings record.
func (c *Client) DeleteResUsersSettings(id int64) error {
	return c.DeleteResUsersSettingss([]int64{id})
}

// DeleteResUsersSettingss deletes existing res.users.settings records.
func (c *Client) DeleteResUsersSettingss(ids []int64) error {
	return c.Delete(ResUsersSettingsModel, ids)
}

// GetResUsersSettings gets res.users.settings existing record.
func (c *Client) GetResUsersSettings(id int64) (*ResUsersSettings, error) {
	russ, err := c.GetResUsersSettingss([]int64{id})
	if err != nil {
		return nil, err
	}
	if russ != nil && len(*russ) > 0 {
		return &((*russ)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users.settings not found", id)
}

// GetResUsersSettingss gets res.users.settings existing records.
func (c *Client) GetResUsersSettingss(ids []int64) (*ResUsersSettingss, error) {
	russ := &ResUsersSettingss{}
	if err := c.Read(ResUsersSettingsModel, ids, nil, russ); err != nil {
		return nil, err
	}
	return russ, nil
}

// FindResUsersSettings finds res.users.settings record by querying it with criteria.
func (c *Client) FindResUsersSettings(criteria *Criteria) (*ResUsersSettings, error) {
	russ := &ResUsersSettingss{}
	if err := c.SearchRead(ResUsersSettingsModel, criteria, NewOptions().Limit(1), russ); err != nil {
		return nil, err
	}
	if russ != nil && len(*russ) > 0 {
		return &((*russ)[0]), nil
	}
	return nil, fmt.Errorf("res.users.settings was not found with criteria %v", criteria)
}

// FindResUsersSettingss finds res.users.settings records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersSettingss(criteria *Criteria, options *Options) (*ResUsersSettingss, error) {
	russ := &ResUsersSettingss{}
	if err := c.SearchRead(ResUsersSettingsModel, criteria, options, russ); err != nil {
		return nil, err
	}
	return russ, nil
}

// FindResUsersSettingsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersSettingsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersSettingsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersSettingsId finds record id by querying it with criteria.
func (c *Client) FindResUsersSettingsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersSettingsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users.settings was not found with criteria %v and options %v", criteria, options)
}
