ALTER TABLE "accounts" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "member_role_related_permissions" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "permission_groups" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "permission_related_routers" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "routers" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "source_data" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "sources" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "user_auth_sources" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "user_related_roles" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "user_roles" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "users" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_categories" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_dependencies" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_durations" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_editors" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_entities" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_grand_totals" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_heart_beats" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_languages" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_project_durations" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_project_infos" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_projects" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatime_systems" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
ALTER TABLE "wakatimes" ALTER COLUMN "id" SET GENERATED BY DEFAULT SET START WITH 1 SET INCREMENT BY 1 RESTART;
