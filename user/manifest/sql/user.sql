DROP TABLE IF EXISTS "public"."user_main";
CREATE TABLE "public"."user_main" (
  "id" int8 NOT NULL,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL,
  "phone" char(11) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(0),
  "updated_at" timestamptz(0),
  "deleted_at" timestamptz(0)
)
;
