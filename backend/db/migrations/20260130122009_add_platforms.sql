-- migrate:up
INSERT INTO platforms (name, icon, color) VALUES
  ('Discord', 'discord', '#5865F2'),
  ('Dribbble', 'dribbble', '#EA4C89'),
  ('Facebook', 'facebook', '#3B5998'),
  ('GitHub', 'github', '#333333'),
  ('Instagram', 'instagram', '#E1306C'),
  ('LinkedIn', 'linkedin', '#0077B5'),
  ('Medium', 'medium', '#000000'),
  ('Reddit', 'reddit', '#FF4500'),
  ('TikTok', 'tiktok', '#010101'),
  ('Website', 'web', '#000000'),
  ('X', 'x', '#000000'),
  ('YouTube', 'youtube', '#FF0000'),
  ('Custom', 'custom', '#000000')

-- migrate:down
DELETE FROM platforms WHERE name IN ('Discord', 'Dribbble', 'Facebook', 'GitHub', 'Instagram', 'LinkedIn', 'Medium', 'Reddit', 'TikTok', 'Website', 'X', 'YouTube', 'Custom');

