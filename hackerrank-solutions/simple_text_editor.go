package main
import (
    "fmt" 
)

type TextEditor struct {
    S string
    ops []string 
}

func main() {
    q := 0 
    fmt.Scanln(&q)
    editor := TextEditor{}
    for i :=0; i<q; i++{
        op := 0

        fmt.Scan(&op)
        
        switch op {
            case 1:
                editor.ops = append(editor.ops, editor.S)
                op2 := ""
                fmt.Scan(&op2)
                editor.S += op2
            case 2:
                editor.ops = append(editor.ops, editor.S)
                op2 := 0
                fmt.Scan(&op2)
                editor.S = editor.S[:len(editor.S)-op2]
            case 3:
                op2 := 0
                fmt.Scan(&op2)
                fmt.Println(string(editor.S[op2-1]))
            case 4:
                editor.S = editor.ops[len(editor.ops)-1]
                editor.ops = editor.ops[:len(editor.ops)-1]
        }
        //fmt.Println(editor.S)
    }
}
