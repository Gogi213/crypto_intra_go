// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package  main

import (
  "github.com/mailru/easyjson/jwriter"
  "github.com/mailru/easyjson/jlexer"
)

func ( Message ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* Message ) UnmarshalJSON([]byte) error { return nil }
func ( Message ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* Message ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_Message *Message
