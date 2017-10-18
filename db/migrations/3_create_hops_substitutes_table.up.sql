CREATE TABLE IF NOT EXISTS`hops_substitutes` (
	`id`		INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`hop_a_id`	INTEGER,
	`hop_b_id`	INTEGER,
	`match`		DECIMAL(3, 1),
	`source` TEXT,
	FOREIGN KEY(`hop_a_id`) REFERENCES hops(id),
	FOREIGN KEY(`hop_b_id`) REFERENCES hops(id)
);
CREATE INDEX hop_a_id ON hops_substitutes(hop_a_id);
CREATE INDEX hop_b_id ON hops_substitutes(hop_b_id);
