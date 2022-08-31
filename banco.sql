CREATE TABLE "equipes" (
  "id_equipe" bigserial PRIMARY KEY,
  "nome_equipe" varchar NOT NULL
);

CREATE TABLE "pessoas" (
  "id_pessoa" bigserial PRIMARY KEY,
  "nome_pessoa" varchar NOT NULL,
  "funcao_pessoa" varchar NOT NULL,
  "equipe_id" bigint,
  "data_contratacao" timestamp NOT NULL DEFAULT (now()),
  "favoritar" int NOT NULL DEFAULT 0
);

CREATE TABLE "projetos" (
  "id_projeto" bigserial PRIMARY KEY,
  "nome_projeto" varchar NOT NULL,
  "descricao_projeto" text NOT NULL,
  "equipe_id" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'Em planejamento',
  "data_inicio" timestamp NOT NULL DEFAULT (now()),
  "data_conclusao" timestamp
);

CREATE TABLE "tasks" (
  "id_task" bigserial PRIMARY KEY,
  "descricao_task" text NOT NULL,
  "pessoa_id" bigint NOT NULL,
  "projeto_id" bigint NOT NULL,
  "status" varchar NOT NULL DEFAULT 'Em planejamento',
  "nivel" varchar NOT NULL
);

CREATE INDEX ON "equipes" ("id_equipe");

CREATE INDEX ON "pessoas" ("id_pessoa");

CREATE INDEX ON "pessoas" ("nome_pessoa");

CREATE INDEX ON "pessoas" ("equipe_id");

CREATE INDEX ON "projetos" ("id_projeto");

CREATE INDEX ON "projetos" ("nome_projeto");

CREATE INDEX ON "projetos" ("equipe_id");

CREATE INDEX ON "tasks" ("id_task");

CREATE INDEX ON "tasks" ("pessoa_id");

CREATE INDEX ON "tasks" ("projeto_id");

COMMENT ON COLUMN "pessoas"."equipe_id" IS 'pode ser nulo';

COMMENT ON COLUMN "projetos"."equipe_id" IS 'um projeto tem que estar obrigatoriamente relacionado a uma equipe';

COMMENT ON COLUMN "projetos"."data_conclusao" IS 'será preenchido automaticamente quando um projeto for marcado como concluído';

ALTER TABLE "pessoas" ADD FOREIGN KEY ("equipe_id") REFERENCES "equipes" ("id_equipe");

ALTER TABLE "projetos" ADD FOREIGN KEY ("equipe_id") REFERENCES "equipes" ("id_equipe");

ALTER TABLE "tasks" ADD FOREIGN KEY ("pessoa_id") REFERENCES "pessoas" ("id_pessoa");

ALTER TABLE "tasks" ADD FOREIGN KEY ("projeto_id") REFERENCES "projetos" ("id_projeto");