CREATE TABLE  IF NOT EXISTS `hops_aromas` (
	`id`			INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`hop_id`		INTEGER,
	`hop_aroma_id`	INTEGER,
	`match`			DECIMAL(3, 1),
	`source`		TEXT,
	FOREIGN KEY(`hop_id`) REFERENCES hops(id),
	FOREIGN KEY(`hop_aroma_id`) REFERENCES hop_aromas(id)
);
CREATE INDEX hop_id ON hops_aromas(hop_id);
CREATE INDEX hop_aroma_id ON hops_aromas(hop_aroma_id);
