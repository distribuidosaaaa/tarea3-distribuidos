# tarea3-distribuidos
tarea 3 de distribuidos


# Logica


2 informant machines.
Connected to fulcrum server, fulcrum server works with 3 replicas with eventual consistency,
consistency between replicas is applied every two minutes

broker redirects to random fulcrum replica to read/write, if there is a consistency conflict redirects to specific replica

informants cand do 4 commands:

* addcity
* updatename
* updatenumber
* deletecity

commands are sent to broker

informants work as read your writes.
Informants must mantain on cache the register that have modified


leia makes request to broker with commands:

* getnumberrebelds


leia works with monotonic reads
