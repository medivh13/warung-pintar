CREATE TABLE warung_pintar.promos (
    id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    deleted_at timestamp with time zone,
    sku character varying(255) NOT NULL,
    value_percentage float DEFAULT 0 NOT NULL,
    minimum_qty integer DEFAULT 0 NOT NULL
);

CREATE SEQUENCE warung_pintar.promos_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE warung_pintar.promos_id_seq OWNED BY warung_pintar.promos.id;

ALTER TABLE ONLY warung_pintar.promos ALTER COLUMN id SET DEFAULT nextval('warung_pintar.promos_id_seq'::regclass);

ALTER TABLE ONLY warung_pintar.promos
    ADD CONSTRAINT promos_pkey PRIMARY KEY (id);

ALTER TABLE ONLY warung_pintar.promos
    ADD CONSTRAINT promos_sku_key UNIQUE (sku);

CREATE INDEX promosindex ON warung_pintar.promos USING btree (deleted_at, sku);


INSERT INTO warung_pintar.promos (sku, value_percentage, minimum_qty)
VALUES ('120P90', 0.05, 2);
INSERT INTO warung_pintar.promos (sku, value_percentage, minimum_qty)
VALUES ('43N23P', 0.1, 2);
INSERT INTO warung_pintar.promos (sku, value_percentage, minimum_qty)
VALUES ('A304SD',  0.15, 2);
INSERT INTO warung_pintar.promos (sku, value_percentage, minimum_qty)
VALUES ('234234',  0.1, 2);