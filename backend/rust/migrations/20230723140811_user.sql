create table user (
    id int unsigned not null auto_increment,
    display_name varchar(255) not null,
    email blob not null,
    password char(60) not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    deleted_at datetime,
    primary key (id),
    key idx_deleted_at (deleted_at),
    unique uni_email (email(255))
) engine = InnoDB default charset = utf8;
