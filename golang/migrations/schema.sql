--
-- PostgreSQL database dump
--

-- Dumped from database version 10.8 (Ubuntu 10.8-0ubuntu0.18.04.1)
-- Dumped by pg_dump version 10.8 (Ubuntu 10.8-0ubuntu0.18.04.1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: comments; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.comments (
    id integer NOT NULL,
    task_id integer NOT NULL,
    user_id integer NOT NULL,
    content character varying(255) DEFAULT ''::character varying NOT NULL,
    rating integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.comments OWNER TO tibi;

--
-- Name: comments_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.comments_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.comments_id_seq OWNER TO tibi;

--
-- Name: comments_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.comments_id_seq OWNED BY public.comments.id;


--
-- Name: languages; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.languages (
    id integer NOT NULL,
    name character varying(255),
    description character varying(255) NOT NULL,
    documentation character varying(255) NOT NULL,
    color character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.languages OWNER TO tibi;

--
-- Name: languages_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.languages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.languages_id_seq OWNER TO tibi;

--
-- Name: languages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.languages_id_seq OWNED BY public.languages.id;


--
-- Name: licenses; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.licenses (
    id integer NOT NULL,
    name character varying(255),
    description character varying(255),
    nickname character varying(255) NOT NULL,
    key character varying(255) NOT NULL,
    body character varying(2550) NOT NULL,
    url character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.licenses OWNER TO tibi;

--
-- Name: licenses_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.licenses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.licenses_id_seq OWNER TO tibi;

--
-- Name: licenses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.licenses_id_seq OWNED BY public.licenses.id;


--
-- Name: platforms; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.platforms (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    home character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.platforms OWNER TO tibi;

--
-- Name: platforms_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.platforms_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.platforms_id_seq OWNER TO tibi;

--
-- Name: platforms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.platforms_id_seq OWNED BY public.platforms.id;


--
-- Name: projects; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.projects (
    id uuid NOT NULL,
    platform_id integer NOT NULL,
    id_on_platform bigint NOT NULL,
    license_id integer,
    name character varying(255) NOT NULL,
    description character varying(255),
    closed boolean DEFAULT false NOT NULL,
    url character varying(255) DEFAULT ''::character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.projects OWNER TO tibi;

--
-- Name: projects_languages; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.projects_languages (
    id integer NOT NULL,
    project_id uuid NOT NULL,
    language_id integer NOT NULL,
    usage numeric NOT NULL,
    is_primary boolean NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.projects_languages OWNER TO tibi;

--
-- Name: projects_languages_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.projects_languages_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.projects_languages_id_seq OWNER TO tibi;

--
-- Name: projects_languages_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.projects_languages_id_seq OWNED BY public.projects_languages.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO tibi;

--
-- Name: tasks; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.tasks (
    id integer NOT NULL,
    task_id integer,
    project_id uuid NOT NULL,
    name character varying(255),
    description character varying(255),
    progress numeric DEFAULT '0'::numeric NOT NULL,
    closed boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tasks OWNER TO tibi;

--
-- Name: tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.tasks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tasks_id_seq OWNER TO tibi;

--
-- Name: tasks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.tasks_id_seq OWNED BY public.tasks.id;


--
-- Name: user_settings; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.user_settings (
    id integer NOT NULL,
    user_id integer NOT NULL,
    avatar character varying(255),
    theme character varying(255) DEFAULT 'default'::character varying NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.user_settings OWNER TO tibi;

--
-- Name: user_settings_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.user_settings_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_settings_id_seq OWNER TO tibi;

--
-- Name: user_settings_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.user_settings_id_seq OWNED BY public.user_settings.id;


--
-- Name: user_verifies; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.user_verifies (
    id integer NOT NULL,
    user_id integer NOT NULL,
    token character varying(255),
    type character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.user_verifies OWNER TO tibi;

--
-- Name: user_verifies_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.user_verifies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.user_verifies_id_seq OWNER TO tibi;

--
-- Name: user_verifies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.user_verifies_id_seq OWNED BY public.user_verifies.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.users (
    id integer NOT NULL,
    name character varying(255),
    password character varying(255),
    email character varying(255),
    remember_token character varying(255),
    slug character varying(255),
    type integer DEFAULT 0 NOT NULL,
    joined_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO tibi;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO tibi;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: users_platforms; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.users_platforms (
    id integer NOT NULL,
    id_on_platform bigint,
    user_id integer,
    platform_id integer,
    username character varying(255),
    token character varying(255),
    token_type character varying(255),
    limit_requests integer,
    url character varying(255),
    last_updated_at timestamp without time zone,
    reset_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users_platforms OWNER TO tibi;

--
-- Name: users_platforms_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.users_platforms_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_platforms_id_seq OWNER TO tibi;

--
-- Name: users_platforms_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.users_platforms_id_seq OWNED BY public.users_platforms.id;


--
-- Name: users_projects; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.users_projects (
    id integer NOT NULL,
    user_id integer NOT NULL,
    project_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users_projects OWNER TO tibi;

--
-- Name: users_projects_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.users_projects_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_projects_id_seq OWNER TO tibi;

--
-- Name: users_projects_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.users_projects_id_seq OWNED BY public.users_projects.id;


--
-- Name: users_tasks; Type: TABLE; Schema: public; Owner: tibi
--

CREATE TABLE public.users_tasks (
    id integer NOT NULL,
    user_id integer NOT NULL,
    task_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users_tasks OWNER TO tibi;

--
-- Name: users_tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: tibi
--

CREATE SEQUENCE public.users_tasks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_tasks_id_seq OWNER TO tibi;

--
-- Name: users_tasks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tibi
--

ALTER SEQUENCE public.users_tasks_id_seq OWNED BY public.users_tasks.id;


--
-- Name: comments id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.comments ALTER COLUMN id SET DEFAULT nextval('public.comments_id_seq'::regclass);


--
-- Name: languages id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.languages ALTER COLUMN id SET DEFAULT nextval('public.languages_id_seq'::regclass);


--
-- Name: licenses id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.licenses ALTER COLUMN id SET DEFAULT nextval('public.licenses_id_seq'::regclass);


--
-- Name: platforms id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.platforms ALTER COLUMN id SET DEFAULT nextval('public.platforms_id_seq'::regclass);


--
-- Name: projects_languages id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.projects_languages ALTER COLUMN id SET DEFAULT nextval('public.projects_languages_id_seq'::regclass);


--
-- Name: tasks id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.tasks ALTER COLUMN id SET DEFAULT nextval('public.tasks_id_seq'::regclass);


--
-- Name: user_settings id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.user_settings ALTER COLUMN id SET DEFAULT nextval('public.user_settings_id_seq'::regclass);


--
-- Name: user_verifies id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.user_verifies ALTER COLUMN id SET DEFAULT nextval('public.user_verifies_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Name: users_platforms id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_platforms ALTER COLUMN id SET DEFAULT nextval('public.users_platforms_id_seq'::regclass);


--
-- Name: users_projects id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_projects ALTER COLUMN id SET DEFAULT nextval('public.users_projects_id_seq'::regclass);


--
-- Name: users_tasks id; Type: DEFAULT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_tasks ALTER COLUMN id SET DEFAULT nextval('public.users_tasks_id_seq'::regclass);


--
-- Name: comments comments_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_pkey PRIMARY KEY (id);


--
-- Name: languages languages_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.languages
    ADD CONSTRAINT languages_pkey PRIMARY KEY (id);


--
-- Name: licenses licenses_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.licenses
    ADD CONSTRAINT licenses_pkey PRIMARY KEY (id);


--
-- Name: platforms platforms_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.platforms
    ADD CONSTRAINT platforms_pkey PRIMARY KEY (id);


--
-- Name: projects_languages projects_languages_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.projects_languages
    ADD CONSTRAINT projects_languages_pkey PRIMARY KEY (id);


--
-- Name: projects projects_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_pkey PRIMARY KEY (id);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- Name: user_settings user_settings_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.user_settings
    ADD CONSTRAINT user_settings_pkey PRIMARY KEY (id);


--
-- Name: user_verifies user_verifies_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.user_verifies
    ADD CONSTRAINT user_verifies_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users_platforms users_platforms_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_platforms
    ADD CONSTRAINT users_platforms_pkey PRIMARY KEY (id);


--
-- Name: users_projects users_projects_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_projects
    ADD CONSTRAINT users_projects_pkey PRIMARY KEY (id);


--
-- Name: users_tasks users_tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_tasks
    ADD CONSTRAINT users_tasks_pkey PRIMARY KEY (id);


--
-- Name: languages_name_idx; Type: INDEX; Schema: public; Owner: tibi
--

CREATE INDEX languages_name_idx ON public.languages USING btree (name);


--
-- Name: platforms_home_idx; Type: INDEX; Schema: public; Owner: tibi
--

CREATE UNIQUE INDEX platforms_home_idx ON public.platforms USING btree (home);


--
-- Name: platforms_name_idx; Type: INDEX; Schema: public; Owner: tibi
--

CREATE UNIQUE INDEX platforms_name_idx ON public.platforms USING btree (name);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: tibi
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: user_verifies_token_idx; Type: INDEX; Schema: public; Owner: tibi
--

CREATE UNIQUE INDEX user_verifies_token_idx ON public.user_verifies USING btree (token);


--
-- Name: users_email_idx; Type: INDEX; Schema: public; Owner: tibi
--

CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);


--
-- Name: users_platforms_id_on_platform_idx; Type: INDEX; Schema: public; Owner: tibi
--

CREATE UNIQUE INDEX users_platforms_id_on_platform_idx ON public.users_platforms USING btree (id_on_platform);


--
-- Name: users_platforms_url_idx; Type: INDEX; Schema: public; Owner: tibi
--

CREATE UNIQUE INDEX users_platforms_url_idx ON public.users_platforms USING btree (url);


--
-- Name: users_remember_token_idx; Type: INDEX; Schema: public; Owner: tibi
--

CREATE UNIQUE INDEX users_remember_token_idx ON public.users USING btree (remember_token);


--
-- Name: users_slug_idx; Type: INDEX; Schema: public; Owner: tibi
--

CREATE UNIQUE INDEX users_slug_idx ON public.users USING btree (slug);


--
-- Name: comments comments_task_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_task_id_fkey FOREIGN KEY (task_id) REFERENCES public.tasks(id) ON DELETE CASCADE;


--
-- Name: comments comments_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.comments
    ADD CONSTRAINT comments_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: projects_languages projects_languages_language_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.projects_languages
    ADD CONSTRAINT projects_languages_language_id_fkey FOREIGN KEY (language_id) REFERENCES public.languages(id) ON DELETE CASCADE;


--
-- Name: projects_languages projects_languages_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.projects_languages
    ADD CONSTRAINT projects_languages_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id) ON DELETE CASCADE;


--
-- Name: projects projects_license_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_license_id_fkey FOREIGN KEY (license_id) REFERENCES public.licenses(id);


--
-- Name: projects projects_platform_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.projects
    ADD CONSTRAINT projects_platform_id_fkey FOREIGN KEY (platform_id) REFERENCES public.platforms(id) ON DELETE CASCADE;


--
-- Name: tasks tasks_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id) ON DELETE CASCADE;


--
-- Name: tasks tasks_task_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_task_id_fkey FOREIGN KEY (task_id) REFERENCES public.tasks(id) ON DELETE CASCADE;


--
-- Name: user_settings user_settings_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.user_settings
    ADD CONSTRAINT user_settings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: user_verifies user_verifies_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.user_verifies
    ADD CONSTRAINT user_verifies_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: users_platforms users_platforms_platform_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_platforms
    ADD CONSTRAINT users_platforms_platform_id_fkey FOREIGN KEY (platform_id) REFERENCES public.platforms(id) ON DELETE CASCADE;


--
-- Name: users_platforms users_platforms_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_platforms
    ADD CONSTRAINT users_platforms_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: users_projects users_projects_project_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_projects
    ADD CONSTRAINT users_projects_project_id_fkey FOREIGN KEY (project_id) REFERENCES public.projects(id) ON DELETE CASCADE;


--
-- Name: users_projects users_projects_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_projects
    ADD CONSTRAINT users_projects_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: users_tasks users_tasks_task_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_tasks
    ADD CONSTRAINT users_tasks_task_id_fkey FOREIGN KEY (task_id) REFERENCES public.tasks(id) ON DELETE CASCADE;


--
-- Name: users_tasks users_tasks_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: tibi
--

ALTER TABLE ONLY public.users_tasks
    ADD CONSTRAINT users_tasks_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

