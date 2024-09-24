ALTER TABLE product_stock_item
    DROP CONSTRAINT IF EXISTS fk_product_code,
    DROP CONSTRAINT IF EXISTS product_stock_item_product_code_fkey;

ALTER TABLE product_stock_item
    ADD CONSTRAINT fk_product_code
        FOREIGN KEY (product_code)
            REFERENCES products(c_code)
            ON DELETE CASCADE;

