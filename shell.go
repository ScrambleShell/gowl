
package gowl

import (
	"bytes"
)

type Shell struct {
	*WlObject
	events []func (s *Shell, msg []byte)
}

//// Requests
func (s *Shell) Get_shell_surface (id *Shell_surface, surface *Surface ) {
	buf := new(bytes.Buffer)
	appendObject(id)
	writeInteger(buf, id.Id())
	writeInteger(buf, surface.Id())

	sendmsg(s, 0, buf.Bytes())
}

//// Events
func (s *Shell) HandleEvent(opcode int16, msg []byte) {
	if s.events[opcode] != nil {
		s.events[opcode](s, msg)
	}
}

func NewShell() (s *Shell) {
	s = new(Shell)

	return
}