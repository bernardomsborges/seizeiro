# Guia de Design, Requisitos e Integrações Técnicas

Este documento centraliza os recursos visuais, especificações de requisitos e documentações de APIs externas necessários para o desenvolvimento do front-end e back-end da plataforma.

---

## 1. Arquivos de Requisitos do Projeto
Para compreender o escopo funcional detalhado de cada módulo, consulte os arquivos de especificação técnica localizados no diretório `/Requisitos`:

* [📄 **dashboard.md**](./dashboard.md): Especificação do Dashboard de Governança (Filtros e listagem de processos).
* [📄 **detalhamento-do-processo.md**](./detalhamento-do-processo.md): Mapeamento de campos da aba de dados gerais, histórico de ciências e anotações.
* [📄 **gestao-equipe.md**](./gestao-equipe.md): Regras para controle de membros, usuários pendentes (caixa do SEI) e gerenciamento de afastamentos.

---

## 2. Recursos de Design e UI/UX (Figma)

> 💡 **Nota para os Desenvolvedores:** Utilize a ferramenta de *Inspect (Dev Mode)* do Figma para extrair os espaçamentos (`padding`/`margin`), raios de borda (`border-radius`) e paletas hexadecimais de cores semânticas das tags e badges de status.

* 🖥️ [**Link do Protótipo no Figma**](https://www.figma.com/design/fSF1dYZM1Qz2CnkL1t2qE3/SEIZEIRO?node-id=0-1&t=9E6zmOl1jahVTrQT-1)
  * **Componentes Principais:** Sidebar de navegação do *Seizeiro Intelligence*, Badges de Status (`Em Triagem`, `Análise Pendente`, `Ativo`, `Pendente`), e Avatares de Usuários.
  * _Obs:_ Para opção de edição, favor solicitar.

---

## 3. Documentação das APIs e Links Úteis

O projeto consome dados integrados diretamente do barramento de serviços do Sistema Eletrônico de Informações (SEI). Utilize as referências abaixo para validação de payloads e autenticação:

### 🔗 Ambientes e Portais Oficiais
* **Portal SEI - Minas Gerais:** [sei.mg.gov.br/sei/](https://www.sei.mg.gov.br/sei/)
  * Referência de ambiente para validação do comportamento padrão da interface oficial e nomenclatura de campos institucionais.

### 🛠️ Documentação Swagger / Swagger UI (Web Services SEI)
* **API de Listagem de Tipos de Processo:** [Documentação de Métodos - WSSEI](https://pengovbr.github.io/mod-wssei/#/Processo/listarTipoProcesso)
  * Utilize este endpoint para cruzar as informações de permissões de documentos sigilosos (`permiteSigiloso`) e nomes amigáveis de procedimentos na abertura de processos.

---

## 4. Convenções de Código para o Front-End (UI-Kit)
* **Ícones:** Utilizar preferencialmente a biblioteca correspondente ao layout do Figma (Ex: *Lucide Icons* ou *Material Icons*).
* **Tratamento de Cores dos Marcadores:** Mapear o campo `idCor` retornado pelo endpoint `/processo/{protocolo}` diretamente com as classes de estilização no arquivo de estilos globais para garantir fidelidade visual com as tags originais do SEI.