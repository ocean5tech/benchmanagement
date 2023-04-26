# benchmanagement

-- Table: benchmanagement.assigement

-- DROP TABLE IF EXISTS benchmanagement.assigement;

CREATE TABLE IF NOT EXISTS benchmanagement.assigement
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    empid text COLLATE pg_catalog."default" NOT NULL,
    taskid integer NOT NULL,
    startdate date,
    enddate date,
    status text COLLATE pg_catalog."default",
    CONSTRAINT assigement_pkey PRIMARY KEY (id),
    CONSTRAINT empid_fkey FOREIGN KEY (empid)
        REFERENCES benchmanagement.employee (serialid) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT taskid_fkey FOREIGN KEY (taskid)
        REFERENCES benchmanagement.tasks (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS benchmanagement.assigement
    OWNER to postgres;


-- Table: benchmanagement.demo

-- DROP TABLE IF EXISTS benchmanagement.demo;

CREATE TABLE IF NOT EXISTS benchmanagement.demo
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name text COLLATE pg_catalog."default" NOT NULL,
    "PM" text COLLATE pg_catalog."default",
    supervisor text COLLATE pg_catalog."default",
    status text COLLATE pg_catalog."default",
    homepage text COLLATE pg_catalog."default",
    CONSTRAINT demo_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS benchmanagement.demo
    OWNER to postgres;

-- Table: benchmanagement.employee

-- DROP TABLE IF EXISTS benchmanagement.employee;

CREATE TABLE IF NOT EXISTS benchmanagement.employee
(
    serialid text COLLATE pg_catalog."default" NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    email text COLLATE pg_catalog."default",
    pemname text COLLATE pg_catalog."default",
    pememail text COLLATE pg_catalog."default",
    benchstartday date,
    benchendday date,
    CONSTRAINT employee_pkey PRIMARY KEY (serialid)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS benchmanagement.employee
    OWNER to postgres;

-- Table: benchmanagement.requirements

-- DROP TABLE IF EXISTS benchmanagement.requirements;

CREATE TABLE IF NOT EXISTS benchmanagement.requirements
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name text COLLATE pg_catalog."default" NOT NULL,
    demoid integer NOT NULL,
    demoname text COLLATE pg_catalog."default" NOT NULL,
    "desc" text COLLATE pg_catalog."default",
    CONSTRAINT requirements_pkey PRIMARY KEY (id),
    CONSTRAINT demoid_fkey FOREIGN KEY (demoid)
        REFERENCES benchmanagement.demo (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS benchmanagement.requirements
    OWNER to postgres;

-- Table: benchmanagement.tasks

-- DROP TABLE IF EXISTS benchmanagement.tasks;

CREATE TABLE IF NOT EXISTS benchmanagement.tasks
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    name text COLLATE pg_catalog."default" NOT NULL,
    requirementid integer NOT NULL,
    CONSTRAINT tasks_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS benchmanagement.tasks
    OWNER to postgres;