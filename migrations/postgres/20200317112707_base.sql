-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE public.users
(
    id          uuid      NOT NULL DEFAULT uuid_generate_v4(),
    login       text      NOT NULL,
    description text      NULL,
    price       numeric   NOT NULL DEFAULT 0,
    summa_one   float4    NOT NULL DEFAULT 0,
    summa_two   float8    NOT NULL DEFAULT 0,
    cnt         int4      NOT NULL DEFAULT 0,
    cnt2        int2      NOT NULL DEFAULT 0,
    cnt4        int4      NOT NULL DEFAULT 0,
    cnt8        int8      NOT NULL DEFAULT 0,
    is_online   bool      NOT NULL DEFAULT false,
    metrika     jsonb     NULL,
    duration    int8      NOT NULL DEFAULT 0,
    "data"      bytea     NULL,
    alias       _text     NULL,
    created_at  timestamp NOT NULL DEFAULT now(),
    updated_at  timestamp NOT NULL DEFAULT now(),
    deleted_at  timestamp NULL,
    CONSTRAINT users_pk PRIMARY KEY (id)
);
CREATE TRIGGER users_updated_at
    BEFORE UPDATE
    ON public.users
    FOR EACH ROW
EXECUTE PROCEDURE moddatetime(updated_at);

CREATE TABLE public.roles
(
    id          uuid NOT NULL DEFAULT uuid_generate_v4(),
    code        text NOT NULL,
    description text NOT NULL,
    CONSTRAINT roles_pk PRIMARY KEY (id)
);

CREATE TABLE public.users_roles
(
    user_id uuid NOT NULL,
    role_id uuid NOT NULL,
    CONSTRAINT users_roles_pk PRIMARY KEY (user_id, role_id)
);
ALTER TABLE public.users_roles
    ADD CONSTRAINT users_roles_fk FOREIGN KEY (user_id) REFERENCES users (id) ON
        UPDATE CASCADE ON DELETE CASCADE;
ALTER TABLE public.users_roles
    ADD CONSTRAINT users_roles_fk_1 FOREIGN KEY (role_id) REFERENCES roles (id) ON
        UPDATE CASCADE ON DELETE CASCADE;

CREATE TABLE public.orders
(
    id         uuid      NOT NULL DEFAULT uuid_generate_v4(),
    user_id    uuid      NULL,
    "number"   int4      NOT NULL DEFAULT 0,
    status     text      NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP NULL,
    CONSTRAINT orders_pk PRIMARY KEY (id)
);
CREATE TRIGGER orders_updated_at
    BEFORE UPDATE
    ON public.orders
    FOR EACH ROW
EXECUTE PROCEDURE moddatetime(updated_at);
ALTER TABLE public.orders
    ADD CONSTRAINT orders_fk FOREIGN KEY (user_id) REFERENCES users (id) ON
        UPDATE CASCADE ON DELETE RESTRICT;

INSERT INTO public.users
(id, login, price, summa_one, summa_two, cnt2, cnt4, cnt8, is_online, metrika, created_at, updated_at)
VALUES (uuid_generate_v4(), 'testLogin', 0, 0, 0, 0, 0, 0, false, '{}', now(), now());

CREATE TABLE public.minio_st
(
    id         uuid        NOT NULL DEFAULT uuid_generate_v4(), -- ИД
    bucket     text        NOT NULL,                            -- папка хранения - тип объекта
    object_id  uuid        NOT NULL,                            -- файл хранения - ид объекта
    "name"     text        NOT NULL,                            -- имя файла
    file_type  text        NOT NULL,                            -- тип файла
    file_size  int4        NOT NULL DEFAULT 0,                  -- размер файла
    "label"    jsonb       NULL,                                -- дополнительные параметры файла
    user_login text        NOT NULL,                            -- пользователь
    created_at timestamptz NOT NULL DEFAULT now(),              -- дата и время создания
    is_confirm bool        NOT NULL DEFAULT false,              -- подтверждение загрузки
    CONSTRAINT minio_st_pk PRIMARY KEY (id)
);
CREATE INDEX minio_st_bucket_idx ON public.minio_st USING btree (bucket);
CREATE UNIQUE INDEX minio_st_object_id_idx ON public.minio_st USING btree (object_id);

-- Column comments

COMMENT ON COLUMN public.minio_st.id IS 'ИД';
COMMENT ON COLUMN public.minio_st.bucket IS 'папка хранения - тип объекта';
COMMENT ON COLUMN public.minio_st.object_id IS 'файл хранения - ид объекта';
COMMENT ON COLUMN public.minio_st."name" IS 'имя файла';
COMMENT ON COLUMN public.minio_st.file_type IS 'тип файла';
COMMENT ON COLUMN public.minio_st.file_size IS 'размер файла';
COMMENT ON COLUMN public.minio_st."label" IS 'дополнительные параметры файла';
COMMENT ON COLUMN public.minio_st.user_login IS 'пользователь';
COMMENT ON COLUMN public.minio_st.created_at IS 'дата и время создания';
COMMENT ON COLUMN public.minio_st.is_confirm IS 'подтверждение загрузки';


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE public.orders;

DROP TABLE public.users_roles;

DROP TABLE public.users;

DROP TABLE public.roles;

DROP TABLE public.minio_st;
