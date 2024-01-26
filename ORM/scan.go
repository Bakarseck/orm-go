package orm

import (
	"fmt"
	"reflect"
	"strings"
)

func (o *ORM) ScanRows(model interface{}) error {
	// Obtenir le type du modèle passé
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	// Construire la requête SQL
	var columns []string
	for i := 0; i < t.NumField(); i++ {
		columns = append(columns, t.Field(i).Name)
	}
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(columns, ", "), t.Name())

	// Exécuter la requête
	rows, err := o.db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Scanner les résultats
	for rows.Next() {
		// Créez une slice de interfaces pour scanner chaque colonne
		values := make([]interface{}, len(columns))
		for i := range values {
			values[i] = reflect.New(t.Field(i).Type).Interface()
		}

		if err := rows.Scan(values...); err != nil {
			return err
		}

		// Ici, vous pouvez traiter les valeurs scannées comme vous le souhaitez
		// Par exemple, afficher les résultats
		for _, value := range values {
			fmt.Print(reflect.ValueOf(value).Elem().Interface())
		}
	}

	return nil
}
