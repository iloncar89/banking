package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Ivan", "Osijek", "31000", "1989-11-21", "1"},
		{"1002", "Josip", "Amsterdam", "10000", "1992-12-01", "1"},
	}
	return CustomerRepositoryStub{customers}
}
