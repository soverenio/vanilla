package promise

import (
	"context"
	"reflect"

	"github.com/soverenio/vanilla/throw"
	"github.com/soverenio/vanilla/zero"
)

func select1[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 1 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	}
}

func select2[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 2 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	}
}

func select3[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 3 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	}
}

func select4[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 4 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	}
}

func select5[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 5 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	}
}

func select6[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 6 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	}
}

func select7[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 7 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	case val, ok := <-inp[6]:
		return 6, val, ok
	}
}

func select8[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 8 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	case val, ok := <-inp[6]:
		return 6, val, ok
	case val, ok := <-inp[7]:
		return 7, val, ok
	}
}

func select9[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 9 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	case val, ok := <-inp[6]:
		return 6, val, ok
	case val, ok := <-inp[7]:
		return 7, val, ok
	case val, ok := <-inp[8]:
		return 8, val, ok
	}
}

func select10[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 10 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	case val, ok := <-inp[6]:
		return 6, val, ok
	case val, ok := <-inp[7]:
		return 7, val, ok
	case val, ok := <-inp[8]:
		return 8, val, ok
	case val, ok := <-inp[9]:
		return 9, val, ok
	}
}

func select11[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 11 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	case val, ok := <-inp[6]:
		return 6, val, ok
	case val, ok := <-inp[7]:
		return 7, val, ok
	case val, ok := <-inp[8]:
		return 8, val, ok
	case val, ok := <-inp[9]:
		return 9, val, ok
	case val, ok := <-inp[10]:
		return 10, val, ok
	}
}

func select12[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 12 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	case val, ok := <-inp[6]:
		return 6, val, ok
	case val, ok := <-inp[7]:
		return 7, val, ok
	case val, ok := <-inp[8]:
		return 8, val, ok
	case val, ok := <-inp[9]:
		return 9, val, ok
	case val, ok := <-inp[10]:
		return 10, val, ok
	case val, ok := <-inp[11]:
		return 11, val, ok
	}
}

func select13[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 13 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	case val, ok := <-inp[6]:
		return 6, val, ok
	case val, ok := <-inp[7]:
		return 7, val, ok
	case val, ok := <-inp[8]:
		return 8, val, ok
	case val, ok := <-inp[9]:
		return 9, val, ok
	case val, ok := <-inp[10]:
		return 10, val, ok
	case val, ok := <-inp[11]:
		return 11, val, ok
	case val, ok := <-inp[12]:
		return 12, val, ok
	}
}

func select14[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 14 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	case val, ok := <-inp[6]:
		return 6, val, ok
	case val, ok := <-inp[7]:
		return 7, val, ok
	case val, ok := <-inp[8]:
		return 8, val, ok
	case val, ok := <-inp[9]:
		return 9, val, ok
	case val, ok := <-inp[10]:
		return 10, val, ok
	case val, ok := <-inp[11]:
		return 11, val, ok
	case val, ok := <-inp[12]:
		return 12, val, ok
	case val, ok := <-inp[13]:
		return 13, val, ok
	}
}

func select15[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 15 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	case val, ok := <-inp[6]:
		return 6, val, ok
	case val, ok := <-inp[7]:
		return 7, val, ok
	case val, ok := <-inp[8]:
		return 8, val, ok
	case val, ok := <-inp[9]:
		return 9, val, ok
	case val, ok := <-inp[10]:
		return 10, val, ok
	case val, ok := <-inp[11]:
		return 11, val, ok
	case val, ok := <-inp[12]:
		return 12, val, ok
	case val, ok := <-inp[13]:
		return 13, val, ok
	case val, ok := <-inp[14]:
		return 14, val, ok
	}
}

func select16[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	if len(inp) != 16 {
		panic(throw.IllegalValue())
	}

	select {
	case <-ctx.Done():
		return -1, zero.Zero[T](), false
	case val, ok := <-inp[0]:
		return 0, val, ok
	case val, ok := <-inp[1]:
		return 1, val, ok
	case val, ok := <-inp[2]:
		return 2, val, ok
	case val, ok := <-inp[3]:
		return 3, val, ok
	case val, ok := <-inp[4]:
		return 4, val, ok
	case val, ok := <-inp[5]:
		return 5, val, ok
	case val, ok := <-inp[6]:
		return 6, val, ok
	case val, ok := <-inp[7]:
		return 7, val, ok
	case val, ok := <-inp[8]:
		return 8, val, ok
	case val, ok := <-inp[9]:
		return 9, val, ok
	case val, ok := <-inp[10]:
		return 10, val, ok
	case val, ok := <-inp[11]:
		return 11, val, ok
	case val, ok := <-inp[12]:
		return 12, val, ok
	case val, ok := <-inp[13]:
		return 13, val, ok
	case val, ok := <-inp[14]:
		return 14, val, ok
	case val, ok := <-inp[15]:
		return 15, val, ok
	}
}

func selectReflect[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	chans := make([]reflect.SelectCase, len(inp)+1)
	chans[0] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ctx.Done())}
	for i, c := range inp {
		chans[i+1] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(c)}
	}
	i, val, ok := reflect.Select(chans)
	return i - 1, val.Interface().(T), ok
}

func SelectSlice[T any](ctx context.Context, inp []<-chan T) (int, T, bool) {
	switch len(inp) {
	case 0:
		panic(throw.IllegalValue())
	case 1:
		return select1(ctx, inp)
	case 2:
		return select2(ctx, inp)
	case 3:
		return select3(ctx, inp)
	case 4:
		return select4(ctx, inp)
	case 5:
		return select5(ctx, inp)
	case 6:
		return select6(ctx, inp)
	case 7:
		return select7(ctx, inp)
	case 8:
		return select8(ctx, inp)
	case 9:
		return select9(ctx, inp)
	case 10:
		return select10(ctx, inp)
	case 11:
		return select11(ctx, inp)
	case 12:
		return select12(ctx, inp)
	case 13:
		return select13(ctx, inp)
	case 14:
		return select14(ctx, inp)
	case 15:
		return select15(ctx, inp)
	case 16:
		return select16(ctx, inp)
	default:
		return selectReflect(ctx, inp)
	}
}

func channelBiToUni[T any](inp []chan T) []<-chan T {
	out := make([]<-chan T, len(inp))
	for i, c := range inp {
		out[i] = c
	}
	return out
}

func SelectSliceBi[T any](ctx context.Context, inp []chan T) (int, T, bool) {
	return SelectSlice(ctx, channelBiToUni(inp))
}
