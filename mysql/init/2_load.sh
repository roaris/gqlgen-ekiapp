mysql -uroot -ppassword --local-infile=1 db_dev \
-e "load data local infile '/docker-entrypoint-initdb.d/station20220314free.csv' into table stations fields terminated by ',';
load data local infile '/docker-entrypoint-initdb.d/company20200619.csv' into table companies fields terminated by ',';
load data local infile '/docker-entrypoint-initdb.d/line20220314free.csv' into table station_lines fields terminated by ',';
load data local infile '/docker-entrypoint-initdb.d/join20220314.csv' into table station_joins fields terminated by ',';"
