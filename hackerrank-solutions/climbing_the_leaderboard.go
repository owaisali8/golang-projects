package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player
 */
func climbingLeaderboard(ranked []int32, player []int32) []int32 {
    
    playerRanks := make([]int32, len(player))
    keys        := make([]int32,0,len(ranked))
    rankMap     := make(map[int32]int32)
    
    var r int32 = 1
    for _,v := range ranked {
        _, exists := rankMap[v]
        if !exists {
            rankMap[v] = r
            keys = append(keys,v)
            r++
        }
    }
    
    
    for i,v := range player{
        
        r := findRank(&rankMap,&keys, v)
        playerRanks[i] = r
    }
    
    return playerRanks
    

}
    func findRank(rankMapPtr *map[int32]int32,keysPtr *[]int32,v int32) int32 {
        keys := *keysPtr
        rankMap := *rankMapPtr
        end := len(keys) - 1
        start := 0 
        middle := (end+start)/2
        var key int32
        
        if v > keys[start] {
            return rankMap[keys[start]] 
        } else if v < keys[end] {
            return rankMap[keys[end]] +1
        }
        
        if _, e:= rankMap[v]; e {
            return rankMap[v]
        }
        
        for start <= end{
            key = keys[middle]
            if v < key{
                start = middle +1
            } else if v > key {
                end = middle -1 
            }
            middle =(end+start)/2
        }
        
        if  v < key{ 
            return rankMap[key] +1
        } else {
            return rankMap[key]
        }
    }
    
func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    rankedCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    rankedTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var ranked []int32

    for i := 0; i < int(rankedCount); i++ {
        rankedItemTemp, err := strconv.ParseInt(rankedTemp[i], 10, 64)
        checkError(err)
        rankedItem := int32(rankedItemTemp)
        ranked = append(ranked, rankedItem)
    }

    playerCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    playerTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

    var player []int32

    for i := 0; i < int(playerCount); i++ {
        playerItemTemp, err := strconv.ParseInt(playerTemp[i], 10, 64)
        checkError(err)
        playerItem := int32(playerItemTemp)
        player = append(player, playerItem)
    }

    result := climbingLeaderboard(ranked, player)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, "\n")
        }
    }

    fmt.Fprintf(writer, "\n")

    defer writer.Flush()
    
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
