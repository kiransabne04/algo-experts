drop table if exists mountain_huts;
create table mountain_huts 
(
	id 			integer not null unique,
	name 		varchar(40) not null unique,
	altitude 	integer not null
);
insert into mountain_huts values (1, 'Dakonat', 1900);
insert into mountain_huts values (2, 'Natisa', 2100);
insert into mountain_huts values (3, 'Gajantut', 1600);
insert into mountain_huts values (4, 'Rifat', 782);
insert into mountain_huts values (5, 'Tupur', 1370);

drop table if exists trails;
create table trails 
(
	hut1 		integer not null,
	hut2 		integer not null
);
insert into trails values (1, 3);
insert into trails values (3, 2);
insert into trails values (3, 5);
insert into trails values (4, 5);
insert into trails values (1, 5);

select * from mountain_huts;
select * from trails;

;with r1 as (
select
		h1.id, h1.name as h1name, h1.altitude as h1altitude, t1.hut1, t1.hut2
	from mountain_huts h1
	join trails t1
		on h1.id = t1.hut1
), r2 as (
select 
		h1name, h1altitude, hut1, hut2, h2.name as h2name, h2.altitude as h2altitude
	from r1
	join mountain_huts h2
		on r1.hut2 = h2.id
), final_route as (
select 
	*,
	case 
		when h1altitude > h2altitude 
		then h1name 
		else h2name 
	end as startpnt,
	case 
		when h1altitude > h2altitude
		then hut1 
		else hut2
	end as starthut,
	case 
		when h1altitude > h2altitude
		then h2name
		else h1name
	end as endpnt,
	case
		when h1altitude > h2altitude 
		then hut2
		else hut1
	end as endhut
from r2
)

select
	t1.startpnt, t2.startpnt as midpnt, t2.endpnt
from final_route t1
join final_route t2
	on t1.endhut = t2.starthut

select * from final_route