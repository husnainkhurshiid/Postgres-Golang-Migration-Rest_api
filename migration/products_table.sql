-- +migrate Up

CREATE TABLE IF NOT EXISTS "Products" (
  id uuid PRIMARY KEY,
  title  TEXT NOT NULL,
  price float NOT NULL,
  info TEXT NOT NULL,
  rating float NOT NULL
);

-- +migrate Down

DROP TABLE "Products";