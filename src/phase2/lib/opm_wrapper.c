#include "opm_wrapper.h"
#include <stdlib.h>
#include <string.h>

// Wrapper functions to make it easier to work with from Go

opm_t* opm_new() {
    opm_t* chip = (opm_t*)malloc(sizeof(opm_t));
    if (chip != NULL) {
        memset(chip, 0, sizeof(opm_t));
        OPM_Reset(chip);
    }
    return chip;
}

void opm_free(opm_t* chip) {
    if (chip != NULL) {
        free(chip);
    }
}

void opm_write(opm_t* chip, uint32_t port, uint8_t data) {
    OPM_Write(chip, port, data);
}

void opm_clock(opm_t* chip, int32_t* output, uint8_t* sh1, uint8_t* sh2, uint8_t* so) {
    OPM_Clock(chip, output, sh1, sh2, so);
}

uint32_t opm_get_cycles(opm_t* chip) {
    return chip->cycles;
}
