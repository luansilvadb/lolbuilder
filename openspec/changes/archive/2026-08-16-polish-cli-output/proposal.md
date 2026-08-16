## Why

Os comandos `runes` e `builds` calculam resultados úteis, mas a saída atual expõe detalhes internos do Go (`map[...]`) e deixa escolhas indiferentes vazias. Isso dificulta a leitura humana, torna casos sem contribuição ambíguos e não oferece um contrato confiável para automação.

## What Changes

- Melhorar a saída textual de `runes` e `builds` com rótulos explícitos, unidades em português e totais legíveis.
- Mostrar slots de runa sem contribuição como `indiferente`, sem transformar desempates em recomendações.
- Destacar o atributo otimizado e separar atributos adicionais nas builds de itens.
- Adicionar `-format text|json`, mantendo texto como padrão e aceitando a flag antes do comando.
- Garantir que JSON seja emitido sozinho no stdout, com metadados de patch, resultado e avisos.
- Tratar objetivos válidos sem contribuição como sucesso explícito, não como erro.
- Preservar os campos canônicos dos objetos internos dentro de `resultado`.
- Adicionar testes de snapshot para texto e testes de contrato para JSON.
- **BREAKING**: consumidores que fazem parsing da formatação textual atual deverão migrar para o formato JSON.

## Capabilities

### New Capabilities

- `cli-query-output`: define os formatos, o conteúdo e o comportamento de saída dos comandos de consulta `runes` e `builds`.

### Modified Capabilities

- Nenhuma.

## Impact

- Afeta `cmd/lolbuilder/main.go` e `cmd/lolbuilder/optimize.go`, incluindo parsing da nova flag e renderização dos resultados.
- Pode introduzir helpers de formatação no pacote canônico para nomes, unidades e vetores de atributos.
- Afeta testes de CLI e o contrato textual/documental dos comandos.
- Não altera os algoritmos de otimização, o dataset exportado, `build`, `export` ou `ingame`.
- Não adiciona dependências externas.
