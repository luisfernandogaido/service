SELECT mcu, nome, tipo, cidade, sigla_dr, uf, cep
FROM orgao
WHERE	nome_dr IS NOT NULL AND
			sigla_dr IS NOT NULL

;
			
SELECT *
FROM orgao
WHERE nome = 'CDD FALCAO';