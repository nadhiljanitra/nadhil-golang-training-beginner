BEGIN;
CREATE TABLE IF NOT EXISTS payments
(
    id              VARCHAR         NOT NULL PRIMARY KEY,
    transaction_id  VARCHAR         NOT NULL,
    payment_code    VARCHAR         NOT NULL,
    name            VARCHAR         NOT NULL,
    amount          NUMERIC         NOT NULL
);
COMMIT;