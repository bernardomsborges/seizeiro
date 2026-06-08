# Documento de Requisitos do Sistema: Dashboard

## 1. Visão Geral do Produto

O **Dashboard** é uma central de monitoramento de processos administrativos focada no acompanhamento de processos, etiquetas, prioridades e responsáveis. A ferramenta centraliza os dados provenientes da API mapeada (WSSEI), fornecendo aos gestores e analistas uma visão analítica global do fluxo de trabalho, permitindo triagem célere, filtragem avançada e atribuição de responsáveis.

---

## 2. Requisitos Funcionais (RF)

### 2.1 Seção Superior: Cabeçalho e Ações Globais

* **RF-001 [Título do Dashboard]:** A interface deve exibir de forma clara o título "Dashboard de Governança" acompanhado do subtítulo descritivo: *"Visão geral dos processos abertos na unidade"*.
* **RF-002 [Criar Novo Processo]:** Deve haver um botão de ação primária destacado (`+ Novo Processo`) posicionado no canto superior direito da seção de busca. Ao ser clicado, deve abrir o fluxo de abertura/cadastro de novos procedimentos.

### 2.2 Seção de Busca Rápida e Filtros

A barra de filtros deve permitir a combinação dinâmica de parâmetros (operação lógica `AND`).

* **RF-003 [Busca por Texto Livre]:** Input de pesquisa que aceita o número completo ou parcial do processo administrativo (Ex: `1260.01.0042881/2026-69`). O gatilho de busca pode ser por *debounce* (300ms) ou ao pressionar *Enter*.
* **RF-004 [Filtro de Origem]:** Componente *Select/Dropdown* para filtrar os processos pela unidade geradora (Ex: `SEE/SRE Januária/Aposentadori`, `SEF/SPGF-DAPE-DCTA`).
* **RF-005 [Filtro de Status]:** Componente *Select/Dropdown* para filtrar pelo estado atual de tramitação do processo (`Análise Pendente`, `Em Triagem`, `Concluído`, etc.).
* **RF-006 [Filtro de Etiquetas]:** Componente *Select/Dropdown* mapeado a partir dos marcadores retornados pela API (`Ouro`, `Idoso`, `Prata`, `Roxo`).
    * _Obs:_ Etiqueta visual. Utilizar `data.atributos.marcador.idCor` para estilizar a tag, além disso ao passar a mão no marcador deve aparecer a sua descrição que é representada por `data.atributos.marcador.texto`
* **RF-007 [Filtro de Prioridade]:** Componente *Select/Dropdown* booleano para isolar processos marcados com prioridade urgente (`Sim` / `Não`).
* **RF-008 [Filtro de Analista]:** Componente *Select/Dropdown* para filtrar por processos atribuídos a usuários específicos (Ex: `Maria Costa`, `João Silva`) ou `Não possui`.
* **RF-009 [Minimizar Painel]:** O bloco de "Busca rápida e filtros" deve ser colapsável através de um gatilho de *accordion* no canto superior direito da seção.

### 2.3 Tabela Principal de Governança

* **RF-010 [Listagem de Dados]:** A tabela deve exibir de forma estruturada as colunas: *Número*, *Unidade de Origem*, *Status*, *Etiquetas*, *Prioritário* e *Analista*.
* **RF-011 [Hiperlink do Processo]:** O número do processo deve ser renderizado como um link clicável (cor azul padrão). Ao clicar, o sistema deve direcionar o usuário para a aba de **Dados Gerais/Detalhes do Processo**.
* **RF-012 [Estilização de Status (Badges)]:** O status deve ser exibido dentro de componentes visuais arredondados (*badges*) com cores semânticas mapeadas:
    * `Análise Pendente`: Fundo cinza claro, texto escuro com marcador circular cinza.
    * `Em Triagem`: Fundo azul claro, borda azul, texto azul com marcador circular azul.
    * `Concluído`: Fundo verde claro, borda verde, texto verde com ícone de *check* ($\checkmark$).
* **RF-013 [Multi-Tags de Etiquetas]:** A coluna "Etiquetas" deve suportar a renderização de múltiplos marcadores empilhados verticalmente por linha, respeitando a cor de background e texto vinda da API (Ex: Tag amarela para `Ouro`, Tag roxa para `Roxo`). Linhas sem etiquetas exibem apenas um traço (`-`). 
* **RF-014 [Destaque de Prioridade]:** Caso o processo possua prioridade (`sinPrioridade` ativo ou regras internas), a palavra **Sim** deve ser exibida em negrito para chamar a atenção visual do operador.
* **RF-015 [Avatar do Analista]:** A coluna de analista deve exibir o nome do usuário acompanhado de suas iniciais dentro de um círculo com fundo azul escuro (Avatar Component) se houver atribuição. Caso contrário, exibe o texto *"Não possui"*.

### 2.4 Paginação de Resultados

* **RF-016 [Paginação]:** O rodapé da tabela deve conter controles de navegação de páginas (`< Anterior` e `Próxima >`), além do indicador centralizado do estado atual da paginação (Ex: *"Página 1 de 6"*). Os botões devem ficar desabilitados (*disabled*) caso o usuário esteja na primeira ou última página, respectivamente.

---

## 3. Requisitos Não-Funcionais (RNF)

* **RNF-001 [Performance/Paginação no Back-end]:** A listagem de processos não deve carregar todos os registros de uma vez. O consumo deve utilizar obrigatoriamente os parâmetros de query string `limit` e `start` já existentes nas rotas da API para garantir respostas abaixo de 200ms.
* **RNF-002 [Design System / Responsividade]:** A interface deve seguir um padrão limpo (estilo clean/corporate), utilizando fontes sem serigrafia (sans-serif), cantos levemente arredondados (`border-radius: 6px` a `8px`) e grids bem espaçados para evitar fadiga visual.
* **RNF-003 [Gerenciamento de Estado local]:** Os filtros aplicados na URL ou no estado global da aplicação (ex: Context API, Redux ou Zustand) devem ser limpos ou mantidos conforme o usuário navega entre os detalhes de um processo e retorna à listagem principal.
* **RNF-004 [Tratamento de Dados Nulos]:** Valores nulos ou vazios retornados pela API (como strings vazias de descrição ou falta de analista atribuído) devem ser sanitizados no front-end para nunca quebrar a estrutura da tabela, exibindo fallbacks visuais como `Não possui` ou `-`.