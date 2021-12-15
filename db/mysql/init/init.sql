-- SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
-- SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
-- SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';
SET CHARSET UTF8;
DROP SCHEMA IF EXISTS `app`;
CREATE SCHEMA IF NOT EXISTS `app` DEFAULT CHARACTER SET utf8;
USE `app`;

-- drop ----
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `users_genres`;
DROP TABLE IF EXISTS `users_info`;
DROP TABLE IF EXISTS `users_list`;
DROP TABLE IF EXISTS `articles_contents`;
DROP TABLE IF EXISTS `articles_tag`;
DROP TABLE IF EXISTS `articles_comments`;
DROP TABLE IF EXISTS `articles_nice_status`;

-- create ----
CREATE TABLE IF NOT EXISTS `app`.`users` (
`id` VARCHAR(32) NOT NULL COMMENT 'ユーザID',
`name` VARCHAR(32) NOT NULL COMMENT 'ユーザ名',
`image` VARCHAR(128) NOT NULL COMMENT 'プロフィール画像',
`year` int NOT NULL COMMENT '生年月日(年)',
`month` int NOT NULL COMMENT '生年月日(月)',
`day` int NOT NULL COMMENT '生年月日(日)',
`gender` int NOT NULL COMMENT '性別',
PRIMARY KEY (`id`),
INDEX `idx_auth_token` (`id` ASC)
)
ENGINE = InnoDB
COMMENT = 'ユーザプロフィール';

CREATE TABLE IF NOT EXISTS `app`.`users_genres` (
`id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
`genre_1` VARCHAR(64) NOT NULL COMMENT '1つ目のジャンル',
`genre_2` VARCHAR(64) NOT NULL COMMENT '2つ目のジャンル',
`genre_3` VARCHAR(64) NOT NULL COMMENT '3つ目のジャンル',
`genre_4` VARCHAR(64) NOT NULL COMMENT '4つ目のジャンル',
PRIMARY KEY (`id`)
)
ENGINE = InnoDB
COMMENT = 'ユーザ選択ジャンル';

CREATE TABLE IF NOT EXISTS `app`.`users_info` (
`id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
`password` VARCHAR(64) NOT NULL COMMENT 'パスワード',
PRIMARY KEY (`id`)
)
ENGINE = InnoDB
COMMENT = 'ユーザ情報';

CREATE TABLE IF NOT EXISTS `app`.`users_list` (
`list_id` VARCHAR(64) NOT NULL COMMENT '識別ID',
`user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
PRIMARY KEY (`list_id`),
FOREIGN KEY (`user_id`)
    REFERENCES `app`.`users` (`id`)
)
ENGINE = InnoDB
COMMENT = 'リスト';

CREATE TABLE IF NOT EXISTS `app`.`articles_contents` (
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
`image_path` VARCHAR(128) NOT NULL COMMENT '画像のパス',
`title` VARCHAR(32) NOT NULL COMMENT '記事のタイトル',
`context` text NOT NULL COMMENT '記事の内容',
`genre` VARCHAR(64) NOT NULL COMMENT '記事のジャンル',
`nice` int NOT NULL COMMENT 'いいね数',
`era_year` int NOT NULL COMMENT '年代(年)',
`era_month` int NOT NULL COMMENT '年代(月)',
PRIMARY KEY (`article_id`)
)
ENGINE = InnoDB
COMMENT = '記事の内容';

CREATE TABLE IF NOT EXISTS `app`.`articles_tag` (
`tag_id` VARCHAR(64) NOT NULL COMMENT 'タグID',
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
`article_tag` VARCHAR(32) NOT NULL COMMENT '記事のタグ',
PRIMARY KEY (`tag_id`)
)
ENGINE = InnoDB
COMMENT = '記事のタグ';

CREATE TABLE IF NOT EXISTS `app`.`articles_comments` (
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
`comments_id` VARCHAR(64) NOT NULL COMMENT 'コメントのID',
`comments_contents` VARCHAR(64) NOT NULL COMMENT 'コメントの内容',
`user_name` VARCHAR(64) NOT NULL COMMENT 'ユーザネーム',
`user_image` VARCHAR(128) NOT NULL COMMENT 'ユーザーの画像',
PRIMARY KEY (`comments_id`)
)
ENGINE = InnoDB
COMMENT = '記事へのコメント';

CREATE TABLE IF NOT EXISTS `app`.`articles_nice_status` (
`nice_id` VARCHAR(64) NOT NULL COMMENT 'NiceID',
`article_id` VARCHAR(64) NOT NULL COMMENT '記事識別ID',
`user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
PRIMARY KEY (`nice_id`),
FOREIGN KEY (`user_id`)
    REFERENCES `app`.`users` (`id`)
)
ENGINE = InnoDB
COMMENT = '記事にいいねした人';

CREATE TABLE IF NOT EXISTS `app`.`articles_requests` (
`request_id` VARCHAR(64) NOT NULL COMMENT 'リクエスト識別ID',
`user_id` VARCHAR(64) NOT NULL COMMENT 'ユーザID',
`genre` VARCHAR(64) COMMENT  'ジャンル',
`title` VARCHAR(128) COMMENT 'タイトル',
`context` TEXT COMMENT '内容',
`year` INT COMMENT '年代(年)',
`month` INT COMMENT '年代(月)',
PRIMARY KEY (`request_id`),
FOREIGN KEY (`user_id`)
    REFERENCES `app`.`users` (`id`)
)
ENGINE = InnoDB
COMMENT = 'リクエスト';

