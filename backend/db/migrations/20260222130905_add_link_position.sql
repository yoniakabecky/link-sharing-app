-- migrate:up
ALTER TABLE links ADD COLUMN position INT NOT NULL DEFAULT 0 AFTER url;

-- migrate:down
ALTER TABLE links DROP COLUMN position;
