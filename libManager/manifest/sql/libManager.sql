-- ----------------------------
-- Table structure for floor
-- ----------------------------
DROP TABLE IF EXISTS "public"."floor";
CREATE TABLE "public"."floor" (
  "id" serial4 NOT NULL,
  "lib_id" int4 NOT NULL,
  "floor_name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL
)
;

-- ----------------------------
-- Table structure for hall
-- ----------------------------
DROP TABLE IF EXISTS "public"."hall";
CREATE TABLE "public"."hall" (
  "id" serial4 NOT NULL,
  "lib_id" int4 NOT NULL,
  "floor_id" int4 NOT NULL,
  "hall_name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamptz(6) NOT NULL,
  "updated_at" timestamptz(6) NOT NULL
)
;

-- ----------------------------
-- Table structure for lib
-- ----------------------------
DROP TABLE IF EXISTS "public"."lib";
CREATE TABLE "public"."lib" (
  "id" serial4 NOT NULL,
  "name" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
  "address" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "active" bool NOT NULL
)
;
COMMENT ON COLUMN "public"."lib"."address" IS '地址';
COMMENT ON COLUMN "public"."lib"."active" IS '是否正在使用';

-- ----------------------------
-- Primary Key structure for table floor
-- ----------------------------
ALTER TABLE "public"."floor" ADD CONSTRAINT "floor_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table hall
-- ----------------------------
ALTER TABLE "public"."hall" ADD CONSTRAINT "hall_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table lib
-- ----------------------------
ALTER TABLE "public"."lib" ADD CONSTRAINT "lib_pkey" PRIMARY KEY ("id");
