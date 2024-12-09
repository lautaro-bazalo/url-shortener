use `url-shortener`;

SET FOREIGN_KEY_CHECKS = 0;
DROP TABLE IF EXISTS `url`;

CREATE TABLE `url` (
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    short_url VARCHAR(255),
    original_url VARCHAR(255) UNIQUE NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    deleted_at DATETIME NULL,
    PRIMARY KEY (id)
);
