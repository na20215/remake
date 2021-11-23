package gol

type distributorChannels struct {
	events     chan<- Event
	ioCommand  chan<- ioCommand
	ioIdle     <-chan bool
	ioFilename chan<- string
	ioOutput   chan<- uint8
	ioInput    <-chan uint8
}

func makenewworld(height, width int) [][]uint8 {
	sides := make([][]uint8, width)
	for i := range sides {
		sides[i] = make([]uint8, height)
	}
	return sides
}

// distributor divides the work between workers and interacts with other goroutines.
func distributor(p Params, c distributorChannels) {

	// TODO: Create a 2D slice to store the world.
	world := makenewworld(p.ImageHeight, p.ImageWidth)
	for h := 0; h < p.ImageHeight; h++ {
		for w := 0; w < p.ImageWidth; w++ {
			world[h][w] = <-c.ioInput
		}
	}
}

////////////////////turn := 0

// TODO: Execute all turns of the Game of Life.

// TODO: Report the final state using FinalTurnCompleteEvent.

// Make sure that the Io has finished any output before exiting.
//////////////c.ioCommand <- ioCheckIdle
///////////////<-c.ioIdle

///////////////c.events <- StateChange{turn, Quitting}

// Close the channel to stop the SDL goroutine gracefully. Removing may cause deadlock.
///////////////close(c.events) //////}
