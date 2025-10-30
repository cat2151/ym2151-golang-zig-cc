package main

/*
#cgo CFLAGS: -I${SRCDIR}/lib
#cgo LDFLAGS: -L${SRCDIR}/lib -lopm
#include "opm_wrapper.h"
*/
import "C"
import (
	"encoding/binary"
	"fmt"
	"os"
	"time"
)

const (
	SampleRate = 62500 // YM2151 native clock / 64 = 4MHz / 64
	Duration   = 3     // seconds
	Frequency  = 440   // Hz (A4 note)
)

// WAV file header structure
type WAVHeader struct {
	ChunkID       [4]byte
	ChunkSize     uint32
	Format        [4]byte
	Subchunk1ID   [4]byte
	Subchunk1Size uint32
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	Subchunk2ID   [4]byte
	Subchunk2Size uint32
}

func main() {
	fmt.Println("Nuked-OPM WAV Generator")
	fmt.Printf("Generating %dHz tone for %d seconds...\n", Frequency, Duration)

	// Create OPM chip instance
	chip := C.opm_new()
	if chip == nil {
		fmt.Fprintln(os.Stderr, "Failed to create OPM chip instance")
		os.Exit(1)
	}
	defer C.opm_free(chip)

	// Initialize the chip
	initializeChip(chip)

	// Generate samples
	numSamples := SampleRate * Duration
	samples := make([]int16, numSamples*2) // stereo

	generateSamples(chip, samples)

	// Write WAV file
	if err := writeWAV("output.wav", samples); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing WAV file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully generated output.wav")
}

func initializeChip(chip *C.opm_t) {
	// Write register with 10ms cycle consumption
	writeReg := func(addr, data byte) {
		// Write address
		C.opm_write(chip, 0, C.uint8_t(addr))
		consumeCycles(chip, 10*time.Millisecond)

		// Write data
		C.opm_write(chip, 1, C.uint8_t(data))
		consumeCycles(chip, 10*time.Millisecond)
	}

	// Reset the chip
	C.OPM_Reset(chip)

	// Configure channel 0 to produce a 440Hz tone
	// YM2151 uses Key Code (KC) and Key Fraction (KF) to set frequency
	
	// For 440Hz on YM2151:
	// The formula is roughly: F = (Master Clock / 64) * 2^((KC-1)/12) * (1 + KF/64) / (2 * MUL)
	// With KC=0x4D (77), KF=0x80 (128) we can get close to 440Hz with MUL=1
	
	kc := byte(0x4D) // Key code for around 440Hz
	kf := byte(0x80) // Key fraction

	// Set connection algorithm (algorithm 0 - simple carrier)
	// RL=11 (both channels), FB=0, CON=7 (all operators as carriers)
	writeReg(0x20, 0xC7) // CH0: RL=11, FB=0, CON=7

	// Configure operator 1 (M1) of channel 0
	// DT1=0, MUL=1
	writeReg(0x40, 0x01) // OP1: DT1=0, MUL=1

	// Total Level = 0 (maximum volume)
	writeReg(0x60, 0x00) // OP1: TL=0

	// Key Scaling and Attack Rate
	// KS=0, AR=31 (fastest attack)
	writeReg(0x80, 0x1F) // OP1: KS=0, AR=31

	// AM Enable and Decay Rate 1
	// AME=0, D1R=0 (no decay)
	writeReg(0xA0, 0x00) // OP1: AME=0, D1R=0

	// Decay Rate 2
	writeReg(0xC0, 0x00) // OP1: DT2=0, D2R=0

	// Decay Level and Release Rate
	// D1L=0, RR=15 (fast release)
	writeReg(0xE0, 0x0F) // OP1: D1L=0, RR=15

	// Set Key Code and Key Fraction
	writeReg(0x28, kc) // CH0: Key Code
	writeReg(0x30, kf) // CH0: Key Fraction

	// Key On - turn on all operators for channel 0
	writeReg(0x08, 0x78) // Key On: CH=0, all operators (M1,C1,M2,C2)
}

func consumeCycles(chip *C.opm_t, duration time.Duration) {
	// YM2151 runs at 4MHz (typical master clock)
	// We need to consume cycles for the specified duration
	masterClock := int64(4000000) // 4MHz
	cyclesToConsume := int64(duration.Nanoseconds()) * masterClock / int64(time.Second)

	var output [2]C.int32_t
	var sh1, sh2, so C.uint8_t

	// Clock the chip for the specified number of cycles
	for i := int64(0); i < cyclesToConsume; i++ {
		C.opm_clock(chip, &output[0], &sh1, &sh2, &so)
	}
}

func generateSamples(chip *C.opm_t, samples []int16) {
	var output [2]C.int32_t
	var sh1, sh2, so C.uint8_t

	// The YM2151 outputs at master clock rate (4MHz)
	// We need to downsample to our desired sample rate (62.5kHz)
	clocksPerSample := 4000000 / SampleRate // 64 clocks per sample

	for i := 0; i < len(samples)/2; i++ {
		// Clock the chip multiple times per sample
		for j := 0; j < clocksPerSample; j++ {
			C.opm_clock(chip, &output[0], &sh1, &sh2, &so)
		}

		// Get stereo output
		left := int16(output[0])
		right := int16(output[1])

		samples[i*2] = left
		samples[i*2+1] = right
	}
}

func writeWAV(filename string, samples []int16) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	numChannels := uint16(2)
	bitsPerSample := uint16(16)
	byteRate := uint32(SampleRate * int(numChannels) * int(bitsPerSample) / 8)
	blockAlign := uint16(numChannels * bitsPerSample / 8)
	dataSize := uint32(len(samples) * 2) // 2 bytes per sample

	header := WAVHeader{
		ChunkID:       [4]byte{'R', 'I', 'F', 'F'},
		ChunkSize:     36 + dataSize,
		Format:        [4]byte{'W', 'A', 'V', 'E'},
		Subchunk1ID:   [4]byte{'f', 'm', 't', ' '},
		Subchunk1Size: 16,
		AudioFormat:   1, // PCM
		NumChannels:   numChannels,
		SampleRate:    uint32(SampleRate),
		ByteRate:      byteRate,
		BlockAlign:    blockAlign,
		BitsPerSample: bitsPerSample,
		Subchunk2ID:   [4]byte{'d', 'a', 't', 'a'},
		Subchunk2Size: dataSize,
	}

	// Write header
	if err := binary.Write(file, binary.LittleEndian, &header); err != nil {
		return err
	}

	// Write samples
	if err := binary.Write(file, binary.LittleEndian, samples); err != nil {
		return err
	}

	return nil
}
