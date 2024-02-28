package conveyor

import (
	"context"
	"io"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	c := New()
	assert.IsType(t, &conveyor{}, c)
	assert.NotNil(t, c.inCh)
	assert.NotNil(t, c.sqCh)
	assert.NotNil(t, c.dbCh)
}

func Test_conveyor_SquareNum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Valid numbers",
			input: []int{2, 3, 5, 7, 10},
			want:  []int{4, 9, 25, 49, 100},
		},
		{
			name:  "Empty input",
			input: []int{},
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			c.wg.Add(1)
			go c.SquareNum()
			go func() {
				for _, v := range tt.input {
					c.inCh <- v
				}
				close(c.inCh)
			}()

			res := make([]int, 0, len(tt.input))
			for v := range c.sqCh {
				res = append(res, v)
			}
			assert.Equal(t, tt.want, res)
		})
	}
}

func Test_conveyor_DoubleNum(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{
			name:  "Valid numbers",
			input: []int{2, 3, 5, 7, 10},
			want:  []int{4, 6, 10, 14, 20},
		},
		{
			name:  "Empty input",
			input: []int{},
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			c.wg.Add(1)
			go c.DoubleNum()
			go func() {
				for _, v := range tt.input {
					c.sqCh <- v
				}
				close(c.sqCh)
			}()

			res := make([]int, 0, len(tt.input))
			for v := range c.dbCh {
				res = append(res, v)
			}
			assert.Equal(t, tt.want, res)
		})
	}
}

func Test_conveyor_Write(t *testing.T) {
	old := os.Stdout // keep backup of the real stdout
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	os.Stdout = w

	defer func() {
		os.Stdout = old // restoring the real stdout

	}()

	tests := []struct {
		name  string
		input []int
		want  string
	}{
		{
			name:  "Valid numbers",
			input: []int{2, 3, 5, 7, 10},
			want:  "Result is: 2\nResult is: 3\nResult is: 5\nResult is: 7\nResult is: 10\n",
		},
		{
			name:  "Empty input",
			input: []int{},
			want:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := New()
			c.wg.Add(1)
			go c.Write()

			go func() {
				for _, v := range tt.input {
					c.dbCh <- v
				}
				close(c.dbCh)
			}()

			c.wg.Wait()
			w.Close()

			out, err := io.ReadAll(r)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(out))
		})
	}
}

func Test_conveyor_Read(t *testing.T) {
	old := os.Stdin // keep backup of the real stdin
	r, w, err := os.Pipe()
	assert.NoError(t, err)

	os.Stdin = r

	defer func() {
		os.Stdin = old // restoring the real stdin
	}()

	tests := []struct {
		name  string
		input []string
		want  []int
	}{
		{
			name:  "Valid numbers",
			input: []string{"2\n", "3\n", "5\n", "7\n", "10\n", "stop\n"},
			want:  []int{2, 3, 5, 7, 10},
		},
		{
			name:  "Ignore non digits",
			input: []string{"2\n", "ab\n", "5\n", "7\n", "$\n", "stop\n"},
			want:  []int{2, 5, 7},
		},
		{
			name:  "Empty input",
			input: []string{},
			want:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := New()
			c.wg.Add(1)
			go c.Read(ctx)

			go func() {
				time.Sleep(10 * time.Millisecond)
				if len(tt.input) == 0 {
					cancel()
					return
				}
				for _, v := range tt.input {
					w.WriteString(v)
				}
			}()

			var res []int
			for v := range c.inCh {
				res = append(res, v)
			}

			assert.Equal(t, tt.want, res)
		})
	}
}

func Test_conveyor_Run(t *testing.T) {
	oldIn := os.Stdin   // keep backup of the real stdin
	oldOut := os.Stdout // keep backup of the real stdout
	defer func() {
		os.Stdin = oldIn   // restoring the real stdin
		os.Stdout = oldOut // restoring the real stdout
	}()

	tests := []struct {
		name  string
		input []string
		want  string
	}{
		{
			name:  "Valid numbers",
			input: []string{"2\n", "3\n", "5\n", "7\n", "10\n", "stop\n"},
			want: `enter a number or 'stop' to exit app: Result is: 8
enter a number or 'stop' to exit app: Result is: 18
enter a number or 'stop' to exit app: Result is: 50
enter a number or 'stop' to exit app: Result is: 98
enter a number or 'stop' to exit app: Result is: 200
enter a number or 'stop' to exit app: Graceful shutdown by 'stop' command!
`,
		},
		{
			name:  "Ignore non digits",
			input: []string{"2\n", "ab\n", "5\n", "7\n", "$\n", "stop\n"},
			want: `enter a number or 'stop' to exit app: Result is: 8
enter a number or 'stop' to exit app: bad input: strconv.Atoi: parsing "ab": invalid syntax
enter a number or 'stop' to exit app: Result is: 50
enter a number or 'stop' to exit app: Result is: 98
enter a number or 'stop' to exit app: bad input: strconv.Atoi: parsing "$": invalid syntax
enter a number or 'stop' to exit app: Graceful shutdown by 'stop' command!
`,
		},
		{
			name:  "Empty input",
			input: []string{},
			want:  "enter a number or 'stop' to exit app: ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rIn, wIn, err := os.Pipe()
			assert.NoError(t, err)

			rOut, wOut, err := os.Pipe()
			assert.NoError(t, err)

			os.Stdin = rIn
			os.Stdout = wOut

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			c := New()

			go func() {
				time.Sleep(10 * time.Millisecond)
				if len(tt.input) == 0 {
					cancel()
					return
				}
				for _, v := range tt.input {
					wIn.WriteString(v)
				}
			}()

			c.Run(ctx)

			wOut.Close()

			out, err := io.ReadAll(rOut)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(out))
		})
	}
}
