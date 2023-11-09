resource "local_file" "sql_script_fdw" {
  content = <<EOF
SHOW hba_file;
SHOW data_directory;
CREATE EXTENSION IF NOT EXISTS postgres_fdw;
SELECT * FROM pg_extension;
CREATE SERVER ${local.fdw_server} FOREIGN DATA WRAPPER postgres_fdw OPTIONS (host '${aws_db_instance.postgres.address}', port '${aws_db_instance.postgres.port}', dbname '${aws_db_instance.postgres.db_name}');
SELECT * FROM pg_foreign_server;
CREATE USER MAPPING FOR ${local.fdw_username} SERVER ${local.fdw_server} OPTIONS (user '${local.fdw_username}', password '${var.db_password}');
SELECT * FROM pg_user_mapping;
SELECT * FROM pg_user_mappings;
GRANT USAGE ON FOREIGN SERVER ${local.fdw_server} TO ${local.fdw_username};
IMPORT FOREIGN SCHEMA public FROM SERVER ${local.fdw_server} INTO public;
EOF

  filename = "${path.module}/fdw-setup.sql"
}