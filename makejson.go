// Write a program which prompts the user to first enter a name, and then enter an address. Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively. Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.

package main

import (
  "bufio"
  "os"
  "fmt"
  "strings"
  "encoding/json"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  
  fmt.Print("Name: ")
  name, _ := reader.ReadString('\n')
  name = strings.TrimRight(name, "\r\n")
  
  fmt.Print("Address: ")
  address, _ := reader.ReadString('\n')
  address = strings.TrimRight(address, "\r\n")
  
  m := map[string]string{"name": name, "address": address}
  
  b, _ := json.Marshal(m)
  fmt.Println(string(b))
}

