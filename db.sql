--
-- PostgreSQL database dump
--

-- Dumped from database version 11.1
-- Dumped by pg_dump version 12.1

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

SET default_tablespace = '';

--
-- Name: specs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.specs (
    status integer NOT NULL,
    filename character varying(50) NOT NULL,
    line_number integer NOT NULL,
    commit_id integer NOT NULL,
    created_at timestamp without time zone NOT NULL
);


ALTER TABLE public.specs OWNER TO postgres;

--
-- Name: commit_id_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX commit_id_idx ON public.specs USING btree (commit_id);


--
-- Name: specs_created_at_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX specs_created_at_idx ON public.specs USING btree (created_at);


--
-- PostgreSQL database dump complete
--

