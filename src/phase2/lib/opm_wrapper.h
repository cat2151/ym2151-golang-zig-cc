#ifndef OPM_WRAPPER_H
#define OPM_WRAPPER_H

#include "nuked-opm/opm.h"

opm_t* opm_new();
void opm_free(opm_t* chip);
void opm_write(opm_t* chip, uint32_t port, uint8_t data);
void opm_clock(opm_t* chip, int32_t* output, uint8_t* sh1, uint8_t* sh2, uint8_t* so);
uint32_t opm_get_cycles(opm_t* chip);

#endif // OPM_WRAPPER_H
