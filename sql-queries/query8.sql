drop table if exists job_skills;
create table job_skills
(
	row_id		int,
	job_role	varchar(20),
	skills		varchar(20)
);
insert into job_skills values (1, 'Data Engineer', 'SQL');
insert into job_skills values (2, null, 'Python');
insert into job_skills values (3, null, 'AWS');
insert into job_skills values (4, null, 'Snowflake');
insert into job_skills values (5, null, 'Apache Spark');
insert into job_skills values (6, 'Web Developer', 'Java');
insert into job_skills values (7, null, 'HTML');
insert into job_skills values (8, null, 'CSS');
insert into job_skills values (9, 'Data Scientist', 'Python');
insert into job_skills values (10, null, 'Machine Learning');
insert into job_skills values (11, null, 'Deep Learning');
insert into job_skills values (12, null, 'Tableau');

select * from job_skills;

;with cte as (
select 
	*, 
	sum(case when job_role is null then 0 else 1 end) over(order by row_id) job_flag
from job_skills
)

select *,
	first_value(job_role) over(partition by job_flag order by row_id)
from cte

--
;with recursive cte as (
	select
		row_id,
		skills,
		job_role
	from job_skills where row_id = 1
	
	union all

	select
		t1.row_id,
		t1.skills,
		coalesce(t1.job_role, t2.job_role) as filled
	from job_skills t1
	join cte t2
		on t1.row_id = t2.row_id + 1
	
)
--select * from cte
select * from cte
