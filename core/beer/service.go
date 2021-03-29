package beer

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Aqui vamos definir as regras de negocio

//Define a interface com as funções que serão usadas pelo restante do projeto
// É basicamente o que representa um serviço para mim

type Reader interface {
	GetAll() ([]*Beer, error)
	Get(ID int64) (beer *Beer)
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

type Service struct {
	DB *sql.DB
}

/*
Vai ser uma especie de construtor
-> Retorna um ponteiro em memória para uma estrutura
-> A função agora recebe uma conexão com o banco de dados

*/

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetAll() ([]*Beer, error) {
	// Result é um slice de ponteiros do tipo Beer
	var result []*Beer

	// vamos sempre usar a conexão que está dentro do Service
	rows, err := s.DB.Query("SELECT id, name, type, style FROM beer")

	// O erro deve ser tratado por quem chamou o pacote
	// Essa é uma boa prática em Go.
	if err != nil {
		return nil, err
	}

	// Garante que o comando rows.Close vai ser executado na saída da func
	// assim não precisamos nos preocupar em fechar a conexão
	defer rows.Close()

	for rows.Next() {
		var beer Beer
		err = rows.Scan(&beer.ID, &beer.Name, &beer.Style, &beer.Type)
		if err != nil {
			return nil, err
		}

		result = append(result, &beer)

	}
	return result, nil

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
