DROP TABLE IF EXISTS "public"."user_lib";
CREATE TABLE "public"."user_lib" (
  "id" int4 NOT NULL,
  "user_id" int8 NOT NULL,
  "lib_id" int4 NOT NULL
);

-- ----------------------------
-- Table structure for user_main
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_main";
CREATE TABLE "public"."user_main" (
  "id" int8 NOT NULL,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL,
  "phone" char(11) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(0),
  "updated_at" timestamptz(0)
)
;

-- ----------------------------
-- Primary Key structure for table user_lib
-- ----------------------------
ALTER TABLE "public"."user_lib" ADD CONSTRAINT "user_lib_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table user_main
-- ----------------------------
CREATE UNIQUE INDEX "phone" ON "public"."user_main" USING btree (
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."bpchar_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "username" ON "public"."user_main" USING btree (
  "username" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
