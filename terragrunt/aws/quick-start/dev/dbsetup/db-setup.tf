
resource "null_resource" "db_setup" {
  depends_on = [ local_file.sql_script_db_setup ]
  triggers = {
    sql_script_content = local_file.sql_script_db_setup.content
  }

  provisioner "local-exec" {
    command = "echo run sql scipt ${local_file.sql_script_db_setup.filename}, exists=${fileexists(local_file.sql_script_db_setup.filename)}"
  }
}

resource "null_resource" "fdw_setup" {
  depends_on = [ local_file.sql_script_fdw ]
  triggers = {
    sql_script_content = local_file.sql_script_fdw.content
  }

  provisioner "local-exec" {
    command = "echo run sql scipt ${local_file.sql_script_fdw.filename}, exists=${fileexists(local_file.sql_script_fdw.filename)}"
  }
}