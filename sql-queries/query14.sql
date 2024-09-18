drop table if exists invoice;
create table invoice
(
	serial_no		int,
	invoice_date	date
);
insert into invoice values (330115, to_date('01-Mar-2024','DD-MON-YYYY'));
insert into invoice values (330120, to_date('01-Mar-2024','DD-MON-YYYY'));
insert into invoice values (330121, to_date('01-Mar-2024','DD-MON-YYYY'));
insert into invoice values (330122, to_date('02-Mar-2024','DD-MON-YYYY'));
insert into invoice values (330125, to_date('02-Mar-2024','DD-MON-YYYY'));

select * from invoice;

;with recursive cte as (
	select min(serial_no) as sn from invoice 
	union all
	select (sn + 1) from cte
	where sn < (select max(serial_no) from invoice)
)

select 
	f.sn
from cte f
left join invoice i
	on f.sn = i.serial_no
where i.serial_no is null
order by 1