ALTER TABLE product_stock
    DROP CONSTRAINT IF EXISTS product_stock_customer_id_fkey;

ALTER TABLE product_stock
    ALTER COLUMN customer_id TYPE TEXT;