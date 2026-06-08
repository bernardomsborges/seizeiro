# Documentação de Interface: Gestão de Equipe

## 1. Visão Geral do Produto

O módulo de **Gestão de Equipe** tem como objetivo centralizar o controle de membros, funções, atribuições de carga de trabalho e gerenciamento de afastamentos da unidade SEI. A tela permite que gestores analisem a capacidade produtiva da equipe, façam a atribuição de função e responsabilidades e ativação de novos membros vindos do SEI e controlem impedimentos de distribuição de processos (férias, licenças, etc.).

---

## 2. Requisitos Funcionais (RF)

### 2.1 Menu Lateral de Navegação

* **RF-001 [Indicação de Contexto]:** O menu lateral deve destacar visualmente o item **"Gestão de Equipe"** como ativo, utilizando background diferenciado e um marcador lateral esquerdo, posicionando o usuário dentro da suíte da unidade.

### 2.2 Cabeçalho e Ações de Comando (Top Bar)

* **RF-002 [Ações de Configuração]:** O painel superior direito da seção de conteúdo deve exibir três botões de ação rápidos:
    * **Editar Atribuições:** Abre a configuração de escopo de análise dos membros (ex: mudar de Análise Cível para Criminal).
    * **Vincular Função:** Permite associar ou alterar o nível hierárquico/cargo (Sênior, Coordenadora, Pleno, Junior).
    * **Gerenciar Afastamentos:** Botão de ação primária (fundo azul) para cadastrar e editar períodos de recesso e pausas de distribuição.


* **RF-003 [Busca Global e Perfil]:** O topo da plataforma deve conter uma barra de pesquisa de membros (`Buscar membros...`), um ícone de notificações e o avatar do usuário logado.

### 2.3 Cards de Indicadores de Governança (KPIs)

A tela deve exibir três cartões analíticos baseados nos dados em tempo real da unidade:

* **RF-004 [Total de Usuários]:** Exibe o somatório total de colaboradores vinculados à unidade (Ex: `12`).
* **RF-005 [Em Afastamento]:** Exibe o número de colaboradores com afastamentos ativos no dia corrente (Ex: `2`), destacado com texto na cor vermelha caso o valor seja maior que zero.
* **RF-006 [Capacidade da Unidade]:** Exibe a porcentagem de ocupação/carga de trabalho atual da unidade (Ex: `82%`) acompanhada de uma barra de progresso visual correspondente.

### 2.4 Lista de Membros (Integração SEI)

A tabela exibe todos os usuários identificados na caixa de entrada/lotação do SEI.

* **RF-007 [Legenda de Status]:** O topo da tabela deve conter uma legenda estática indicando os três estados possíveis: `Ativo` (Verde), `Ausente` (Vermelho) e `Pendente` (Amarelo).
* **RF-008 [Listagem Estruturada]:** A tabela deve conter as colunas: *Nome* (com Avatar e iniciais do funcionário), *SEI ID* (identificador de login do sistema integrado), *Função*, *Atribuição*, *Status* e *Ações*.
* **RF-009 [Regra de Negócio - Usuários Pendentes / Em Análise]:** Os usuários identificados com o status **Pendente** (Tag amarela) representam colaboradores que constam na caixa do SEI, mas **ainda não foram ativados no sistema**. Eles permanecem nesta listagem aguardando a aprovação do gestor para começarem a receber fluxos automatizados de processos.
* **RF-010 [Tags de Função]:** A coluna "Função" deve estilizar as atribuições com backgrounds coloridos específicos para identificação rápida. Exemplo:
    * `Sênior`: Cinza/Azul escuro.
    * `Coordenadora`: Roxo.
    * `Pleno`: Amarelo/Laranja claro.
    * `Estagiário` (Estagiário): Laranja.


* **RF-011 [Ações por Membro]:** Cada linha da tabela deve finalizar com um botão de contexto (três pontos verticais `⋮`) para disparar ações individuais (Ativar, Suspender, Alterar Permissões).

### 2.5 Painel de Afastamentos Ativos (Impedimento de Atribuição)

Seção dedicada a listar quem está impedido de receber novos volumes de trabalho.

* **RF-012 [Listagem de Bloqueio]:** Deve exibir em formato de card de alerta os colaboradores em afastamento vigente. O bloco deve conter um ícone de avião/férias, o nome do colaborador, o motivo e o período exato (Ex: `Isabella Costa - Férias | 14/10/2023 até 25/10/2023`).
* **RF-013 [Regra de Distribuição Pausada]:** O sistema deve exibir explicitamente a mensagem explicativa: *"Distribuição automática pausada."* dentro do card. O motor de automação (back-end) deve ignorar este usuário na triagem de novos processos enquanto o período for válido.
* **RF-014 [Ação de Edição]:** O card de afastamento deve disponibilizar um botão `Editar` no canto direito para correções ou encerramento antecipado do período de pausa.

---

## 3. Requisitos Não-Funcionais (RNF)

* **RNF-001 [Consistência de Identidade Visual]:** A interface de Gestão de Equipe deve manter estrito alinhamento com o *Design System* do *Dashboard*, utilizando a mesma família tipográfica, cantos arredondados, bordas sutis e paleta de cores semânticas para os componentes de status.
* **RNF-002 [Sincronização com o SEI]:** A listagem de novos membros que entram em estado "Pendente" (Aguardando ativação) deve ser atualizada de forma assíncrona ou por gatilho de atualização diária ou por solicitação do usuário gestor, evitando sobrecarga nas requisições da API do SEI.
* **RNF-003 [Segurança e Nível de Acesso]:** As ações de `Editar Atribuições`, `Vincular Função`, `Gerenciar Afastamentos` e a ativação de usuários pendentes devem ser exclusivas para usuários com perfil de **Gestor/Administrador**. Analistas comuns devem visualizar a tela apenas em modo de leitura (*Read-Only*).