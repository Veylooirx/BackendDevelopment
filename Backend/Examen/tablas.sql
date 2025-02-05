CREATE TABLE clientes (
    cliente_id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    fecha_inicio DATE NOT NULL,
    fecha_fin DATE NOT NULL,
    monto_pago NUMERIC(10, 2) NOT NULL
);

CREATE TABLE pagos (
    pago_id SERIAL PRIMARY KEY,
    cliente_id INT REFERENCES clientes(cliente_id),
    fecha_cobro DATE NOT NULL,
    pagado BOOLEAN DEFAULT FALSE
);

