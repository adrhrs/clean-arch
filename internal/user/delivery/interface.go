package delivery

type UserDelivery interface {
	List() error
	Create() error
}
