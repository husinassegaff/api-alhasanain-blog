-- Type: roles

-- DROP TYPE IF EXISTS public.roles;

CREATE TYPE public.roles AS ENUM
    ('admin', 'user');

ALTER TYPE public.roles
    OWNER TO postgres;

-- Type: status_post

-- DROP TYPE IF EXISTS public.status_post;

CREATE TYPE public.status_post AS ENUM
    ('draft', 'publish', 'delete');

ALTER TYPE public.status_post
    OWNER TO postgres;



-- Table: public.user

-- DROP TABLE IF EXISTS public."user";

CREATE TABLE IF NOT EXISTS public."user"
(
    id character(36) COLLATE pg_catalog."default" NOT NULL,
    name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    email character varying(100) COLLATE pg_catalog."default" NOT NULL,
    created_at date NOT NULL,
    updated_at date,
    role roles NOT NULL,
    password character varying(100) COLLATE pg_catalog."default" NOT NULL,
    token character varying(256) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT id_user PRIMARY KEY (id)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."user"
    OWNER to postgres;

-- Table: public.post

-- DROP TABLE IF EXISTS public.post;

CREATE TABLE IF NOT EXISTS public.post
(
    id character(36) COLLATE pg_catalog."default" NOT NULL,
    id_user character(36) COLLATE pg_catalog."default" NOT NULL,
    title character varying(100) COLLATE pg_catalog."default" NOT NULL,
    meta_title character varying(125) COLLATE pg_catalog."default",
    slug character varying(100) COLLATE pg_catalog."default" NOT NULL,
    content text COLLATE pg_catalog."default" NOT NULL,
    summary text COLLATE pg_catalog."default",
    status status_post NOT NULL,
    created_at date NOT NULL,
    updated_at date,
    published_at date,
    CONSTRAINT id_post PRIMARY KEY (id),
    CONSTRAINT id_user FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.post
    OWNER to postgres;

-- Table: public.category

-- DROP TABLE IF EXISTS public.category;

CREATE TABLE IF NOT EXISTS public.category
(
    id character(36) COLLATE pg_catalog."default" NOT NULL,
    title character varying(100) COLLATE pg_catalog."default" NOT NULL,
    meta_title character varying(125) COLLATE pg_catalog."default" NOT NULL,
    slug character varying(100) COLLATE pg_catalog."default" NOT NULL,
    content character varying(100) COLLATE pg_catalog."default",
    created_at date NOT NULL,
    updated_at date,
    CONSTRAINT id_category PRIMARY KEY (id)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.category
    OWNER to postgres;

-- Table: public.comment

-- DROP TABLE IF EXISTS public.comment;

CREATE TABLE IF NOT EXISTS public.comment
(
    id character(36) COLLATE pg_catalog."default" NOT NULL,
    id_user character(36) COLLATE pg_catalog."default" NOT NULL,
    id_post character(36) COLLATE pg_catalog."default" NOT NULL,
    content text COLLATE pg_catalog."default" NOT NULL,
    created_at date NOT NULL,
    CONSTRAINT id_comment PRIMARY KEY (id),
    CONSTRAINT id_post FOREIGN KEY (id_post)
        REFERENCES public.post (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
        ,
    CONSTRAINT id_user FOREIGN KEY (id_user)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
        
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.comment
    OWNER to postgres;

-- Table: public.meta

-- DROP TABLE IF EXISTS public.meta;

CREATE TABLE IF NOT EXISTS public.meta
(
    id character(36) COLLATE pg_catalog."default" NOT NULL,
    key character varying(100) COLLATE pg_catalog."default" NOT NULL,
    content character varying(256) COLLATE pg_catalog."default" NOT NULL,
    created_at date NOT NULL,
    updated_at date,
    CONSTRAINT id_meta PRIMARY KEY (id)
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.meta
    OWNER to postgres;

-- Table: public.tag

-- DROP TABLE IF EXISTS public.tag;

CREATE TABLE IF NOT EXISTS public.tag
(
    id character(36) COLLATE pg_catalog."default" NOT NULL,
    id_post character(36) COLLATE pg_catalog."default" NOT NULL,
    title character varying(100) COLLATE pg_catalog."default" NOT NULL,
    meta_title character varying(125) COLLATE pg_catalog."default",
    created_at date NOT NULL,
    updated_at date,
    CONSTRAINT id_tag PRIMARY KEY (id),
    CONSTRAINT id_post FOREIGN KEY (id_post)
        REFERENCES public.post (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tag
    OWNER to postgres;

-- Table: public.post_category

-- DROP TABLE IF EXISTS public.post_category;

CREATE TABLE IF NOT EXISTS public.post_category
(
    id_category character(36) COLLATE pg_catalog."default" NOT NULL,
    id_post character(36) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT id_category FOREIGN KEY (id_category)
        REFERENCES public.category (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
        ,
    CONSTRAINT id_post FOREIGN KEY (id_post)
        REFERENCES public.post (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
        
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.post_category
    OWNER to postgres;

-- Table: public.post_meta

-- DROP TABLE IF EXISTS public.post_meta;

CREATE TABLE IF NOT EXISTS public.post_meta
(
    id_meta character(36) COLLATE pg_catalog."default" NOT NULL,
    id_post character(36) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT id_meta FOREIGN KEY (id_meta)
        REFERENCES public.meta (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT,
    CONSTRAINT id_post FOREIGN KEY (id_post)
        REFERENCES public.post (id) MATCH SIMPLE
        ON UPDATE RESTRICT
        ON DELETE RESTRICT
)

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.post_meta
    OWNER to postgres;