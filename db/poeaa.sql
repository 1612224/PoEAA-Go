--
-- PostgreSQL database dump
--

-- Dumped from database version 11.3
-- Dumped by pg_dump version 11.4 (Ubuntu 11.4-0ubuntu0.19.04.1)

-- Started on 2019-08-06 16:35:13 +07

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
-- TOC entry 197 (class 1259 OID 16393)
-- Name: contracts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.contracts (
    id integer NOT NULL,
    product integer,
    revenue numeric,
    datesigned date
);


--
-- TOC entry 200 (class 1259 OID 16417)
-- Name: keys; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.keys (
    name character varying NOT NULL,
    nextid integer
);


--
-- TOC entry 199 (class 1259 OID 16409)
-- Name: people; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.people (
    id integer NOT NULL,
    lastname character varying,
    firstname character varying,
    numberofdependents integer
);


--
-- TOC entry 196 (class 1259 OID 16385)
-- Name: products; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.products (
    id integer NOT NULL,
    name character varying,
    type character varying
);


--
-- TOC entry 198 (class 1259 OID 16401)
-- Name: revenuerecognitions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.revenuerecognitions (
    contract integer NOT NULL,
    amount numeric,
    recognizedon date NOT NULL
);


--
-- TOC entry 2913 (class 0 OID 16393)
-- Dependencies: 197
-- Data for Name: contracts; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.contracts (id, product, revenue, datesigned) FROM stdin;
\.


--
-- TOC entry 2916 (class 0 OID 16417)
-- Dependencies: 200
-- Data for Name: keys; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.keys (name, nextid) FROM stdin;
people	1
products	1
\.


--
-- TOC entry 2915 (class 0 OID 16409)
-- Dependencies: 199
-- Data for Name: people; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.people (id, lastname, firstname, numberofdependents) FROM stdin;
\.


--
-- TOC entry 2912 (class 0 OID 16385)
-- Dependencies: 196
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.products (id, name, type) FROM stdin;
\.


--
-- TOC entry 2914 (class 0 OID 16401)
-- Dependencies: 198
-- Data for Name: revenuerecognitions; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.revenuerecognitions (contract, amount, recognizedon) FROM stdin;
\.


--
-- TOC entry 2784 (class 2606 OID 16400)
-- Name: contracts contracts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.contracts
    ADD CONSTRAINT contracts_pkey PRIMARY KEY (id);


--
-- TOC entry 2790 (class 2606 OID 16424)
-- Name: keys keys_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.keys
    ADD CONSTRAINT keys_pkey PRIMARY KEY (name);


--
-- TOC entry 2788 (class 2606 OID 16416)
-- Name: people people_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.people
    ADD CONSTRAINT people_pkey PRIMARY KEY (id);


--
-- TOC entry 2782 (class 2606 OID 16392)
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- TOC entry 2786 (class 2606 OID 16408)
-- Name: revenuerecognitions revenuerecognitions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.revenuerecognitions
    ADD CONSTRAINT revenuerecognitions_pkey PRIMARY KEY (contract, recognizedon);


-- Completed on 2019-08-06 16:35:14 +07

--
-- PostgreSQL database dump complete
--

