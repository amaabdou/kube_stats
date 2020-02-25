package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func WriteToHttpServerResponse(data PodsData) error {
	jsonData,err := json.Marshal(data.Data)
	if err != nil {
		log.Fatal(err)
	}


	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handling incoming request ")
		_,err := fmt.Fprint(w, `<html>
			<head>
				<link rel="stylesheet" type="text/css" href="https://cdn.datatables.net/1.10.20/css/jquery.dataTables.min.css">
				<script src="https://code.jquery.com/jquery-3.3.1.js"></script>
				<script src="https://cdn.datatables.net/1.10.20/js/jquery.dataTables.min.js"></script>
				<script>
					var dataSet = `,string(jsonData) ,`;
					 
					$(document).ready(function() {
						$('#example').DataTable( {
							data: dataSet,
							columns: `,printHeaders(data) ,`
						} );
					} );
				</script>
			`+
			"" +
			""+`
			</head>
			<body>
					<table id="example" class="order-column dataTable" style="width:100%!" (missing)=""></table>
			</body>
		</html>`)
		if err != nil {
			log.Println("Could not handle request ", err)
		}
	})

	log.Println("Listening on 8080 ...")
	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}

func printHeaders(data PodsData) string {
	headerString := "["
	for _,name := range data.Headers {
		headerString += "{ title: \""+ name + "\" },"
	}
	headerString += "]"
	return  headerString
}
