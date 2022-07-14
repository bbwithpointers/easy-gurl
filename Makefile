run:
	windows:
		go run .\cmd\main.go -method POST -url https://www.google.com -data "{"your": "json"}" -persist YES -header X-adzerk-ApiKey:adasdasda
	linux:
		go run cmd\main.go -method POST -url https://www.google.com -data "{"your": "json"}" -persist YES -header X-adzerk-ApiKey:adasdasda