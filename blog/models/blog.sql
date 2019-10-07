CREATE TABLE  `blog_tag`(
	`id` int(10)  NOT NULL AUTO_INCREMENT,
	`name` varchar(100)  DEFAULT '' COMMENT 'label',
	`create_on` int(10) unsigned DEFAULT '0' COMMENT 'create_time',
	`create_by` varchar(100)  DEFAULT '' COMMENT 'author', 
	`modify_on`	 int(10) unsigned DEFAULT '0' COMMENT 'create_time',
	`modify_by`  varchar(100)  DEFAULT '' COMMENT 'modifier', 
	`delete_on`  int(10) unsigned DEFAULT '0' COMMENT 'delete_time',
	`state`    tinyint(3) unsigned DEFAULT '1'  COMMENT 'state 0/1'
	PRIMARY KEY(`id`)

)ENGINE=InnoDB  DEFAULT CHARSET=utf8  COMMENT=`tag management`;


CREATE TABLE `blog_article`(
    `id` int(10)  NOT NULL AUTO_INCREMENT,
    `tag_id` int(10)  NOT NULL  DEFAULT '0' COMMENT 'arttcle',
    `title` varchar(100)  DEFAULT '' COMMENT 'title', 
    `content` text,
    `desc` varchar(100) DEFAULT '' COMMENT '简述',
    `create_on` int(10) unsigned DEFAULT '0' COMMENT 'create_time',
	`create_by` varchar(100)  DEFAULT '' COMMENT 'author', 
	`modify_on`	 int(10) unsigned DEFAULT '0' COMMENT 'create_time',
	`modify_by`  varchar(100)  DEFAULT '' COMMENT 'modifier', 
	`delete_on`  int(10) unsigned DEFAULT '0' COMMENT 'delete_time',
	`state`    tinyint(3) unsigned DEFAULT '1'  COMMENT 'state 0/1'
	PRIMARY KEY(`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8  COMMENT=`article management`;



CREATE  TABLE `blog_auth`(
     `id` int(10)  NOT NULL AUTO_INCREMENT,
     `username` varchar(100) DEFAULT '' COMMENT `username`,
     `password` varchar(100) DEFAULT '' COMMENT `password`,
     PRIMARY KEY(`id`)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8  COMMENT=`auth management`;