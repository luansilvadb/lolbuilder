## Context

Consulte `proposal.md` para a motivação e `specs/cli-query-output/spec.md` para o contrato observável. Atualmente, `cmd/lolbuilder/optimize.go` imprime os nomes das escolhas diretamente e aplica `%v` a `canon.Vector`, que é um mapa. Os resultados calculados já possuem campos JSON canônicos, mas os comandos não oferecem um envelope nem identificam explicitamente o patch resolvido.

## Goals / Non-Goals

**Goals:**

- Separar cálculo de apresentação para que texto e JSON consumam o mesmo resultado.
- Tornar a saída textual estável, legível e explícita sobre indiferença, objetivo e limitações.
- Produzir JSON puro no stdout, com metadados suficientes para reproduzir a consulta.
- Centralizar ordenação, rótulos, unidades e formatação numérica.
- Cobrir os casos normais e os casos de borda com testes determinísticos.

**Non-Goals:**

- Alterar os algoritmos exatos de runas ou itens.
- Alterar o dataset, o export, o build ou a verificação in-game.
- Adicionar dependências externas, cores ou uma interface gráfica.
- Renomear os campos JSON existentes dentro de `resultado`.
- Mudar a semântica dos objetivos ou criar um modelo de combate.

## Decisions

### 1. Um envelope de saída, dois renderizadores

Os comandos devem resolver o dataset e o resultado uma única vez, depois encaminhar o mesmo resultado para um renderizador textual ou JSON. O envelope JSON será um tipo próprio com `comando`, `patch`, `objetivo`, `resolucao_adaptativa`, `resultado` e `avisos`; `resultado` manterá `optimize.Pagina` ou `optimize.Build` sem renomear seus campos.

Alternativa considerada: criar modelos paralelos para texto e JSON. Foi rejeitada porque duplicaria a semântica do resultado e permitiria que os formatos divergissem entre si.

### 2. Formatação centralizada de vetores e estatísticas

Os nomes humanos, unidades, ordem canônica e regras de números devem ficar em um único helper de apresentação. O texto usa rótulos em português e locale brasileiro; o JSON usa as chaves e números já definidos pelo modelo canônico. Nenhum renderizador deve imprimir diretamente um mapa Go.

Alternativa considerada: formatar cada linha diretamente em `imprimirPaginas` e `imprimirBuilds`. Foi rejeitada porque repetiria regras e manteria o acoplamento com a representação interna.

### 3. A flag de formato permanece global nesta mudança

`-format` será interpretado no mesmo lugar das flags existentes e usado somente pelos comandos `runes` e `builds`. A sintaxe convencional de flags depois do subcomando fica fora deste change, evitando uma refatoração do parser junto com a apresentação.

Alternativa considerada: introduzir subparsers por comando agora. Foi rejeitada por ampliar o risco e o escopo da primeira melhoria.

### 4. Patch resolvido como metadado de execução

O fluxo que carrega o dataset deve preservar o patch explicitamente, inclusive quando ele é escolhido como snapshot mais recente. O cabeçalho textual e o envelope JSON usarão esse valor, evitando que o usuário precise inferi-lo do conteúdo.

### 5. Diagnósticos fora do stdout JSON

Resultados JSON serão emitidos apenas depois de validações concluídas. Erros de formato e falhas de execução continuarão no stderr pelo fluxo principal do programa; avisos de resultado serão dados estruturados no campo `avisos`. Isso permite consumo direto por ferramentas como `jq`.

### 6. Testes de contrato para os dois formatos

Os testes devem exercitar os helpers de apresentação e os comandos com fixtures determinísticas. O texto terá snapshots para escolhas contribuintes, slots indiferentes, objetivo sem contribuição, decimais e atributos adicionais. O JSON será validado como documento, verificando envelope, campos canônicos, avisos e ausência de texto extra.

## Risks / Trade-offs

- [Mudança de layout textual pode afetar scripts existentes] → declarar a mudança como breaking para parsers de texto e oferecer JSON como contrato estável.
- [Nomes ou unidades podem ser traduzidos incorretamente] → usar uma tabela única baseada no vocabulário canônico e cobrir cada stat exibido nos testes.
- [Logs acidentais podem invalidar JSON] → separar explicitamente o caminho JSON dos diagnósticos e testar que stdout é parseável integralmente.
- [Patch mais recente pode mudar entre execuções] → sempre incluir o patch resolvido no resultado e permitir `-patch` explícito como já suportado.
- [Mudança do renderizador pode esconder informação] → manter todos os campos calculados no JSON e listar atributos adicionais no texto.

## Migration Plan

1. Implementar o novo contrato e os testes.
2. Atualizar README e exemplos para mostrar o texto padrão e `-format json`.
3. Validar que `go test ./...` passa e comparar a saída dos casos representativos.
4. Consumidores que fazem parsing de texto devem migrar para o envelope JSON.

Rollback: reverter o change de apresentação; os algoritmos e dados permanecem inalterados.
