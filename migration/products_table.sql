-- +migrate Up

CREATE TABLE IF NOT EXISTS "Products" (
  title  TEXT NOT NULL,
  price INT NOT NULL,
  info TEXT NOT NULL,
  rating Float NOT NULL
);

-- +migrate Down

DROP TABLE "Products";