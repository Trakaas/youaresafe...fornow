package main

/*type blah interface {
  poke(person)
  bother() time
}*/
type ReadWriter interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
}

func main() {

}
