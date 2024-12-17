package validation

import "github.com/go-playground/validator/v10"

// ValidatorPlayground обертка над go-playground/validator
type ValidatorPlayground struct {
	validate *validator.Validate
}

// NewPlayground создает новый экземпляр ValidatorPlayground
func NewPlayground() *ValidatorPlayground {
	return &ValidatorPlayground{
		validate: validator.New(),
	}
}

// Validate выполняет проверку структуры
func (v *ValidatorPlayground) Validate(dto interface{}) error {
	return v.validate.Struct(dto)
}
