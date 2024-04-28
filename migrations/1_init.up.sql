
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS isseus (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "isseu_target_uri" varchar NOT NULL,
    "isseu_image64" varchar NOT NULL,
    "isseu_description" varchar NOT NULL,
    "comment" varchar NULL,
    "client_id" varchar NULL,
    "client_name" varchar NULL,
    "isseu_status" varchar DEFAULT('NEW'),
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS statuses (
    "value" varchar NOT NULL,
    "label" varchar NOT NULL
);
