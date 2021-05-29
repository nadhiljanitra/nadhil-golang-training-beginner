BEGIN;
CREATE TABLE IF NOT EXISTS payment_codes
(
    id              VARCHAR         NOT NULL PRIMARY KEY,
    payment_code    VARCHAR         NOT NULL,
    name            VARCHAR         NOT NULL,
    status          VARCHAR         NOT NULL,
    expiration_date TIMESTAMPTZ     NOT NULL,
    created_at      TIMESTAMPTZ     NOT NULL,
    updated_at      TIMESTAMPTZ     NOT NULL
);
COMMIT;