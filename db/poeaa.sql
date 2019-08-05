--
-- PostgreSQL database dump
--

-- Dumped from database version 11.3
-- Dumped by pg_dump version 11.4 (Ubuntu 11.4-0ubuntu0.19.04.1)

-- Started on 2019-08-05 11:13:00 +07

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

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 16385)
-- Name: people; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.people (
    id integer NOT NULL,
    lastname character varying,
    firstname character varying,
    number_of_dependents integer
);


--
-- TOC entry 2884 (class 0 OID 16385)
-- Dependencies: 196
-- Data for Name: people; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.people (id, lastname, firstname, number_of_dependents) FROM stdin;
1	Bob	John	1
2	Jack	Ace	2
\.


--
-- TOC entry 2762 (class 2606 OID 16392)
-- Name: people people_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.people
    ADD CONSTRAINT people_pkey PRIMARY KEY (id);


-- Completed on 2019-08-05 11:13:01 +07

--
-- PostgreSQL database dump complete
--

