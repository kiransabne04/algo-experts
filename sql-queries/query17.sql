-- Given is user login table for , identify dates where a user has logged in for 5 or more consecutive days.
-- Return the user id, start date, end date and no of consecutive days, sorting based on user id.
-- If a user logged in consecutively 5 or more times but not spanning 5 days then they should be excluded.

/*
-- Output:
USER_ID		START_DATE		END_DATE		CONSECUTIVE_DAYS
1			10/03/2024		14/03/2024		5
1 			25/03/2024		30/03/2024		6
3 			01/03/2024		05/03/2024		5
*/


-- PostgreSQL Dataset
drop table if exists user_login;
create table user_login
(
	user_id		int,
	login_date	date
);
insert into user_login values(1, to_date('01/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('02/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('03/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('04/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('06/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('10/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('11/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('12/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('13/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('14/03/2024','dd/mm/yyyy'));

insert into user_login values(1, to_date('20/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('25/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('26/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('27/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('28/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('29/03/2024','dd/mm/yyyy'));
insert into user_login values(1, to_date('30/03/2024','dd/mm/yyyy'));

insert into user_login values(2, to_date('01/03/2024','dd/mm/yyyy'));
insert into user_login values(2, to_date('02/03/2024','dd/mm/yyyy'));
insert into user_login values(2, to_date('03/03/2024','dd/mm/yyyy'));
insert into user_login values(2, to_date('04/03/2024','dd/mm/yyyy'));

insert into user_login values(3, to_date('01/03/2024','dd/mm/yyyy'));
insert into user_login values(3, to_date('02/03/2024','dd/mm/yyyy'));
insert into user_login values(3, to_date('03/03/2024','dd/mm/yyyy'));
insert into user_login values(3, to_date('04/03/2024','dd/mm/yyyy'));
insert into user_login values(3, to_date('04/03/2024','dd/mm/yyyy'));
insert into user_login values(3, to_date('04/03/2024','dd/mm/yyyy'));
insert into user_login values(3, to_date('05/03/2024','dd/mm/yyyy'));

insert into user_login values(4, to_date('01/03/2024','dd/mm/yyyy'));
insert into user_login values(4, to_date('02/03/2024','dd/mm/yyyy'));
insert into user_login values(4, to_date('03/03/2024','dd/mm/yyyy'));
insert into user_login values(4, to_date('04/03/2024','dd/mm/yyyy'));
insert into user_login values(4, to_date('04/03/2024','dd/mm/yyyy'));


select * from user_login;

;with without_duplicate as (
	select
		user_id,
		login_date
	from user_login
		group by user_id, login_date
), dataset as (
select  
	user_id,
	login_date,
	--dense_rank() over(partition by user_id order by user_id, login_date) as drn,
	--extract(day from login_date) as day,
	login_date - cast (dense_rank() over (partition by user_id order by user_id, login_date) as int) as dateg
	--row_number() over(partition by user_id order by user_id, login_date) as rn
from without_duplicate t1
)
	

select 
	user_id,
	dateg,
	min(login_date) as start_date,
	max(login_date) as end_date, count(*)
from dataset
group by user_id, dateg
having (count(*)) > 4


