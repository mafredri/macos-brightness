package main

/*
#cgo LDFLAGS: -framework CoreDisplay -framework CoreGraphics

#include <CoreGraphics/CoreGraphics.h>

// Private APIs for brightness.
double CoreDisplay_Display_GetUserBrightness(CGDirectDisplayID id);
void CoreDisplay_Display_SetUserBrightness(CGDirectDisplayID id, double b);
*/
import "C"

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type Display struct {
	ID int
}

func (d Display) cid() C.CGDirectDisplayID {
	return C.CGDirectDisplayID(d.ID)
}

func (d Display) Brightness() float64 {
	v := float64(C.CoreDisplay_Display_GetUserBrightness(d.cid()))
	return v * 100
}

func (d Display) SetBrightness(v float64) {
	v /= 100
	C.CoreDisplay_Display_SetUserBrightness(d.cid(), C.double(v))
}

func (d Display) Builtin() bool {
	v := C.CGDisplayIsBuiltin(d.cid())
	return v != 0
}

func displays() ([]Display, error) {
	num := C.uint32_t(0)
	var online [16]C.uint32_t
	var cgerr C.CGError = C.CGGetOnlineDisplayList(16, &online[0], &num)
	if cgerr != 0 {
		return nil, fmt.Errorf("CGGetOnlineDisplayList failed with %d", cgerr)
	}

	var disp []Display
	for _, id := range online[:num] {
		disp = append(disp, Display{
			ID: int(id),
		})
	}

	return disp, nil
}

func findBuiltin(d []Display) (Display, error) {
	for _, dd := range d {
		if dd.Builtin() {
			return dd, nil
		}
	}

	return Display{}, errors.New("builtin display not found")
}

func main() {
	brightness := flag.Float64("b", -1, "Set the brightness (between 0 - 100)")
	flag.Parse()

	d, err := displays()
	if err != nil {
		panic(err)
	}

	bd, err := findBuiltin(d)
	if err != nil {
		panic(err)
	}

	if *brightness == -1 {
		fmt.Println(bd.Brightness())
		os.Exit(0)
	}

	bd.SetBrightness(*brightness)
}
