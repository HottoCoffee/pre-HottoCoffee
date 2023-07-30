create table workspace (
    id int unsigned not null auto_increment,
    name varchar(255) not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    deleted_at datetime,
    primary key (id),
    key idx_deleted_at (deleted_at)
) engine = InnoDB default charset = utf8;

create table batch (
    id int unsigned not null auto_increment,
    batch_name varchar(255) not null,
    server_name varchar(255) not null,
    cron_setting varchar(255) not null,
    initial_date date not null,
    time_limit int unsigned not null,
    estimated_duration int unsigned not null,
    workspace_id int unsigned not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    deleted_at datetime,
    primary key (id),
    key idx_deleted_at (deleted_at),
    foreign key fk_workspace_id (workspace_id) references workspace (id)
) engine = InnoDB default charset = utf8;

create table history (
    id int unsigned not null auto_increment,
    batch_id int unsigned not null,
    status varchar(255) not null,
    start_datetime datetime not null,
    finish_datetime datetime not null,
    workspace_id int unsigned not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    deleted_at datetime,
    primary key (id),
    key idx_deleted_at (deleted_at),
    key idx_created_at (created_at),
    foreign key fk_batch_id (batch_id) references batch (id),
    foreign key fk_workspace_id (batch_id) references workspace (id)
) engine = InnoDB default charset = utf8;

create table workspace_user (
    id int unsigned not null auto_increment,
    workspace_id int unsigned not null,
    user_id int unsigned not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    deleted_at datetime,
    primary key (id),
    key idx_deleted_at (deleted_at),
    foreign key fk_workspace_id (workspace_id) references workspace (id),
    foreign key fk_user_id (user_id) references user (id)
) engine = InnoDB default charset = utf8;

create table workspace_access_key (
    id int unsigned not null auto_increment,
    access_key varchar(255) not null,
    display_name varchar(255) not null,
    workspace_id int unsigned not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    deleted_at datetime,
    primary key (id),
    key idx_deleted_at (deleted_at),
    foreign key fk_workspace_id (workspace_id) references workspace (id)
) engine = InnoDB default charset = utf8;
