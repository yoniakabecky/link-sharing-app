-- migrate:up
UPDATE platforms SET icon = 'social_x' WHERE name = 'X';
UPDATE platforms SET icon = 'link' WHERE name = 'Custom';

-- migrate:down
UPDATE platforms SET icon = 'x' WHERE name = 'X';
UPDATE platforms SET icon = 'custom' WHERE name = 'Custom';