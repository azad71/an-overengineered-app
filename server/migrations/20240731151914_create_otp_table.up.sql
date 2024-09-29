
BEGIN;


CREATE TABLE IF NOT EXISTS otp_codes (
  id SERIAL PRIMARY KEY,
  otp VARCHAR(20) NOT NULL,
  email VARCHAR(300)  NOT NULL,
  otp_type VARCHAR(20) NOT NULL,
  retry_count INTEGER DEFAULT 0,
  expires_at TIMESTAMPTZ DEFAULT  NOT NULL,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMPTZ NULL
);

COMMIT;
