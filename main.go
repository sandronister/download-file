package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

func downloadFile(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	files := map[string]string{
		"pdf/file1.pdf": "https://www.ufms.br/wp-content/uploads/2017/09/PDF-teste.pdf",
		"pdf/file2.pdf": "https://mescti.gov.ao/fotos/frontend_22/gov_documentos/input_file2_18225733755fb3f5c257344.pdf",
		"pdf/file3.pdf": "https://araucariageneticabovina.com.br/arquivos/servico/pdfServico_57952bf8ca7af_24-07-2016_17-58-32.pdf",
		"pdf/file4.pdf": "https://jucisrs.rs.gov.br/upload/arquivos/201710/30150625-criacao-de-pdf-a.pdf",
	}
	wg := sync.WaitGroup{}

	wg.Add(len(files))

	for file, url := range files {

		go func(file string, url string) {
			defer wg.Done()
			err := downloadFile(url, file)
			if err != nil {
				fmt.Println(err)
			}

		}(file, url)

	}

	wg.Wait()

}
