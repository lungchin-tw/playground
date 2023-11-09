
resource "local_file" "sql_script_db_setup" {
  content = <<EOF
DROP TABLE IF EXISTS "user_data";
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

DROP TABLE IF EXISTS "class";
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

CREATE USER ${local.fdw_username} WITH PASSWORD '${var.db_password}';
GRANT USAGE ON SCHEMA PUBLIC TO ${local.fdw_username};
GRANT SELECT ON class TO ${local.fdw_username};
GRANT SELECT ON user_data TO ${local.fdw_username};
EOF

  filename = "${path.module}/db-setup.sql"
}