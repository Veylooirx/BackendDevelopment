package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Cliente struct {
	Nombre      string
	FechaInicio time.Time
	FechaFin    time.Time
	MontoPago   float64
}

func createCliente(db *sql.DB, cliente Cliente) (int, error) {
	var clienteID int
	err := db.QueryRow(`
        INSERT INTO clientes (nombre, fecha_inicio, fecha_fin, monto_pago) 
        VALUES ($1, $2, $3, $4) RETURNING cliente_id
    `, cliente.Nombre, cliente.FechaInicio, cliente.FechaFin, cliente.MontoPago).Scan(&clienteID)
	if err != nil {
		return 0, fmt.Errorf("createCliente: %v", err)
	}
	return clienteID, nil
}

func generarPagos(db *sql.DB, clienteID int, fechaInicio, fechaFin time.Time) error {
	fechaCobro := fechaInicio
	for fechaCobro.Before(fechaFin) || fechaCobro.Equal(fechaFin) {
		_, err := db.Exec(`
            INSERT INTO pagos (cliente_id, fecha_cobro) 
            VALUES ($1, $2)
        `, clienteID, fechaCobro)
		if err != nil {
			return fmt.Errorf("generarPagos: %v", err)
		}
		fechaCobro = fechaCobro.AddDate(0, 0, 7) // Incremento semanal
	}
	return nil
}

func aplicarPago(db *sql.DB, clienteID int, fechaCobro time.Time) error {
	_, err := db.Exec(`
        UPDATE pagos SET pagado = TRUE 
        WHERE cliente_id = $1 AND fecha_cobro = $2
    `, clienteID, fechaCobro)
	if err != nil {
		return fmt.Errorf("aplicarPago: %v", err)
	}
	return nil
}

func reportePagos(db *sql.DB, clienteID int) error {
	rows, err := db.Query(`
        SELECT fecha_cobro, pagado 
        FROM pagos 
        WHERE cliente_id = $1
        ORDER BY fecha_cobro
    `, clienteID)
	if err != nil {
		return fmt.Errorf("reportePagos: %v", err)
	}
	defer rows.Close()

	fmt.Printf("Reporte de pagos del cliente %d:\n", clienteID)
	for rows.Next() {
		var fechaCobro time.Time
		var pagado bool
		if err := rows.Scan(&fechaCobro, &pagado); err != nil {
			return fmt.Errorf("reportePagos (scan): %v", err)
		}
		estado := "pendiente"
		if pagado {
			estado = "pagado"
		}
		fmt.Printf("%s - %s\n", fechaCobro.Format("02 January 2006"), estado)
	}
	return nil
}

func main() {
	connStr := "host=localhost port=5434 user=postgres password=miercoles133 dbname=cobatron sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al abrir la conexión:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	cliente := Cliente{
		Nombre:      "pako",
		FechaInicio: time.Date(2023, time.January, 12, 0, 0, 0, 0, time.UTC),
		FechaFin:    time.Date(2024, time.January, 12, 0, 0, 0, 0, time.UTC),
		MontoPago:   100.0,
	}

	clienteID, err := createCliente(db, cliente)
	if err != nil {
		log.Fatal(err)
	}

	err = generarPagos(db, clienteID, cliente.FechaInicio, cliente.FechaFin)
	if err != nil {
		log.Fatal(err)
	}

	err = aplicarPago(db, clienteID, time.Date(2023, time.January, 12, 0, 0, 0, 0, time.UTC))
	if err != nil {
		log.Fatal(err)
	}

	err = reportePagos(db, clienteID)
	if err != nil {
		log.Fatal(err)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
