CREATE TYPE public.task_status AS ENUM ('NEW', 'IN_PROGRESS', 'DONE');

CREATE TABLE IF NOT EXISTS public.tasks
(
    id              bigserial PRIMARY KEY,
    user_id         bigint NOT NULL REFERENCES public.users(id),
    title           varchar(100) NOT NULL,
    description     text,
    status          public.task_status NOT NULL DEFAULT 'NEW',
    deadline        timestamptz,
    created_date    timestamptz NOT NULL DEFAULT now(),
    updated_date    timestamptz NOT NULL DEFAULT now(),
    deleted_date    timestamptz
);