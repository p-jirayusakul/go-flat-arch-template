CREATE SCHEMA IF NOT EXISTS public;

CREATE TABLE IF NOT EXISTS public.accounts
(
    id text NOT NULL DEFAULT uuid_generate_v4(),
    email text NOT NULL,
    password text NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT now(),
    updated_at timestamp with time zone,
    CONSTRAINT accounts_pkey PRIMARY KEY (id),
    CONSTRAINT unique_email UNIQUE (email)
);