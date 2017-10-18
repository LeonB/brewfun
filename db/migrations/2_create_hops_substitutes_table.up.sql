CREATE TABLE  IF NOT EXISTS`hop_substitutes` (
	`hop_a_id`	INTEGER,
	`hop_b_id`	INTEGER,
	`match`		DECIMAL(3, 1),
	`source` TEXT,
	FOREIGN KEY(`hop_a_id`) REFERENCES hops,
	FOREIGN KEY(`hop_b_id`) REFERENCES hops
);
CREATE INDEX hop_a_id ON hops(id);
CREATE INDEX hop_b_id ON hops(id);
