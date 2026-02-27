-- migrate:up
ALTER TABLE profiles ADD COLUMN nickname VARCHAR(255) NOT NULL DEFAULT '' AFTER user_id;

-- migrate:down
ALTER TABLE profiles DROP COLUMN nickname;
