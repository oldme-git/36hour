-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS "public"."user";
CREATE TABLE "public"."user" (
  "id" int8 NOT NULL,
  "username" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "password" char(32) COLLATE "pg_catalog"."default" NOT NULL,
  "phone" char(11) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(0),
  "updated_at" timestamptz(0)
)
;

-- ----------------------------
-- Table structure for user_lib
-- ----------------------------
DROP TABLE IF EXISTS "public"."user_lib";
CREATE TABLE "public"."user_lib" (
  "user_id" int8 NOT NULL,
  "lib_id" int4 NOT NULL
)
;

-- ----------------------------
-- Indexes structure for table user
-- ----------------------------
CREATE UNIQUE INDEX "phone" ON "public"."user" USING btree (
  "phone" COLLATE "pg_catalog"."default" "pg_catalog"."bpchar_ops" ASC NULLS LAST
);
CREATE UNIQUE INDEX "username" ON "public"."user" USING btree (
  "username" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);
