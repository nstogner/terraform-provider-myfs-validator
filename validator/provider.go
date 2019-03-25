package validator

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/pkg/errors"
)

type validateFunc func(*schema.ResourceData) error

var validations = map[string]validateFunc{
	"myfs_text_file": validateTextFile,
}

func New(p *schema.Provider) *schema.Provider {
	for name, res := range p.ResourcesMap {
		if validate, ok := validations[name]; ok {
			p.ResourcesMap[name] = resourceWithValidator(res, validate)
		}
	}
	return p
}

func resourceWithValidator(r *schema.Resource, f validateFunc) *schema.Resource {
	origCreate := r.Create
	origUpdate := r.Update

	r.Create = func(d *schema.ResourceData, x interface{}) error {
		if err := f(d); err != nil {
			return errors.Wrap(err, "validating")
		}
		return origCreate(d, x)
	}

	r.Update = func(d *schema.ResourceData, x interface{}) error {
		if err := f(d); err != nil {
			return errors.Wrap(err, "validating")
		}
		return origUpdate(d, x)
	}

	return r
}
