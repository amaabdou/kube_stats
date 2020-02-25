package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func WriteToHtmlServerResponse(data PodsData) error {
	jsonData,err := json.Marshal(data.Data)
	if err != nil {
		log.Fatal(err)
	}
	htmlPage := fmt.Sprint(`<html>
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
	return ioutil.WriteFile("./output.html", []byte(htmlPage), 0644)
}

func printHeaders(data PodsData) string {
	headerString := "["
	for _,name := range data.Headers {
		headerString += "{ title: \""+ name + "\" },"
	}
	headerString += "]"
	return  headerString
}
