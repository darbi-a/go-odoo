package api

import (
	"github.com/skilld-labs/go-odoo/types"
)

type DecimalPrecisionService struct {
	client *Client
}

func NewDecimalPrecisionService(c *Client) *DecimalPrecisionService {
	return &DecimalPrecisionService{client: c}
}

func (svc *DecimalPrecisionService) GetIdsByName(name string) ([]int, error) {
	return svc.client.getIdsByName(types.DecimalPrecisionModel, name)
}

func (svc *DecimalPrecisionService) GetByIds(ids []int) (*types.DecimalPrecisions, error) {
	d := &types.DecimalPrecisions{}
	return d, svc.client.getByIds(types.DecimalPrecisionModel, ids, d)
}

func (svc *DecimalPrecisionService) GetByName(name string) (*types.DecimalPrecisions, error) {
	d := &types.DecimalPrecisions{}
	return d, svc.client.getByName(types.DecimalPrecisionModel, name, d)
}

func (svc *DecimalPrecisionService) GetByField(field string, value string) (*types.DecimalPrecisions, error) {
	d := &types.DecimalPrecisions{}
	return d, svc.client.getByName(types.DecimalPrecisionModel, field, value, d)
}

func (svc *DecimalPrecisionService) GetAll() (*types.DecimalPrecisions, error) {
	d := &types.DecimalPrecisions{}
	return d, svc.client.getAll(types.DecimalPrecisionModel, d)
}

func (svc *DecimalPrecisionService) Create(fields map[string]interface{}, relations *types.Relations) (int, error) {
	return svc.client.create(types.DecimalPrecisionModel, fields, relations)
}

func (svc *DecimalPrecisionService) Update(ids []int, fields map[string]interface{}, relations *types.Relations) error {
	return svc.client.update(types.DecimalPrecisionModel, ids, fields, relations)
}

func (svc *DecimalPrecisionService) Delete(ids []int) error {
	return svc.client.delete(types.DecimalPrecisionModel, ids)
}
