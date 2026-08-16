## 1. Contrato de execução e metadados

- [x] 1.1 Adicionar a flag global `-format` com valores `text` e `json`, validando valores desconhecidos antes de executar `runes` ou `builds`.
- [x] 1.2 Preservar o patch efetivamente carregado — explícito ou mais recente — para uso no cabeçalho textual e no envelope JSON.
- [x] 1.3 Criar o envelope de saída com `comando`, `patch`, `objetivo`, `resolucao_adaptativa`, `resultado` e `avisos`, preservando os campos canônicos do resultado interno.

## 2. Renderização textual

- [x] 2.1 Criar um formatador centralizado para nomes humanos, unidades, ordenação de estatísticas e números em locale brasileiro.
- [x] 2.2 Atualizar a saída de `runes` para mostrar todos os slots, marcar escolhas sem contribuição como `indiferente` e exibir contribuições, slots livres e total.
- [x] 2.3 Tratar objetivos válidos sem contribuição de runas como resultado bem-sucedido com total zero e aviso informativo.
- [x] 2.4 Atualizar a saída de `builds` para destacar o atributo otimizado, mostrar gasto/orçamento, separar atributos adicionais e exibir a limitação sobre passivas e ativas.
- [x] 2.5 Adicionar o cabeçalho textual com comando, patch e snapshot local, sem imprimir estruturas internas ou cores.

## 3. Renderização JSON e erros

- [x] 3.1 Implementar a serialização JSON dos resultados de `runes` e `builds` usando o envelope definido, mantendo números e chaves canônicas.
- [x] 3.2 Garantir que o caminho JSON escreva somente o documento JSON no stdout e envie erros para stderr pelo fluxo existente.
- [x] 3.3 Representar condições informativas em `avisos`, incluindo ausência de contribuição e a limitação das builds, sem transformar consultas válidas em falhas.

## 4. Testes e documentação

- [x] 4.1 Adicionar testes unitários para rótulos, unidades, ordenação, números decimais e vetores vazios.
- [x] 4.2 Adicionar snapshots ou testes de saída para runas com contribuições, runas sem contribuição e builds com atributos adicionais.
- [x] 4.3 Adicionar testes de contrato que validem o envelope JSON, preservação dos campos canônicos, stdout parseável e rejeição de formato inválido.
- [x] 4.4 Atualizar README com a nova flag, exemplos em texto e JSON e a recomendação de usar JSON para automação.
- [x] 4.5 Executar `gofmt`, `go test ./...` e os exemplos representativos dos dois formatos; confirmar que algoritmos e arquivos exportados não mudaram.
