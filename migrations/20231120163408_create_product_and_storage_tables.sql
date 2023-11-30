-- +goose Up
-- +goose StatementBegin
CREATE TABLE storages
(
    id            uuid         not null unique,
    name          varchar(255) not null,
    accessibility varchar(255) not null
);

CREATE INDEX idx_storages_id ON storages (id);

CREATE TABLE products
(
    id         uuid         not null unique,
    size       int          not null,
    storage_id int          not null,
    count      int          not null,
    name       varchar(255) not null,
    status     varchar(255) not null
);

CREATE INDEX idx_products_storage_id_status ON products (storage_id, status);
CREATE INDEX idx_products_id ON products (id);

CREATE TABLE product_storages
(
    id         uuid                                            not null unique,
    storage_id uuid references storages (id) on delete cascade not null,
    product_id uuid references products (id) on delete cascade not null
);

CREATE INDEX idx_product_storages_product_id ON product_storages (product_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_product_storages_product_id;
DROP TABLE product_storages;
DROP INDEX idx_products_storage_id_status;
DROP INDEX idx_products_id;
DROP TABLE products;
DROP INDEX idx_storages_id;
DROP TABLE storages;
-- +goose StatementEnd
