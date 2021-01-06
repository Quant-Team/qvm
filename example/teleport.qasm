OPENQASM 3;

include "stdgates.inc";

gate post q { }

qubit qb[3];

bit c0;
bit c1;
bit c2;

reset qb;

h qb[1];
cx qb[1], qb[2];

barrier qb;

cx qb[0], qb[1];
h qb[0];

c0 = measure qb[0];
c1 = measure qb[1];

if(c0==1) { z qb[2]; }
if(c1==1) { x qb[2]; }

post qb[2];

c2 = measure qb[2];
