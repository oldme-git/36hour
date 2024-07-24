-- ----------------------------
-- Sequence structure for layout_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."layout_id_seq";
CREATE SEQUENCE "public"."layout_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for policy_common_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."policy_common_id_seq";
CREATE SEQUENCE "public"."policy_common_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for policy_layout_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."policy_layout_id_seq";
CREATE SEQUENCE "public"."policy_layout_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for policy_prepare_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."policy_prepare_id_seq";
CREATE SEQUENCE "public"."policy_prepare_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for seat_log_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."seat_log_id_seq";
CREATE SEQUENCE "public"."seat_log_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 2147483647
START 1
CACHE 1;

-- ----------------------------
-- Table structure for layout
-- ----------------------------
DROP TABLE IF EXISTS "public"."layout";
CREATE TABLE "public"."layout" (
  "id" int4 NOT NULL DEFAULT nextval('layout_id_seq'::regclass),
  "location_id" int4 NOT NULL,
  "policy_c_id" int4,
  "policy_l_id" int4,
  "layout_name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "map" text COLLATE "pg_catalog"."default",
  "status" int2,
  "sort" int4,
  "seats" int4,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL
)
;
COMMENT ON COLUMN "public"."layout"."policy_c_id" IS '公共策略id，优先使用';
COMMENT ON COLUMN "public"."layout"."policy_l_id" IS '属于自己的策略id，最后使用';
COMMENT ON COLUMN "public"."layout"."map" IS '布局信息';
COMMENT ON COLUMN "public"."layout"."status" IS '是否正常启用，1启用 2不启用';
COMMENT ON COLUMN "public"."layout"."sort" IS '排序，越小越靠前';
COMMENT ON COLUMN "public"."layout"."seats" IS '座位总数';

-- ----------------------------
-- Table structure for policy_common
-- ----------------------------
DROP TABLE IF EXISTS "public"."policy_common";
CREATE TABLE "public"."policy_common" (
  "id" int4 NOT NULL DEFAULT nextval('policy_common_id_seq'::regclass),
  "name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "info" text COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."policy_common"."info" IS '策略信息';
COMMENT ON TABLE "public"."policy_common" IS '公共策略';

-- ----------------------------
-- Table structure for policy_layout
-- ----------------------------
DROP TABLE IF EXISTS "public"."policy_layout";
CREATE TABLE "public"."policy_layout" (
  "id" int4 NOT NULL DEFAULT nextval('policy_layout_id_seq'::regclass),
  "info" text COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."policy_layout"."info" IS '策略信息';
COMMENT ON TABLE "public"."policy_layout" IS '策略预设';

-- ----------------------------
-- Table structure for policy_prepare
-- ----------------------------
DROP TABLE IF EXISTS "public"."policy_prepare";
CREATE TABLE "public"."policy_prepare" (
  "id" int4 NOT NULL DEFAULT nextval('policy_prepare_id_seq'::regclass),
  "name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
  "info" text COLLATE "pg_catalog"."default" NOT NULL
)
;
COMMENT ON COLUMN "public"."policy_prepare"."info" IS '策略信息';
COMMENT ON TABLE "public"."policy_prepare" IS '策略预设';

-- ----------------------------
-- Table structure for seat_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."seat_log";
CREATE TABLE "public"."seat_log" (
  "id" int4 NOT NULL DEFAULT nextval('seat_log_id_seq'::regclass),
  "location_id" int4 NOT NULL,
  "layout_id" int4 NOT NULL,
  "no" int4 NOT NULL,
  "type" int2 NOT NULL,
  "uid" int4 NOT NULL
)
;
COMMENT ON COLUMN "public"."seat_log"."type" IS '1预约 2签到 3退坐';

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."layout_id_seq"
OWNED BY "public"."layout"."id";
SELECT setval('"public"."layout_id_seq"', 16, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."policy_common_id_seq"
OWNED BY "public"."policy_common"."id";
SELECT setval('"public"."policy_common_id_seq"', 5, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."policy_layout_id_seq"
OWNED BY "public"."policy_layout"."id";
SELECT setval('"public"."policy_layout_id_seq"', 3, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."policy_prepare_id_seq"
OWNED BY "public"."policy_prepare"."id";
SELECT setval('"public"."policy_prepare_id_seq"', 3, true);

-- ----------------------------
-- Primary Key structure for table layout
-- ----------------------------
ALTER TABLE "public"."layout" ADD CONSTRAINT "layout_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table policy_common
-- ----------------------------
CREATE UNIQUE INDEX "policy_common_name_ununique" ON "public"."policy_common" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table policy_common
-- ----------------------------
ALTER TABLE "public"."policy_common" ADD CONSTRAINT "policy_common_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table policy_layout
-- ----------------------------
ALTER TABLE "public"."policy_layout" ADD CONSTRAINT "policy_layout_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table policy_prepare
-- ----------------------------
CREATE UNIQUE INDEX "policy_prepare_name_unique" ON "public"."policy_prepare" USING btree (
  "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
);

-- ----------------------------
-- Primary Key structure for table policy_prepare
-- ----------------------------
ALTER TABLE "public"."policy_prepare" ADD CONSTRAINT "policy_prepare_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table seat_log
-- ----------------------------
ALTER TABLE "public"."seat_log" ADD CONSTRAINT "seat_log_pkey" PRIMARY KEY ("id");
