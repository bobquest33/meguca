select insert_thread(
	$1::varchar(100),
	$2::bigint,
	$3::bool,
	$4::bool,
	$5::bigint,
	$6::varchar(3),
	$7::bigint,
	$8::bigint,
	$9::varchar(2000),
	$10::varchar(50),
	$11::char(10),
	$12::varchar(20),
	$13::bytea,
	$14::inet,
	$15::char(40),
	$16::varchar(200),
	$17::bigint[][2],
	$18::bigint[][2],
	$19::json[]
);
