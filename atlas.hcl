data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./migrations/sql/loader",
  ]
}

env "gorm" {
  src = data.external_schema.gorm.url
  dev = "docker://postgres/14"
  migration {
    dir = "file://migrations/sql/postgres"
  }
  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}