# package under development
# nifi-api-client-go
The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and stop processors, 
monitor queues, query provenance data, and more. Each endpoint below includes a description, 
definitions of the expected input and output, potential response codes, and the authorizations required to invoke each service.

## Getting Started

```go
    package main
    
    import (
    	"fmt"
    	"github.com/mirkos-vf/go_nifi_api"
    )
    
    func main() {
    	nifi := go_nifi_api.NewNiFi("https://localhost:8080")
    	
    	token, _ := nifi.AccessToken(os.Getenv("NIFI_USERNAME"), os.Getenv("NIFI_PASSWORD"))
    	
    	if err != nil {
    		panic(err)
    	}
    
    	fmt.Println(token)
    
    }
```