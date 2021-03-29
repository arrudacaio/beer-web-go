package beer

// Aqui vamos definir as regras de negocio

//Define a interface com as funções que serão usadas pelo restante do projeto
// É basicamente o que representa um serviço para mim

type Reader interface {
	GetAll() ([]*Beer, error)
	Get(ID int) (beer *Beer)
}

type Writer interface {
	Store(beer *Beer) error
	Update(beer *Beer) error
	Remove(ID int) error
}
type UseCase interface {
	Reader
	Writer
}

/*
Em Go qualquer coisa que implementa as funções de uma interface
passa a ser uma implementação válida, desde que implemente todas as funções
*/

type Service struct{}

/*
Vai ser uma especie de construtor
Toda vez que eu tiver uma variavel do tipo Service,
eu poderei chamar essas funções que estou associando à esta variavel
Ex:
	var service Service
	service.NewService()

*/

func NewService() *Service {
	return &Service{}
}

// Implementar posteriormente
func (s *Service) GetAll() ([]*Beer, error) {
	return nil, nil
}

func (s *Service) Get(ID int) (*Beer, error) {
	return nil, nil
}

func (s *Service) Store(beer *Beer) error {
	return nil
}

func (s *Service) Remove(ID int) error {
	return nil
}
