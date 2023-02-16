package odoo

import (
	"fmt"
)

// ResUsersSettingsVolumes represents res.users.settings.volumes model.
type ResUsersSettingsVolumes struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	GuestId       *Many2One `xmlrpc:"guest_id,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	PartnerId     *Many2One `xmlrpc:"partner_id,omptempty"`
	UserSettingId *Many2One `xmlrpc:"user_setting_id,omptempty"`
	Volume        *Float    `xmlrpc:"volume,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// ResUsersSettingsVolumess represents array of res.users.settings.volumes model.
type ResUsersSettingsVolumess []ResUsersSettingsVolumes

// ResUsersSettingsVolumesModel is the odoo model name.
const ResUsersSettingsVolumesModel = "res.users.settings.volumes"

// Many2One convert ResUsersSettingsVolumes to *Many2One.
func (rusv *ResUsersSettingsVolumes) Many2One() *Many2One {
	return NewMany2One(rusv.Id.Get(), "")
}

// CreateResUsersSettingsVolumes creates a new res.users.settings.volumes model and returns its id.
func (c *Client) CreateResUsersSettingsVolumes(rusv *ResUsersSettingsVolumes) (int64, error) {
	return c.Create(ResUsersSettingsVolumesModel, rusv)
}

// UpdateResUsersSettingsVolumes updates an existing res.users.settings.volumes record.
func (c *Client) UpdateResUsersSettingsVolumes(rusv *ResUsersSettingsVolumes) error {
	return c.UpdateResUsersSettingsVolumess([]int64{rusv.Id.Get()}, rusv)
}

// UpdateResUsersSettingsVolumess updates existing res.users.settings.volumes records.
// All records (represented by ids) will be updated by rusv values.
func (c *Client) UpdateResUsersSettingsVolumess(ids []int64, rusv *ResUsersSettingsVolumes) error {
	return c.Update(ResUsersSettingsVolumesModel, ids, rusv)
}

// DeleteResUsersSettingsVolumes deletes an existing res.users.settings.volumes record.
func (c *Client) DeleteResUsersSettingsVolumes(id int64) error {
	return c.DeleteResUsersSettingsVolumess([]int64{id})
}

// DeleteResUsersSettingsVolumess deletes existing res.users.settings.volumes records.
func (c *Client) DeleteResUsersSettingsVolumess(ids []int64) error {
	return c.Delete(ResUsersSettingsVolumesModel, ids)
}

// GetResUsersSettingsVolumes gets res.users.settings.volumes existing record.
func (c *Client) GetResUsersSettingsVolumes(id int64) (*ResUsersSettingsVolumes, error) {
	rusvs, err := c.GetResUsersSettingsVolumess([]int64{id})
	if err != nil {
		return nil, err
	}
	if rusvs != nil && len(*rusvs) > 0 {
		return &((*rusvs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users.settings.volumes not found", id)
}

// GetResUsersSettingsVolumess gets res.users.settings.volumes existing records.
func (c *Client) GetResUsersSettingsVolumess(ids []int64) (*ResUsersSettingsVolumess, error) {
	rusvs := &ResUsersSettingsVolumess{}
	if err := c.Read(ResUsersSettingsVolumesModel, ids, nil, rusvs); err != nil {
		return nil, err
	}
	return rusvs, nil
}

// FindResUsersSettingsVolumes finds res.users.settings.volumes record by querying it with criteria.
func (c *Client) FindResUsersSettingsVolumes(criteria *Criteria) (*ResUsersSettingsVolumes, error) {
	rusvs := &ResUsersSettingsVolumess{}
	if err := c.SearchRead(ResUsersSettingsVolumesModel, criteria, NewOptions().Limit(1), rusvs); err != nil {
		return nil, err
	}
	if rusvs != nil && len(*rusvs) > 0 {
		return &((*rusvs)[0]), nil
	}
	return nil, fmt.Errorf("res.users.settings.volumes was not found with criteria %v", criteria)
}

// FindResUsersSettingsVolumess finds res.users.settings.volumes records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersSettingsVolumess(criteria *Criteria, options *Options) (*ResUsersSettingsVolumess, error) {
	rusvs := &ResUsersSettingsVolumess{}
	if err := c.SearchRead(ResUsersSettingsVolumesModel, criteria, options, rusvs); err != nil {
		return nil, err
	}
	return rusvs, nil
}

// FindResUsersSettingsVolumesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersSettingsVolumesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersSettingsVolumesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersSettingsVolumesId finds record id by querying it with criteria.
func (c *Client) FindResUsersSettingsVolumesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersSettingsVolumesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users.settings.volumes was not found with criteria %v and options %v", criteria, options)
}
