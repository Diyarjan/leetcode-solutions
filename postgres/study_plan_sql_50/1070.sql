-- 1070. Product Sales Analysis III

-- Table: Sales
-- +-------------+-------+
-- | Column Name | Type  |
-- +-------------+-------+
-- | sale_id     | int   |
-- | product_id  | int   |
-- | year        | int   |
-- | quantity    | int   |
-- | price       | int   |
-- +-------------+-------+
-- (sale_id, year) is the primary key (combination of columns with unique values) of this table.
-- product_id is a foreign key (reference column) to Product table.
-- Each row records a sale of a product in a given year.
-- A product may have multiple sales entries in the same year.
-- Note that the per-unit price.

-- Write a solution to find all sales that occurred in the first year each product was sold.
-- For each product_id, identify the earliest year it appears in the Sales table.
-- Return all sales entries for that product in that year.
-- Return a table with the following columns: product_id, first_year, quantity, and price.
-- Return the result in any order.

with result as (select product_id, min(year) as min_year
                from Sales
                group by product_id)
select s.product_id, s.year as first_year, s.quantity, s.price
from Sales s
         join result r on r.min_year = s.year and r.product_id = s.product_id;