resource "local_file" "sql_script" {
  content = <<EOF
CREATE TABLE IF NOT EXISTS "user_data" (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(80) NOT NULL,
    updated TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO "user_data" (name)
VALUES ('Alan'),
    ('Cyan'),
    ('Dora'),
    ('Feynman'),
    ('Hanna'),
    ('Jacky'),
    ('Jason'),
    ('Kevin'),
    ('Nicky'),
    ('Valen'),
    ('Zoey');

CREATE TABLE IF NOT EXISTS "class" (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(16) NOT NULL,
    created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO "class" (title)
VALUES ('Sorcerer'),
    ('Barbarian'),
    ('Rogue'),
    ('Necromancer'),
    ('Druid');
EOF

  filename = "${path.module}/init.sql"
}