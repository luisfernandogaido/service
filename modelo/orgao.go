package modelo

import (
	"service/mysql"
	"strings"
)

type Orgao struct {
	Mcu    string
	Nome   string
	Tipo   string
	Cidade string
	Dr     string
	Uf     string
	Cep    string
}

func Orgaos() ([]Orgao, error) {
	query := `
	SELECT mcu, nome, tipo, cidade, sigla_dr, uf, cep
	FROM orgao
	WHERE	nome_dr IS NOT NULL AND
			sigla_dr IS NOT NULL	
	`
	rows, err := mysql.Ect.Query(query)
	if err != nil {
		return nil, err
	}
	orgaos := make([]Orgao, 0)
	for rows.Next() {
		orgao := Orgao{}
		err = rows.Scan(
			&orgao.Mcu,
			&orgao.Nome,
			&orgao.Tipo,
			&orgao.Cidade,
			&orgao.Dr,
			&orgao.Uf,
			&orgao.Cep,
		)
		orgao.Nome = strings.TrimSpace(orgao.Nome)
		orgao.Cidade = strings.TrimSpace(orgao.Cidade)
		orgao.Dr = strings.TrimSpace(orgao.Dr)
		if err != nil {
			return nil, err
		}
		orgaos = append(orgaos, orgao)
	}
	return orgaos, nil
}
