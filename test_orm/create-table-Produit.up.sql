CREATE TABLE IF NOT EXISTS Produit (
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
	Name_produit TEXT NOT NULL,
	Prix INTEGER ,
	User_id INTEGER NOT NULL,
	FOREIGN KEY (User_id) REFERENCES User (Id)
)