
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS issues (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
    "issue_target_uri" varchar NOT NULL,
    "issue_image64" varchar NOT NULL,
    "issue_description" varchar NOT NULL,
    "comment" varchar NULL,
    "client_id" varchar NULL,
    "client_name" varchar NULL,
    "issue_status" varchar DEFAULT('NEW'),
    "created_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    "updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS statuses (
    "value" varchar NOT NULL,
    "label" varchar NOT NULL
);
