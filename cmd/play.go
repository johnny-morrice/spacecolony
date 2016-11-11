// Copyright © 2016 John Morrice <john@functorama.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"errors"
        "fmt"
        "os"

	"github.com/spf13/cobra"

        "github.com/johnny-morrice/spacecolony/lib-colony"
)

// playCmd represents the play command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: "Play Space Colony",
	Long: `Space Colony is a gentle roguelike about settling on a distant planet.`,
	Run: func(cmd *cobra.Command, args []string) {
                opts, err := getgameoptions(cmd)

                if err != nil {
                        fmt.Fprintf(os.Stderr, "Error in command line: %v", err)

                        os.Exit(1)
                }

		colony.Play(*opts)
	},
}

func getgameoptions(cmd *cobra.Command) (*colony.GameOptions, error) {
        persistent := cmd.PersistentFlags()

        height, err := persistent.GetUint("height")

        if err != nil {
                return nil, err
        }

        width, err := persistent.GetUint("width")

        if err != nil {
                return nil, err
        }

        samples, err := persistent.GetUint("samples")

        if err != nil {
                return nil, err
        }

        fps, err := persistent.GetUint("fps")

        if err != nil {
                return nil, err
        }

        fullscreen, err := persistent.GetBool("fullscreen")

        if err != nil {
                return nil, err
        }

        vsync, err := persistent.GetBool("vsync")

        if err != nil {
                return nil, err
        }

        opts := &colony.GameOptions{}
        opts.Width = width
        opts.Height = height
        opts.Samples = samples
        opts.FPS = fps
        opts.Fullscreen = fullscreen
        opts.Vsync = vsync

	if opts.Width < 500 || opts.Height < 500 {
		return nil, errors.New("Window too small (Must be >500 pixels)")
	}

        return opts, nil
}

func init() {
	RootCmd.AddCommand(playCmd)

        persistent := playCmd.PersistentFlags()

        persistent.Uint("width", 1000, "Window width")
        persistent.Uint("height", 1000, "Window height")
        persistent.Uint("samples", 1, "Multisample count")
        persistent.Uint("fps", 60, "Maximum frames-per-second")

        persistent.Bool("fullscreen", false, "Full screen on desktop machines")
        persistent.Bool("vsync", true, "Enable vertical sync")
}
