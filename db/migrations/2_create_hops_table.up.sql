CREATE TABLE  IF NOT EXISTS`hops` (
	`id`							INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
	`name`							TEXT NOT NULL,
	`country`						TEXT,
	`description`					TEXT,
	`source`						TEXT,
	`hop_use_id`					INTEGER,
	`hop_storage_index`				DECIMAL(3, 1),
	`alpha_acid_percentage_min`		DECIMAL(3, 1),
	`alpha_acid_percentage_max`		DECIMAL(3, 1),
	`beta_acid_percentage_min`		DECIMAL(3, 1),
	`beta_acid_percentage_max`		DECIMAL(3, 1),
	`co_humolone_percentage_max`	DECIMAL(3, 1),
	`co_humolone_percentage_min`	DECIMAL(3, 1),
	`total_oil_min`					DECIMAL(3, 1),
	`total_oil_max`					DECIMAL(3, 1),
	/* `b_pinene_percentage_min`		DECIMAL(3, 1), */
	/* `b_pinene_percentage_max`		DECIMAL(3, 1), */
	`myrcene_percentage_min`		DECIMAL(3, 1),
	`myrcene_percentage_max`		DECIMAL(3, 1),
	/* `linalool_percentage_min`		DECIMAL(3, 1), */
	/* `linalool_percentage_max`		DECIMAL(3, 1), */
	`caryophyllene_percentage_min`	DECIMAL(3, 1),
	`caryophyllene_percentage_max`	DECIMAL(3, 1),
	`farnesene_percentage_min`		DECIMAL(3, 1),
	`farnesene_percentage_max`		DECIMAL(3, 1),
	`humulene_percentage_min`		DECIMAL(3, 1),
	`humulene_percentage_max`		DECIMAL(3, 1),
	/* `geraniol_percentage_min`		DECIMAL(3, 1), */
	/* `geraniol_percentage_max`		DECIMAL(3, 1), */
	FOREIGN KEY(`hop_use_id`) REFERENCES hop_uses(id)
);
CREATE INDEX hop_use_id ON hops(hop_use_id);
