--
-- PostgreSQL database dump
--
-- Dibuat menggunakan phpPgMyAdmin

SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = off;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET escape_string_warning = off;

SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: Canteen; Type: TABLE; Schema: public; Owner: sanwidii; Tablespace:
--

CREATE TABLE "Canteen" (
    id integer NOT NULL,
    name text,
    location text
);


ALTER TABLE public."Canteen" OWNER TO sanwidii;

--
-- Name: Drink; Type: TABLE; Schema: public; Owner: sanwidii; Tablespace:
--

CREATE TABLE "Drink" (
    id integer NOT NULL,
    name text,
    price integer
);


ALTER TABLE public."Drink" OWNER TO sanwidii;

--
-- Name: Food; Type: TABLE; Schema: public; Owner: sanwidii; Tablespace:
--

CREATE TABLE "Food" (
    id integer NOT NULL,
    name text,
    price integer
);


ALTER TABLE public."Food" OWNER TO sanwidii;

--
-- Name: Canteen_pkey; Type: CONSTRAINT; Schema: public; Owner: sanwidii; Tablespace:
--

ALTER TABLE ONLY "Canteen"
    ADD CONSTRAINT "Canteen_pkey" PRIMARY KEY (id);


--
-- Name: Food_pkey; Type: CONSTRAINT; Schema: public; Owner: sanwidii; Tablespace:
--

ALTER TABLE ONLY "Food"
    ADD CONSTRAINT "Food_pkey" PRIMARY KEY (id);


--
-- Name: public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM postgres;
GRANT ALL ON SCHEMA public TO postgres;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--
