package data

import "github.com/maxbrain0/react-go-graphql/server/logger"

type product struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

// Data holds some mock product data
type Data struct {
	products []product
}

var startingData = []product{
	{
		ID:   1,
		Name: "Wax seal",
	},
	{
		ID:   2,
		Name: "Toilet Flap",
	},
	{
		ID:   3,
		Name: "Bidet Seat",
	},
}

var ctxLogger = logger.CtxLogger

// InitData initializes DataStruct with mock data
func (d *Data) InitData() {
	for _, val := range startingData {
		d.products = append(d.products, val)
	}

	ctxLogger.Debug("Products array has been filled")
}
