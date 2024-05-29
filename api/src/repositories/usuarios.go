package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *usuarios {
	return &usuarios{db}
}

func (repositorio usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("INSERT INTO usuarios (nome, nick, email, senha) VALUES (?, ?, ?, ?)")

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}

func (repositorio usuarios) Buscar(nomeOuNickname string) ([]models.Usuario, error) {
	nomeOuNick := fmt.Sprintf("%%%s%%", nomeOuNickname) //%nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, email, criadoEm FROM usuarios WHERE nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)

	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio usuarios) BuscarPorId(id uint64) (models.Usuario, error) {

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?",
		id,
	)

	if erro != nil {
		return models.Usuario{}, nil
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {

		erro = linhas.Scan(&usuario.Nome, &usuario.Nick, &usuario.Email, &usuario.CriadoEm)

		if erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio usuarios) Atualizar(id uint64, usuario models.Usuario) error {

	statement, erro := repositorio.db.Prepare(
		"UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?",
	)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) Deletar(id uint64) error {

	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id = ?")

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"SELECT id, senha FROM usuarios WHERE email = ?",
		email,
	)

	if erro != nil {
		return models.Usuario{}, nil
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {

		erro = linhas.Scan(&usuario.ID, &usuario.Senha)

		if erro != nil {
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}
