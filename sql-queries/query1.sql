--- VIDEO_Q1 ---

/* Problem Statement:
- For pairs of brands in the same year (e.g. apple/samsung/2020 and samsung/apple/2020) 
    - if custom1 = custom3 and custom2 = custom4 : then keep only one pair

- For pairs of brands in the same year 
    - if custom1 != custom3 OR custom2 != custom4 : then keep both pairs

- For brands that do not have pairs in the same year : keep those rows as well
*/


DROP TABLE IF EXISTS brands;
CREATE TABLE brands 
(
    brand1      VARCHAR(20),
    brand2      VARCHAR(20),
    year        INT,
    custom1     INT,
    custom2     INT,
    custom3     INT,
    custom4     INT
);
INSERT INTO brands VALUES ('apple', 'samsung', 2020, 1, 2, 1, 2);
INSERT INTO brands VALUES ('samsung', 'apple', 2020, 1, 2, 1, 2);
INSERT INTO brands VALUES ('apple', 'samsung', 2021, 1, 2, 5, 3);
INSERT INTO brands VALUES ('samsung', 'apple', 2021, 5, 3, 1, 2);
INSERT INTO brands VALUES ('google', NULL, 2020, 5, 9, NULL, NULL);
INSERT INTO brands VALUES ('oneplus', 'nothing', 2020, 5, 9, 6, 3);

SELECT * FROM brands;

;with base1 as (
	select 
		*, 
		case 
			when brand1 < brand2 
			then concat(cast(brand1 as varchar(10)) , (cast(brand2 as varchar(10))) , (cast(year as varchar(4))) ) 
			else concat(cast(brand2 as varchar(10)) , (cast(brand1 as varchar(10))) , (cast(year as varchar(4))) )
		end as str
	from brands	
), base2 as (
	select
		*, row_number() over(partition by str) as rn
	from base1	
)
	
select * from base2 where rn = 1 or (custom1 <> custom3 and custom2 <> custom4)



