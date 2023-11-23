-- +goose Up
-- +goose StatementBegin
CREATE TABLE storages
(
    id            serial       not null unique,
    name          varchar(255) not null,
    accessibility varchar(255) not null
);

CREATE TABLE products
(
    id         serial       not null unique,
    size       int          not null,
    storage_id int          not null,
    count      int          not null,
    name       varchar(255) not null,
    status     varchar(255) not null
);

CREATE TABLE product_storages
(
    id         serial                                            not null unique,
    storage_id serial references storages (id) on delete cascade not null,
    product_id serial references products (id) on delete cascade not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product_storages;
DROP TABLE products;
DROP TABLE storages;
-- +goose StatementEnd
