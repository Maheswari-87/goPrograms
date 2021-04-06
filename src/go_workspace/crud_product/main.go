package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
}

var tpl *template.Template

var db *sql.DB

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	var err error
	db, err := sql.Open("mysql", "root:Mahik87@@tcp(127.0.0.1:3306)/godb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/browse", browsetHandler)
	http.HandleFunc("/update/", updateHandler)
	http.HandleFunc("/updateresult/", updateResultHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe("localhost:8080", nil)
}
func browsetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*********browseHandler running*******")
	stmt := "SELECT * FROM products"
	rows, err := db.Query(stmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
		if err != nil {
			panic(err)
		}
		products = append(products, p)
	}
	tpl.ExecuteTemplate(w, "select.html", products)
}
func insertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("******inserHandler running*****")
	if r.Method == "Get" {
		tpl.ExecuteTemplate(w, "insert.html", nil)
		return
	}
	r.ParseForm()
	name := r.FormValue("nameName")
	price := r.FormValue("priceName")
	descr := r.FormValue("descrName")

	var err error
	if name == "" || price == "" || descr == "" {
		fmt.Println("Error inserting row:", err)
		tpl.ExecuteTemplate(w, "insert.html", "Error inserting data, please check all fields.")
		return
	}
	var ins *sql.Stmt

	ins, err = db.Prepare("INSERT INTO godb.products (name,price,description) VALUES(?,?,?);")
	if err != nil {
		panic(err)
	}
	defer ins.Close()
	res, err := ins.Exec(name, price, descr)

	rowsAffec, _ := res.RowsAffected()
	if err != nil || rowsAffec != 1 {
		fmt.Println("Error inserting row:", err)
		tpl.ExecuteTemplate(w, "insert.html", "Error inserting data,please check all fields.")
		return
	}
	lastInserted, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()
	fmt.Println("ID of last row inserted:", lastInserted)
	fmt.Println("number of rows affected:", rowsAffected)
	tpl.ExecuteTemplate(w, "insert.html", "Product succesfully inserted")
}
func updateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****update Handler running*****")
	r.ParseForm()
	id := r.FormValue("id")
	row := db.QueryRow("SELECT * FROM godb.products WHERE id=?;", id)
	var p Product
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/browse", http.StatusTemporaryRedirect)
		return
	}
	tpl.ExecuteTemplate(w, "update.html", p)
}
func updateResultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("****update resulthandler running******")
	r.ParseForm()
	id := r.FormValue("id")
	name := r.FormValue("nameName")
	price := r.FormValue("priceName")
	description := r.FormValue("description")
	upStmt := "UPDATE godb.products SET 'name'=?, 'price'=?,'description'=? WHERE ('id'=?)"
	stmt, err := db.Prepare(upStmt)
	if err != nil {
		fmt.Println("error preparing statement")
		panic(err)
	}
	fmt.Println("db.Preppare err:", err)
	fmt.Println("db.Prepare stmt:", stmt)
	defer stmt.Close()
	var res sql.Result
	res, err = stmt.Exec(name, price, description, id)
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		fmt.Println(err)
		tpl.ExecuteTemplate(w, "result.html", "There was a problem updating product")
		return
	}
	tpl.ExecuteTemplate(w, "result.html", "Product was succesfully updated")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****delete Handler running*****")
	r.ParseForm()
	id := r.FormValue("id")
	del, err := db.Prepare("DELETE FROM godb.products WHERE ('id'=?);")
	if err != nil {
		panic(err)
	}
	defer del.Close()
	var res sql.Result
	res, err = del.Exec(id)
	rowsAff, _ := res.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	if err != nil || rowsAff != 1 {
		fmt.Fprint(w, "Error deleting product")
		return
	}

	/*if err!=nil {
		fmt.Fprint(w, "Error deleting product")
		return
	}*/
	fmt.Println("err:", err)
	tpl.ExecuteTemplate(w, "result.html", "Product was successfully deleted")
}
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/browse", http.StatusTemporaryRedirect)
}
