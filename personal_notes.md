# implement:

-> cmd/starnamed
-> x/starname
    -> icb3 compatible?
    -> verify all dependencies? 

-> app/app.go
    -> split file of wasmapp from the other parameters
    -> add starname modules
        -> modules manager
        -> keeper
        -> upgrade register
        -> add routes
        -> setBegginerOreder and setEndOrder 
        

-> app/upgrade.go
    -> verify if the all modules was inside of upgrade handler version map. If not add.