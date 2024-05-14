# Primeira API usando GO puro

## Descrição

Este projeto é uma aplicação do back-end de fórum que serve como uma excelente oportunidade para aplicar e consolidar os conhecimentos adquiridos em Golang. A aplicação foi projetada para demonstrar a forma de criar endpoints, estruturar e manipular um banco de dados. Além disso, fazer o CRUD de um sistema simples que simula uma "rede social".

## Visão Geral do Projeto

**Back-End**

O back-end foi construído utilizando Golang (Go) e é responsável por:

* Gerenciamento de Rotas: Controlar as rotas de navegação e garantir que as requisições dos usuários sejam direcionadas corretamente.
* Processamento de Dados: Manipular e processar os dados recebidos e enviados ao front-end.
* Autenticação e Autorização: Implementar o sistema de cadastro e login de usuários, garantindo segurança e privacidade.
* Interação com o Banco de Dados: Gerenciar a persistência de dados, incluindo postagens, curtidas e seguidores.

## End points

* /login (POST): Fazer o processo de login do usuário.
* /publication (POST): Criar uma postagem.
* /publications (GET): Listagem de todas as postagens.
* /publication/{id} (GET): Requisição de uma única postagem.
* /publication/{id} (PUT): Atualização de uma postagem.
* /publication/{id} (DELETE): Deletar uma postagem.
* /users/{id}/publications (GET): Listagem das postagens de um usuário.
* /publication/{id}/like (POST): Like em uma postagem.
* /publication/{id}/dislike (POST): Dislike em uma postagem.
* /users (POST): Criar um usuário.
* /users (GET): Listagem dos usuários.
* /users/{id} (GET): Listar um usuário.
* /users/{id} (PUT): Atualizar um usuário.
* /users/{id} (DELETE): Deletar um usuário.
* /users/{id}/follow (POST): Seguir um usuário.
* /users/{id}/unfollow (POST): Deixar de seguir um usuário.
* /users/{id}/followers (GET): Listagem dos seguidores de um usuário.
* /users/{id}/following (GET): Listagem de quem o usuário está seguindo.
* /users/{id}/update-password (POST): Atualizar a senha.

## Referências

- [Golang Doc URL](https://go.dev/doc/)
- [Udemy](https://www.udemy.com/)
- [Chat GPT](https://chat.openai.com/)
- [TabNews](https://www.tabnews.com.br/) 