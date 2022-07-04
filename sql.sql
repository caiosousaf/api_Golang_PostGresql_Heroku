/*
alter table tasks 
ADD COLUMN status varchar(20)

alter table projetos 
ADD COLUMN descricao_projeto varchar(500)

select * from projetos

select * from tasks

insert into projetos (status) values ('Em Andamento')

update projetos set status = 'Em Andamento'

update tasks set status = 'Em Andamento'

update tasks set status = 'Concluido' where id_task = 1



select tk.id_task, tk.descricao_task, tk.pessoa_id, pe.nome_pessoa, tk.projeto_id, pr.nome_projeto, tk.status
 from tasks as tk inner join pessoas as pe on tk.pessoa_id = pe.id_pessoa inner join projetos as pr on tk.projeto_id = pr.id_projeto
 where tk.id_task = 1

select * from tasks where status = 'Em Andamento'
*/