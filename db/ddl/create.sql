CREATE DATABASE vote CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;


create table if not exists m_votes (
    id bigint(20) unsigned AUTO_INCREMENT not null comment 'ID',
    `name` varchar(128) not null comment '名称',
    created_at datetime not null comment '作成日時',
    updated_at datetime not null comment '更新日時',
    deleted_at datetime null comment '削除日時',
    primary key (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='投票対象マスタ';

create table if not exists u_votes (
    id bigint(20) unsigned AUTO_INCREMENT not null comment 'ID',
    m_vote_id bigint(20) unsigned not null comment '投票対象ID',
    created_at datetime not null comment '作成日時',
    updated_at datetime not null comment '更新日時',
    deleted_at datetime null comment '削除日時',
    primary key (id),
    CONSTRAINT fk_vote_id
        FOREIGN KEY (m_vote_id)
        REFERENCES m_votes (id)
        ON DELETE RESTRICT ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci comment='投票';


insert into m_votes values
(NULL, "りんご", NOW(), NOW(), NULL),
(NULL, "バナナ", NOW(), NOW(), NULL),
(NULL, "ぶどう", NOW(), NOW(), NULL);
