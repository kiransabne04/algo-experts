create table lifts
(
      id         	  int
    , capacity_kg     int
);

insert into lifts values (1, 300);
insert into lifts values (2, 350);

create table lift_passengers
(
      passenger_name    varchar(50)
    , weight_kg     	int
	, lift_id			int
);

insert into lift_passengers values ('Rahul', 85, 1);
insert into lift_passengers values ('Adarsh', 73, 1);
insert into lift_passengers values ('Riti', 95, 1);
insert into lift_passengers values ('Dheeraj', 80, 1);
insert into lift_passengers values ('Vimal', 83, 2);
insert into lift_passengers values ('Neha', 77, 2);
insert into lift_passengers values ('Priti', 73, 2);
insert into lift_passengers values ('Himanshi', 85, 2);

select * from lifts;
select * from lift_passengers;

;with passenger_weights as (
select 
	*,
	sum(weight_kg) over(partition by lift_id order by weight_kg) as sum_weight
from lift_passengers
)

select
	l.id as lift_id,
	string_agg(p.passenger_name, ' , ') as passengers
from lifts l
join passenger_weights p
	on l.id = p.lift_id
	and l.capacity_kg >= p.sum_weight
group by l.id
