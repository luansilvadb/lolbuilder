# cli-query-output Specification

## Purpose

Oferecer resultados de consulta legíveis, determinísticos e consumíveis por automação nos comandos `runes` e `builds`, sem expor estruturas internas nem perder o contexto do patch.

## Requirements

### Requirement: Formato textual legível por padrão

Os comandos `runes` e `builds` SHALL usar texto humano em português quando `-format` não for informado. A saída SHALL ser determinística, não SHALL expor representações internas como `map[...]` e não SHALL usar cores ou mensagens de progresso no stdout.

#### Scenario: Consulta textual de runas

- **WHEN** o usuário executa `lolbuilder -objective armor runes`
- **THEN** o comando imprime um cabeçalho com o patch usado, a página completa e um resumo em linguagem humana

#### Scenario: Consulta textual de builds

- **WHEN** o usuário executa `lolbuilder -objective armor -gold 10000 builds`
- **THEN** o comando imprime o objetivo maximizado, o valor obtido, o gasto, os itens e os atributos adicionais em formato legível

### Requirement: Apresentação explícita de páginas de runas

A saída de `runes` SHALL mostrar todos os slots da página. Um slot em que nenhuma opção contribui para o objetivo SHALL ser identificado como `indiferente`, e não por um nome vazio ou por um desempate arbitrário. A saída SHALL informar as runas contribuintes, a quantidade de slots livres e o total do objetivo.

#### Scenario: Página com slots indiferentes

- **WHEN** uma página possui slots sem qualquer contribuição para o objetivo
- **THEN** cada um desses slots aparece como `indiferente` e o resumo informa quantos slots estão livres

#### Scenario: Objetivo sem contribuição de runas

- **WHEN** o objetivo é válido, mas nenhuma runa do catálogo contribui para ele
- **THEN** o comando termina com sucesso, informa que nenhuma runa contribui, mostra total zero na unidade correta e não escolhe uma runa arbitrária

### Requirement: Apresentação explícita de builds de itens

A saída de `builds` SHALL destacar o atributo otimizado e seu valor antes dos demais atributos. Também SHALL informar orçamento, gasto, resolução da força adaptativa e uma nota de que passivas e ativas dos itens não entram no cálculo.

#### Scenario: Build com atributos secundários

- **WHEN** os itens escolhidos concedem atributos além do objetivo
- **THEN** o objetivo aparece em destaque e os demais atributos são listados separadamente como atributos adicionais

#### Scenario: Build que não gasta todo o orçamento

- **WHEN** a combinação ótima usa menos ouro que o orçamento disponível
- **THEN** a saída informa gasto e orçamento lado a lado, sem sugerir que o ouro restante foi utilizado

### Requirement: Formato JSON estruturado

Os comandos `runes` e `builds` SHALL aceitar `-format json`. Nesse formato, o stdout SHALL conter exclusivamente um único documento JSON válido, com os campos `comando`, `patch`, `objetivo`, `resolucao_adaptativa`, `resultado` e `avisos`. O campo `resultado` SHALL preservar os campos canônicos dos objetos calculados; os nomes canônicos SHALL permanecer disponíveis no JSON mesmo quando o texto usar rótulos em português.

#### Scenario: JSON de uma consulta de runas

- **WHEN** o usuário executa uma consulta de runas com `-format json`
- **THEN** o stdout contém somente JSON, o envelope identifica o comando e o patch, e `resultado` contém a página calculada

#### Scenario: JSON de uma consulta de builds

- **WHEN** o usuário executa uma consulta de builds com `-format json`
- **THEN** o stdout contém somente JSON, o envelope identifica objetivo e resolução adaptativa, e `resultado` contém a build calculada

#### Scenario: Aviso sem falha

- **WHEN** uma consulta válida produz uma condição informativa, como nenhum slot contribuinte
- **THEN** o comando termina com sucesso e registra a mensagem em `avisos`, sem misturá-la ao JSON no stdout

### Requirement: Validação do formato solicitado

O comando SHALL aceitar somente `text` e `json` como valores de `-format`. Um valor desconhecido SHALL produzir erro explícito em stderr e código de saída diferente de zero, sem emitir resultado parcial no stdout.

#### Scenario: Formato desconhecido

- **WHEN** o usuário informa `-format yaml`
- **THEN** o comando informa os formatos aceitos, não executa a consulta e termina com falha

### Requirement: Formatação localizada sem alterar o contrato canônico

A saída textual SHALL usar nomes humanos em português e formatação brasileira para percentuais, milhares e decimais, removendo zeros decimais desnecessários. A saída JSON SHALL manter números como números e estatísticas com suas chaves canônicas.

#### Scenario: Valor decimal no texto e no JSON

- **WHEN** um resultado contém um valor decimal como 3.5
- **THEN** o texto exibe `3,5` na unidade correspondente e o JSON exibe o número JSON `3.5`
