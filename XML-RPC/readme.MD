# Calculadora en Python Usando XML-RPC

Este documento explica cómo crear una calculadora simple usando el protocolo XML-RPC en Python. La calculadora puede realizar operaciones aritméticas básicas: suma, resta, multiplicación y división.

## Requisitos


## Guía SERVIDOR (server.py).

### 1. Importar la Biblioteca necesaria

Comienza importando la clase `SimpleXMLRPCServer` del módulo `xmlrpc.server`.

```python 
from xmlrpc.server import SimpleXMLRPCServer
``` 




## Guía CLIENTE (client.py)

### 1.  Importar la Biblioteca Necesaria

Primero, importa la biblioteca `xmlrpc.client`, que te permite comunicarte con un servidor XML-RPC.

```python
import xmlrpc.client
```

### 2.  Creamos la funcion main 

```python
def main():
    proxy = xmlrpc.client.ServerProxy("http://localhost:9000/")
```
En ella, creamos un objeto proxy que se conecta al servidor XML-RPC que se ejecuta en localhost en el puerto 9000.


### 3.  Definir la funcion de operaciones

```python
def switch_operation(proxy):
    flag = True
    while flag:
```
La función switch_operation toma el proxy como argumento.
Se inicializa una variable ```flag``` en ```True``` para controlar un bucle while, el cual sera el bucle principal del cliente. 
El bucle se ejecutara siempre que ```flag``` sea verdadero, es decir, tendremos que detener su ejecucion para terminar el proceso, y esto solo se lograra cambiando ```flag``` a ```False```.

### 4.  Manejo de entradas del usuario

```python
try: 
    number_a = float(input("Give the first number (a): "))
    number_b = float(input("Give the second number (b): "))
```
Usamos try para manejar excepciones. Se solicitan dos números al usuario y los parseamos a tipo float.
Recordemos que el input toma por defecto en tipo string, por ello es importante el parseo. 

### 5.  Mostrar opciones de operacion

```python
    print("1: Add   2: Subtract   3: Multiply   4: Division")
    print("Just select one. Write the number.")
    operation = int(input("Which operation do you want to do? "))
```
Se imprimen las opciones para que el usuario elija. La entrada del usuario se convierte a un número entero para determinar la operacion.

### 6.  Realizar la Operacion Seleccionada

```python
if operation == 1:
    result = proxy.add(number_a, number_b)
    print(f"Result of {number_a} + {number_b} is: {result}")
```
Dependiendo del valor de ```operation```, se llama al metodo correspondiente desde el proxy ```(add, rest, mul, div)```. Los metodos del proxy devuelven (como ya vimos en la configuracion del server), el resultado de una operacion determinada. 

El valor devuelto de dicha operacion, se almacena en la variable ```result```. Y es impresa en consola junto con los valores dados por el usuario, por supuesto, debidamente formateados.
Imitamos el comportamiento de un ```switch``` con una de cadena de ```elif```, con la misma declaracion para cada caso, solo haciendo una excepcion en el caso de una division, puesto que debemos validar la division entre 0. 

```python
elif operation == 4:
    if number_b == 0:
        print("Error: Division by zero is not allowed.")
    else:
        result = proxy.div(number_a, number_b)
        print(f"Result of {number_a} / {number_b} is: {result}")
```

### Continuacion o cierre del bucle principal 

Preguntamos al usuario si desea hacer otra operacion

```python
    print("Do you want to make another operation?")
```

Creamos otro bucle para controlar la salida del bucle principal.

```python
flag_keep = True
while flag_keep:
    keep = input("y/n? ").lower()  
    if keep == "y":
        flag = True
        flag_keep = False  
    elif keep == "n":
        flag = False
        flag_keep = False 
    else:
        print("You did not choose a valid option. This isn't the way.")
        flag_keep = True  
```
```flag_keep``` es una variable booleana utilizada para validar la entrada del teclado, para asegurar que el usuario solo introduzca 


