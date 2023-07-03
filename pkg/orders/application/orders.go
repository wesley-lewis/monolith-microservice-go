package application

type productsService interface {
}

type paymentsService interface {
}

type OrderService struct {
}

func NewOrdersService() {

}

type PlaceOrderCommand struct {
}

func (s OrderService) PlaceOrder(cmd PlaceOrderCommand) error {

}

type MarkorderAsPaidCommand struct {
}

func (s OrderService) MarkorderAsPaid(cmd MarkorderAsPaidCommand) error {

}

func (s OrderService) OrderByID(id orders.ID) (orders.Order, error) {

}
