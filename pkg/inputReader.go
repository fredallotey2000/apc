package pkg

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

//input reader interface to be implemented
type InputReader interface {
	ReadInteger() int
	ReadSlice() []int
	Readinputs() InputsModel
}


type inputReader struct {
	reader *bufio.Reader
}

//Creates a new input reader
func NewReader() InputReader {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	return &inputReader{reader}
}

//Reads an integer from the console
func (m *inputReader) ReadInteger() int {
	nTemp, err := strconv.ParseInt(strings.TrimSpace(m.readLine(m.reader)), 10, 64)
	m.checkError(err)
	return int(nTemp)
}

//Reads a slice from the console
func (m *inputReader) ReadSlice() []int {
	sTemp := strings.Split(strings.TrimSpace(m.readLine(m.reader)), " ")
	var s []int
	for _, v := range sTemp {
		sItemTemp, err := strconv.ParseInt(v, 10, 64)
		m.checkError(err)
		sItem := int(sItemTemp)
		s = append(s, sItem)
	}
	return s

}

//Checks if an error occured during reading
func (m *inputReader) checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//Reads a line from the console
func (m *inputReader) readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
//Reads all inputs and puts them into a struct for processing
func (m *inputReader) Readinputs() InputsModel {

	inputs := InputsModel{}
	//read grid width
	gridSize := m.ReadSlice()
	inputs.GridSize = gridSize

	startFinish := m.ReadSlice()
	inputs.StartFinish = startFinish

	obstacles := m.ReadInteger()
	inputs.Obstacles = obstacles

	var newObstacles [][]int
	for i := 0; i < obstacles; i++ {
		newObstacles = append(newObstacles, m.ReadSlice())
	}
	inputs.ObstacleSlice = newObstacles
	return inputs
}
