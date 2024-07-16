# gitopsctl (Go)
Dies ist die *professionelle* Neuentwicklung des ursprünglichen Go-clients für die gitops-REST-API (siehe ../client/).

Das verwendete CLI-Framework heißt `cobra-cli` ist super einfach zu verwenden und dient als Basis für Tools wie `kubectl`!

## ohne Args
```
> go run main.go 
gitops-client written in Go and Cobra - the long version!

Usage:
  gitopsctl [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      create a new instance
  delete      delete a instance by id
  help        Help about any command
  list        list all instances
  read        read an instance by id
  update      delete a instance by id

Flags:
      --d           enable verbose output
  -h, --help        help for gitopsctl
      --su string   server URI (default "http://localhost:8000/instances")
  -t, --toggle      Help message for toggle

Use "gitopsctl [command] --help" for more information about a command.
```

## Instance anlegen
```
> go run main.go --d create --ba 1234 --in cobra-instance01 --oid dominik --rc 5 --sid 4321 --sv blubb --v "3.4.*"

Response Info:
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 495.095083ms
  Received At: 2024-06-24 15:58:57.303466 +0200 CEST m=+0.498291001
  Body       :
   {"instance_name":"cobra-instance01","orderer_id":"dominik","bits_account":1234,"service_id":4321,"replica_count":5,"version":"3.4.*","some_value":"blubb","order_time":"Mon Jun 24 15:58:56 2024","instance_id":"ce13df2a-1d5a-4694-b9ee-7fe2a68a9180","stage":"prod"}

New instance:  ce13df2a-1d5a-4694-b9ee-7fe2a68a9180
```

## Alle Instanzen anzeigen
```
> go run main.go list

Instances:
*  09c2b985-9691-4e07-a097-cae90f34ec14
*  9d9f4950-ab19-477a-866d-27be1dd43145
*  ab4280dd-2491-44a8-a648-86d917988587
*  c9621b1e-241e-4d1b-b69d-2160abb5ff52
*  ce13df2a-1d5a-4694-b9ee-7fe2a68a9180
```

## Instance anzeigen
```
> go run main.go --d read --iid 09c2b985-9691-4e07-a097-cae90f34ec14

Response Info:
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 113.263792ms
  Received At: 2024-06-24 16:01:03.548178 +0200 CEST m=+0.116372709
  Body       :
   {"instance_name":"cobra-instance01","orderer_id":"dominik","bits_account":1234,"service_id":4321,"replica_count":5,"version":"3.4.*","some_value":"blubb","order_time":"Mon Jun 24 15:58:25 2024","instance_id":"09c2b985-9691-4e07-a097-cae90f34ec14","stage":"prod"}

INSTANCE:  {09c2b985-9691-4e07-a097-cae90f34ec14 Mon Jun 24 15:58:25 2024 prod cobra-instance01 dominik 1234 4321 5 3.4.* blubb
```

## Instance updaten
```
> go run main.go --d update --iid 09c2b985-9691-4e07-a097-cae90f34ec14 --ba 1234 --in cobra-instance01 --rc 3 --sid 43214 --sv blubbx --v "3.5.*"

Response Info:
  Status Code: 200
  Status     : 200 OK
  Proto      : HTTP/1.1
  Time       : 1.8347015s
  Received At: 2024-06-24 16:02:48.87428 +0200 CEST m=+1.838269292
  Body       :
   {"instance_name":"cobra-instance01","orderer_id":"dominik","bits_account":0,"service_id":0,"replica_count":0,"version":"","some_value":"","order_time":"Mon Jun 24 15:58:25 2024","instance_id":"09c2b985-9691-4e07-a097-cae90f34ec14","stage":"prod"}

INSTANCE:  {09c2b985-9691-4e07-a097-cae90f34ec14 Mon Jun 24 15:58:25 2024 prod cobra-instance01 dominik 0 0 0  }
```

## Instance löschen
```
> go run main.go delete --iid 09c2b985-9691-4e07-a097-cae90f34ec14
```

## Kompilieren und in ${GOBIN} ablegen
```
> go install .
```
Wenn nicht anders definiert zeigt ${GOBIN} ins Leere. Der Default-Wert ist dann `~/go/bin` -> Dort wird das Kompilat unter dem Namen aus go.mod (`module gitopsctl` -> `gitopsctl`) abgelegt.# gitopsctl
