package gadget

import "fmt"

type TapePlayer struct{
	Batteries string
}

func (t TapePlayer) Play(song string){
	fmt.Println("playing", song)
}

func (t TapePlayer) Stop(){
	fmt.Println("stopped!")
}

type TapeRecorder struct{
	Microphones int
}
func (t TapeRecorder) Play(song string){
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Stop(){
	fmt.Println("Stopped!")
}
