# Documentation SQLBuilder

## English

`SQLBuilder` is a structure used to build SQL queries in a programmatic and flexible manner. It supports various SQL operations like INSERT, UPDATE, DELETE, and SELECT.

- `NewSQLBuilder()`: Creates and returns a new instance of `SQLBuilder`.
- `Build()`: Returns the constructed SQL query string along with any parameters.
- `Clear()`: Clears the current query and parameters from the builder.
- `Insert(table *Table, values []interface{})`: Constructs an INSERT query for the given table and values.
- `Update(updates *Modifier)`: Constructs an UPDATE query using the provided modifications.
- `Delete()`: Starts a DELETE query construction.
- `Select(columns ...string)`: Begins a SELECT query with specified columns.
- `SelectAll()`: Begins a SELECT query that selects all columns.
- `From(table *Table)`: Specifies the table for a SELECT query.
- `Where(column string, value interface{})`: Adds a WHERE clause with the given condition.
- `And(column string, value interface{})`: Adds an AND condition to the WHERE clause.
- `Or(column string, value interface{})`: Adds an OR condition to the WHERE clause.

## Français

`SQLBuilder` est une structure utilisée pour construire des requêtes SQL de manière programmatique et flexible. Elle prend en charge diverses opérations SQL telles que INSERT, UPDATE, DELETE et SELECT.

- `NewSQLBuilder()`: Crée et retourne une nouvelle instance de `SQLBuilder`.
- `Build()`: Renvoie la chaîne de requête SQL construite ainsi que les paramètres.
- `Clear()`: Efface la requête actuelle et les paramètres du constructeur.
- `Insert(table *Table, values []interface{})`: Construit une requête INSERT pour la table et les valeurs données.
- `Update(updates *Modifier)`: Construit une requête UPDATE en utilisant les modifications fournies.
- `Delete()`: Commence la construction d'une requête DELETE.
- `Select(columns ...string)`: Commence une requête SELECT avec les colonnes spécifiées.
- `SelectAll()`: Commence une requête SELECT qui sélectionne toutes les colonnes.
- `From(table *Table)`: Spécifie la table pour une requête SELECT.
- `Where(column string, value interface{})`: Ajoute une clause WHERE avec la condition donnée.
- `And(column string, value interface{})`: Ajoute une condition AND à la clause WHERE.
- `Or(column string, value interface{})`: Ajoute une condition OR à la clause WHERE.
