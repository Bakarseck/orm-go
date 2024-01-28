# Documentation de SQLBuilder

## Vue d'ensemble

Ce code Go implémente un constructeur de requêtes SQL (SQL Builder) pour interagir avec une base de données. Il permet de créer des requêtes SQL de manière programmative et flexible.

## Structure : SQLBuilder

- **Champs** :

  - `query` : Une chaîne de caractères pour stocker la requête SQL en cours de construction.

  - `parameters` : Un tableau pour stocker les paramètres de la requête, utilisé pour la prévention des injections SQL.

- **Méthodes** :

  - `NewSQLBuilder` : Constructeur qui initialise et retourne une nouvelle instance de `SQLBuilder`.

  - `Build` : Retourne la requête SQL construite et les paramètres associés.

  - `Clear` : Réinitialise le constructeur pour commencer une nouvelle requête.

## Fonctions de Construction de Requêtes

Ces méthodes modifient la requête SQL et les paramètres en fonction des actions spécifiques (INSERT, UPDATE, SELECT, etc.) :

- **Insert** :

  - Construit une requête d'insertion pour une table donnée avec les valeurs spécifiées.

  - Utilise `table.Name` pour le nom de la table et `table.GetFieldName()` pour les noms des champs.

  - Ajoute des placeholders `?` pour les valeurs à insérer.

- **Update** :

  - Construit une requête de mise à jour pour une table spécifiée par `updates.Model.Name`.

  - Utilise `updates.field` et `updates.value` pour construire la clause SET.

- **Delete** :

  - Commence la construction d'une requête de suppression.

- **Select** :

  - Construit une requête SELECT avec les colonnes spécifiées.

- **SelectAll** :

  - Construit une requête SELECT pour toutes les colonnes (`SELECT *`).

- **From** :

  - Spécifie la table à utiliser dans la requête SELECT.

- **Where**, **And**, **Or** :

  - Ajoutent des conditions à la requête en utilisant les opérateurs WHERE, AND, et OR respectivement.

  - Utilisent des placeholders `?` pour les valeurs de condition.

## Points Clés

- **Sécurité** : L'utilisation de placeholders `?` dans les requêtes aide à prévenir les injections SQL.

- **Flexibilité** : Les méthodes chaînables permettent de construire des requêtes complexes de manière lisible et maintenable.

- **Abstraction** : Le code cache les détails de la construction de requêtes SQL, rendant le code qui l'utilise plus propre et plus facile à comprendre.

## Exemple d'Utilisation

```go

builder := NewSQLBuilder()

query, params := builder.Select("name", "age").From(&Table{Name: "users"}).Where("id", 1).Build()

// Ceci construit une requête "SELECT name, age FROM users WHERE id = ?", avec params contenant [1].
```
