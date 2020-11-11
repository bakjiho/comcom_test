/**
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// [START container]
package main

import (
	"strconv"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// register hello function to handle all requests
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	// use PORT environment variable, or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// start the web server on port and accept requests
	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func hello(w http.ResponseWriter, r *http.Request) {
	// get input parameter
	keys, ok := r.URL.Query()["input"]
    
	// when input parameter is blank
   	 if !ok || len(keys[0]) < 1 {
        	fmt.Fprintf(w, "Please input \"input\" parameter with number.\n")
        	return
   	 }

    	key := keys[0]
	
	var input int
	var err error
	max := 1
	// input string to integer
	input, err = strconv.Atoi(string(key))
	// if string to int error occurs
	if err != nil {
		fmt.Fprintf(w, "Please input right number. 1~1000000\n")
		return
	}
	// if out of range
	if input > 1000000 || input < 1 {
		fmt.Fprintf(w, "Please check input range. 1~1000000\n")
		return
	}
	// get largest prime factor
	for i:= 2; i<=input; i++ {
		for input % i == 0 {
			input /= i
			if max < i {
				max = i
			}
			
		}
	}
	// print prime factor
	fmt.Fprintf(w, "%d\n", max)
}
// END
