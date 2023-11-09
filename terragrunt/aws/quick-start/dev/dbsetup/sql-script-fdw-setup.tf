locals {
  fdw_server = "fdwserver"
}

resource "local_file" "sql_script_fdw" {
  content = <<EOF
SHOW hba_file;
SHOW data_directory;
CREATE EXTENSION IF NOT EXISTS postgres_fdw;
SELECT * FROM pg_extension;
CREATE SERVER ${local.fdw_server} FOREIGN DATA WRAPPER postgres_fdw OPTIONS (host '127.0.0.1', port '5432', dbname 'gamedb');
SELECT * FROM pg_foreign_server;
CREATE USER MAPPING FOR ${var.fdwdb_username} SERVER ${local.fdw_server} OPTIONS (user '${var.fdwdb_username}', password '${var.db_password}');
SELECT * FROM pg_user_mapping;
SELECT * FROM pg_user_mappings;
GRANT USAGE ON FOREIGN SERVER ${local.fdw_server} TO ${var.fdwdb_username};
IMPORT FOREIGN SCHEMA public FROM SERVER ${local.fdw_server} INTO public;
EOF

  filename = "${path.module}/fdw-setup.sql"
}