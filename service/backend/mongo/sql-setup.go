package mongo

import "data-handler/service/errors"

func resourceSetupTables(runner QueryRunner) errors.ServiceError {
	_, err := runner.Exec(`
		create table if not exists public.resource (
		  id uuid NOT NULL PRIMARY KEY,
		  name character varying(64) not null,
		  workspace character varying(64) not null,
		  type smallint not null,
		  source_data_source character varying(64) not null,
		  source_mapping character varying(64) not null,
		  read_only_records boolean not null,
		  unique_record boolean not null,
		  keep_history boolean not null,
		  auto_created boolean not null,
		  disable_migration boolean not null,
		  disable_audit boolean not null,
		  do_primary_key_lookup boolean not null,
		  created_on timestamp without time zone not null,
		  updated_on timestamp without time zone,
		  created_by character varying(64) not null,
		  updated_by character varying(64),
		  version integer not null,
		  CONSTRAINT name_workspace_uniq unique (workspace, name)
		);
		
		create table if not exists public.resource_property (
		  workspace     character varying(64) not null,
		  resource_name character varying(64) not null,
		  property_name character varying(64) not null,
		  type smallint,
		  source_type smallint,
		  source_mapping character varying(64),
		  source_def character varying(64),
		  source_primary bool,
		  source_auto_generation smallint,
		  required boolean,
		  "unique" boolean,
		  length integer,
		  primary key (workspace, resource_name, property_name),
		  foreign key (workspace, resource_name) references public.resource (workspace, name) match simple on update cascade on delete cascade
		);
		
		create table if not exists public.resource_reference (
		  workspace     character varying(64) not null,
		  resource_name character varying(64) not null,
		  property_name character varying(64) not null unique,
		  referenced_resource character varying(64) not null,
		  "cascade" bool not null,
		  primary key (workspace, resource_name, property_name),
		  foreign key (workspace, resource_name) references public.resource (workspace, name) match simple on update cascade on delete cascade
		);
`)

	return handleDbError(err)
}
