DROP TABLE IF EXISTS company;
CREATE TABLE company
(
	employee	varchar(10) primary key,
	manager		varchar(10)
);

INSERT INTO company values ('Elon', null);
INSERT INTO company values ('Ira', 'Elon');
INSERT INTO company values ('Bret', 'Elon');
INSERT INTO company values ('Earl', 'Elon');
INSERT INTO company values ('James', 'Ira');
INSERT INTO company values ('Drew', 'Ira');
INSERT INTO company values ('Mark', 'Bret');
INSERT INTO company values ('Phil', 'Mark');
INSERT INTO company values ('Jon', 'Mark');
INSERT INTO company values ('Omid', 'Earl');

SELECT * FROM company;

;with  recursive base_team as (
select 
	c.employee, c.manager,
	 'team ' || row_number() over (order by c.employee) as teams
from company c where c.manager in (select employee from company where manager is null)
), t1 as (
	select 
		*
	from base_team bt

	union all

	select
		c.employee, c.manager, t1.teams
	from t1
	join company c
		on c.manager = t1.employee
), t2 as (

	select
		c.employee, c.manager, bt.teams
	from company c
	join base_team bt
	on 1 = 1
	where c.manager is null
	
	union all
	
	select
		employee, manager, teams
	from t1
)

select
	teams,
	string_agg(employee, ',') as members
	from t2
	group by teams;
	