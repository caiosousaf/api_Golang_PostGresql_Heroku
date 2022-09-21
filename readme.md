<div align="center" display="flex">
  <img height="70px"  src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original-wordmark.svg" />
  <img height="70px" src='https://upload.wikimedia.org/wikipedia/commons/thumb/2/29/Postgresql_elephant.svg/1200px-Postgresql_elephant.svg.png'>
  <img height="70px" src='https://avatars.githubusercontent.com/u/53864671?v=4'>
  <h1 align="center">API com Golang, Gin-Gonic e PostgreSQL. Documentação e testes com Swagger</h1>
</div>

![Badge Concluido](http://img.shields.io/static/v1?label=STATUS&message=EM-DESENVOLVIMENTO&color=GREEN&style=for-the-badge)
<img height="27px" src= 'https://camo.githubusercontent.com/f5d55c1699aa4ae2a226b74cb942c1d1efc8e9d5b2461ae78f65a116c989c2c6/68747470733a2f2f636972636c6563692e636f6d2f67682f434952434c4543492d4757502f676f6c616e672d636f6d70616e792d6170692f747265652f6d61696e2e7376673f7374796c653d737667'>

<p>Disponível para acesso em: https://golang-posgre-brisanet.herokuapp.com/</p>
<br>
<p>Disponível para testes com Swagger em: https://golang-posgre-brisanet.herokuapp.com/swagger/index.html#/</p>

# Desenvolvedor

<ul>
  <li><a href="https://github.com/caiosousaf">Caio Sousa</a></li>
</ul>

# Sobre o Sistema

- Um sistema para manter projetos, o sistema deve cadastrar projetos e equipes, um projeto possui uma equipe e deve ter tarefas dentro do projeto onde os membros da equipe podem se atribuir;

- Foi desenvolvido com [Golang](https://go.dev/), [Gin-Gonic](https://gin-gonic.com/) e [PostgreSQL](https://www.postgresql.org);
- A API e o Banco de dados estão estão sendo mantidos no [Heroku](https://www.heroku.com).
- A documentação e os testes das rotas foram feitas com o [Swagger](https://swagger.io/)

# Funcionalidades

| Funcionalidade                         | Estado |
| -------------------------------------- | :----: |
| Manter equipe                          |   ✔️   |
| Manter projeto                         |   ✔️   |
| Associar equipe a projeto              |   ✔️   |
| Criar tarefa no projeto                |   ✔️   |
| Atribuir tarefa                        |   ✔️   |
| Utilizar algum SGBD (PostgreSQL/MySQL) |   ✔️   |
| Criar Usuário                          |   ✔️   |
| Login/Autenticar Usuário               |   ✔️   |
| Logout Usuário                         |   X   |

# Relacionamento dos Elementos

<img   src="BD.png" />

# Avisos

- O tratamento correto dos erros ainda está em andamento.
- End Points de cadastro e login de usuarios ainda estão em fase de testes
- A documentação ainda está em constante atualização