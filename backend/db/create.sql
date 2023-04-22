create table batch (
    id int unsigned not null auto_increment,
    batch_name varchar(255) not null,
    server_name varchar(255) not null,
    cron_setting varchar(255) not null,
    initial_date date not null,
    time_limit int unsigned not null,
    estimated_duration int unsigned not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    deleted_at datetime,
    primary key (id),
    key idx_deleted_at (deleted_at)
) engine = InnoDB default charset = utf8;

create table history (
    id int unsigned not null auto_increment,
    batch_id int unsigned not null,
    status varchar(255) not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    deleted_at datetime,
    primary key (id),
    key idx_deleted_at (deleted_at),
    key idx_created_at (created_at),
    foreign key fk_batch_id (batch_id) references batch (id)
) engine = InnoDB default charset = utf8;
