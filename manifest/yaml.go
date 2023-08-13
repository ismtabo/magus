package manifest

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/ismtabo/magus/context"
	"gopkg.in/yaml.v3"
)

var v = validator.New()

func init() {
	v.RegisterValidation("required_with_any", requiredWithAny)
}

func requiredWithAny(fl validator.FieldLevel) bool {
	param := fl.Param()
	others := strings.Split(param, " ")
	for _, other := range others {
		if !fl.Parent().FieldByName(other).IsZero() {
			return !fl.Field().IsZero()
		}
	}
	return true
}

func UnmarshalYAML(ctx context.Context, data []byte, manifest *Manifest) error {
	if err := yaml.Unmarshal(data, &manifest); err != nil {
		// TODO: Wrap error
		return err
	}
	if err := v.Struct(manifest); err != nil {
		// TODO: Wrap error
		return err
	}
	return nil
}
