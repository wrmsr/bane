#include <stdlib.h>

#include "foo.h"

struct point_s *get_point() {
    struct point_s *p = (struct point_s *)malloc(sizeof(point));
    p->x = 420;
    p->y = 69;
    return p;
}