-- insert ----
INSERT INTO `articles_contents` VALUES ('1', 'ImageURL', 'Title_sample1','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',46,2020,6);
INSERT INTO `articles_contents` VALUES ('6', 'ImageURL', 'Title_sample2','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',4,2019,12);
INSERT INTO `articles_contents` VALUES ('7', 'ImageURL', 'Title_sample3','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',6,2018,11);
INSERT INTO `articles_contents` VALUES ('8', 'ImageURL', 'Title_sample4','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',146,2017,3);
INSERT INTO `articles_contents` VALUES ('9', 'ImageURL', 'Title_sample5','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',416,2016,7);
INSERT INTO `articles_contents` VALUES ('10', 'ImageURL', 'Title_sample6','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',11,2015,8);
INSERT INTO `articles_contents` VALUES ('11', 'ImageURL', 'Title_sample7','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',43,2014,9);
INSERT INTO `articles_contents` VALUES ('12', 'ImageURL', 'Title_sample8','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',73,2013,1);
INSERT INTO `articles_contents` VALUES ('13', 'ImageURL', 'Title_sample9','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',2,2012,2);
INSERT INTO `articles_contents` VALUES ('14', 'ImageURL', 'Title_sample10','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',12,2011,3);
INSERT INTO `articles_contents` VALUES ('15', 'ImageURL', 'Title_sample11','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',73,2010,7);
INSERT INTO `articles_contents` VALUES ('2', 'ImageURL','Title_sample5', 'sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2sample2','スポーツ',100,2000,5 );
INSERT INTO `articles_contents` VALUES ('3', 'ImageURL','Title_sample10', 'sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3sample3','国際',2,1999,10 );
INSERT INTO `articles_contents` VALUES ('4', 'ImageURL', 'Title_sample15','sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4sample4','音楽',0,1940,6);
INSERT INTO `articles_contents` VALUES ('5', 'ImageURL', 'Title_sample20','sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5sample5','グルメ',123,2006,2 );
INSERT INTO `articles_contents` VALUES ('16', 'ImageURL', 'Title_sample1','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',46,2020,7);
INSERT INTO `articles_contents` VALUES ('17', 'ImageURL', 'Title_sample1','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',46,2020,5);
INSERT INTO `articles_contents` VALUES ('18', 'ImageURL', 'Title_sample1','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',46,2020,5);
INSERT INTO `articles_contents` VALUES ('19', 'ImageURL', 'Title_sample1','五稜郭で飲んだ帰りsample1sample1sample1sample1sample1sample1sample1sample1','ドラマ',46,2020,4);

INSERT INTO `users_info` VALUES ('111111', 'ede840d10d2d313a870bc131a4e2c311d7ad09bdf32b3418147221f51a6e2');

INSERT INTO `users` VALUES ('111111','user1','ImageURL',1999,8,30,1);

INSERT INTO `articles_tag` VALUES ('1','1', '1_sampleTag1');
INSERT INTO `articles_tag` VALUES ('2','1', '1_sampleTag2');
INSERT INTO `articles_tag` VALUES ('3','1', '1_sampleTag3');
INSERT INTO `articles_tag` VALUES ('4','1', '1_sampleTag4');
INSERT INTO `articles_tag` VALUES ('5','2', '2_sampleTag1');
INSERT INTO `articles_tag` VALUES ('6','2', '2_sampleTag2');
INSERT INTO `articles_tag` VALUES ('7','2', '2_sampleTag3');
INSERT INTO `articles_tag` VALUES ('8','2', '2_sampleTag4');
INSERT INTO `articles_tag` VALUES ('9','3', '3_sampleTag1');
INSERT INTO `articles_tag` VALUES ('10','3', '3_sampleTag2');
INSERT INTO `articles_tag` VALUES ('11','3', '3_sampleTag3');
INSERT INTO `articles_tag` VALUES ('12','3', '3_sampleTag4');
INSERT INTO `articles_tag` VALUES ('13','4', '4_sampleTag1');
INSERT INTO `articles_tag` VALUES ('14','4', '4_sampleTag2');
INSERT INTO `articles_tag` VALUES ('15','4', '4_sampleTag3');
INSERT INTO `articles_tag` VALUES ('16','4', '4_sampleTag4');
INSERT INTO `articles_tag` VALUES ('17','5', '5_sampleTag1');
INSERT INTO `articles_tag` VALUES ('18','5', '5_sampleTag2');
INSERT INTO `articles_tag` VALUES ('19','5', '5_sampleTag3');
INSERT INTO `articles_tag` VALUES ('20','5', '5_sampleTag4');

INSERT INTO `users_list` VALUES ('1','111111','1');
INSERT INTO `users_list` VALUES ('2','111111','4');
INSERT INTO `users_list` VALUES ('3','111111','19');


INSERT INTO `articles_nice_status` VALUES ('1','1','111111');

INSERT INTO `articles_comments` VALUES ('1','2v43d6ef-a83d-57d0-f33d-b5d78ncj4a58','面白い','user1','ImageURL');
INSERT INTO `articles_comments` VALUES ('1','feeafa01-38ed-48a1-be98-4d926db542e2','1コメ','user2','ImageURL');
INSERT INTO `articles_comments` VALUES ('1','389489a7-c6cc-42f1-a0d3-cfd497941585','2コメ','user3','ImageURL');
INSERT INTO `articles_comments` VALUES ('1','3bee2b4e-e179-47cd-a4b6-acd8092c4e4a','3コメ','user4','ImageURL');
INSERT INTO `articles_comments` VALUES ('1','6c928198-1419-4b03-877e-99b37e1eb460','4コメ','user5','ImageURL');
INSERT INTO `articles_comments` VALUES ('1','ce3137ea-1e10-43a2-884f-fa0b13971f03','5コメ','user6','ImageURL');
