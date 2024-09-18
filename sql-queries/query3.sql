DROP TABLE IF EXISTS FOOTER;
CREATE TABLE FOOTER 
(
	id 			INT PRIMARY KEY,
	car 		VARCHAR(20), 
	length 		INT, 
	width 		INT, 
	height 		INT
);

INSERT INTO FOOTER VALUES (1, 'Hyundai Tucson', 15, 6, NULL);
INSERT INTO FOOTER VALUES (2, NULL, NULL, NULL, 20);
INSERT INTO FOOTER VALUES (3, NULL, 12, 8, 15);
INSERT INTO FOOTER VALUES (4, 'Toyota Rav4', NULL, 15, NULL);
INSERT INTO FOOTER VALUES (5, 'Kia Sportage', NULL, NULL, 18); 

SELECT * FROM FOOTER;

;with car_cte as (
select car from footer where car is not null order by id desc limit 1
), length_cte as (
select length from footer where length is not null order by id desc limit 1
), width_cte as (
select width from footer where width is not null order by id desc limit 1
), height_cte as (
select height from footer where height is not null order by id desc limit 1
)

select
	*
from car_cte 
join length_cte on 1 = 1
join width_cte on 1 = 1
join height_cte on 1 = 1