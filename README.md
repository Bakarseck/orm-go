# ORM Package for Go

## Overview

This ORM (Object-Relational Mapping) package provides a convenient and efficient way to interact with databases in Go. It abstracts the database interactions, allowing developers to work with Go structs instead of direct SQL queries.

## Features

- **Auto Migration**: Easily create and update database tables based on Go struct definitions.
- **CRUD Operations**: Simplified Create, Read, Update, and Delete operations using Go structs.
- **Query Scanning**: Directly scan SQL query results into Go structs for easy data manipulation.
- **Tag Parsing**: Use custom `orm-go` tags in structs for additional field specifications like foreign keys and SQL attributes.

## Installation

To install this package, use the following `go get` command:

```bash
    go get github.com/Bakarseck/orm-go
```

## Usage

To get started, import the package and initialize it with your database connection.

```go
    package main

    import (
        "fmt"

        orm "github.com/Bakarseck/orm-go/ORM"
        _ "github.com/mattn/go-sqlite3"
    )

    type User struct {
        orm.Model
        Username string `orm-go:"NOT NULL UNIQUE"`
        Email    string `orm-go:"NOT NULL UNIQUE"`
    }

    type Produit struct {
        orm.Model
        Name_produit string `orm-go:"NOT NULL"`
        Prix         int64
        User_id      int64 `orm-go:"NOT NULL FOREIGN_KEY:User:Id"`
    }

    func NewUser(name, email string) User {
        return User{
            Username: name,
            Email:    email,
        }
    }

    func NewProduit(name string, p int64) Produit {
        return Produit{
            Name_produit: name,
            Prix:         p,
        }
    }

    func main() {
        user := User{}
        produit := Produit{}

        orm := orm.NewORM()
        orm.InitDB("test.db")
        orm.AutoMigrate(user, produit)

        var produits []interface{}
        for i := 1; i <= 100; i++ {
            name := fmt.Sprintf("Produit%d", i)
            produit := NewProduit(name, int64(i*5))
            produits = append(produits, produit)
        }

        orm.Insert(produits...)

        orm.SetModel("Email", "abdou@gmail.com", "User").UpdateField("moussa@gmail.com").Update(orm.Db)

        orm.Delete(User{}, "Id", 2)
        u := NewUser("Mouhamed Sylla", "syllamouhamed99@gmail.com")
        u1 := NewUser("Abdou", "abdou@gmail.com")
        u2 := NewUser("Sidi", "sidi@gmail.com")

        orm.Insert(u, u1, u2)

        users := orm.Scan(User{}, "Id", "CreatedAt", "Email", "Username").([]User)

        for _, user := range users {
            fmt.Println("User: ", user)
        }

    }
```

## Contributing
Contributions to this project are welcome! If you're interested in contributing, please follow these steps:

1. **Fork the Repository**: Create your own fork of the repository by clicking the 'Fork' button on the GitHub page.

2. **Clone the Fork**: Clone your fork to your local machine.

    ```bash
        git clone https://github.com/Bakarseck/orm-go.git
    ```

3. **Create a New Branch**: Create a new branch for your feature or bug fix.

    ```bash
        git checkout -b feature/your-new-feature
    ```

4. **Make Your Changes**: Implement your feature or bug fix.

5. **Run the Tests**: Ensure all the tests pass with your changes.

6. **Commit Your Changes**: Commit your changes with a clear commit message.

    ```bash
    git commit -am 'Add some feature'
    ```

7. **Push to the Branch**: Push your changes to your fork.

    ```bash
    git push origin feature/your-new-feature
    ```

8. **Submit a Pull Request**: Go to your fork on GitHub and submit a pull request.

Please make sure to update tests as appropriate and follow the code style of the project. We also recommend including documentation with your changes.


## Authors
* [Bakarseck](https://github.com/Bakarseck)
* [Mouhamadou Sylla](https://github.com/mouhamsylla)