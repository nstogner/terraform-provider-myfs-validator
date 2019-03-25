package validator

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
)

func validateTextFile(d *schema.ResourceData) error {
	// This validation logic could be farmed off somewhere else.
	if len(d.Get("text").(string)) < 5 {
		return errors.New("len(text) < 5")
	}

	return nil
}
