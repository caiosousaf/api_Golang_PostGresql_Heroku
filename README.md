![Badge em Desenvolvimento](http://img.shields.io/static/v1?label=STATUS&message=EM%20DESENVOLVIMENTO&color=GREEN&style=for-the-badge)
<br>
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Postgres](https://img.shields.io/badge/postgres-%23316192.svg?style=for-the-badge&logo=postgresql&logoColor=white)
![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)
![Heroku](https://img.shields.io/badge/heroku-%23430098.svg?style=for-the-badge&logo=heroku&logoColor=white)
![Vercel](https://img.shields.io/badge/vercel-%23000000.svg?style=for-the-badge&logo=vercel&logoColor=white)
![cypress](https://img.shields.io/badge/-cypress-%23E5E5E5?style=for-the-badge&logo=cypress&logoColor=058a5e)
## API com Banco de Dados e Front-End - Aprendizes Brisanet

Sistema que mantém projetos. Cada projeto recebe uma equipe, e cada equipe pode receber um número não delimitado de pessoas. Cada projeto tem tasks,
cada task pode ser atribuída a uma determinada pessoa que está na equipe do projeto.

## Índice 

* [Detalhes](#detalhes)
* [Equipe](#equipe)
* [Andamento do Projeto](#andamento-do-projeto)
* [Rotas](#rotas)
* [Projetos](#projetos)
* [Pessoas](#pessoas)
* [Equipes](#equipes)
* [Tasks](#tasks)

## Detalhes

- Utilizando linguagem `Go` com `Gin` para desenvolvimento, e o software `Insomnia` para testes;
- `GORM` para conexão com banco de dados `PostgreSQL`;
- Disponível no `Heroku`: https://sistema-aprendizes-brisanet-go.herokuapp.com/
- Front-End em `React`(Em desenvolvimento)


## Equipe:
- <a href="https://github.com/Brun0Nasc"> Bruno do Nascimento</a>: `Reestruturação da API` `Implementação do Banco de Dados` `Revisão e Reestruturação de Rotas e Funções` `Criação do BD` `Definição das Consultas do BD.`
- <a href="https://github.com/Lucasmartinsn"> Lucas Martins:</a> `Criação do Front-End (em desenvolvimento)`
- <a href="https://github.com/IaraFV"> Iara Ferreira:</a> `Criação do Front-End (em desenvolvimento)`


## Andamento do projeto

| Funcionalidade        | Estado |
| ------------- |:-------------:|
| Manter equipe      | ✔️ |
| Manter projeto      | ✔️ |
| Associar equipe a projeto | ✔️ | 
| Criar tarefa no projeto | ✔️ | 
| Atribuir tarefa | ✔️ | 
| Manter dados no Banco de Dados | ✔️ | 
| Front-End | ⌛ |
| Testes e2e com Cypress | ❌ | 

## ROTAS

<h4>PROJETOS</h4>

```
https://sistema-aprendizes-brisanet-go.herokuapp.com/projetos
```

<table border=2>
<tr>
 <th>GET</th>
 <th>POST</th>
 <th>PUT</th>
 <th>DELETE</th>
</tr>
<tr>
  <td>
   <div>/projetos</div>
  </td>
  <td>
   <div>/projetos</div>
  </td>
  <td>
   <div>/projetos/:id</div>
  </td>
  <td>
   <div>/projetos/:id</div>
  </td>
 </tr>
 <tr>
  <td>
   <div>/projetos/:id</div>
  </td>
  <td></td>
  <td>
   <div>/projetos/:id/status</div>
  </td>
  <td></td>
 </tr>
</table>

<table border=1>
<th>
 DETALHES
</th>
<tr>
<td>
<p>As funções <b>GET</b> retornam os seguintes dados:</p>
<ul>
<li>ID do projeto</li>
<li>Nome do projeto</li>
<li>Status do projeto</li>
<li>Data de início</li>
<li>Data de conclusão</li>
<li>Equipe responsável pelo projeto</li>
<li>Tasks atribuídas ao projeto</li>
</ul>
</td>
</tr>

<tr>
<td>
<p>Para realizar um novo registro com a função <b>POST</b>, os seguintes dados deverão ser informados:</p>
<ul>
 <li>Nome do projeto</li>
 <li>ID da equipe que ficará responsável pelo projeto</li>
</ul>
<p>Dados como Status, Data de Inicio e Data de Conclusão são automáticos, todo projeto inicia automáticamente com o status
"Em planejamento". É possível alterar o status de desenvolvimento do projeto, mas apenas um projeto por equipe pode estar
com o status "Em desenvolvimento".</p>
</td>
</tr>

<tr>
<td>
<p>Nas funções <b>PUT</b>, dois tipos de atualização podem ser feitas:</p>
<ul>
 <li>Mudança de nome e/ou equipe do projeto</li>
 <li>Atualização do status do projeto</li>
</ul>
</td>
</tr>

<tr>
<td>
<p>Para deletar um projeto, através do <b>DELETE</b>, será necessário apenas passar o ID do projeto como parâmetro na rota.</p>
</td>
</tr>
</table>
<hr>

<h4>PESSOAS</h4>

```
https://sistema-aprendizes-brisanet-go.herokuapp.com/pessoas
```

<table border=2>
<tr>
 <th>GET</th>
 <th>POST</th>
 <th>PUT</th>
 <th>DELETE</th>
</tr>

<tr>
 <td>/pessoas</td>
 <td>/pessoas</td>
 <td>/pessoas/:id</td>
 <td>/pessoas/:id</td>
</tr>

<tr>
 <td>/pessoas/:id</td>
 <td>/pessoas/:id/tasks</td>
 <td></td>
 <td></td>
</tr>

<tr>
 <td>/pessoas/:id/tasks</td>
 <td></td>
 <td></td>
 <td></td>
</tr>
</table>

<table border=1>
 <tr>
  <th>
   DETALHES 
  </th>
 </tr>
 <tr>
  <td>
   <p>As funções <b>GET</b> retornam as seguintes informações</p>
   <ul>
    <li>ID da pessoa</li>
    <li>Nome da pessoa</li>
    <li>Data de contratação</li>
    <li>Informações da equipe em que ela está</li>
   </ul>
  </td>
 </tr>
 <tr>
  <td>
   <p>As funções <b>POST</b> permitem realizar as seguintes ações:</p>
   <ul>
    <li>Registrar uma nova pessoa no Banco de Dados</li>
    <li>Registrar novas tasks para uma pessoa</li>
   </ul>
   <p>Para cadastrar uma nova pessoa, os seguintes dados são necessários:</p>
   <ul>
    <li>Nome da pessoa</li>
    <li>Função da pessoa</li>
    <li>Equipe que ela fará parte (pode ser registrada sem equipe)</li>
   </ul>
   <p>Para atribuir uma nova task a uma pessoa, será necessário passar apenas a descrição da Task.</p>
   <p>- Informações como ID do Projeto e Status da task serão automáticas. Para que uma task seja atribuída a uma pessoa, essa pessoa precisa estar em uma equipe, e, por sua vez, esta equipe precisa estar atrelada a um projeto que esteja em desenvolvimento.</p>
  </td>
 </tr>
 <tr>
  <td>
   <p>A função <b>PUT</b> permite alterar os seguintes dados:</p>
   <ul>
    <li>Nome da pessoa</li>
    <li>Função da pessoa</li>
    <li>ID da equipe (podendo ser nulo)</li>
   </ul>
  </td>
 </tr>
 <tr>
  <td>
   <p>Para deletar uma pessoa, através do <b>DELETE</b>, será necessário apenas passar o ID da pessoa como parâmetro na rota.</p>
  </td>
 </tr>
</table>

<hr>

<h4>EQUIPES</h4>

```
https://sistema-aprendizes-brisanet-go.herokuapp.com/equipes
```

<table border=2>
  <tr>
    <th>GET</th>
    <th>POST</th>
    <th>PUT</th>
    <th>DELETE</th>
  </tr>
  <tr>
    <td>/equipes</td>
    <td>/equipes</td>
    <td>/equipes/:id</td>
    <td>/equipes/:id</td>
  </tr>
  <tr>
    <td>/equipes/:id</td>
    <td></td>
    <td></td>
    <td></td>
  </tr>
  <tr>
    <td>/equipes/:id/projetos</td>
    <td></td>
    <td></td>
    <td></td>
  </tr>
</table>
 
<table border = 1>
  <tr>
    <th>DETALHES</th>
  </tr>

  <tr>
    <td>
      <p>As funções <b>GET</b> retornam as seguintes informações:</p>
      <ul>
        <li>ID da equipe</li>
        <li>Nome da equipe</li>
        <li>Pessoas que estão na equipe</li>
      </ul>
      <p>O GET de projetos de uma equipe retorna todos os projetos associados à equipe do ID informado na rota.</p>
    </td>
  </tr>
  <tr>
    <td>
     <p>Para realizar um novo registro com a função <b>POST</b>, apenas o nome da equipe deverá ser informado.</p>
    </td>
  </tr>
  <tr>
    <td>
      <p>Assim como no POST, a função <b>PUT</b> irá alterar apenas o nome da equipe, recebendo o id na rota como parâmetro.</p>
    </td>
  </tr>
  <tr>
    <td>
      <p>Para deletar uma equipe, através do <b>DELETE</b>, será necessário apenas passar o ID da equipe como parâmetro na rota.</p>
    </td>
 </tr>
</table>

<hr>

<h4>TASKS</h4>

```
https://sistema-aprendizes-brisanet-go.herokuapp.com/tasks
```

<table border=2>
<tr>
 <th>GET</th>
 <th>POST</th>
 <th>PUT</th>
 <th>DELETE</th>
</tr>

<tr>
 <td>/tasks</td>
 <td>/tasks</td>
 <td>/tasks/:id</td>
 <td>/tasks/:id</td>
<tr>

<tr>
 <td>/tasks/:id</td>
 <td></td>
 <td>/tasks/:id/status</td>
 <td></td>
<tr>
</table>
 
 <table border = 1>
  <tr>
    <th>DETALHES</th>
  </tr>

  <tr>
    <td>
      <p>As funções <b>GET</b> retornam as seguintes informações:</p>
      <ul>
        <li>ID da task</li>
        <li>Descrição</li>
        <li>Status</li>
        <li>ID da pessoa responsável pela task</li>
        <li>Nome da pessoa responsável</li>
        <li>ID do projeto que a task está associada</li>
        <li>Nome do projeto</li>
      </ul>
    </td>
  </tr>
  <tr>
    <td>
     <p>Para realizar um novo registro com a função <b>POST</b>, os seguintes dados devem ser informados:</p>
     <ul>
       <li>Descrição da Task</li>
       <li>ID da pessoa que ficará responsável</li>
       <li>ID do projeto associado</li>
     </ul>
     <p>Para atribuir uma task a uma pessoa, essa pessoa precisa estar na equipe que está responsável pela pelo projeto que a task será atribuída, e esse projeto precisa estar com o status "Em desenvolvimento".</p>
    </td>
  </tr>
  <tr>
    <td>
      <p>A função <b>PUT</b> apenas com ID poderá ser usada para alterar informações de Descrição, Pessoa e Projeto.</p>
      <p>O segundo PUT, relacionado ao status, só altera o Status da Task.</p>
    </td>
  </tr>
  <tr>
    <td>
      <p>Para deletar uma Task, através do <b>DELETE</b>, será necessário apenas passar o ID da Task como parâmetro na rota.</p>
    </td>
 </tr>
</table>
