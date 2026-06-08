# Documentação de Interface: Aba de Detalhes do Processo

Este documento detalha os campos que compõem a aba de **Detalhes do Processo / Dados Gerais** na interface do sistema, bem como os métodos e endpoints da API necessários para consumi-los e renderizá-los, usando como base o WSSEI.

---

## 1. Visão Geral da Arquitetura de Chamadas
Para carregar completamente esta aba, o front-end deverá realizar chamadas paralelas aos endpoints principais utilizando o ID do processo (`{protocolo}`):

1. `GET /processo/{protocolo}` (Dados estruturais, status, ciências base e anotações)
2. `GET /processo/{protocolo}/interessados/listar` (Lista de participantes)
3. `GET /processo/acompanhamento/consultar` (Observações internas da unidade)
4. `GET /processo/{protocolo}/ciencia/listar` (Histórico completo de Ciências/Leituras)

A ideia é carregar essas informações para um banco de dados.

---

## 2. Mapeamento de Campos por Seção

### Seção A: Identificação do Processo
Campos principais de cabeçalho e indexação legal do processo.

| Nome do Campo | Origem do Dado (Endpoint) | Caminho no JSON | Tipo | Exemplo / Observação |
| :--- | :--- | :--- | :--- | :--- |
| **Número do Processo** | `GET /processo/{protocolo}` | `data.atributos.numero` | String | `99990.000007/2019-00` |
| **Tipo do Processo** | `GET /processo/{protocolo}` | `data.atributos.tipoProcesso` | String | `Acesso à Informação: Demanda do e-SIC` |
| **Especificação / Descrição** | `GET /processo/{protocolo}` | `data.atributos.descricao` | String | Resumo textual do objeto do processo |
| **Classificação por Assunto** | `GET /processo/pesquisar` ou Rotas de Assunto | `data[].codigoestruturadoformatado` | String | `001 - MODERNIZAÇÃO E REFORMA...` |

### Seção B: Responsáveis e Lotação Atual
Indica onde e com quem o processo se encontra no momento da consulta.

| Nome do Campo | Origem do Dado (Endpoint) | Caminho no JSON | Tipo | Exemplo / Observação |
| :--- | :--- | :--- | :--- | :--- |
| **Unidade Atual** | `GET /processo/{protocolo}` | `data.atributos.unidade.sigla` | String | `TESTE` |
| **Usuário Atribuído** | `GET /processo/{protocolo}` | `data.atributos.usuarioAtribuido.nomeformatado` | String | `nome - sigla` |
| **Unidade de Origem** | `GET /processo/pesquisar` | `data[].siglaUnidadeGeradora` | String | Unidade que realizou a abertura do processo |

### Seção C: Controle de Acesso e Temporalidade
Informações de segurança e alertas de prazos para a unidade jurídica/administrativa.

| Nome do Campo | Origem do Dado (Endpoint) | Caminho no JSON | Tipo | Exemplo / Observação |
| :--- | :--- | :--- | :--- | :--- |
| **Nível de Acesso** | `GET /processo/{protocolo}` | `data.atributos.status.nivelAcessoGlobal` | String | Pode ser inferido via flags `documentoSigiloso` / `documentoRestrito` |
| **Status de Tramitação**| `GET /processo/{protocolo}` | `data.atributos.status.processoEmTramitacao` | Booleano| Mapear `true` para "Em Tramitação" |
| **Retorno Programado** | `GET /processo/{protocolo}` | `data.atributos.status.retornoData` | Object/Date| Data limite acordada para retorno |
| **Alerta de Atraso** | `GET /processo/{protocolo}` | `data.atributos.status.retornoAtrasado` | Booleano| Se `true`, exibir badge visual de alerta |

### Seção D: Interessados e Contexto Local
Dados dinâmicos sobre os envolvidos e marcações específicas da unidade atual.

| Nome do Campo | Origem do Dado (Endpoint) | Caminho no JSON | Tipo | Exemplo / Observação |
| :--- | :--- | :--- | :--- | :--- |
| **Interessados** | `GET /processo/{protocolo}/interessados/listar` | `data[].nomeformatado` | Array | Lista de interessados (Ex: `INTRANET (INTRANET)`) |
| **Marcador / Tag** | `GET /processo/{protocolo}` | `data.atributos.marcador.nome` (para o nome) | String | Etiqueta visual. Utilizar `data.atributos.marcador.idCor` para estilizar a tag, além disso ao passar a mão no marcador deve aparecer a sua descrição que é representada por `data.atributos.marcador.texto`|
| **Observações da Unidade**| `GET /processo/acompanhamento/consultar` | `data.observacao` | String | Anotação do grupo de acompanhamento |

---

## 3. Seção E: Histórico e Linha do Tempo (Timeline)
Esta seção reconstrói os passos de auditoria internos, divididos em **Ciências de Recebimento** (quem visualizou) e **Histórico de Despachos/Anotações**.

### 3.1 Linha do Tempo de Ciências (Quem tomou conhecimento)
Ideal para renderizar em um componente de *Timeline* vertical ordenado por data decrescente.

* **Método para obter:** `GET /processo/{protocolo}/ciencia/listar`
* **Campos para renderização:**

| Informação na Timeline | Caminho no JSON | Tipo | Exemplo / Observação |
| :--- | :--- | :--- | :--- |
| **Data/Hora da Ciência** | `data[].data` | String/Date | `09/10/2019 18:16:01` |
| **Usuário que Visualizou** | `data[].siglaUsuario` / `nomeUsuario` | String | `teste` |
| **Unidade do Usuário** | `data[].siglaUnidade` / `unidade` | String | `TESTE` - `Unidade de Teste 1` |
| **Descrição do Evento** | `data[].descricao` | String | `Ciência no processo` |

### 3.2 Histórico de Anotações Internas
Lista de notas de trabalho inseridas pelos usuários que tramitaram o processo.

* **Método para obter:** `GET /processo/{protocolo}`
* **Campos para renderização:**

| Informação na Lista | Caminho no JSON | Tipo | Exemplo / Observação |
| :--- | :--- | :--- | :--- |
| **Texto da Nota** | `data.atributos.anotacoes[].descricao` | String | Despacho interno ou lembrete de pendência |
| **Data da Anotação** | `data.atributos.anotacoes[].dthAnotacao` | String/Date | Data em que a nota foi salva |
| **Sinalizador de Prioridade**| `data.atributos.anotacoes[].sinPrioridade` | String | Se ativo, renderizar um ícone de "Alerta/Estrela" |

---

## 4. Regras de Negócio para Renderização da Interface

1. **Badge de Nível de Acesso:**
   * Se `documentoSigiloso == "string/true"`, renderizar Tag vermelha **Sigiloso**.
   * Se `documentoRestrito == "string/true"`, renderizar Tag amarela **Restrito**.
   * Caso contrário, renderizar Tag verde **Público**.

2. **Estilização de Marcadores:**
   * O objeto `marcador` possui os campos `nome`, `texto` e `descricaoCor`. Utilize a cor retornada para preencher o background do componente de Tag na interface, simulando o comportamento nativo do SEI.

3. **Renderização do Histórico (Timeline):**
   * Una os dados de `ciências` e `anotações` se desejar uma linha do tempo unificada, ordenando-os pela data (`data` e `dthAnotacao` respectivamente) para que o usuário veja a ordem cronológica exata dos acontecimentos do processo.


Referências:

https://pengovbr.github.io/mod-wssei/#/Processo/listarTipoProcesso

https://www.sei.mg.gov.br/sei/
