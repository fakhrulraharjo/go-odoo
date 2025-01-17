package odoo

import (
	"fmt"
)

// ImageMixin represents image.mixin model.
type ImageMixin struct {
	Image1024 *String `xmlrpc:"image_1024,omptempty"`
	Image128  *String `xmlrpc:"image_128,omptempty"`
	Image1920 *String `xmlrpc:"image_1920,omptempty"`
	Image256  *String `xmlrpc:"image_256,omptempty"`
	Image512  *String `xmlrpc:"image_512,omptempty"`
}

// ImageMixins represents array of image.mixin model.
type ImageMixins []ImageMixin

// ImageMixinModel is the odoo model name.
const ImageMixinModel = "image.mixin"

// Many2One convert ImageMixin to *Many2One.
func (im *ImageMixin) Many2One() *Many2One {
	return NewMany2One(im.Id.Get(), "")
}

// CreateImageMixin creates a new image.mixin model and returns its id.
func (c *Client) CreateImageMixin(im *ImageMixin) (int64, error) {
	return c.Create(ImageMixinModel, im)
}

// UpdateImageMixin updates an existing image.mixin record.
func (c *Client) UpdateImageMixin(im *ImageMixin) error {
	return c.UpdateImageMixins([]int64{im.Id.Get()}, im)
}

// UpdateImageMixins updates existing image.mixin records.
// All records (represented by ids) will be updated by im values.
func (c *Client) UpdateImageMixins(ids []int64, im *ImageMixin) error {
	return c.Update(ImageMixinModel, ids, im)
}

// DeleteImageMixin deletes an existing image.mixin record.
func (c *Client) DeleteImageMixin(id int64) error {
	return c.DeleteImageMixins([]int64{id})
}

// DeleteImageMixins deletes existing image.mixin records.
func (c *Client) DeleteImageMixins(ids []int64) error {
	return c.Delete(ImageMixinModel, ids)
}

// GetImageMixin gets image.mixin existing record.
func (c *Client) GetImageMixin(id int64) (*ImageMixin, error) {
	ims, err := c.GetImageMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if ims != nil && len(*ims) > 0 {
		return &((*ims)[0]), nil
	}
	return nil, fmt.Errorf("id %v of image.mixin not found", id)
}

// GetImageMixins gets image.mixin existing records.
func (c *Client) GetImageMixins(ids []int64) (*ImageMixins, error) {
	ims := &ImageMixins{}
	if err := c.Read(ImageMixinModel, ids, nil, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

// FindImageMixin finds image.mixin record by querying it with criteria.
func (c *Client) FindImageMixin(criteria *Criteria) (*ImageMixin, error) {
	ims := &ImageMixins{}
	if err := c.SearchRead(ImageMixinModel, criteria, NewOptions().Limit(1), ims); err != nil {
		return nil, err
	}
	if ims != nil && len(*ims) > 0 {
		return &((*ims)[0]), nil
	}
	return nil, fmt.Errorf("image.mixin was not found with criteria %v", criteria)
}

// FindImageMixins finds image.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImageMixins(criteria *Criteria, options *Options) (*ImageMixins, error) {
	ims := &ImageMixins{}
	if err := c.SearchRead(ImageMixinModel, criteria, options, ims); err != nil {
		return nil, err
	}
	return ims, nil
}

// FindImageMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImageMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ImageMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindImageMixinId finds record id by querying it with criteria.
func (c *Client) FindImageMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImageMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("image.mixin was not found with criteria %v and options %v", criteria, options)
}
