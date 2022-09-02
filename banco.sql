CREATE TABLE "equipes" (
  "id_equipe" bigserial PRIMARY KEY,
  "nome_equipe" varchar NOT NULL
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
  "nome_projeto" varchar NOT NULL,
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
  "status" varchar NOT NULL DEFAULT 'Em planejamento',
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