set global local_infile = 1;

create table stations(
    station_cd int not null,
    station_g_cd int not null,
    station_name varchar(100) not null,
    station_name_k varchar(100),
    station_name_r varchar(100),
    line_cd int,
    pref_cd int,
    post varchar(100),
    address varchar(100),
    lon float,
    lat float,
    open_ymd varchar(100),
    close_ymd varchar(100),
    e_status int,
    e_sort int,
    PRIMARY KEY(station_cd)
);

create table companies(
    company_cd int not null,
    rr_cd int not null,
    company_name varchar(100) not null,
    company_name_k varchar(100),
    company_name_h varchar(100),
    company_name_r varchar(100),
    company_url varchar(100),
    company_type int,
    e_status int,
    e_sort int,
    PRIMARY KEY(company_cd)
);

create table station_lines(
    line_cd int not null,
    company_cd int not null,
    line_name varchar(100) not null,
    line_name_k varchar(100),
    line_name_h varchar(100),
    line_color_c varchar(100),
    line_color_t varchar(100),
    line_type int,
    lon float,
    lat float,
    zoom int,
    e_status int,
    e_sort int,
    PRIMARY KEY(line_cd)
);

create table station_joins(
    line_cd int not null,
    station_cd1 int not null,
    station_cd2 int not null,
    PRIMARY KEY(line_cd, station_cd1, station_cd2)
);
