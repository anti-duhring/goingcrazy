CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE "people" (
    "id" uuid DEFAULT gen_random_uuid(),
    "created_at" timestamptz,
    "updated_at" timestamptz,
    "deleted_at" timestamptz,
    "apelido" varchar(32) NOT NULL UNIQUE,
    "nome" varchar(100) NOT NULL,
    "search_index" text NOT NULL,
    "nascimento" date NOT NULL,
    "stack" JSONB,
    PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS idx_people_search ON public.people USING gist (search_index public.gist_trgm_ops (siglen='64'));

ALTER DATABASE postgres SET synchronous_commit=OFF;
-- using 25% of memory as suggested in the docs:
--    https://www.postgresql.org/docs/9.1/runtime-config-resource.html
ALTER SYSTEM SET shared_buffers TO "425MB";
-- debug slow queries, run \d pg_stat_statements
-- docs: 
--    https://www.postgresql.org/docs/current/pgstatstatements.html
-- CREATE EXTENSION pg_stat_statements;
-- ALTER SYSTEM SET shared_preload_libraries = 'pg_stat_statements';