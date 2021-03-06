package gather

// Dummy is a dummy object that satisfies Gatherer interface
import (
	"context"
	"fmt"

	"github.com/themotion/ladder/types"
)

const (
	quantityOpt = "quantity"
)

// Dummy is a dummy object that satisfies Gatherer interface
type Dummy struct {
	quantity types.Quantity
}

// NewDummy creates a dummy object
func NewDummy(opts map[string]interface{}) (d *Dummy, err error) {
	// Recover from wrong type assertions
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	v, ok := opts[quantityOpt]
	// Set each option with the correct type
	if !ok {
		return nil, fmt.Errorf("%s configuration opt is required", quantityOpt)
	}
	q := types.I2Int64(v)

	return &Dummy{quantity: types.Quantity{Q: q}}, nil
}

// Gather does nothing, is dummy, returns 0
func (d *Dummy) Gather(ctx context.Context) (types.Quantity, error) {
	return d.quantity, nil
}

// DummyCreator is a dummy creation object that satisfies Creator interface
type DummyCreator struct{}

// Create creates a a Gatherer
func (d *DummyCreator) Create(ctx context.Context, opts map[string]interface{}) (Gatherer, error) {
	return NewDummy(opts)
}
