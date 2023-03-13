package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'isValid' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func isValid(s string) string {
    // Write your code here
    charCounter := make(map[string]int)
    
    for _, c := range s{
        charCounter[string(c)] += 1
    }
    
    
    max := 0
    for _, v := range charCounter{
        if v > max{
            max = v
        }
    }
    
    allEqual := checkAllEqual(charCounter, max)
    
    if(allEqual){
        return "YES"
    }
    
    for k := range charCounter{
        
        charCounter[k] -= 1
        new_max := charCounter[k]
        if(charCounter[k] == 0){
            delete(charCounter, k)
            new_max = max
        }

        allEqual = checkAllEqual(charCounter, new_max)
        if allEqual{
            return "YES"
        }
        charCounter[k] += 1
        
    }
    
    return "NO"

}

func checkAllEqual(charCounter map[string]int, max int) bool{
    for _,v := range charCounter{
        if v != max{
            return  false
        }
    }
    return true
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := isValid(s)

    fmt.Fprintf(writer, "%s\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
