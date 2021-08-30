BEGIN;
CREATE TABLE IF NOT EXISTS inquiries
(
    id              VARCHAR         NOT NULL PRIMARY KEY,
    transaction_id  VARCHAR         NOT NULL,
    payment_code    VARCHAR         NOT NULL,

    UNIQUE (transaction_id)
);
COMMIT;