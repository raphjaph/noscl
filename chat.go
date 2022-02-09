package main

import (
  "fmt"

	"github.com/docopt/docopt-go"
	"github.com/fiatjaf/go-nostr"
	"github.com/fiatjaf/go-nostr/nip04"
)

func chat(opts docopt.Opts) {
    
	initNostr()
  
  me := getPubKey(config.PrivateKey)
  chatParticipant := opts["<id>"].(string)
  fmt.Println("Starting chat with ", chatParticipant)

	sub := pool.Sub(nostr.EventFilters{
    {
      Kinds: []int{4},
      Authors: []string{opts["<id>"].(string)},
      TagP: []string{me},
    },
    { 
      Kinds: []int{4},
      Authors: []string{me},
      TagP: []string{chatParticipant},
    },
  })
  
  key, err := nip04.ComputeSharedSecret(config.PrivateKey, chatParticipant)
  if err != nil {
    fmt.Println(err.Error())
  }

	for event := range sub.UniqueEvents {
    message, err := nip04.Decrypt(event.Content, key)
    if err != nil {
      fmt.Println(err.Error())
      continue
    }

    if message == "" {
      continue
    }

    if event.PubKey == me {
      fmt.Println("\t\t\t\t\t\t", message)
    } else {
      fmt.Println(message)
    }
  }
}

