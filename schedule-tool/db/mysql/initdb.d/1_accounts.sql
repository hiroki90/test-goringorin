CREATE TABLE IF NOT EXISTS accounts(
    id   INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    age  INT UNSIGNED NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- TODO: shedule追加