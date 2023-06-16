CREATE TABLE clients_data_raw
(
    id                   UUID        NOT NULL PRIMARY KEY,
    document             VARCHAR(19) NOT NULL,
    private              VARCHAR(12) NOT NULL,
    incomplete           VARCHAR(12) NOT NULL,
    last_purchase_date   VARCHAR(22) NOT NULL,
    ticket_average       VARCHAR(22) NOT NULL,
    ticket_last_purchase VARCHAR(24) NOT NULL,
    store_most_frequent  VARCHAR(20) NOT NULL,
    store_last_purchase  VARCHAR(24) NOT NULL
);

CREATE TABLE clients
(
    id                   UUID        NOT NULL PRIMARY KEY,
    document             VARCHAR(14) NOT NULL,
    document_type        VARCHAR(4)  NOT NULL,
    private              BOOLEAN     NOT NULL,
    incomplete           BOOLEAN     NOT NULL,
    last_purchase_date   DATE,
    ticket_average       BIGINT,
    ticket_last_purchase BIGINT,
    store_most_frequent  VARCHAR(14),
    store_last_purchase  VARCHAR(14),
    status               VARCHAR(20),
    created_at           TIMESTAMP,
    updated_at           TIMESTAMP
);
