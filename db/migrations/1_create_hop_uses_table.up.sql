CREATE TABLE IF NOT EXISTS `hop_uses` (
	`id`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`usage`	TEXT
);

INSERT INTO `hop_uses` VALUES
	(1, 'bittering'),
	(2, 'aroma'),
	(3, 'both')
;
