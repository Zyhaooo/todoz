package data

const todos = `
CREATE TABLE IF NOT EXISTS "todos" (
	"id" INTEGER NOT NULL UNIQUE,
	-- 标题
	"title" VARCHAR NOT NULL,
	-- 详情
	"description" VARCHAR,
	-- 重要等级
	"level" INTEGER NOT NULL,
	-- 创建时间
	"create_time" TIME NOT NULL,
	-- 过期时间
	"expiration_time" TIME NOT NULL,
	-- 是否完成
	"finish" BOOLEAN,
	PRIMARY KEY("id")
);
`
