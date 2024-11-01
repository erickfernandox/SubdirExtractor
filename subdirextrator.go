package main

import (
    "bufio"
    "flag"
    "fmt"
    "net/url"
    "os"
    "sort"
)

func extractUniqueSuburls(filename string) ([]string, error) {
    // Usar um map para evitar duplicados
    uniqueSuburls := make(map[string]bool)

    // Abrir o arquivo
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Ler o arquivo linha por linha
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        // Extrair o caminho da URL
        parsedUrl, err := url.Parse(line)
        if err != nil {
            continue // Ignora URLs malformadas
        }
        path := parsedUrl.Path
        uniqueSuburls[path] = true
    }

    if err := scanner.Err(); err != nil {
        return nil, err
    }

    // Converter o mapa para uma lista
    var suburls []string
    for path := range uniqueSuburls {
        // Ignorar caminhos vazios ou apenas "/"
        if path != "" && path != "/" {
            suburls = append(suburls, path)
        }
    }

    // Ordenar a lista
    sort.Strings(suburls)
    return suburls, nil
}

func main() {
    // Configurar argumentos
    listFile := flag.String("l", "", "Arquivo txt com lista de URLs")
    flag.Parse()

    if *listFile == "" {
        fmt.Println("Por favor, forneça o caminho para o arquivo de URLs usando -l")
        os.Exit(1)
    }

    // Executar a função principal
    suburls, err := extractUniqueSuburls(*listFile)
    if err != nil {
        fmt.Println("Erro:", err)
        os.Exit(1)
    }

    // Imprimir os caminhos únicos
    for _, suburl := range suburls {
        fmt.Println(suburl)
    }
}
