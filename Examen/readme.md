Explicacion del Código

1.-     Definición de Estructura Cliente:

 ```go
type Cliente struct {
    Nombre      string
    FechaInicio time.Time
    FechaFin    time.Time
    MontoPago   float64
}
 ```
Esta estructura define al cliente, con campos para su nombre, fecha de inicio y fin del periodo de cobros, y el monto de cada pago.

2.      Función createCliente
 ```go
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
 ```

Esta función inserta los datos del cliente en la tabla clientes y devuelve su ID. Usa una consulta SQL INSERT INTO y devuelve el cliente_id generado para identificar al cliente en la base de datos.

3.      Función generarPagos
 ```go
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
 ```

Esta función genera las fechas de cobro para un cliente. Inicia desde fechaInicio y va añadiendo nuevas fechas de cobro en la tabla pagos, hasta alcanzar fechaFin. La línea fechaCobro = fechaCobro.AddDate(0, 0, 7) agrega 7 días (una semana) a fechaCobro, por lo que actualmente se configura para pagos semanales.

Cambios para otras Frecuencias:
Para ajustar la frecuencia de los pagos, se debe cambiar la última línea dentro del bucle for:

Para pagos mensuales: Cambia fechaCobro.AddDate(0, 0, 7) a fechaCobro.AddDate(0, 1, 0), lo cual agrega un mes a cada fechaCobro.
Para pagos anuales: Cambia fechaCobro.AddDate(0, 0, 7) a fechaCobro.AddDate(1, 0, 0), lo cual agrega un año a cada fechaCobro.

4.      Función aplicarPago

 ```go
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
 ```
Esta función marca un pago como “pagado” en la tabla pagos. Recibe el clienteID y una fecha de cobro (fechaCobro), actualizando el campo pagado a TRUE para la fecha específica.

5.      Función reportePagos

 ```go
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
 ```
Esta función genera el reporte de pagos de un cliente. Consulta la tabla pagos para obtener todas las fechas de cobro y su estado de pago (pagado). Luego imprime el estado de cada fecha, mostrando si está “pendiente” o “pagado”.

