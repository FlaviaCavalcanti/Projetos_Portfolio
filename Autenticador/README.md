###Autenticador

Este projeto implementa uma API simples de login e registro de usuários utilizando Go. O objetivo é praticar conceitos como autenticação, manipulação de mapas e servidores HTTP.

##Funcionalidades:

Login: Verifica se o email e a senha fornecidos são válidos.
Registro: Permite o cadastro de novos usuários com email e senha.

##Endpoints:

POST /login
Requer os parâmetros email e senha. Retorna status 200 se as credenciais estiverem corretas, ou 401 se forem inválidas.

POST /register
Requer os parâmetros email e senha. Retorna status 200 se o registro for bem-sucedido, ou 409 se o usuário já existir.

##Como rodar:

Clone o repositório.
Execute go run main.go.
Acesse o servidor local em http://localhost:8080.
