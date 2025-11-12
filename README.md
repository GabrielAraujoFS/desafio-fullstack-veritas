#  Desafio Fullstack - Veritas Consultoria

Este projeto consiste em um **Mini Kanban de Tarefas**, desenvolvido com **Go (Golang)** no **backend** e **React** no **frontend**, como parte do **Desafio Fullstack da Veritas Consultoria**.

O sistema permite **criar**, **editar**, **excluir** e **mover tarefas** entre colunas, simulando um ambiente real de organização de atividades.

---

##  Tecnologias Utilizadas

###  Backend
- **Linguagem:** Go
- **Framework:** net/http nativo
- **Rotas:** `GET`, `POST`, `PUT`, `DELETE` em `/tasks`
- **Armazenamento:** Em memória (sem banco de dados)
- **Porta:** `8080`

###  Frontend
- **Linguagem:** JavaScript (ES6+)
- **Framework:** React
- **Gerenciamento de Estado:** useState / useEffect
- **Integração com API:** fetch
- **Estilização:** CSS puro com layout responsivo baseado em Kanban

---

##  Como Rodar o Projeto

###  Backend
```bash
cd backend
go run main.go

Servidor disponível em:
----> http://localhost:8080
```
### Frontend
```bash
cd frontend
npm install
npm start

Aplicação disponível em:
----> http://localhost:3000
```
## Funcionalidades Implementadas

-  Criar novas tarefas com título e descrição
- Validar título obrigatório
- Exibir tarefas nas colunas A Fazer, Em Progresso e Concluído
- Editar tarefas existentes
- Excluir tarefas
- Arrastar e soltar (drag and drop) entre colunas
- Atualização automática da interface
- Comunicação com API em Go via requisições HTTP
  
## Fluxo de Uso (User Flow)

O diagrama abaixo representa o fluxo principal de uso do sistema, desde o acesso à página até a manipulação das tarefas no quadro Kanban:

### <img width="561" height="681" alt="Diagrama veritas" src="https://github.com/user-attachments/assets/cdff3641-8588-4732-8443-eb58a91fcd01" />

## Melhorias Futuras

- Persistência de dados com banco de dados (ex: SQLite ou PostgreSQL)

- Autenticação de usuários

- Interface mais personalizada e responsiva

- Deploy do projeto completo na web

## 👨‍💻  Autor


### Gabriel Araujo farias de santana
### Estudante de Ciência da Computação | Desenvolvedor em formação

## Observação:
O projeto foi desenvolvido para fins acadêmicos e de avaliação técnica no Desafio Fullstack da Veritas Consultoria, com foco em boas práticas de desenvolvimento e integração entre frontend e backend.
