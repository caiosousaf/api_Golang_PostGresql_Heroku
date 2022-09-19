DO $$ DECLARE
    r RECORD;
BEGIN
    -- if the schema you operate on is not "current", you will want to
    -- replace current_schema() in query with 'schematodeletetablesfrom'
    -- *and* update the generate 'DROP...' accordingly.
    FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = current_schema()) LOOP
        EXECUTE 'DROP TABLE IF EXISTS ' || quote_ident(r.tablename) || ' CASCADE';
    END LOOP;
END $$;

CREATE TABLE "equipes" (
  "id_equipe" bigserial PRIMARY KEY,
  "nome_equipe" varchar NOT NULL UNIQUE,
  "data_criacao" date NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE "pessoas" (
  "id_pessoa" bigserial PRIMARY KEY,
  "nome_pessoa" varchar NOT NULL,
  "funcao_pessoa" varchar NOT NULL,
  "equipe_id" bigint,
  "data_contratacao" date NOT NULL DEFAULT CURRENT_DATE
);

CREATE TABLE "projetos" (
  "id_projeto" bigserial PRIMARY KEY,
  "nome_projeto" varchar NOT NULL UNIQUE,
  "descricao_projeto" text NOT NULL,
  "equipe_id" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'A Fazer',
  "data_criacao" date NOT NULL DEFAULT CURRENT_DATE,
  "data_conclusao" date, 
  "prazo_entrega" date 
);

CREATE TABLE "tasks" (
  "id_task" bigserial PRIMARY KEY,
  "descricao_task" text NOT NULL,
  "pessoa_id" bigint NOT NULL,
  "projeto_id" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'A Fazer',
  "prioridade" int NOT NULL, 
  "data_criacao" date NOT NULL DEFAULT CURRENT_DATE,
  "data_conclusao" date, 
  "prazo_entrega" date 
);

CREATE TABLE "users" (
    "id_usuario" bigserial PRIMARY KEY,
    "nome" varchar NOT NULL,
    "email" varchar NOT NULL,
    "password" varchar NOT NULL,
    "data_criacao" DATE NOT NULL DEFAULT CURRENT_DATE
);



ALTER TABLE "pessoas" ADD FOREIGN KEY ("equipe_id") REFERENCES "equipes" ("id_equipe");

ALTER TABLE "projetos" ADD FOREIGN KEY ("equipe_id") REFERENCES "equipes" ("id_equipe");

ALTER TABLE "tasks" ADD FOREIGN KEY ("pessoa_id") REFERENCES "pessoas" ("id_pessoa");

ALTER TABLE "tasks" ADD FOREIGN KEY ("projeto_id") REFERENCES "projetos" ("id_projeto");

INSERT INTO equipes(nome_equipe) VALUES('Komanda');
INSERT INTO equipes(nome_equipe) VALUES('Cariri Inovacao');
INSERT INTO equipes(nome_equipe) VALUES('Rapid Buffalo');
INSERT INTO equipes(nome_equipe) VALUES('Cavaliers');
INSERT INTO equipes(nome_equipe) VALUES('Magic Beasts');
INSERT INTO equipes(nome_equipe) VALUES('Mean Crabs');
INSERT INTO equipes(nome_equipe) VALUES('Swans');
INSERT INTO equipes(nome_equipe) VALUES('Wicked Blitzes');
INSERT INTO equipes(nome_equipe) VALUES('Alpha');
INSERT INTO equipes(nome_equipe) VALUES('Grupo Futurista');

INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES('Caio Sousa', 'Back-End', 1);
INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES('Vinicius Guedes', 'Back-End', 2);
INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES('Larissa Cardoso', 'Front-End', 3);
INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES('Pedro Miguel', 'Front-End', 4);
INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES('Bruno Nascimento', 'Back-End', 5);
INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES('Kellie Phelps', 'Analista', 6);
INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES('Ivor Morris', 'Tester', 7);
INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES('Len Kerr', 'Back-End', 8);
INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES('Althea Campos', 'Front-End', 9);
INSERT INTO pessoas(nome_pessoa, funcao_pessoa, equipe_id) VALUES('Dawn Edwards', 'Back-End', 10);

INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES('Blue', 'Class rutrum euismod nisl ac vitae hendrerit quam vitae aliquam primis, elit nulla vel euismod urna mauris nullam sit ultrices. ', 1, current_date+20);
INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES('Diamond', 'Lorem scelerisque ultrices ad mollis platea tellus auctor, aliquam aliquet curabitur potenti phasellus adipiscing, orci ante donec lacinia et interdum. ', 2, current_date+18);
INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES('Falcon', 'Iaculis lectus elit sem platea commodo aenean netus, ornare nunc etiam fermentum augue mi, duis pretium per massa tincidunt ad. ', 3, current_date+16);
INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES('Phoenix', 'Quis interdum tincidunt vivamus porta magna, eros pharetra nibh nunc. nisi pretium ornare venenatis suscipit himenaeos, aliquet vivamus massa at. ', 4, current_date+14);
INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES('Eagle', 'Mattis sed nulla mattis rhoncus vitae risus euismod auctor cubilia nam, convallis metus ad rutrum orci dapibus inceptos facilisis nunc. ', 5, current_date+12);
INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES('Lion', 'Est imperdiet mollis aliquet lectus orci faucibus cras, taciti nulla venenatis augue sagittis euismod, maecenas magna velit nulla sed tristique. ', 6, current_date+10);
INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES('X', 'Ad nullam accumsan donec nisi dolor etiam nisl gravida urna augue, inceptos sem hac porta viverra vivamus vestibulum tincidunt vulputate. ', 7, current_date+15);
INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES('Green', 'Proin ut magna et placerat aliquam magna congue, fames curabitur senectus torquent nulla nisi, aliquet eleifend mollis aptent hendrerit semper. ', 8, current_date+17);
INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES('Sapphire', 'Dictumst mollis porttitor ultricies dictumst dolor sapien netus, in ipsum vel lectus orci at, litora at luctus nam egestas lobortis. ', 9, current_date+13);
INSERT INTO projetos(nome_projeto, descricao_projeto, equipe_id, prazo_entrega) VALUES('Panther', 'Primis eu elementum lorem iaculis metus lorem taciti, primis nunc mi lectus scelerisque egestas, sem eu dapibus torquent potenti ut. ', 10, current_date+9);


INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) VALUES('Class rutrum euismod nisl ac vitae hendrerit quam vitae aliquam primis, elit nulla vel euismod urna mauris nullam sit ultrices.', 1, 10, 2, current_date+20);
INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) VALUES('Lorem scelerisque ultrices ad mollis platea tellus auctor, aliquam aliquet curabitur potenti phasellus adipiscing, orci ante donec lacinia et interdum.', 2, 9, 2, current_date+18);
INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) VALUES('Iaculis lectus elit sem platea commodo aenean netus, ornare nunc etiam fermentum augue mi, duis pretium per massa tincidunt ad.', 3, 8, 2, current_date+16);
INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) VALUES('Quis interdum tincidunt vivamus porta magna, eros pharetra nibh nunc. nisi pretium ornare venenatis suscipit himenaeos, aliquet vivamus massa at.', 4, 7, 2, current_date+14);
INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) VALUES('Mattis sed nulla mattis rhoncus vitae risus euismod auctor cubilia nam, convallis metus ad rutrum orci dapibus inceptos facilisis nunc.', 5, 6, 2, current_date+12);
INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) VALUES('Est imperdiet mollis aliquet lectus orci faucibus cras, taciti nulla venenatis augue sagittis euismod, maecenas magna velit nulla sed tristique.', 6, 5, 1, current_date+15);
INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) VALUES('Ad nullam accumsan donec nisi dolor etiam nisl gravida urna augue, inceptos sem hac porta viverra vivamus vestibulum tincidunt vulputate.', 7, 4, 1, current_date+13);
INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) VALUES('Proin ut magna et placerat aliquam magna congue, fames curabitur senectus torquent nulla nisi, aliquet eleifend mollis aptent hendrerit semper.', 8, 3, 1, current_date+11);
INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) VALUES('Dictumst mollis porttitor ultricies dictumst dolor sapien netus, in ipsum vel lectus orci at, litora at luctus nam egestas lobortis.', 9, 2, 2, current_date+9);
INSERT INTO tasks(descricao_task, pessoa_id, projeto_id, prioridade, prazo_entrega) VALUES('Primis eu elementum lorem iaculis metus lorem taciti, primis nunc mi lectus scelerisque egestas, sem eu dapibus torquent potenti ut.', 10, 1, 1, current_date+7);

INSERT INTO users(nome, email, password) VALUES('Caio Sousa', 'caio.admin@email.com', 'root');
INSERT INTO users(nome, email, password) VALUES('Usuario Root', 'root@root.com', 'salmo34');
INSERT INTO users(nome, email, password) VALUES('Usuario Comum', 'comum@comum.com', 'usuario');