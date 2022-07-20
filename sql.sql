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

    select pr.id_projeto, pr.nome_projeto from projetos as pr
	inner join pessoas as pe
	on pe.equipe_id = pr.equipe_id 
	where pe.id_pessoa = 4






	11/07/2022

select pr.id_projeto, pr.nome_projeto, pr.equipe_id, eq.nome_equipe, pr.status, pr.descricao_projeto, 
pr.data_criacao, pr.data_conclusao
from projetos as pr inner join equipes as eq on pr.equipe_id = eq.id_equipe

select * from tasks

select pr.id_projeto, pr.nome_projeto, eq.nome_equipe,tk.id_task, tk.descricao_task, tk.status, 
pe.nome_pessoa from 
projetos as pr inner join tasks as tk on pr.id_projeto = tk.projeto_id inner join
equipes as eq on pr.equipe_id = eq.id_equipe inner join
pessoas as pe on pe.id_pessoa = tk.pessoa_id where id_projeto = 3 


20/07/2022

select current_date + interval '26 DAYS' AS Data;
*/ 