package main

import (
  "net/http"
  "sync"
  "net/url"
  "fmt"
)


var keyValueStore map[string]string
var kVStoreMutex sync.RWMutex

func main(){
  keyValueStore = make(map[string]string)
  kVStoreMutex = sync.RWMutex{}
  http.HandleFunc("/get", get)
  http.HandleFunc("/set", set)
  http.HandleFunc("/remove", remove)
  http.HandleFunc("/list", list)
  http.HandleFunc("/about/", about)
  http.ListenAndServe(":3000", nil);
}

func get(w http.ResponseWriter, r *http.Request){
  if(r.Method == http.MethodGet){
    values, err := url.ParseQuery(r.URL.RawQuery)
    if err != nil{
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprint(w, "Error:", err)
      return
    }
    if len(values.Get("key")) == 0 {
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprint(w, "Error:", err)
      return
    }
    kVStoreMutex.RLock()
    value := keyValueStore[string(values.Get("key"))]
    kVStoreMutex.RUnlock()

    fmt.Fprint(w, value)
  }else {
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(w, "Error: Only GET accepted.")
  }
}

func set(w http.ResponseWriter, r *http.Request){
  if(r.method == http.MethodPost){
    values, err := url.ParseQuery(r.URL.RawQuery)
    if err != nil{
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprint(w, "Error:", err)
      return
    }
    if len(values.Get("key")) == 0 {
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprint(w, "Error:", "wrong input key.")
      return
    }
    if len(values.Get("value")) == 0 {
      w.WriteHeader(http.StatusBadRequest)
      fmt.Fprint(w, "Error:", "Wrong input value.")
      returns
    }
    kvStoreMutex.Lock()
    keyValueStore[string.Get("key")] =
    string(values.Get("value"))
    kVStoreMutex.Unlock()

    fmt.Fprint(w, "success")
  }else{
    w.WriteHeader(http.StatusBadRequest)
    fmt.Fprint(w, "Error: Only POST accepted.")
  }
}

func remove(w http.ResponseWriter, r *http.Request){

}

func list(w http.ResponseWriter, r *http.Request){

}
