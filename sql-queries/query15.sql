DROP TABLE IF EXISTS Friends;

CREATE TABLE Friends
(
	Friend1 	VARCHAR(10),
	Friend2 	VARCHAR(10)
);
INSERT INTO Friends VALUES ('Jason','Mary');
INSERT INTO Friends VALUES ('Mike','Mary');
INSERT INTO Friends VALUES ('Mike','Jason');
INSERT INTO Friends VALUES ('Susan','Jason');
INSERT INTO Friends VALUES ('John','Mary');
INSERT INTO Friends VALUES ('Susan','Mary');

select * from Friends;

select
	*
from Friends f1
join Friends f2
	on f1.friend1 = f1.friend2

;with friend_all as (
	select friend1, friend2 from friends
	union 
	select friend2, friend1 from friends
)
select
	*
from friend_all fl
left join friends fa
	on fa.friend1 = fl.friend1
	and fl.friend2 in (select friend2 from friend_all where friend1 = fa.friend2)
where fa.friend1 = 'Jason'
