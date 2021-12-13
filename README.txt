Integrantes:
Ignacio Lizana 201673542-9
Lukas Gutierrez 201673059-1
Nicolás Miranda 201673583-6

Instrucciones:

Para cada máquina, ejecutar el comando "cd tarea3-distribuidos" para ingresar a la carpeta con los archivos.
1) Las ubicaciones de las entidades son las siguientes:
-Leia se ejecuta en la maquina 62
-Los informantes en las máquinas 63 y 64.
-servidores fulcrum en las maquinas 62,63,64
-Broker en la maquina 61.

2) El orden de ejecucion y los comandos son los siguientes:

1- Primero ejecutar "make runfulcrum" en las maquinas 62,63,64.
2- Ejecutar "make runbroker" en la máquina 61.
3- Ejecutar "make runleia" en la maquina 62.
4- Ejecutar "make runinformante" en las máquinas 63,64.
Todas estas entidades deben ser ejecutadas en ventanas distintas y los comandos a usar son exactamente
los mismos pedidos en el enunciado. Ej: "AddCity planeta ciudad 30"


